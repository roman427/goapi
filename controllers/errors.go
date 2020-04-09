package controllers

import (
	"net/http"
)

func ServeInternalServerError(w http.ResponseWriter, err error) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func AuthenticationFailed(w http.ResponseWriter, err error) {
	http.Error(w, "Authentication failed", http.StatusUnauthorized)
}
func ServeStatusForbiddenError(w http.ResponseWriter, err error) {
	http.Error(w, "You don't have permission to access on this server", http.StatusForbidden)
}
func RecordNotFound(w http.ResponseWriter, err error) {
	http.Error(w, "Record not found", http.StatusNotFound)
}
