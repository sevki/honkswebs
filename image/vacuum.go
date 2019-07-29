//
// Copyright (c) 2019 Ted Unangst <tedu@tedunangst.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// basic image manipulation
package image

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"math"

	_ "golang.org/x/image/webp"
)

type Image struct {
	Data   []byte
	Format string
	Width  int
	Height int
}

type Params struct {
	MaxWidth  int
	MaxHeight int
	MaxSize   int
}

const dirLeft = 1
const dirRight = 2

func fixrotation(img image.Image, dir int) image.Image {
	w, h := img.Bounds().Max.X, img.Bounds().Max.Y
	newimg := image.NewRGBA(image.Rectangle{Max: image.Point{X: h, Y: w}})
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			c := img.At(i, j)
			if dir == dirLeft {
				newimg.Set(j, w-i-1, c)
			} else {
				newimg.Set(h-j-1, i, c)
			}
		}
	}
	return newimg
}

var rotateLeftSigs = [][]byte{
	{0x01, 0x12, 0x00, 0x03, 0x00, 0x00, 0x00, 0x01, 0x00, 0x08},
	{0x12, 0x01, 0x03, 0x00, 0x01, 0x00, 0x00, 0x00, 0x08, 0x00},
}
var rotateRightSigs = [][]byte{
	{0x12, 0x01, 0x03, 0x00, 0x01, 0x00, 0x00, 0x00, 0x06, 0x00},
	{0x01, 0x12, 0x00, 0x03, 0x00, 0x00, 0x00, 0x01, 0x00, 0x06},
}

