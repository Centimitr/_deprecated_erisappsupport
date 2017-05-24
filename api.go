package main

import (
	"archive/tar"
	"encoding/json"
	"erisapp.com/entity"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var r entity.Reader

type API struct{}

func init() {
	r.Init()
}

func (a *API) GetBook(locator string, w io.Writer) {
	book, err := r.Get(locator)
	if err != nil {
		fmt.Println(500)
		return
	}
	content, _ := json.Marshal(book.Meta)
	w.Write(content)
}

func (a *API) GetBookPage(locator string, pageLocator string, w io.Writer) {
	book, err := r.Get(locator)
	if err != nil {
		fmt.Println(500)
		return
	}
	pm := book.Meta.GetPage(pageLocator)
	img, _ := pm.GetFile()
	io.Copy(w, img)
	img.Close()
}

func (a *API) MakeErisFromFolder(path string, dst string) {

	name := ""
	text, err := ioutil.ReadFile(filepath.Join(path, "book.json"))
	bm := &entity.BookMeta{}
	if err == nil {
		if err := json.Unmarshal(text, &bm); err == nil {
			name = bm.Name
		}
	}
	if name == "" {
		name = string(time.Now().UnixNano())
	}
	name += ".eris"
	fmt.Println(name)

	os.MkdirAll(dst, 0777)
	f, _ := os.Create(filepath.Join(dst, name))
	defer f.Close()
	tw := tar.NewWriter(f)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name(),
			Mode: 0600,
			Size: file.Size(),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		bytes, err := ioutil.ReadFile(filepath.Join(path, file.Name()))
		if err != nil {
			log.Fatalln(err)
		}
		if _, err := tw.Write(bytes); err != nil {
			log.Fatalln(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}
}
