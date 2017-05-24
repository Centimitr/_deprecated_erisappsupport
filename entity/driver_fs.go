package entity

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FSDriver struct {
}

func (d FSDriver) readKey(b *Book) (string, error) {
	return "", nil
}

func (d FSDriver) readList(b *Book) (paths []string, err error) {
	filepath.Walk(b.Locator, func(path string, info os.FileInfo, err error) error {
		if isImageSupported(path) {
			if !info.IsDir() {
				p, err := filepath.Rel(b.Locator, path)
				if err == nil {
					paths = append(paths, p)
				}
			}
		}
		return nil
	})
	return
}

func (d FSDriver) readMeta(b *Book) (m BookMeta, err error) {
	data, err := ioutil.ReadFile(filepath.Join(b.Locator, "book.json"))
	if err == nil {
		err = json.Unmarshal(data, &m)
	}
	return
}

func (d FSDriver) readPage(b *Book, id string) (io.ReadCloser, error) {
	return os.Open(filepath.Join(b.Locator, id))
}
