package models

type CatalogTemplate struct {
	Header     Header
	Title      string
	Content    []Game
	Pagination []string
	Footer     []string
}
