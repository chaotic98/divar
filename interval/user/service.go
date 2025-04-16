package user

import (
	"strings"

	"github.com/chaotic98/divar/interval/models"
	"github.com/chaotic98/divar/pkg/utils"
)

type InMemoryUserManager struct {
	users map[string]*models.User
	ads   map[string]models.Ad // shared with adManager
}

func NewUserManager(ads map[string]models.Ad) *InMemoryUserManager {
	return &InMemoryUserManager{
		users: make(map[string]*models.User),
		ads:   ads,
	}
}

func (m *InMemoryUserManager) Register(username string) string {
	if _, exists := m.users[username]; exists {
		return "invalid username"
	}
	m.users[username] = &models.User{
		Username:  username,
		PostedAds: []string{},
		Favorites: []string{},
		FavSet:    make(map[string]bool),
	}
	return "registered successfully"
}

func (m *InMemoryUserManager) AddFavorite(username, title string) string {
	user, ok := m.users[username]
	if !ok {
		return "invalid username"
	}
	if _, ok := m.ads[title]; !ok {
		return "invalid title"
	}
	if user.FavSet[title] {
		return "already favorite"
	}
	user.Favorites = append(user.Favorites, title)
	user.FavSet[title] = true
	return "added successfully"
}

func (m *InMemoryUserManager) RemoveFavorite(username, title string) string {
	user, ok := m.users[username]
	if !ok {
		return "invalid username"
	}
	if _, ok := m.ads[title]; !ok {
		return "invalid title"
	}
	if !user.FavSet[title] {
		return "already not favorite"
	}
	user.Favorites = utils.RemoveStringFromSlice(user.Favorites, title)
	delete(user.FavSet, title)
	return "removed successfully"
}

func (m *InMemoryUserManager) ListFavorites(username, tag string) string {
	user, ok := m.users[username]
	if !ok {
		return "invalid username"
	}
	var result []string
	for _, t := range user.Favorites {
		ad, ok := m.ads[t]
		if ok && (tag == "" || ad.Tag == tag) {
			result = append(result, t)
		}
	}
	return strings.Join(result, " ")
}

func (m *InMemoryUserManager) Exists(username string) bool {
	_, ok := m.users[username]
	return ok
}

func (m *InMemoryUserManager) GetAllUsernames() []string {
	var names []string
	for k := range m.users {
		names = append(names, k)
	}
	return names
}
func NewUserManagerWithStore(users map[string]*models.User, ads map[string]models.Ad) *InMemoryUserManager {
	return &InMemoryUserManager{
		users: users,
		ads:   ads,
	}
}
