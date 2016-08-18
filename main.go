package rfpl

import (
	"fmt"
)

const BASEURL string = "https://api.vk.com/method/wall.get?"

var GROUPSNAMES = []Group{
	{domain: "zenit", people: 809707},
	{domain: "pfc_cska", people: 481131},
	{domain: "fckrasnodar", people: 69241},
	{domain: "fctomtomsk", people: 18161},
	{domain: "fc_ue", people: 809707},
	{domain: "pfc_arsenal", people: 16301},
	{domain: "fcorenburg", people: 16301},
	{domain: "fc_anji_ru", people: 51508},
	{domain: "ufafc", people: 34700},
	{domain: "amkar", people: 49144},
	{domain: "fcrk", people: 87994},
	{domain: "ksamara", people: 43200},
	{domain: "fcterekgrozny", people: 19048},
	{domain: "fclokomotivmoscow", people: 250000},
	{domain: "fcrostov", people: 45000},
}

func main() {
	var channel_json chan VKApiTtem = make(chan VKApiTtem)

	go Fetch(GROUPSNAMES, channel_json)
	go Parse(channel_json)

	var input string
	fmt.Scanln(&input)
}
