package controllers

import (
	"github.com/canerakdas/game_crawl/models"
	"github.com/canerakdas/game_crawl/utils"
	"strings"
)

func GetHeader() models.Header {
	var header models.Header

	utils.CatalogDb.Find(nil).One(&header)
	delete(header, "_id")
	return header
}

func CollectHeader() {
	var games []models.Game
	utils.GameDb.Find(nil).All(&games)

	header := make(models.Header, 10)
	for i := 0; i < len(games); i++ {
		genres := games[i].Genres
		for _, v := range genres {
			header[v.ID] = models.HeaderList{
				Name: v.Description,
				URL:  strings.Replace(v.Description, " ", "-", -1),
				Rank: 0,
			}
		}
	}

	utils.CatalogDb.DropCollection()
	utils.CatalogDb.Insert(header)
}
