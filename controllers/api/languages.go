package api

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/models/language"
)

func Languages(w http.ResponseWriter, r *http.Request) {
	lngs, err := language.Search(bson.M{
		"name": bson.RegEx{
			Pattern: r.URL.Query().Get("q"),
			Options: "i",
		},
	})

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]interface{}{
			"languages": lngs,
		},
		Errors: "",
	})
}
