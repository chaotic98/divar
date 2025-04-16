package user

import (
	"github.com/chaotic98/divar/interval/models"
	"github.com/chaotic98/divar/pkg/utils"
	"strings"
)

var users = make(map[string]*models.User)
var ads = make(map[string]models.Ad)

func Register(username string) string {
	if _, exists := users[username]; exists {
		return "invalid username"
	}
	users[username] = &models.User{
		Username:  username,
		PostedAds: []string{},
		Favorites: []string{},
		FavSet:    make(map[string]bool),
	}
	return "registered successfully"
}

func AddFavorite(username, title string) string {
	user, exists := users[username]
	if !exists {
		return "invalid username"
	}

	if _, exists := ads[title]; !exists {
		return "invalid title"
	}

	if user.FavSet[title] {
		return "already favorite"
	}
	user.Favorites = append(user.Favorites, title)
	user.FavSet[title] = true
	return "added successfully"
}

func RemFavorite(username, title string) string {
	user, exists := users[username]
	if !exists {
		return "invalid username"
	}

	if _, exists := ads[title]; !exists {
		return "invalid title"
	}

	if !user.FavSet[title] {
		return "already not favorite"
	}
	user.Favorites = utils.RemoveStringFromSlice(user.Favorites, title)
	delete(user.FavSet, title)
	return "removed successfully"
}

func ListFavoriteAdvertises(username string, tagFilter string) string {
	user, exists := users[username]
	if !exists {
		return "invalid username"
	}
	var result []string
	for _, t := range user.Favorites {
		if tagFilter != "" {
			ad, ok := ads[t]
			if ok && ad.Tag == tagFilter {
				result = append(result, t)
			}
		} else {
			result = append(result, t)
		}
	}
	return strings.Join(result, " ")
}
