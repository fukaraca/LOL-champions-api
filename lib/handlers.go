package lib

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"strings"
)

func GetHeroList(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), TIMEOUT)
	defer cancel()
	retList := getHeroListFromDB(ctx)
	c.JSON(http.StatusOK, gin.H{
		"Heroes": retList,
	})

}

//GetRestRequest function handles get request for REST API
func GetRestRequest(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), TIMEOUT)
	defer cancel()
	uRl := c.Request.URL
	//  eg: /rest/champions? name=Jinx & q=stats.hp,stats.mp,spells    makes query to get HP, MP and spell informations of named hero Jinx
	if heroName := uRl.Query()["name"]; heroName != nil {
		queryList := []string{}
		for _, queryItem := range uRl.Query()["q"] {
			queryList = strings.Split(queryItem, ",")
			break
		}
		retJSON := getFromDBByName(ctx, heroName[0], queryList)
		c.JSON(http.StatusOK, retJSON)
	} else {
		// eg:  /rest/champions? key=hp & op=gt & val=510 makes  query to get hero names which has more hp than 510
		key := uRl.Query()["key"]
		op := uRl.Query()["op"]
		val := uRl.Query()["value"]
		if len(key) == 0 || len(op) == 0 || len(val) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"example:":        "/rest/champions?key=[SOME_KEY_FIELD]'ampersand_sign'op=[CONDITIONAL_LIKE_EQ_or_GT]'ampersand_sign'val=[CONDITIONAL_VALUE]",
				"_status message": "missing parameter like key,op or val",
			})
			return
		}
		retJSON := getFromDBWithConditional(ctx, key[0], op[0], val[0])
		if len(retJSON) > 0 {
			c.JSON(http.StatusOK, retJSON)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status message:": "no match found",
			})
		}

	}

}

//PostRestRequest function handles new hero creation requests with form-data.
func PostRestRequestAsForm(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), TIMEOUT)
	defer cancel()

	multiP, err := c.MultipartForm()
	if err != nil {
		log.Println("form couldn't be parsed:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"detail":         " form key,op and val parameters needed as post-form key",
			"status message": "missing paramter",
		})
		return
	}
	if nameField, ok := multiP.Value["name"]; !ok {
		log.Println("name fied is needed:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"detail":         " at least one key parameter must be 'name' to indicate hero name",
			"status message": "missing parameter",
		})
		return
	} else {
		for _, name := range nameField {
			if heroExist(name) {
				c.JSON(http.StatusNotModified, gin.H{
					"detail":         " there must a new hero to be added. we already have hero: " + name,
					"status message": "already have",
				})
				return
			}
			break
		}
	}
	document := bson.M{}

	for key, vals := range multiP.Value {
		for _, val := range vals {
			document[key] = val
		}

	}
	result := false
	if result = insertNewHero(ctx, document); !result {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail":         " insertion to db failed, consult db admin",
			"status message": "error",
		})
		return
	}

	c.JSON(http.StatusCreated, result)
}

/*PostRestRequestAsJSON is used for a better-errorless hero creation with json. Example raw json input:

{
    "name": "yepisyeni hero",
    "stats": {
        "hp": 650,
        "mp": 260
    }
}

*/
func PostRestRequestAsJSON(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), TIMEOUT)
	defer cancel()
	heroObj := HeroInfoStruct{}
	err := c.BindJSON(&heroObj)
	if err != nil {
		log.Println("bind json for creation failed:", err)
		return
	}
	if heroExist(heroObj.Name) || heroObj.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail":         "either name field is missing or named hero is already exist",
			"status message": "bad request",
		})
		return
	}
	document, err := bson.Marshal(heroObj)
	if err != nil {
		log.Println("bson marshalling failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail":         "",
			"status message": "json couldn't be parsed",
		})
		return
	}
	res, err := Coll.InsertOne(ctx, document)
	if err != nil {
		log.Println("new hero insertion failed:", err, "result:", res)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail":         "",
			"status message": "json couldn't be inserted",
		})
		return
	}
	c.JSON(http.StatusOK, true)
}

//UpdateRestHero handles update hero information
func UpdateRestHero(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), TIMEOUT)
	defer cancel()

	heroName := c.PostForm("name")
	key := c.PostForm("key")
	op := c.PostForm("op")
	val := c.PostForm("val")

	if !heroExist(heroName) || heroName == "" || key == "" || op == "" || val == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail":         "either key,op,name or val fields are missing or invalid hero name",
			"status message": "bad request",
		})
		return
	}
	res, ok := updateHero(ctx, heroName, key, op, val)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail":         "hero informations couldn't be updated, consult db admin",
			"status message": "error",
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

//DeleteHero function handles for deleting a hero from DB.
func DeleteRestHero(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), TIMEOUT)
	defer cancel()
	heroName := c.PostForm("name")
	if !heroExist(heroName) {
		log.Println("hero to be deleted is not exist:")
		c.JSON(http.StatusBadRequest, gin.H{
			"detail":         "invalid hero name",
			"status message": "bad request",
		})
		return
	}
	result := deleteHeroFromDB(ctx, heroName)
	if !result {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail":         "hero couldn't be deleted, consult db admin",
			"status message": "error",
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
