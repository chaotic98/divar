package models

type User struct {
	Username  string
	PostedAds []string
	Favorites []string
	FavSet    map[string]bool
}
