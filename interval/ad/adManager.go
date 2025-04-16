package ad

import "github.com/chaotic98/divar/interval/models"

type AdManager interface {
	Add(username, title, tag string) string
	Remove(username, title string) string
	ListByUser(username, tag string) string
	Get(title string) (models.Ad, bool)
	ListAll() map[string]models.Ad
}
