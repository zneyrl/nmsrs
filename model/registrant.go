package model

import (
	"mime/multipart"
	"time"

	"github.com/emurmotol/nmsrs/database"
	"github.com/emurmotol/nmsrs/helper"
	"github.com/emurmotol/nmsrs/lang"
)

type Registrant struct {
	ID         uint64     `json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	RegistInfo RegistInfo
	RegistEmp  RegistEmp
}

type CreateRegistrantForm struct {
	FamilyName        string                `schema:"family_name" validate:"required"`
	GivenName         string                `schema:"given_name" validate:"required"`
	MiddleName        string                `schema:"middle_name" validate:"required"`
	Birthdate         time.Time             `schema:"birthdate" validate:"required"`
	Password          string                `schema:"password"`
	PhotoFile         multipart.File        `schema:"-"`
	PhotoHeader       *multipart.FileHeader `schema:"-"`
	StSub             string                `schema:"st_sub" validate:"required"`
	CityMunID         uint                  `schema:"city_mun_id" validate:"required"`
	ProvID            uint                  `schema:"prov_id"`
	BrgyID            uint                  `schema:"brgy_id" validate:"required"`
	PlaceOfBirth      string                `schema:"place_of_birth" validate:"required"`
	ReligionID        uint                  `schema:"religion_id" validate:"required"`
	CivilStatID       uint                  `schema:"civil_stat_id" validate:"required"`
	CivilStatOther    string                `schema:"civil_stat_other" validate:"required"`
	SexID             uint                  `schema:"sex_id" validate:"required"`
	Age               int                   `schema:"age"`
	Height            float32               `schema:"height" validate:"required"`
	Weight            float32               `schema:"weight"`
	LandlineNo        string                `schema:"landline_no"`
	MobileNo          string                `schema:"mobile_no"`
	Email             string                `schema:"email" validate:"email"`
	EmpStatID         uint                  `schema:"emp_stat_id" validate:"required"`
	UnEmpStatID       uint                  `schema:"un_emp_stat_id"`
	TocID             uint                  `schema:"toc_id"`
	Alfw              bool                  `schema:"alfw"`
	PrefOccIDs        []int                 `schema:"pref_occ_ids" validate:"required"`
	PrefLocalLocID    uint                  `schema:"pref_local_loc_id" validate:"required"`
	PrefOverseasLocID uint                  `schema:"pref_overseas_loc_id" validate:"required"`
	PassportNo        string                `schema:"passport_no"`
	Pned              time.Time             `schema:"pned"`
	Disabled          bool                  `schema:"disabled"`
	DisabilityID      uint                  `schema:"disability_id"`
	DisabilityOther   uint                  `schema:"disability_other"`
	LanguageIDs       []int                 `schema:"language_ids"`
	RegDate           time.Time             `schema:"reg_date"`
	Errors            map[string]string     `schema:"-"`
}

func (form *CreateRegistrantForm) IsValid() bool {
	form.Errors = make(map[string]string)

	if errs := helper.ValidateForm(form); len(errs) != 0 {
		form.Errors = errs
	}

	if taken := RegistrantEmailTaken(form.Email); taken {
		form.Errors["Email"] = lang.Get("email_taken")
	}

	if form.PhotoFile != nil {
		if err := helper.ValidateImage(form.PhotoHeader); err != nil {
			form.Errors["Photo"] = err.Error()
		}
	}
	return len(form.Errors) == 0
}

func RegistrantByID(id uint64) *Registrant {
	db := database.Con()
	defer db.Close()
	registrant := new(Registrant)

	if notFound := db.First(registrant, id).RecordNotFound(); notFound {
		return nil
	}
	return registrant
}

func RegistrantByEmail(email string) *Registrant {
	db := database.Con()
	defer db.Close()
	registInfo := RegistInfo{}

	if notFound := db.Where("email = ?", email).First(&registInfo).RecordNotFound(); notFound {
		return nil
	}
	return RegistrantByID(registInfo.RegistrantID)
}

func RegistrantEmailTaken(email string) bool {
	registrant := RegistrantByEmail(email)

	if registrant != nil {
		return true
	}
	return false
}
