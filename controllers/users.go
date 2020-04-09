package controllers

import (
	"github.com/nafisfaysal/goapi/models"
	"github.com/nafisfaysal/goapi/views"
	"net/http"

)

func NewUsers(us models.UserService) *Users {
	return &Users{
		SignupTempl: views.NewView("bootstrap", "users/signup"),
		LoginTempl:  views.NewView("bootstrap", "users/login"),
		UserService: us,
	}
}

type Users struct {
	models.UserService

	SignupTempl *views.View
	LoginTempl  *views.View
}

func (u *Users) ServeSignupForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := u.SignupTempl.Render(w, nil)
	if err != nil {
		ServeInternalServerError(w, err)
		return
	}
}

func (u *Users) RegisterAccount(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Email    string `schema:"email"`
		Password string `schema:"password"`
	}{}
	if err := parseForm(r, &body); err != nil {
		panic(err)
	}
	user := models.User{
		Email:    body.Email,
		Password: body.Password,
	}
	if err := u.UserService.Create(&user); err != nil {
		ServeInternalServerError(w, err)
		return
	}
	err := saveToSession(r, w, "userID", user.ID)
	if err != nil {
		ServeInternalServerError(w, err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (u *Users) ServeLoginForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := u.LoginTempl.Render(w, nil)
	if err != nil {
		ServeInternalServerError(w, err)
		return
	}
}

func (u *Users) HandleLogin(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Email    string `schema:"email"`
		Password string `schema:"password"`
	}{}
	if err := parseForm(r, &body); err != nil {
		panic(err)
	}
	user, err := u.UserService.Authenticate(body.Email, body.Password)
	if err != nil {
		AuthenticationFailed(w, err)
		return
	}

	if user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = saveToSession(r, w, "userID", user.ID)
	if err != nil {
		ServeInternalServerError(w, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)

}

func (u *Users) HandleLogout(w http.ResponseWriter, r *http.Request) {
	seesion, _ := store.Get(r, "goapi")
	delete(seesion.Values, "userID")
	seesion.Options.MaxAge = -1
	_ = seesion.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
