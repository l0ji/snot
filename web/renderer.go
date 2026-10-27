package web

import (
	"embed"
	"html/template"
)

//go:embed templates/*.tmpl templates/layout/*.tmpl
var templateFS embed.FS

type Renderer struct {
	templates *template.Template
}

func NewRenderer() (*Renderer, error) {
	t, err := template.ParseFS(templateFS,
		"templates/*.tmpl",
		"templates/layout/*.tmpl",
	)
	if err != nil {
		return nil, err
	}

	return &Renderer{templates: t}, nil
}

func (r *Renderer) Render(
	w interface{ Write([]byte) (int, error) },
	name string,
	data any,
) error {
	return r.templates.ExecuteTemplate(w, name, data)
}
