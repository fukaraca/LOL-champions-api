package lib

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//GetRestRequest function handles get request for REST API
func GetRestRequest(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), TIMEOUT)
	defer cancel()
	uRl := c.Request.URL
	//  eg: /rest/champions? name=Jinx & q=stats.hp,stats.mp,spells    makes query to get HP, MP and spell informations of named hero Jinx
	if heroName := uRl.Query()["name"]; heroName != nil {
		queryList := strings.Split(uRl.Query()["q"][0], ",")
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
