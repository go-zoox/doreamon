package image

import "net/http"

type ImageSVG struct {
	Path string
}

func (i ImageSVG) Type() string {
	return "image/svg+xml"
}

func (i ImageSVG) File() File {
	return File{Type: i.Type(), Path: i.Path}
}

func (i ImageSVG) Render(w http.ResponseWriter) error {
	return i.File().Render(w)
}

func (i ImageSVG) WriteContentType(w http.ResponseWriter) {
	i.File().WriteContentType(w)
}
