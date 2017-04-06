package home

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Home",
	}
	tmpl.RenderWithFunc(w, "main", "home.index", data, nil)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Welcome",
	}
	tmpl.RenderWithFunc(w, "main", "home.welcome", data, nil)
}
