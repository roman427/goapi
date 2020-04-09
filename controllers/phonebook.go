package controllers

import (
	"github.com/nafisfaysal/goapi/models"
	"github.com/nafisfaysal/goapi/views"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func NewPhoneBooks(ps models.PhoneBookService) *PhoneBooks {
	return &PhoneBooks{
		PhoneBooksListTempl: views.NewView("bootstrap", "phonebooks/phonebooks"),
		NewFormTempl:        views.NewView("bootstrap", "phonebooks/newphonebooks"),
		EditPhoneBooksTempl: views.NewView("bootstrap", "phonebooks/editphonebooks"),

		PhoneBookService: ps,
	}
}

type PhoneBooks struct {
	models.PhoneBookService

	PhoneBooksListTempl *views.View
	NewFormTempl        *views.View
	EditPhoneBooksTempl *views.View
}

func (p PhoneBooks) ServePhoneBookList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	userID := readUserIDFromSession(r)
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	islogged := IsLoggedIn(r)

	phones, err := p.PhoneBookService.ListByUserID(userID)
	if err != nil {
		ServeInternalServerError(w, err)
		return
	}

	err = p.PhoneBooksListTempl.Render(w, struct {
		PhoneBooks []models.PhoneBook
		IsLoggedIn bool
	}{
		PhoneBooks: phones,
		IsLoggedIn: islogged,
	})
	if err != nil {
		ServeInternalServerError(w, err)
		return
	}
}

func (p PhoneBooks) ServeNewPhoneBookForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	userID := readUserIDFromSession(r)
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	islogged := IsLoggedIn(r)

	err := p.NewFormTempl.Render(w, struct {
		IsLoggedIn bool
	}{
		IsLoggedIn: islogged,
	})
	if err != nil {
		ServeInternalServerError(w, err)
		return
	}
}

func (p PhoneBooks) CreatePhoneBook(w http.ResponseWriter, r *http.Request) {
	userID := readUserIDFromSession(r)
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	body := struct {
		Name  string `schema:"name"`
		Phone string `schema:"phone"`
	}{}
	if err := parseForm(r, &body); err != nil {
		panic(err)
	}

	pb := models.PhoneBook{
		UserID: userID,
		Name:   body.Name,
		Phone:  body.Phone,
	}

	if err := p.PhoneBookService.Create(&pb); err != nil {
		ServeInternalServerError(w, err)
		return
	}

	http.Redirect(w, r, "/phonebooks", http.StatusSeeOther)
}

func (p PhoneBooks) ServeUpdatePhoneBookForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	userID := readUserIDFromSession(r)
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	islogged := IsLoggedIn(r)

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ServeInternalServerError(w, err)
		return
	}

	ph, err := p.PhoneBookService.ByID(uint(id))
	if err != nil {
		RecordNotFound(w, err)
		return
	}
	err = p.EditPhoneBooksTempl.Render(w, struct {
		PhoneBook  *models.PhoneBook
		IsLoggedIn bool
	}{
		PhoneBook:  ph,
		IsLoggedIn: islogged,
	})
	if err != nil {
		ServeInternalServerError(w, err)
		return
	}
}

func (p PhoneBooks) UpdatePhoneBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ServeInternalServerError(w, err)
		return
	}
	ph, err := p.PhoneBookService.ByID(uint(id))
	if err != nil {
		RecordNotFound(w, err)
		return
	}

	body := struct {
		Name  string `schema:"name"`
		Phone string `schema:"phone"`
	}{}
	if err := parseForm(r, &body); err != nil {
		panic(err)
	}
	ph.Name = body.Name
	ph.Phone = body.Phone
	err = p.PhoneBookService.Update(ph)
	if err != nil {
		ServeInternalServerError(w, err)
		return
	}
	http.Redirect(w, r, "/phonebooks", http.StatusSeeOther)
}

func (c PhoneBooks) DeletePhoneBook(w http.ResponseWriter, r *http.Request) {
	userID := readUserIDFromSession(r)
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ServeInternalServerError(w, err)
		return
	}
	err = c.PhoneBookService.Delete(uint(id))
	if err != nil {
		ServeInternalServerError(w, err)
		return
	}
	http.Redirect(w, r, "/phonebooks", http.StatusSeeOther)
}
