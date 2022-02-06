package lib

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var R *gin.Engine
var Client = &mongo.Client{}
var Coll *mongo.Collection
var TIMEOUT = 10 * time.Second

type LolHeroInfo struct {
	Type    string                    `json:"type"  `
	Format  string                    `json:"format"`
	Version string                    `json:"version"`
	Data    map[string]HeroInfoStruct `json:"data"`
}

type HeroInfoStruct struct {
	ID    string `json:"id" bson:"id"`
	Key   string `json:"key" bson:"key"`
	Name  string `json:"name" bson:"name"`
	Title string `json:"title" bson:"title"`
	Image struct {
		Full   string `json:"full" bson:"full"`
		Sprite string `json:"sprite" bson:"sprite"`
		Group  string `json:"group" bson:"group"`
		X      int    `json:"x" bson:"x"`
		Y      int    `json:"y" bson:"y"`
		W      int    `json:"w" bson:"w"`
		H      int    `json:"h" bson:"h"`
	} `json:"image" bson:"image"`
	Skins []struct {
		ID      string `json:"id" bson:"id"`
		Num     int    `json:"num" bson:"num"`
		Name    string `json:"name" bson:"name"`
		Chromas bool   `json:"chromas" bson:"chromas"`
	} `json:"skins" bson:"skins"`
	Lore      string   `json:"lore" bson:"lore"`
	Blurb     string   `json:"blurb" bson:"blurb"`
	Allytips  []string `json:"allytips" bson:"allytips"`
	Enemytips []string `json:"enemytips" bson:"enemytips"`
	Tags      []string `json:"tags" bson:"tags"`
	Partype   string   `json:"partype" bson:"partype"`
	Info      struct {
		Attack     int `json:"attack" bson:"attack"`
		Defense    int `json:"defense" bson:"defense"`
		Magic      int `json:"magic" bson:"magic"`
		Difficulty int `json:"difficulty" bson:"difficulty"`
	} `json:"info" bson:"info"`
	Stats struct {
		Hp                   float64 `json:"hp" bson:"hp"`
		Hpperlevel           float64 `json:"hpperlevel" bson:"hpperlevel"`
		Mp                   float64 `json:"mp" bson:"mp"`
		Mpperlevel           float64 `json:"mpperlevel" bson:"mpperlevel"`
		Movespeed            float64 `json:"movespeed" bson:"movespeed"`
		Armor                float64 `json:"armor" bson:"armor"`
		Armorperlevel        float64 `json:"armorperlevel" bson:"armorperlevel"`
		Spellblock           float64 `json:"spellblock" bson:"spellblock"`
		Spellblockperlevel   float64 `json:"spellblockperlevel" bson:"spellblockperlevel"`
		Attackrange          float64 `json:"attackrange" bson:"attackrange"`
		Hpregen              float64 `json:"hpregen" bson:"hpregen"`
		Hpregenperlevel      float64 `json:"hpregenperlevel" bson:"hpregenperlevel"`
		Mpregen              float64 `json:"mpregen" bson:"mpregen"`
		Mpregenperlevel      float64 `json:"mpregenperlevel" bson:"mpregenperlevel"`
		Crit                 float64 `json:"crit" bson:"crit"`
		Critperlevel         float64 `json:"critperlevel" bson:"critperlevel"`
		Attackdamage         float64 `json:"attackdamage" bson:"attackdamage"`
		Attackdamageperlevel float64 `json:"attackdamageperlevel" bson:"attackdamageperlevel"`
		Attackspeedperlevel  float64 `json:"attackspeedperlevel" bson:"attackspeedperlevel"`
		Attackspeed          float64 `json:"attackspeed" bson:"attackspeed"`
	} `json:"stats" bson:"stats"`
	Spells []struct {
		ID          string `json:"id" bson:"id"`
		Name        string `json:"name" bson:"name"`
		Description string `json:"description" bson:"description"`
		Tooltip     string `json:"tooltip" bson:"tooltip"`
		Leveltip    struct {
			Label  []string `json:"label" bson:"label"`
			Effect []string `json:"effect" bson:"effect"`
		} `json:"leveltip" bson:"leveltip"`
		Maxrank      int       `json:"maxrank" bson:"maxrank"`
		Cooldown     []float64 `json:"cooldown" bson:"cooldown"`
		CooldownBurn string    `json:"cooldownBurn" bson:"cooldown_burn"`
		Cost         []int     `json:"cost" bson:"cost"`
		CostBurn     string    `json:"costBurn" bson:"cost_burn"`
		Datavalues   struct {
		} `json:"datavalues" bson:"datavalues"`
		Effect     []interface{} `json:"effect" bson:"effect"`
		EffectBurn []interface{} `json:"effectBurn" bson:"effect_burn"`
		Vars       []interface{} `json:"vars" bson:"vars"`
		CostType   string        `json:"costType" bson:"cost_type"`
		Maxammo    string        `json:"maxammo" bson:"maxammo"`
		Range      []int         `json:"range" bson:"range"`
		RangeBurn  string        `json:"rangeBurn" bson:"range_burn"`
		Image      struct {
			Full   string `json:"full" bson:"full"`
			Sprite string `json:"sprite" bson:"sprite"`
			Group  string `json:"group" bson:"group"`
			X      int    `json:"x" bson:"x"`
			Y      int    `json:"y" bson:"y"`
			W      int    `json:"w" bson:"w"`
			H      int    `json:"h" bson:"h"`
		} `json:"image" bson:"image"`
		Resource string `json:"resource" bson:"resource"`
	} `json:"spells" bson:"spells"`
	Passive struct {
		Name        string `json:"name" bson:"name"`
		Description string `json:"description" bson:"description"`
		Image       struct {
			Full   string `json:"full" bson:"full"`
			Sprite string `json:"sprite" bson:"sprite"`
			Group  string `json:"group" bson:"group"`
			X      int    `json:"x" bson:"x"`
			Y      int    `json:"y" bson:"y"`
			W      int    `json:"w" bson:"w"`
			H      int    `json:"h" bson:"h"`
		} `json:"image" bson:"image"`
	} `json:"passive" bson:"passive"`
	Recommended []interface{} `json:"recommended" bson:"recommended"`
}
