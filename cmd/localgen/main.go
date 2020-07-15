package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	ls "github.com/iochen/local-search"
)

var (
	FlagPath = flag.String("p","./","path to walk")
	FlagOut = flag.String("o","./ls.json","output path")
	FlagClean = flag.Bool("c",true,"drop duplicated keys")
	FlagDir = flag.String("d","/post","url dir base")
	FlagLower = flag.Bool("l",true,"key to lower")
)

func Load(path string) (*ls.Entry,error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return &ls.Entry{}, err
	}
	e := &ls.Entry{}
	if err := e.Parse(file); err != nil {
		return &ls.Entry{}, err
	}
	p := filepath.Base(path)
	p = strings.TrimSuffix(p,filepath.Ext(p))
	e.URL = filepath.Join(*FlagDir,p)
	if *FlagClean {
		e.DropDuplicated()
	}
	if *FlagLower {
		e.ToLower()
	}
	return e,nil
}

func main() {
	flag.Parse()
	data := new(ls.Data)

	err := filepath.Walk(*FlagPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) !=".md" {
			return nil
		}
		entry, err := Load(path)
		if err != nil {
			return err
		}
		data.Entries = append(data.Entries,entry)
		return nil
	})
	if err != nil {
		log.Println(err)
		return
	}

	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(*FlagOut, b, 0644)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("data has written to %s",filepath.Clean(*FlagOut))
}
