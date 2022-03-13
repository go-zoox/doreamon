package image

import "net/http"

type ImageWEBP struct {
	Path string
}

func (i ImageWEBP) Type() string {
	return "image/webp"
}

func (i ImageWEBP) File() File {
	return File{Type: i.Type(), Path: i.Path}
}

func (i ImageWEBP) Render(w http.ResponseWriter) error {
	return i.File().Render(w)
}

func (i ImageWEBP) WriteContentType(w http.ResponseWriter) {
	i.File().WriteContentType(w)
}
