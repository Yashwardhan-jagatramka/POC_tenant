package dataAccess

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"

	"project-tenant/pkg/configs"
)

var TenantCollection *mongo.Collection = configs.GetCollection(configs.DB, "tenants")
var Cache = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})
