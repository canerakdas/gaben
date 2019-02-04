package controllers

import (
	"fmt"
	"github.com/canerakdas/gaben/models"
	"github.com/canerakdas/gaben/utils"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// GET -- 200 (OK), single game. 404 (Not Found), if ID not found or invalid.
func GetCatalog(w http.ResponseWriter, r *http.Request) {
	CheckUserStatus(w,r)

	vars := mux.Vars(r)
	id := vars["id"]
	name := vars["name"]
	page := r.URL.Query()["p"]

	var skip int

	// query
	multipleId := strings.Split(id, "-")

	andQuery := []bson.M{}
	for _, v := range multipleId {
		andQuery = append(andQuery, bson.M{"genres.id": v})
	}

	// pagination
	if page != nil {
		skip, _ = strconv.Atoi(page[0])
	}

	limit := 10
	skip = limit * skip
	count, _ := utils.GameDb.Find(bson.M{"$and": andQuery}).Count()
	pageCount := ((count - (count % limit)) / limit) + 1

	pagination := make([]string, pageCount, pageCount)
	for i, _ := range pagination {
		queryTemp := []string{name, "?p=", strconv.Itoa(i)}
		paginationTemp := []string{"", "catalog", id, strings.Join(queryTemp, "")}
		pagination[i] = strings.Join(paginationTemp, "/")
	}
	var catalog = make([]models.Game, 10, 100)

	utils.GameDb.Find(bson.M{"$and": andQuery}).Skip(skip).Limit(limit).All(&catalog)

	if len(catalog) == 0 {
		fmt.Println("TODO: EMPTY CATALOG PAGE")
	}
	//json.NewEncoder(w).Encode(game)

	tpl = template.Must(template.ParseFiles("./template/catalog.gohtml", "./template/header.gohtml", "./template/footer.gohtml"))

	template := models.CatalogTemplate{
		Header:     GetHeader(),
		Title:      name,
		Content:    catalog,
		Pagination: pagination,
		Footer:     []string{"haha"},
		Status:		Status,
	}

	tpl.Execute(w, template)
}

func PostCatalog() {
}

func PatchCatalog() {
}

func DeleteCatalog() {
}
