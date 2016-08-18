package rfpl

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type VKApiResponse struct {
	Response struct {
		Count int         `json:"count"`
		Items []VKApiTtem `json:"items"`
	} `json:"response"`
}

type VKApiTtem struct {
	ID       int    `json:"id"`
	FromID   int    `json:"from_id"`
	OwnerID  int    `json:"owner_id"`
	Date     int    `json:"date"`
	PostType string `json:"post_type"`
	Text     string `json:"text"`
	IsPinned int    `json:"is_pinned,omitempty"`
	Domein   string
	People   int
	Comments struct {
		Count int `json:"count"`
	} `json:"comments"`
	Likes struct {
		Count int `json:"count"`
	} `json:"likes"`
	Reposts struct {
		Count int `json:"count"`
	} `json:"reposts"`
}

type Group struct {
	domain    string
	parsetime int64
	people    int
}

func (g *Group) Get_url() string {
	params := url.Values{}
	params.Add("filter", "owner")
	params.Add("count", "20")
	params.Add("extended", "0")
	params.Add("v", "5.52")
	params.Add("domain", g.domain)
	return BASEURL + string(params.Encode())
}

func (g *Group) FeathNews() []VKApiTtem {

	finalUrl := g.Get_url()

	response, err := http.Get(finalUrl) // Коннектимся по урлу
	if err != nil {                     // Если ошибка то ну его нафиг
		log.Fatal(err)
	}

	defer response.Body.Close() // В конце закрываем респ

	body, err := ioutil.ReadAll(response.Body) // Читаем тело ответа
	if err != nil {                            // Если ошибка то ну его нафиг
		log.Fatal(err)
		os.Exit(1)
	}

	data := bytes.NewBuffer(body) // Читаем данные в буфер

	var result VKApiResponse
	json.NewDecoder(data).Decode(&result) // Маршуем данные в структуру

	return result.Response.Items
}
