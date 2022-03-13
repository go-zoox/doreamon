package image

import (
	"io"
	"net/http"
	"os"
)

type File struct {
	Type string
	Path string
}

func (i File) Render(w http.ResponseWriter) error {
	// @TODO force write content type by self
	i.WriteContentType(w)

	file, err := os.Open(i.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(w, file)
	return err
}

func (i File) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{i.Type}
	}

	// writeContentType(w, []string{i.Type})
}
