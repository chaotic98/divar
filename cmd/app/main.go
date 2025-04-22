package main

import (
	"bufio"
	"fmt"
	"github.com/chaotic98/divar/interval/models"
	"os"
	"strings"

	adpkg "github.com/chaotic98/divar/interval/ad"
	userpkg "github.com/chaotic98/divar/interval/user"
)

func main() {
	users := make(map[string]*models.User)
	ads := make(map[string]models.Ad)

	userManager := userpkg.NewUserManagerWithStore(users, ads)
	adManager := adpkg.NewAdManagerWithStore(ads, users)

	var userIntf userpkg.UserManager = userManager
	var adIntf adpkg.AdManager = adManager

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		tokens := strings.Fields(line)
		cmd := tokens[0]
		var result string

		switch cmd {
		case "register":
			if len(tokens) == 2 {
				result = userIntf.Register(tokens[1])
			} else {
				result = "invalid input"
			}

		case "add_advertise":
			if len(tokens) >= 3 {
				username := tokens[1]
				title := tokens[2]
				tag := ""
				if len(tokens) >= 4 {
					tag = tokens[3]
				}
				result = adIntf.Add(username, title, tag)
			} else {
				result = "invalid input"
			}

		case "rem_advertise":
			if len(tokens) == 3 {
				result = adIntf.Remove(tokens[1], tokens[2])
			} else {
				result = "invalid input"
			}

		case "list_my_advertises":
			if len(tokens) >= 2 {
				username := tokens[1]
				tag := ""
				if len(tokens) >= 3 {
					tag = tokens[2]
				}
				result = adIntf.ListByUser(username, tag)
			} else {
				result = "invalid input"
			}

		case "add_favorite":
			if len(tokens) == 3 {
				result = userIntf.AddFavorite(tokens[1], tokens[2])
			} else {
				result = "invalid input"
			}

		case "rem_favorite":
			if len(tokens) == 3 {
				result = userIntf.RemoveFavorite(tokens[1], tokens[2])
			} else {
				result = "invalid input"
			}

		case "list_favorite_advertises":
			if len(tokens) >= 2 {
				username := tokens[1]
				tag := ""
				if len(tokens) >= 3 {
					tag = tokens[2]
				}
				result = userIntf.ListFavorites(username, tag)
			} else {
				result = "invalid input"
			}

		default:
			result = "unknown command"
		}

		fmt.Println(result)
	}
}
