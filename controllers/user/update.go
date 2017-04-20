package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zneyrl/nmsrs-lookup/helpers/res"
	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
	"github.com/zneyrl/nmsrs-lookup/helpers/trans"
	"github.com/zneyrl/nmsrs-lookup/models/user"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	usr, err := user.Find(v["id"])

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	data := map[string]interface{}{
		"Title": "Edit User",
		"User":  usr,
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "dashboard", "user.edit", data, funcMap)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil { // TODO: Must be multipart
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	var profile user.Profile

	if err := decoder.Decode(&profile, r.PostForm); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	errs := trans.StructHasError(profile)

	if len(errs) != 0 {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: errs,
		})
		return
	}
	v := mux.Vars(r)
	id := v["id"]
	sameAsOld, err := user.CheckEmailIfSameAsOld(id, profile.Email)

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}

	if !sameAsOld {
		if err := user.CheckEmailIfTaken(profile.Email); err != nil {
			res.JSON(w, res.Make{
				Status: http.StatusForbidden,
				Data:   "",
				Errors: map[string]interface{}{
					"email": err.Error(),
				},
			})
			return
		}
	}

	if err := user.UpdateProfile(id, profile); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]string{
			"message": "User has been successfully updated",
		},
		Errors: "",
	})
	return
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	var resetPassword user.ResetPassword

	if err := decoder.Decode(&resetPassword, r.PostForm); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	errs := trans.StructHasError(resetPassword)

	if len(errs) != 0 {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: errs,
		})
		return
	}
	v := mux.Vars(r)
	id := v["id"]

	if err := user.UpdatePassword(id, resetPassword); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]string{
			"message": "Password has been successfully updated",
		},
		Errors: "",
	})
	return
}