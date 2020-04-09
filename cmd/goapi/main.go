package main

import (
	"github.com/nafisfaysal/goapi/controllers"
	"github.com/nafisfaysal/goapi/models"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func main() {
	errCh := make(chan error)

	db, err := gorm.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}

	userGORM := models.NewUserGORM(db)
	if err != nil {
		log.Fatal(err)
	}

	phoneBookGORM := models.NewPhoneBookGORM(db)
	if err != nil {
		log.Fatal(err)
	}

	userGORM.AutoMigrate()
	phoneBookGORM.AutoMigrate()

	indexC := controllers.NewIndex()
	usersC := controllers.NewUsers(userGORM)
	phoneBooksC := controllers.NewPhoneBooks(phoneBookGORM)

	r := mux.NewRouter()

	r.NewRoute().
		Name("HomePage").
		Methods("GET").
		Path("/").
		Handler(indexC.Homepage)

	//  User related routes

	r.NewRoute().
		Name("ServeSignupForm").
		Methods("GET").
		Path("/register").
		HandlerFunc(usersC.ServeSignupForm)

	r.NewRoute().
		Name("RegisterAccount").
		Methods("POST").
		Path("/register").
		HandlerFunc(usersC.RegisterAccount)

	r.NewRoute().
		Name("ServeLoginForm").
		Methods("GET").
		Path("/login").
		HandlerFunc(usersC.ServeLoginForm)

	r.NewRoute().
		Name("Login").
		Methods("POST").
		Path("/login").
		HandlerFunc(usersC.HandleLogin)

	r.NewRoute().
		Name("Logout").
		Methods("GET").
		Path("/logout").
		HandlerFunc(usersC.HandleLogout)

	r.NewRoute().
		Name("ServePhoneBookList").
		Methods("GET").
		Path("/phonebooks").
		HandlerFunc(phoneBooksC.ServePhoneBookList)
	r.NewRoute().
		Name("ServeNewPhoneBookForm").
		Methods("GET").
		Path("/phonebooks/new").
		HandlerFunc(phoneBooksC.ServeNewPhoneBookForm)

	r.NewRoute().
		Name("CreatePhoneBook").
		Methods("POST").
		Path("/phonebooks/new").
		HandlerFunc(phoneBooksC.CreatePhoneBook)

	r.NewRoute().
		Name("ServeUpdatePhoneBookForm").
		Methods("GET").
		Path("/phonebooks/{id:[0-9]+}/edit").
		HandlerFunc(phoneBooksC.ServeUpdatePhoneBookForm)

	r.NewRoute().
		Name("UpdatePhoneBook").
		Methods("POST").
		Path("/phonebooks/{id:[0-9]+}/edit").
		HandlerFunc(phoneBooksC.UpdatePhoneBook)

	r.NewRoute().
		Name("DeletePhoneBook").
		Methods("POST").
		Path("/phonebooks/{id:[0-9]+}/delete").
		HandlerFunc(phoneBooksC.DeletePhoneBook)

	go func() {
		err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
		if err != nil {
			errCh <- err
		}
	}()

	signCh := make(chan os.Signal)
	signal.Notify(signCh, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errCh:
		log.Fatal(err)

	case sign := <-signCh:
		log.Printf("Server gracefully %s", sign)
	}
}
