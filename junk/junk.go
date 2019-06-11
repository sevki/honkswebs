package junk

import (
	"encoding/json"
	"io"
	"strconv"
)

type Junk map[string]interface{}

func New() Junk {
	return make(map[string]interface{})
}

func (j Junk) Write(w io.Writer) error {
	e := json.NewEncoder(w)
	e.SetEscapeHTML(false)
	e.SetIndent("", "  ")
	err := e.Encode(j)
	return err
}

func Read(r io.Reader) (Junk, error) {
	decoder := json.NewDecoder(r)
	var j Junk
	err := decoder.Decode(&j)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func jsonfindinterface(ii interface{}, keys []string) interface{} {
	for _, key := range keys {
		idx, err := strconv.Atoi(key)
		if err == nil {
			m := ii.([]interface{})
			if idx >= len(m) {
				return nil
			}
			ii = m[idx]
		} else {
			m, ok := ii.(map[string]interface{})
			if !ok {
				m = ii.(Junk)
			}
			ii = m[key]
			if ii == nil {
				return nil
			}
		}
	}
	return ii
}
func (j Junk) FindString(keys []string) (string, bool) {
	s, ok := jsonfindinterface(j, keys).(string)
	return s, ok
}
func (j Junk) FindArray(keys []string) ([]interface{}, bool) {
	a, ok := jsonfindinterface(j, keys).([]interface{})
	if ok {
		for i, ii := range a {
			j, ok := ii.(map[string]interface{})
			if ok {
				a[i] = Junk(j)
			}
		}
	}
	return a, ok
}
func (j Junk) FindMap(keys []string) (Junk, bool) {
	ii := jsonfindinterface(j, keys)
	m, ok := ii.(map[string]interface{})
	if !ok {
		m, ok = ii.(Junk)
	}
	return m, ok
}
func (j Junk) GetString(key string) (string, bool) {
	return j.FindString([]string{key})
}
func (j Junk) GetArray(key string) ([]interface{}, bool) {
	return j.FindArray([]string{key})
}
func (j Junk) GetMap(key string) (Junk, bool) {
	return j.FindMap([]string{key})
}
