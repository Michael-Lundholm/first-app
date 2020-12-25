package main

import (
	"fmt"
	"math/rand"
	str "strings"
)

// struct guildName {
// 	name string
// 	votes int
// }
var votesMap = map[string]int{}

func processResponse(response string) {
	responseItems := str.Split(response, ", ")
	extraItems := len(responseItems) - 10
	for i := extraItems; i > 0; i-- {
		responseItems[rand.Intn(len(responseItems))] = "extra vote"
	}
	for i := 0; i < len(responseItems); i++ {
		if _, ok := votesMap[responseItems[i]]; !ok {
			votesMap[responseItems[i]] = 1
		} else {
			votesMap[responseItems[i]]++
		}
	}
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println("change test!")

	// name := "Oh Hi Mark [Hi], Nice Outfit [Dork], Bad At Guildwars [BAG], Keyboard Gremlins [POG], Five Finger Discount [Mine], The Danny Devitos [Egg], Todd Howard Ate My Son [WTF], Here comes treble  [Clef], I Want a Real Mount ANet   [HORS], We love Max   [Gay]"
	// fmt.Println(name)

	// my response (not in spreadsheet)
	processResponse("Oh Hi Mark [Hi], Nice Outfit [Dork], Bad At Guildwars [BAG], Keyboard Gremlins [POG], Five Finger Discount [Mine], The Danny Devitos [Egg], Todd Howard Ate My Son [WTF], Here comes treble  [Clef], I Want a Real Mount ANet   [HORS], We love Max   [Gay]")

	// responses in spreadsheet
	// processResponse("The Dinkster [Wink], Order of the White Lotus [Pai], Braincells United [ONE], Chili's 3 for 10 [Yum], Here comes treble  [Clef], The Dimmadomes [DOUG], The Rum Tum Tuggers [CATS], Bells Deep [BELL], Your Little Pogchamps [UGH], The Jade Choirey Jinglers [TJCJ]")
	// processResponse("The Fungis [Mush], Order of the White Lotus [Pai], Jesters Of Keyboard Encounters [JOKE], Give Us Horses [HORS], Braincells United [ONE], Keyboard Smashers [KS], Yump Gremlins [YUMP], Low Frame Raders [AOE], Keyboard Gremlins [POG]")

	rawInput :=
		`The Dinkster [Wink], Order of the White Lotus [Pai], Braincells United [ONE], Chili's 3 for 10 [Yum], Here comes treble  [Clef], The Dimmadomes [DOUG], The Rum Tum Tuggers [CATS], Bells Deep [BELL], Your Little Pogchamps [UGH], The Jade Choirey Jinglers [TJCJ]
The Fungis [Mush], Order of the White Lotus [Pai], Jesters Of Keyboard Encounters [JOKE], Give Us Horses [HORS], Braincells United [ONE], Keyboard Smashers [KS], Yump Gremlins [YUMP], Low Frame Raders [AOE], Keyboard Gremlins [POG]
Pokemon Rescue Team [Red/Blue], Nice Outfit [Dork], Braincells United [ONE], Low Frame Raders [AOE], Washed League Players Association [FEED], Ring Dinglers [DING], I Want a Real Mount ANet   [HORS], Bells Deep [BELL], Your Little Pogchamps [UGH], The Jade Choirey Jinglers [TJCJ]
Nice Outfit [Dork], The Dinkster [Wink], Order of the White Lotus [Pai], Give Us Horses [HORS], Rockin Rockin N Rollin [Yoda], Here comes treble  [Clef], Ring Dinglers [DING], I Want a Real Mount ANet   [HORS], Bells Deep [BELL], Your Little Pogchamps [UGH]`

	arrResponses := str.Split(rawInput, "\n")

	// arrResponses := []string{
	// 	"Pokemon Rescue Team [Red/Blue], Nice Outfit [Dork], Braincells United [ONE], Low Frame Raders [AOE], Washed League Players Association [FEED], Ring Dinglers [DING], I Want a Real Mount ANet   [HORS], Bells Deep [BELL], Your Little Pogchamps [UGH], The Jade Choirey Jinglers [TJCJ]",
	// 	"Nice Outfit [Dork], The Dinkster [Wink], Order of the White Lotus [Pai], Give Us Horses [HORS], Rockin Rockin N Rollin [Yoda], Here comes treble  [Clef], Ring Dinglers [DING], I Want a Real Mount ANet   [HORS], Bells Deep [BELL], Your Little Pogchamps [UGH]",
	// 	"The Fungis [Mush], Order of the White Lotus [Pai], Jesters Of Keyboard Encounters [JOKE], Give Us Horses [HORS], Braincells United [ONE], Keyboard Smashers [KS], Yump Gremlins [YUMP], Low Frame Raders [AOE], Keyboard Gremlins [POG]",
	// 	"The Dinkster [Wink], Order of the White Lotus [Pai], Braincells United [ONE], Chili's 3 for 10 [Yum], Here comes treble  [Clef], The Dimmadomes [DOUG], The Rum Tum Tuggers [CATS], Bells Deep [BELL], Your Little Pogchamps [UGH], The Jade Choirey Jinglers [TJCJ]",
	// }
	for _, val := range arrResponses {
		processResponse(val)
	}

	// responseItems := str.Split(name, ", ")
	// for i := 0; i < len(responseItems); i++ {
	// 	fmt.Println(responseItems[i])
	// }

	fmt.Println(votesMap)

}
