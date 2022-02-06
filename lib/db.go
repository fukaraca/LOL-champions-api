package lib

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
)

//getFromDBByName makes a query to mongoDB to fetch hero informations.
func getFromDBByName(ctx context.Context, heroName string, queryList []string) (ret gin.H) {
	projection := bson.M{"_id": 0}
	for _, param := range queryList {
		projection[param] = 1
	}
	opts := options.Find().SetProjection(projection)
	filter := bson.M{"name": heroName}
	cursor, err := Coll.Find(ctx, filter, opts)
	defer cursor.Close(ctx)
	if err != nil {
		log.Println("rest get request by name from db failed:", err)
		return nil
	}
	for cursor.Next(ctx) {
		cursor.Decode(&ret)
	}
	return ret
}

//getFromDBWithConditional makes a query with conditionals like $eq or $gt
func getFromDBWithConditional(ctx context.Context, key, op, val string) gin.H {
	projection := bson.M{"_id": 0, "name": 1}
	opts := options.Find().SetProjection(projection)
	filter := bson.M{}
	num, err := strconv.Atoi(val)
	if err != nil {
		filter = bson.M{key: bson.M{fmt.Sprintf("$%s", op): val}}
	} else {
		filter = bson.M{key: bson.M{fmt.Sprintf("$%s", op): num}}
	}
	cursor, err := Coll.Find(ctx, filter, opts)
	defer cursor.Close(ctx)
	if err != nil {
		log.Println("rest get request by condition from db failed:", err)
		return nil
	}
	count := cursor.RemainingBatchLength()
	res := []bson.M{}
	ret := make(gin.H)
	temp := []string{}
	if err := cursor.All(ctx, &res); err != nil {
		log.Println("conditional get req for rest failed:", err)
		return nil
	}
	for _, re := range res {
		for _, v := range re {
			temp = append(temp, v.(string))
		}
	}
	if len(temp) > 0 {
		ret["result"] = temp
		ret["count"] = count
	}

	return ret
}
