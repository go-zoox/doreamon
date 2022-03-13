package image

import "net/http"

type ImagePNG struct {
	Path string
}

func (i ImagePNG) Type() string {
	return "image/png"
}

func (i ImagePNG) File() File {
	return File{Type: i.Type(), Path: i.Path}
}

func (i ImagePNG) Render(w http.ResponseWriter) error {
	return i.File().Render(w)
}

func (i ImagePNG) WriteContentType(w http.ResponseWriter) {
	i.File().WriteContentType(w)
}
