package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var db *mgo.Database

var session, err = mgo.Dial("localhost")
var Conn = session.DB("game-db").C("steam-game-db")
var ConnGameList = session.DB("game-db").C("steam-game-list-new")

type App struct {
	Applist struct {
		string
		Apps []struct {
			Appid int    `json:"appid"`
			Name  string `json:"name"`
		} `json:"apps"`
	} `json:"applist"`
}

type Game struct {
	Success bool `json:"success"`
	Data    struct {
		Type                string `json:"type"`
		Name                string `json:"name"`
		SteamAppid          int    `json:"steam_appid"`
		RequiredAge         int    `json:"required_age"`
		IsFree              bool   `json:"is_free"`
		Dlc                 []int  `json:"dlc"`
		DetailedDescription string `json:"detailed_description"`
		AboutTheGame        string `json:"about_the_game"`
		ShortDescription    string `json:"short_description"`
		SupportedLanguages  string `json:"supported_languages"`
		Reviews             string `json:"reviews"`
		HeaderImage         string `json:"header_image"`
		Website             string `json:"website"`
		PcRequirements      struct {
			Minimum     string `json:"minimum"`
			Recommended string `json:"recommended"`
		} `json:"pc_requirements"`
		MacRequirements struct {
			Minimum     string `json:"minimum"`
			Recommended string `json:"recommended"`
		} `json:"mac_requirements"`
		LinuxRequirements []interface{} `json:"linux_requirements"`
		Developers        []string      `json:"developers"`
		Publishers        []string      `json:"publishers"`
		Demos             []struct {
			Appid       int    `json:"appid"`
			Description string `json:"description"`
		} `json:"demos"`
		PriceOverview struct {
			Currency         string `json:"currency"`
			Initial          int    `json:"initial"`
			Final            int    `json:"final"`
			DiscountPercent  int    `json:"discount_percent"`
			InitialFormatted string `json:"initial_formatted"`
			FinalFormatted   string `json:"final_formatted"`
		} `json:"price_overview"`
		Packages      []int `json:"packages"`
		PackageGroups []struct {
			Name                    string `json:"name"`
			Title                   string `json:"title"`
			Description             string `json:"description"`
			SelectionText           string `json:"selection_text"`
			SaveText                string `json:"save_text"`
			DisplayType             int    `json:"display_type"`
			IsRecurringSubscription string `json:"is_recurring_subscription"`
			Subs                    []struct {
				Packageid                int    `json:"packageid"`
				PercentSavingsText       string `json:"percent_savings_text"`
				PercentSavings           int    `json:"percent_savings"`
				OptionText               string `json:"option_text"`
				OptionDescription        string `json:"option_description"`
				CanGetFreeLicense        string `json:"can_get_free_license"`
				IsFreeLicense            bool   `json:"is_free_license"`
				PriceInCentsWithDiscount int    `json:"price_in_cents_with_discount"`
			} `json:"subs"`
		} `json:"package_groups"`
		Platforms struct {
			Windows bool `json:"windows"`
			Mac     bool `json:"mac"`
			Linux   bool `json:"linux"`
		} `json:"platforms"`
		Metacritic struct {
			Score int    `json:"score"`
			URL   string `json:"url"`
		} `json:"metacritic"`
		Categories []struct {
			ID          int    `json:"id"`
			Description string `json:"description"`
		} `json:"categories"`
		Genres []struct {
			ID          string `json:"id"`
			Description string `json:"description"`
		} `json:"genres"`
		Screenshots []struct {
			ID            int    `json:"id"`
			PathThumbnail string `json:"path_thumbnail"`
			PathFull      string `json:"path_full"`
		} `json:"screenshots"`
		Movies []struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Thumbnail string `json:"thumbnail"`
			Webm      struct {
				Num480 string `json:"480"`
				Max    string `json:"max"`
			} `json:"webm"`
			Highlight bool `json:"highlight"`
		} `json:"movies"`
		Recommendations struct {
			Total int `json:"total"`
		} `json:"recommendations"`
		Achievements struct {
			Total       int `json:"total"`
			Highlighted []struct {
				Name string `json:"name"`
				Path string `json:"path"`
			} `json:"highlighted"`
		} `json:"achievements"`
		ReleaseDate struct {
			ComingSoon bool   `json:"coming_soon"`
			Date       string `json:"date"`
		} `json:"release_date"`
		SupportInfo struct {
			URL   string `json:"url"`
			Email string `json:"email"`
		} `json:"support_info"`
		Background         string `json:"background"`
		ContentDescriptors struct {
			Ids   []interface{} `json:"ids"`
			Notes interface{}   `json:"notes"`
		} `json:"content_descriptors"`
	} `json:"data"`
}

func getGameInfo(url string, id string) {
	var target map[string]Game
	r, err := http.Get(strings.Replace(url, "gameid", id, 1))
	if err != nil {
		fmt.Printf("server not responding %s", err.Error())
		os.Exit(1)
	}

	t, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(t, &target)

	for _, v := range target {
		if v.Data.Type == "game" {
			err = ConnGameList.Insert(v.Data)
		}
	}

	if err == nil {
	} else {
		fmt.Printf("server not responding %s", err.Error())
		os.Exit(1)
	}

	defer r.Body.Close()

}

func main() {
	result := make([]App, 100, 1000) // Here you can specify a len and a cap.

	err := Conn.Find(nil).All(&result)

	if err != nil {
		fmt.Println("whooops")
		os.Exit(1)
	}
	for i := 0; i < len(result[0].Applist.Apps); i++ {
		var appid = result[0].Applist.Apps[i].Appid
		fmt.Println("Index:", i, "App id:", appid)
		getGameInfo("https://store.steampowered.com/api/appdetails?appids=gameid&cc=us&l=en", strconv.Itoa(appid))
	}
}
