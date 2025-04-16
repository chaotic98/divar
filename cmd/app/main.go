package main

import (
	"bufio"
	"fmt"
	"github.com/chaotic98/divar/interval/ad"
	"github.com/chaotic98/divar/interval/user"
	"os"
	"strings"
)

func main() {
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
			if len(tokens) != 2 {
				result = ""
			} else {
				result = user.Register(tokens[1])
			}

		case "add_advertise":
			if len(tokens) < 3 {
				result = ""
			} else {
				username := tokens[1]
				title := tokens[2]
				tag := ""
				if len(tokens) >= 4 {
					tag = tokens[3]
				}
				result = ad.AddAdvertise(username, title, tag)
			}
		case "rem_advertise":
			if len(tokens) != 3 {
				result = ""
			} else {
				result = ad.RemAdvertise(tokens[1], tokens[2])
			}
		case "list_my_advertises":
			if len(tokens) < 2 {
				result = ""
			} else {
				username := tokens[1]
				tag := ""
				if len(tokens) >= 3 {
					tag = tokens[2]
				}
				result = ad.ListMyAdvertises(username, tag)
			}
		case "add_favorite":
			if len(tokens) != 3 {
				result = ""
			} else {
				result = user.AddFavorite(tokens[1], tokens[2])
			}
		case "rem_favorite":
			if len(tokens) != 3 {
				result = ""
			} else {
				result = user.RemFavorite(tokens[1], tokens[2])
			}
		case "list_favorite_advertises":
			if len(tokens) < 2 {
				result = ""
			} else {
				username := tokens[1]
				tag := ""
				if len(tokens) >= 3 {
					tag = tokens[2]
				}
				result = user.ListFavoriteAdvertises(username, tag)
			}
		default:
			result = ""
		}
		fmt.Println(result)
	}
}
