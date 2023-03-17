package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"

	"project-tenant/pkg/dataAccess/mongoDB"
	"project-tenant/pkg/dataAccess/redisCache"
	"project-tenant/pkg/models"
)

// wrapper--
// var col mongo.DbMethods
var db = mongoDB.NewMongoF()
var cache = redisCache.NewredisF()

func CreateTenant(ctx context.Context, reqBody models.Tenant) (*mongo.InsertOneResult, error) {

	ans, err := db.TotalDocs()
	if err != nil {
		return nil, err
	}
	reqBody.UniqueId = ans
	return db.InsertOne(ctx, reqBody)
}
func GetAllTenants(c echo.Context, users []models.Tenant, ctx context.Context) ([]models.Tenant, error) {
	results, err := db.Find(ctx, bson.M{})
	if err != nil {
		return nil, c.JSON(http.StatusInternalServerError, models.TenantResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.Tenant
		if err = results.Decode(&singleUser); err != nil {
			return nil, c.JSON(http.StatusInternalServerError, models.TenantResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}
		users = append(users, singleUser)
	}
	return users, nil
}
func GetATenant(c echo.Context, ctx context.Context, key int) (models.Tenant, error) {
	user2, err3 := cache.Get(ctx, c.QueryParam("tenantId")).Result()
	if err3 != redis.Nil {
		var user models.Tenant
		err := db.FindOne(ctx, bson.M{"uniqueid": key}).Decode(&user)
		if err != nil {
			return user, err
		}
		userV2, _ := json.Marshal(user)
		cacheErr := cache.Set(ctx, strconv.Itoa(key), userV2, 0)
		if cacheErr != nil {
			return user, cacheErr
		}
		return user, nil

	} else {
		var userV3 models.Tenant
		err := json.Unmarshal([]byte(user2), &userV3)
		if err != nil {
			return userV3, err
		}
		return userV3, err
	}
}
func UpdateTenant(ctx context.Context, userId int, reqTenant models.Tenant) (*mongo.UpdateResult, error) {
	var findTenant models.Tenant
	err1 := db.FindOne(ctx, bson.M{"uniqueid": userId}).Decode(&findTenant)
	if err1 != nil {
		return nil, err1
	}

	update := bson.M{"fname": reqTenant.TenantFirstName, "lname": reqTenant.TenantLastName, "location": reqTenant.Country, "domain": reqTenant.BusinessDomain, "email": reqTenant.OfficialEmail, "phone": reqTenant.OfficialPhone}

	result, err := db.UpdateOne(ctx, bson.M{"uniqueid": findTenant.UniqueId}, bson.M{"$set": update})
	reqTenant.UniqueId = userId

	if err != nil {
		return nil, err
	}

	userV2, _ := json.Marshal(reqTenant)
	cacheErr := cache.Set(ctx, strconv.Itoa(findTenant.UniqueId), userV2, 0)
	if cacheErr != nil {
		return nil, cacheErr
	}
	return result, nil
}
