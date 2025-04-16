package ad

import (
	"github.com/chaotic98/divar/interval/models"
	"github.com/chaotic98/divar/pkg/utils"
	"strings"
)

var ads = make(map[string]models.Ad)
var users = make(map[string]*models.User)

func AddAdvertise(username, title, tag string) string {
	user, exists := users[username]
	if !exists {
		return "invalid username"
	}

	if _, exists := ads[title]; exists {
		return "invalid title"
	}

	ads[title] = models.Ad{
		Owner: username,
		Title: title,
		Tag:   tag,
	}
	user.PostedAds = append(user.PostedAds, title)
	return "posted successfully"
}

func RemAdvertise(username, title string) string {
	user, exists := users[username]
	if !exists {
		return "invalid username"
	}

	ad, exists := ads[title]
	if !exists {
		return "invalid title"
	}

	if ad.Owner != username {
		return "access denied"
	}

	delete(ads, title)

	user.PostedAds = utils.RemoveStringFromSlice(user.PostedAds, title)

	for _, u := range users {
		if u.FavSet[title] {
			u.Favorites = utils.RemoveStringFromSlice(u.Favorites, title)
			delete(u.FavSet, title)
		}
	}
	return "removed successfully"
}

func ListMyAdvertises(username string, tagFilter string) string {
	user, exists := users[username]
	if !exists {
		return "invalid username"
	}
	var result []string
	for _, t := range user.PostedAds {
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