// Read an image and shrink it down to web scale
func Vacuum(reader io.Reader, params Params) (*Image, error) {
	var tmpbuf bytes.Buffer
	io.CopyN(&tmpbuf, reader, 256)
	peek := tmpbuf.Bytes()
	img, format, err := image.Decode(io.MultiReader(&tmpbuf, reader))
	if err != nil {
		return nil, err
	}

	if params.MaxWidth == 0 {
		params.MaxWidth = 16000
	}
	if params.MaxHeight == 0 {
		params.MaxHeight = 16000
	}
	if params.MaxSize == 0 {
		params.MaxSize = 512 * 1024
	}

	for img.Bounds().Max.X > params.MaxWidth || img.Bounds().Max.Y > params.MaxHeight {
		switch oldimg := img.(type) {
		case *image.NRGBA:
			w, h := oldimg.Rect.Max.X/2, oldimg.Rect.Max.Y/2
			newimg := image.NewNRGBA(image.Rectangle{Max: image.Point{X: w, Y: h}})
			for j := 0; j < h; j++ {
				for i := 0; i < w; i++ {
					p := newimg.Stride*j + i*4
					q1 := oldimg.Stride*(j*2+0) + i*4*2
					q2 := oldimg.Stride*(j*2+1) + i*4*2
					newimg.Pix[p+0] = blend(oldimg.Pix, q1+0, q1+4, q2+0, q2+4)
					newimg.Pix[p+1] = blend(oldimg.Pix, q1+1, q1+5, q2+1, q2+5)
					newimg.Pix[p+2] = blend(oldimg.Pix, q1+2, q1+6, q2+2, q2+6)
					newimg.Pix[p+3] = squish(oldimg.Pix, q1+3, q1+7, q2+3, q2+7)
				}
			}
			img = newimg
		case *image.RGBA:
			w, h := oldimg.Rect.Max.X/2, oldimg.Rect.Max.Y/2
			newimg := image.NewRGBA(image.Rectangle{Max: image.Point{X: w, Y: h}})
			for j := 0; j < h; j++ {
				for i := 0; i < w; i++ {
					p := newimg.Stride*j + i*4
					q1 := oldimg.Stride*(j*2+0) + i*4*2
					q2 := oldimg.Stride*(j*2+1) + i*4*2
					newimg.Pix[p+0] = blend(oldimg.Pix, q1+0, q1+4, q2+0, q2+4)
					newimg.Pix[p+1] = blend(oldimg.Pix, q1+1, q1+5, q2+1, q2+5)
					newimg.Pix[p+2] = blend(oldimg.Pix, q1+2, q1+6, q2+2, q2+6)
					newimg.Pix[p+3] = squish(oldimg.Pix, q1+3, q1+7, q2+3, q2+7)
				}
			}
			img = newimg
		case *image.YCbCr:
			w, h := oldimg.Rect.Max.X/2, oldimg.Rect.Max.Y/2
			newimg := image.NewYCbCr(image.Rectangle{Max: image.Point{X: w, Y: h}},
				oldimg.SubsampleRatio)
			for j := 0; j < h; j++ {
				for i := 0; i < w; i++ {
					p := newimg.YStride*j + i
					q1 := oldimg.YStride*(j*2+0) + i*2
					q2 := oldimg.YStride*(j*2+1) + i*2
					newimg.Y[p+0] = blend(oldimg.Y, q1+0, q1+1, q2+0, q2+1)
				}
			}
			switch newimg.SubsampleRatio {
			case image.YCbCrSubsampleRatio444:
				w, h = w, h
			case image.YCbCrSubsampleRatio422:
				w, h = w/2, h
			case image.YCbCrSubsampleRatio420:
				w, h = w/2, h/2
			case image.YCbCrSubsampleRatio440:
				w, h = w, h/2
			case image.YCbCrSubsampleRatio411:
				w, h = w/4, h
			case image.YCbCrSubsampleRatio410:
				w, h = w/4, h/2
			}
			for j := 0; j < h; j++ {
				for i := 0; i < w; i++ {
					p := newimg.CStride*j + i
					q1 := oldimg.CStride*(j*2+0) + i*2
					q2 := oldimg.CStride*(j*2+1) + i*2
					newimg.Cb[p+0] = blend(oldimg.Cb, q1+0, q1+1, q2+0, q2+1)
					newimg.Cr[p+0] = blend(oldimg.Cr, q1+0, q1+1, q2+0, q2+1)
				}
			}
			img = newimg
		default:
			// convert to RGBA, then loop and try again
			w, h := oldimg.Bounds().Max.X, oldimg.Bounds().Max.Y
			newimg := image.NewRGBA(image.Rectangle{Max: image.Point{X: w, Y: h}})
			for j := 0; j < h; j++ {
				for i := 0; i < w; i++ {
					c := oldimg.At(i, j)
					newimg.Set(i, j, c)
				}
			}
			img = newimg
		}
	}
	if format == "jpeg" {
		for _, sig := range rotateLeftSigs {
			if bytes.Contains(peek, sig) {
				img = fixrotation(img, dirLeft)
				break
			}
		}
		for _, sig := range rotateRightSigs {
			if bytes.Contains(peek, sig) {
				img = fixrotation(img, dirRight)
				break
			}
		}
	}
	quality := 80
	var buf bytes.Buffer
	for {
		switch format {
		case "gif":
			format = "png"
			png.Encode(&buf, img)
		case "png":
			png.Encode(&buf, img)
		case "webp":
			format = "jpeg"
			fallthrough
		case "jpeg":
			jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality})
		default:
			return nil, fmt.Errorf("can't encode format: %s", format)
		}
		if buf.Len() > params.MaxSize && quality > 30 {
			switch format {
			case "png":
				format = "jpeg"
			case "jpeg":
				quality -= 10
			}
			buf.Reset()
			continue
		}
		break
	}
	rv := &Image{
		Data:   buf.Bytes(),
		Format: format,
		Width:  img.Bounds().Max.X,
		Height: img.Bounds().Max.Y,
	}
	return rv, nil
}

func lineate(s uint8) float64 {
	x := float64(s)
	x /= 255.0
	if x < 0.04045 {
		x /= 12.92
	} else {
		x += 0.055
		x /= 1.055
		x = math.Pow(x, 2.4)
	}
	return x
}

func delineate(x float64) uint8 {
	if x > 0.0031308 {
		x = math.Pow(x, 1/2.4)
		x *= 1.055
		x -= 0.055
	} else {
		x *= 12.92
	}
	x *= 255.0
	return uint8(x)
}

func blend(d []byte, s1, s2, s3, s4 int) byte {
	l1 := lineate(d[s1])
	l2 := lineate(d[s2])
	l3 := lineate(d[s3])
	l4 := lineate(d[s4])
	return delineate((l1 + l2 + l3 + l4) / 4.0)
}

func squish(d []byte, s1, s2, s3, s4 int) byte {
	return uint8((uint32(d[s1]) + uint32(d[s2]) + uint32(d[s3]) + uint32(d[s4])) / 4)
}
