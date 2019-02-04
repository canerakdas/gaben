package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id						 bson.ObjectId `json:"id" bson:"_id"`
	Steamid                  string `json:"steamid"`
	Communityvisibilitystate int    `json:"communityvisibilitystate"`
	Profilestate             int    `json:"profilestate"`
	Personaname              string `json:"personaname"`
	Lastlogoff               int    `json:"lastlogoff"`
	Profileurl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	Avatarmedium             string `json:"avatarmedium"`
	Avatarfull               string `json:"avatarfull"`
	Personastate             int    `json:"personastate"`
	Realname                 string `json:"realname"`
	Primaryclanid            string `json:"primaryclanid"`
	Timecreated              int    `json:"timecreated"`
	Personastateflags        int    `json:"personastateflags"`
	Loccountrycode           string `json:"loccountrycode"`
	Locstatecode             string `json:"locstatecode"`
	Loccityid                int    `json:"loccityid"`
	Email                    string `json:"email" bson:"email,omitempty"`
	Password                 string `json:"password" bson:"password,omitempty"`
	LastVisitedGames		 []int `json:"lastvisitedgames"`
}

type Status struct {
	Online bool
}