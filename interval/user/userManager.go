package user

type UserManager interface {
	Register(username string) string
	AddFavorite(username, title string) string
	RemoveFavorite(username, title string) string
	ListFavorites(username, tag string) string
	Exists(username string) bool
	GetAllUsernames() []string
}
