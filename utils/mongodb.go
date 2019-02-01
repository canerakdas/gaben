package utils

import (
	"gopkg.in/mgo.v2"
)

var db *mgo.Database

var session, err = mgo.Dial("localhost")

var GameDb = session.DB("game-db").C("steam-game-list-new")
var GameListDb = session.DB("game-db").C("steam-game-db")
var CatalogDb = session.DB("game-db").C("catalog")
var UserDb = session.DB("game-db").C("user-collection")