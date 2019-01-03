package controllers

import "net/http"

// GET -- 200 (OK), single game. 404 (Not Found), if ID not found or invalid.
func GetUserr(w http.ResponseWriter, r *http.Request) {
}

// POST -- 404 (Not Found), 409 (Conflict) if resource already exists.
func PostUser(w http.ResponseWriter, r *http.Request) {
}

// PATCH -- 200 (OK) or 204 (No Content). 404 (Not Found), if ID not found or invalid.
func PatchUser(w http.ResponseWriter, r *http.Request) {
}

// DELETE -- 200 (OK). 404 (Not Found), if ID not found or invalid.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
}
