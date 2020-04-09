package controllers

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("very-secret-much-wow"))

func saveToSession(r *http.Request, w http.ResponseWriter, k string, v interface{}) error {
	session, err := store.Get(r, "goapi")
	if err != nil {
		return err
	}

	session.Values[k] = v
	return session.Save(r, w)
}

func readFromSession(r *http.Request, k string) (interface{}, error) {
	session, err := store.Get(r, "goapi")
	if err != nil {
		return nil, err
	}

	return session.Values[k], nil
}

func readUserIDFromSession(r *http.Request) uint {
	v, _ := readFromSession(r, "userID")
	userID, ok := v.(uint)
	if !ok {
		return 0
	}
	return userID
}

func IsLoggedIn(r *http.Request) bool {
	v, _ := readFromSession(r, "userID")
	userID := readUserIDFromSession(r)
	if userID == v {
		return true
	}
	return false
}
