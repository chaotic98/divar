package ad

import (
	"strings"

	"github.com/chaotic98/divar/interval/models"
	"github.com/chaotic98/divar/pkg/utils"
)

type InMemoryAdManager struct {
	ads   map[string]models.Ad
	users map[string]*models.User // shared with userManager
}

func (m *InMemoryAdManager) Add(username, title, tag string) string {
	user, ok := m.users[username]
	if !ok {
		return "invalid username"
	}
	if _, exists := m.ads[title]; exists {
		return "invalid title"
	}
	m.ads[title] = models.Ad{Owner: username, Title: title, Tag: tag}
	user.PostedAds = append(user.PostedAds, title)
	return "posted successfully"
}

func (m *InMemoryAdManager) Remove(username, title string) string {
	user, ok := m.users[username]
	if !ok {
		return "invalid username"
	}
	ad, ok := m.ads[title]
	if !ok {
		return "invalid title"
	}
	if ad.Owner != username {
		return "access denied"
	}
	delete(m.ads, title)
	user.PostedAds = utils.RemoveStringFromSlice(user.PostedAds, title)
	for _, u := range m.users {
		if u.FavSet[title] {
			u.Favorites = utils.RemoveStringFromSlice(u.Favorites, title)
			delete(u.FavSet, title)
		}
	}
	return "removed successfully"
}

func (m *InMemoryAdManager) ListByUser(username, tag string) string {
	user, ok := m.users[username]
	if !ok {
		return "invalid username"
	}
	var result []string
	for _, t := range user.PostedAds {
		ad, ok := m.ads[t]
		if ok && (tag == "" || ad.Tag == tag) {
			result = append(result, t)
		}
	}
	return strings.Join(result, " ")
}

func (m *InMemoryAdManager) Get(title string) (models.Ad, bool) {
	ad, ok := m.ads[title]
	return ad, ok
}

func (m *InMemoryAdManager) ListAll() map[string]models.Ad {
	return m.ads
}

func NewAdManagerWithStore(ads map[string]models.Ad, users map[string]*models.User) *InMemoryAdManager {
	return &InMemoryAdManager{
		ads:   ads,
		users: users,
	}
}
