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

//getHeroListFromDB function gets hero list from db
func getHeroListFromDB(ctx context.Context) []string {
	projection := bson.M{"_id": 0, "name": 1}
	opts := options.Find().SetProjection(projection)
	filter := bson.D{{}}
	cursor, err := Coll.Find(ctx, filter, opts)
	defer cursor.Close(ctx)
	if err != nil {
		log.Println("rest get request by name from db failed:", err)
		return nil
	}

	tempM := []bson.M{}
	tempL := []string{}
	err = cursor.All(ctx, &tempM)
	if err != nil {
		log.Println("cursor.all failed at hero list:", err)
	}
	for _, tempMap := range tempM {
		for _, v := range tempMap {
			tempL = append(tempL, v.(string))
		}
	}

	return tempL
}

//heroExist checks if hero is in database already
func heroExist(hero string) bool {
	heroList := getHeroListFromDB(context.Background())
	for _, name := range heroList {
		if hero == name {
			return true
		}
	}
	return false

}

//getFromDBByName makes a query to mongoDB to fetch hero informations.
func getFromDBByName(ctx context.Context, heroName string, queryList []string) (ret gin.H) {
	projection := bson.M{"_id": 0}
	for _, param := range queryList {
		if param == "" {
			continue
		}
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

//insertNewHero does insert prepared document to DB and returns boolean result
func insertNewHero(ctx context.Context, document bson.M) bool {

	res, err := Coll.InsertOne(ctx, document)
	if err != nil {
		log.Println("new hero insertion failed:", err, "result:", res)
		return false
	}
	return true
}

func updateHero(ctx context.Context, name, key, op, val string) (gin.H, bool) {
	opts := options.Update()
	opt := opts.SetUpsert(false)
	update := bson.M{}
	num, err := strconv.Atoi(val)
	if err != nil {
		update = bson.M{"$" + op: bson.M{key: val}}
	} else {
		update = bson.M{"$" + op: bson.M{key: num}}
	}
	filter := bson.M{"name": name}
	res, err := Coll.UpdateOne(ctx, filter, update, opt)
	if err != nil {
		log.Println("update failed:", err, res)
		return nil, false
	}
	ret := getFromDBByName(ctx, name, []string{key})
	return ret, true

}

//deleteHeroFromDB is function for deleting hero from DB
func deleteHeroFromDB(ctx context.Context, hero string) bool {

	res, err := Coll.DeleteOne(ctx, bson.M{"name": hero})
	if err != nil {
		log.Println("delete hero failed:", err, res)
		return false
	}
	return true
}
