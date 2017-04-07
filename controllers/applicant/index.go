package applicant

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Applicants",
		"R":     r,
	}
	tmpl.RenderWithFunc(w, "dashboard", "applicant.index", data, nil)
}