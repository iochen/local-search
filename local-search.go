package local_search

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"strings"
)

type Entry struct {
	Title string
	URL   string
	Key   []string
}

type Data struct {
	Entries []*Entry
}

var FieldsSep = map[rune]bool{
	// ascii spaces
	'\t': true, '\n': true, '\v': true,
	'\f': true, '\r': true, ' ': true,

	// english signs
	'~': true, '!': true, '@': true,
	'#': true, '$': true, '%': true,
	'^': true, '&': true, '*': true,
	'(': true, ')': true, '_': true,
	'+': true, '|': true, '}': true,
	'{': true, ':': true, '"': true,
	'<': true, '>': true, '?': true,
	'`': true, '-': true, '=': true,
	'[': true, ']': true, '\\': true,
	';': true, '\'': true, ',': true,
	'.': true, '/': true,

	// chinese signs
	'～': true, '！': true, '￥': true,
	'…': true, '×': true, '（': true,
	'）': true, '：': true, '“': true,
	'”': true, '《': true, '》': true,
	'？': true, '、': true, '·': true,
	'」': true, '；': true, '，': true,
	'。': true,
}

// Parse parses the file passed in
// and set Title and Key fields in Entry
//
// NOTE: set URL field out of this function
// NOTE: title would be included in Key field
func (e *Entry) Parse(r io.Reader) error {
	reader := bufio.NewReader(r)

	// parse data in
	// ---
	// title: my app
	// meta: data
	// ---
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return errors.New(`cannot detect "---"`)
			}
			return err
		}
		line = strings.TrimSpace(line)
		lineS := line
		if lineS == "---" {
			break
		}
	}
	var titleDetected bool
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return errors.New(`cannot detect end "---"`)
			}
			return err
		}
		line = strings.TrimSpace(line)
		lineS := line
		if lineS == "---" {
			break
		}
		if strings.HasPrefix(lineS,"title:") {
			titleDetected = true
			line = strings.TrimPrefix(line,"title:")
			e.Title = strings.TrimSpace(line)
		}
	}
	if !titleDetected {
		return errors.New("cannot detect title in meta")
	}

	// parse main body
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	e.Key = strings.FieldsFunc(string(body), func(r rune) bool {
		return FieldsSep[r]
	})
	e.Key = append(e.Key,e.Title)
	return nil
}

func (e *Entry) DropDuplicated() {
	var cleaned []string
	m := map[string]bool{}
	for i :=0;i<len(e.Key);i++ {
		if m[e.Key[i]] {
			continue
		}
		cleaned = append(cleaned,e.Key[i])
		m[e.Key[i]] = true
	}
	e.Key = cleaned
}

func (e *Entry)ToLower() {
	for i:=0;i<len(e.Key);i++ {
		e.Key[i] = strings.ToLower(e.Key[i])
	}
}

func (e *Entry)Json() ([]byte,error) {
	return json.Marshal(e)
}
