package image

import "net/http"

type ImageJPG struct {
	Path string
}

func (i ImageJPG) Type() string {
	return "image/jpeg"
}

func (i ImageJPG) File() File {
	return File{Type: i.Type(), Path: i.Path}
}

func (i ImageJPG) Render(w http.ResponseWriter) error {
	return i.File().Render(w)
}

func (i ImageJPG) WriteContentType(w http.ResponseWriter) {
	i.File().WriteContentType(w)
}
