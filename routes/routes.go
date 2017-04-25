package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zneyrl/nmsrs/controllers"
	"github.com/zneyrl/nmsrs/controllers/auth"
	"github.com/zneyrl/nmsrs/controllers/check"
	"github.com/zneyrl/nmsrs/controllers/home"
	"github.com/zneyrl/nmsrs/controllers/registrant"
	"github.com/zneyrl/nmsrs/controllers/reports"
	"github.com/zneyrl/nmsrs/controllers/search"
	"github.com/zneyrl/nmsrs/controllers/user"
	"github.com/zneyrl/nmsrs/middlewares"
)

func Register() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Auth routes
	router.Path("/logout").Methods("GET").Handler(middlewares.Auth(auth.Logout))
	router.Path("/check/file/image/{id}").Methods("POST").Handler(middlewares.Auth(check.Image))
	router.Path("/reports").Methods("GET").Handler(middlewares.Auth(reports.Index))
	router.Path("/search").Methods("GET").Handler(middlewares.Auth(search.Index))
	router.Path("/results").Methods("GET").Handler(middlewares.Auth(search.Results))

	// Admin routes
	users := router.PathPrefix("/users").Subrouter()
	users.Path("/").Methods("GET").Handler(middlewares.Admin(user.Index))
	users.Path("/").Methods("POST").Handler(middlewares.Admin(user.Store))
	users.Path("/create").Methods("GET").Handler(middlewares.Admin(user.Create))
	users.Path("/ids").Methods("POST").Handler(middlewares.Admin(user.DestroyMany))
	users.Path("/{id}").Methods("GET").Handler(middlewares.Admin(user.Show))
	users.Path("/{id}/edit").Methods("GET").Handler(middlewares.Admin(user.Edit))
	users.Path("/{id}/photo").Methods("GET").Handler(middlewares.Admin(user.Photo))
	users.Path("/{id}").Methods("PUT").Handler(middlewares.Admin(user.UpdateProfile))
	users.Path("/{id}").Methods("DELETE").Handler(middlewares.Admin(user.Destroy))
	users.Path("/{id}/reset-password").Methods("POST").Handler(middlewares.Admin(user.ResetPassword))
	registrants := router.PathPrefix("/registrants").Subrouter()
	registrants.Path("/").Methods("GET").Handler(middlewares.Admin(registrant.Index))
	registrants.Path("/").Methods("POST").Handler(middlewares.Admin(registrant.Store))
	registrants.Path("/ids").Methods("POST").Handler(middlewares.Admin(registrant.DestroyMany))
	registrants.Path("/{id}").Methods("GET").Handler(middlewares.Admin(registrant.Show))
	registrants.Path("/{id}/photo").Methods("GET").Handler(middlewares.Admin(registrant.Photo))
	registrants.Path("/{id}").Methods("DELETE").Handler(middlewares.Admin(registrant.Destroy))
	registrants.Path("/{id}/profile").Methods("GET").Handler(middlewares.Admin(registrant.Profile))
	registrants.Path("/{id}/profile").Methods("PUT").Handler(middlewares.Admin(registrant.UpdateProfile))
	registrants.Path("/{id}/formal-education").Methods("GET").Handler(middlewares.Admin(registrant.FormalEducation))
	registrants.Path("/{id}/formal-education").Methods("PUT").Handler(middlewares.Admin(registrant.UpdateFormalEducation))
	registrants.Path("/{id}/professional-license").Methods("GET").Handler(middlewares.Admin(registrant.ProfessionalLicense))
	registrants.Path("/{id}/professional-license").Methods("PUT").Handler(middlewares.Admin(registrant.UpdateProfessionalLicense))
	registrants.Path("/{id}/eligibility").Methods("GET").Handler(middlewares.Admin(registrant.Eligibility))
	registrants.Path("/{id}/eligibility").Methods("PUT").Handler(middlewares.Admin(registrant.UpdateEligibility))
	registrants.Path("/{id}/technical-training-and-relevant-experience").Methods("GET").Handler(middlewares.Admin(registrant.TechnicalTrainingAndRelevantExperience))
	registrants.Path("/{id}/technical-training-and-relevant-experience").Methods("PUT").Handler(middlewares.Admin(registrant.UpdateTechnicalTrainingAndRelevantExperience))
	registrants.Path("/{id}/certificate-of-competence").Methods("GET").Handler(middlewares.Admin(registrant.CertificateOfCompetence))
	registrants.Path("/{id}/certificate-of-competence").Methods("PUT").Handler(middlewares.Admin(registrant.UpdateCertificateOfCompetence))
	registrants.Path("/{id}/work-experience").Methods("GET").Handler(middlewares.Admin(registrant.WorkExperience))
	registrants.Path("/{id}/work-experience").Methods("PUT").Handler(middlewares.Admin(registrant.UpdateWorkExperience))
	registrants.Path("/{id}/other-skills-aquired-without-formal-training").Methods("GET").Handler(middlewares.Admin(registrant.OtherSkillsAquiredWithoutFormalTraining))
	registrants.Path("/{id}/other-skills-aquired-without-formal-training").Methods("PUT").Handler(middlewares.Admin(registrant.UpdateOtherSkillsAquiredWithoutFormalTraining))
	registrants.Path("/{id}/certification-authorization").Methods("GET").Handler(middlewares.Admin(registrant.CertificationAuthorization))
	registrants.Path("/{id}/certification-authorization").Methods("PUT").Handler(middlewares.Admin(registrant.UpdateCertificationAuthorization))

	// Web routes
	router.Path("/").Methods("GET").Handler(middlewares.Web(home.Index))
	router.Path("/welcome").Methods("GET").Handler(middlewares.Web(home.Welcome))
	login := router.PathPrefix("/login").Subrouter()
	login.Methods("GET").Handler(middlewares.Web(auth.ShowLoginForm))
	login.Methods("POST").Handler(middlewares.Web(auth.Login))

	router.NotFoundHandler = http.HandlerFunc(controllers.PageNotFound)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	return router
}