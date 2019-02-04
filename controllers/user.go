package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/canerakdas/gaben/models"
	"github.com/canerakdas/gaben/utils"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
)
/*
var url = 'http://localhost:8080/user';

fetch(url, {
  method: 'GET', // or 'PUT'
  headers:{
    'Content-Type': 'application/json',
	'Email': 'canerakdas@gmail.com',
	'Password': '1234'
  }
}).then(res => res.json())
.then(response => localStorage.setItem("user",JSON.stringify(response)))
.catch(error => console.error('Error:', error));
*/
var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)
var Status models.Status

func CheckUserStatus(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "cookie-name")

	if session.Values["authenticated"] == true {
		Status.Online = true
	}else{
		Status.Online = false
	}
}

func Secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	// Check if user is authenticated
	if session.Values["authenticated"] == false{
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	//if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
	//}

	// Print secret message
	fmt.Fprintln(w, session.Values["user"])
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}
// GET -- 200 (OK), single user. 404 (Not Found), if ID not found or invalid.
func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	email := r.Header.Get("Email")
	password := r.Header.Get("Password")

	err := utils.UserDb.Find(bson.M{"email":email}).One(&user)

	if err != nil {
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	user.Password = ""

	session, _ := store.Get(r, "cookie-name")

	if err == nil{
		// Set user as authenticated
		var userId = user.Id.Hex()
		session.Values["authenticated"] = true
		session.Values["user"] = userId
		session.Save(r, w)

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			panic(err)
		}
	}
}

// POST -- 404 (Not Found), 409 (Conflict) if resource already exists.
func PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)
	user.Id = bson.NewObjectId()
	err = utils.UserDb.Insert(user)

	if err != nil {
		panic(err)
	}
}

// PATCH -- 200 (OK) or 204 (No Content). 404 (Not Found), if ID not found or invalid.
func PatchUser(w http.ResponseWriter, r *http.Request) {
}

// DELETE -- 200 (OK). 404 (Not Found), if ID not found or invalid.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
}
