package controller

import (
	"net/http"

	"github.com/emurmotol/nmsrs/model"
)

func BarangayIndex(w http.ResponseWriter, r *http.Request) {
	// todo
}

func CertificateIndex(w http.ResponseWriter, r *http.Request) {
	// todo
}

func CityMunWithProvinceIndex(w http.ResponseWriter, r *http.Request) {
	rd.JSON(w, http.StatusOK, model.CityMunWithProvince(r.URL.Query().Get("q")))
}

func CityMunIndex(w http.ResponseWriter, r *http.Request) {
	// todo
}

func CountryIndex(w http.ResponseWriter, r *http.Request) {
	// todo
}

func CourseIndex(w http.ResponseWriter, r *http.Request) {
	// todo
}

func EduLevelIndex(w http.ResponseWriter, r *http.Request) {
	// todo
}

func EligibilityIndex(w http.ResponseWriter, r *http.Request) {
	// todo
}

func IndustryIndex(w http.ResponseWriter, r *http.Request) {
	// todo
}

func LanguageIndex(w http.ResponseWriter, r *http.Request) {
	// todo
}

func LicenseIndex(w http.ResponseWriter, r *http.Request) {
	// todo
}

func OtherSkillIndex(w http.ResponseWriter, r *http.Request) {
	// todo
}

func PositionIndex(w http.ResponseWriter, r *http.Request) {
	// todo
}

func ProvinceIndex(w http.ResponseWriter, r *http.Request) {
	// todo
}

func RegionIndex(w http.ResponseWriter, r *http.Request) {
	// todo
}

func ReligionIndex(w http.ResponseWriter, r *http.Request) {
	// todo
}

func SchoolIndex(w http.ResponseWriter, r *http.Request) {
	// todo
}

func SkillIndex(w http.ResponseWriter, r *http.Request) {
	// todo
}
