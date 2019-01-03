package models

type HeaderList struct {
	Name string
	URL  string
	Rank int
}

type Header map[string]HeaderList
