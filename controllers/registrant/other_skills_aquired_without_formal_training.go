package registrant

import (
	"net/http"

	"github.com/zneyrl/nmsrs/helpers/tpl"
)

func OtherSkillsAquiredWithoutFormalTraining(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Other Skills Aquired Without Formal Training",
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "dashboard", "registrant.create.other_skills_aquired_without_formal_training", data, funcMap)
}

func UpdateOtherSkillsAquiredWithoutFormalTraining(w http.ResponseWriter, r *http.Request) {
}
