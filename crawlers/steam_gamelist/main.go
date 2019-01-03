package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"net/http"
	"os"
)

var db *mgo.Database

var session, _ = mgo.Dial("localhost")
var Conn = session.DB("game-db").C("steam-game-db")

type App struct {
	Applist struct {
		Apps []struct {
			Appid int    `json:"appid"`
			Name  string `json:"name"`
		} `json:"apps"`
	} `json:"applist"`
}

func getJson(url string) {
	var target App
	r, err := http.Get(url)
	if err != nil {
		fmt.Printf("server not responding %s", err.Error())
		os.Exit(1)
	}

	err = Conn.DropCollection()
	if err != nil{
		fmt.Println(err)
	}

	t, _ := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(t, &target)

	if err != nil{
		fmt.Println(err)
	}

	err = Conn.Insert(&target)
	if err != nil{
		fmt.Println(err)
	}

	defer r.Body.Close()
}

func main() {
	getJson("http://api.steampowered.com/ISteamApps/GetAppList/v0002/?key=STEAMKEY&format=json")
}
