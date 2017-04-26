package registrant

import (
	"net/http"

	"github.com/zneyrl/nmsrs/helpers/tpl"
)

func CertificationAuthorization(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Certification/Authorization",
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "dashboard", "registrant.create.certification_authorization", data, funcMap)
}

func UpdateCertificationAuthorization(w http.ResponseWriter, r *http.Request) {
}
