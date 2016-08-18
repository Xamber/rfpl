package rfpl

import (
	"fmt"
	"github.com/Xamber/rfpl/db"
	"time"
)

func Fetch(groups []Group, channel chan<- VKApiTtem) {
	for {
		for _, group := range groups {
			if group.parsetime <= time.Now().Unix()-60 {
				vk := group.FeathNews()
				for _, record := range vk {
					record.Domein = group.domain
					record.People = group.people
					channel <- record
				}
				group.parsetime = time.Now().Unix()
				fmt.Println(fmt.Sprintf("fetching complite: %s", group.domain))
				time.Sleep(time.Millisecond * 1500)
			}
		}
	}
}

func Parse(channel <-chan VKApiTtem) {
	for {
		jsondata := <-channel

		rating := (float64(jsondata.Likes.Count) + float64(jsondata.Comments.Count)*1.5 + float64(jsondata.Reposts.Count)*2.0) / float64(jsondata.People) * 10000

		element := db.News{
			ID:       jsondata.ID,
			Text:     jsondata.Text,
			Date:     jsondata.Date,
			OwnerID:  jsondata.OwnerID,
			Domein:   jsondata.Domein,
			People:   jsondata.People,
			Likes:    jsondata.Likes.Count,
			Reposts:  jsondata.Reposts.Count,
			Comments: jsondata.Comments.Count,
			Rating:   rating,
		}

		err := element.UpdateOrInsert()
		err = element.UpdateHistory()
		if err != nil { // Если ошибка, то ну его
			continue
		}
	}
}
