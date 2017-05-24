package entity

import (
	"io"
	"path/filepath"
	"strings"
)

type Driver interface {
	readKey(b *Book) (string, error)
	readList(b *Book) ([]string, error)
	readMeta(b *Book) (BookMeta, error)
	readPage(b *Book, id string) (io.ReadCloser, error)
}

const SUPPORTED_BOOK_EXT = ".eris.rar.cbr.zip.cbz"
const SUPPORTED_IMAGE_EXT = ".jpg.png.gif.jpeg.webp"

func isFormatSupported(path string, candidates string) bool {
	ext := filepath.Ext(path)
	return ext != "" && strings.Contains(candidates, ext)
}

func isBookSupported(path string) bool {
	return isFormatSupported(path, SUPPORTED_BOOK_EXT)
}

func isImageSupported(path string) bool {
	return isFormatSupported(path, SUPPORTED_IMAGE_EXT)
}
