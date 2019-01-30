package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/canerakdas/gaben/models"
	"github.com/canerakdas/gaben/utils"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"net/http"
	"strconv"
)

type Game models.Game
type GameDetailTemplate struct {
	Header  models.Header
	Title   string
	Content Game
	Footer  []string
}

var tpl *template.Template

// GET -- 200 (OK), single game. 404 (Not Found), if ID not found or invalid.
func IGetGame(w http.ResponseWriter, r *http.Request) {
	var game Game
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := utils.GameDb.Find(bson.M{"steamappid": id}).One(&game)

	var templates = []string{"./template/game.gohtml", "./template/header.gohtml", "./template/footer.gohtml"}
	var page = GameDetailTemplate{GetHeader(), game.Name, game, []string{}}

	ExecuteGameTemplate(w, templates, page)
	game.view("myset")

	if err != nil {
		panic(err)
	}
}

func GetGame(w http.ResponseWriter, r *http.Request) {
	var gameCollection = make([]models.Game, 1, 1000)
	err := utils.GameDb.Find(nil).All(&gameCollection)

	err = json.NewEncoder(w).Encode(gameCollection)

	if err != nil {
		panic(err)
	}
}

func PostGame() {
}

func PatchGame() {
}

func DeleteGame() {
}

// Execute template
func ExecuteGameTemplate(w http.ResponseWriter, templates []string, page GameDetailTemplate) {
	tpl = template.Must(template.ParseFiles(templates...))

	tpl.Execute(w, page)
}

// Redis ordered list example
func (g Game) view(key string) {
	id := strconv.Itoa(g.SteamAppid)
	keys, cursor, err := utils.RedisClient.ZScan(key, 0, id, 1).Result()
	if len(keys) != 0 {
		utils.RedisClient.ZIncrBy(key, float64(1), id)
	} else {
		utils.RedisClient.ZAdd(key, redis.Z{
			Score:  float64(1),
			Member: id,
		})
	}

	if err != nil {
		panic(err)
	}
	fmt.Println(keys, cursor, err)
}
