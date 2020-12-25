package main

import (
	"fmt"
	"math/rand"
	str "strings"
)

var votes = map[string]int{}

func processResponse(response string) {
	responseItems := str.Split(response, ", ")
	extraItems := len(responseItems) - 10
	for i := extraItems; i > 0; i-- {
		responseItems[rand.Intn(len(responseItems))] = "extra vote"
	}
	for i := 0; i < len(responseItems); i++ {
		if votes[responseItems[i]] == 0 {
			votes[responseItems[i]] = 1
		} else {
			votes[responseItems[i]]++
		}
	}
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println("change test!")

	// name := "Oh Hi Mark [Hi], Nice Outfit [Dork], Bad At Guildwars [BAG], Keyboard Gremlins [POG], Five Finger Discount [Mine], The Danny Devitos [Egg], Todd Howard Ate My Son [WTF], Here comes treble  [Clef], I Want a Real Mount ANet   [HORS], We love Max   [Gay]"
	// fmt.Println(name)
	processResponse("Oh Hi Mark [Hi], Nice Outfit [Dork], Bad At Guildwars [BAG], Keyboard Gremlins [POG], Five Finger Discount [Mine], The Danny Devitos [Egg], Todd Howard Ate My Son [WTF], Here comes treble  [Clef], I Want a Real Mount ANet   [HORS], We love Max   [Gay]")
	processResponse("The Dinkster [Wink], Order of the White Lotus [Pai], Braincells United [ONE], Chili's 3 for 10 [Yum], Here comes treble  [Clef], The Dimmadomes [DOUG], The Rum Tum Tuggers [CATS], Bells Deep [BELL], Your Little Pogchamps [UGH], The Jade Choirey Jinglers [TJCJ]")
	processResponse("The Fungis [Mush], Order of the White Lotus [Pai], Jesters Of Keyboard Encounters [JOKE], Give Us Horses [HORS], Braincells United [ONE], Keyboard Smashers [KS], Yump Gremlins [YUMP], Low Frame Raders [AOE], Keyboard Gremlins [POG]")

	// responseItems := str.Split(name, ", ")
	// for i := 0; i < len(responseItems); i++ {
	// 	fmt.Println(responseItems[i])
	// }

	fmt.Println(votes)
}
