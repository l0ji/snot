package web

import (
	"net/http"
	"snot/config"
)

type HomeData struct {
	Title        string
	InstanceName string
}

func (r *Renderer) Home(c config.Config, w http.ResponseWriter, _ *http.Request) {
	data := HomeData{
		Title:        "snot",
		InstanceName: c.InstanceName,
	}

	if err := r.Render(w, "base", data); err != nil {
		http.Error(w, "Internal server error",
			http.StatusInternalServerError)
	}
}
