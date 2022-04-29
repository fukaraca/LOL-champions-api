package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"
)

//FetchHeroesFromRiotApi fetches hero infos from Riot API and inserts to DB
func FetchHeroesFromRiotApi() {
	heroes := ExcelToSlice()
	for _, hero := range heroes {
		time.Sleep(500 * time.Millisecond)
		getHero := GetHeroInfo(hero)
		if !CreateDocument(getHero) {
			fmt.Printf("%s couldn't be added\n", hero)
			break
		} else {
			fmt.Println(hero, " has been added to db")
		}
	}
}

//get hero names from the excel file
func ExcelToSlice() []string {
	filePath := "lolheroes.xlsx"
	f, err := excelize.OpenFile(filepath.ToSlash(filePath))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var lolHeroes []string
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}

	}()
	n := 158 //total count of lol heroes
	for i := 1; i <= n; i++ {
		heroName, err := f.GetCellValue("Sayfa1", fmt.Sprintf("E%d", i))
		if err != nil {
			fmt.Println("read from xlsx failed:", err)
		}
		heroName = strings.TrimSpace(heroName)
		lolHeroes = append(lolHeroes, heroName)

	}
	return lolHeroes
}

func GetHeroInfo(heroName string) *HeroInfoStruct {

	getUrl := fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/12.2.1/data/en_US/champion/%s.json", url.PathEscape(heroName))
	resp, err := http.Get(getUrl)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	read, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	heroGoverted := LolHeroInfo{}
	err = json.Unmarshal(read, &heroGoverted)
	if err != nil {
		fmt.Println("unmarshall failed:", err)
		return nil
	}
	tempHeroInfo := HeroInfoStruct{}
	for _, v := range heroGoverted.Data {
		tempHeroInfo = v
	}
	return &tempHeroInfo

}

func CreateDocument(hero *HeroInfoStruct) bool {
	document, err := bson.Marshal(hero)
	if err != nil {
		fmt.Println("bson marshalling failed:", err)
		return false
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = Coll.InsertOne(ctx, document)
	if err != nil {
		fmt.Println("insert to collection failed:", err)
		return false
	}
	return true
}
