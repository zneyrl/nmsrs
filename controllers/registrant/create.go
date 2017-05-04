package registrant

import (
	"net/http"
	"strconv"

	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/helpers/str"
	"github.com/emurmotol/nmsrs/helpers/tpl"
	"github.com/emurmotol/nmsrs/models/citymunicipality"
	"github.com/emurmotol/nmsrs/models/civilstatus"
	"github.com/emurmotol/nmsrs/models/employmentstatus"
	"github.com/emurmotol/nmsrs/models/province"
	"github.com/emurmotol/nmsrs/models/religion"
	"github.com/emurmotol/nmsrs/models/sex"
)

func Create(w http.ResponseWriter, r *http.Request) {
	empStats, err := employmentstatus.All()

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	sexs, err := sex.All()

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	civStats, err := civilstatus.All()

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	religs, err := religion.All()

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	cityMuns, err := citymunicipality.All()

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	provs, err := province.All()

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}

	data := map[string]interface{}{
		"Title":            "Create registrant",
		"EmploymentStatus": empStats,
		"Sex":              sexs,
		"CivilStatus":      civStats,
		"Religion":         religs,
		"CityMunicipality": cityMuns,
		"Province":         provs,
	}
	funcMap := map[string]interface{}{
		"SentenceCaseToSnakeCase": str.SentenceCaseToSnakeCase,
		"AllCapsToSentenceCase":   str.AllCapsToSentenceCase,
	}
	tpl.Render(w, r, "wizard", "registrant.create", data, funcMap)
}

func Store(w http.ResponseWriter, r *http.Request) {
	step, err := strconv.Atoi(r.URL.Query()["step"][0])

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}

	switch step {
	case 1:
		res.JSON(w, res.Make{
			Status: http.StatusOK,
			Data: map[string]interface{}{
				"proceed": true,
			},
			Errors: "",
		})
		return
	case 2:
		return
	case 3:
		return
	case 4:
		return
	case 5:
		return
	case 6:
		return
	case 7:
		return
	case 8:
		return
	case 9:
		return
	default:
		return
	}
}
