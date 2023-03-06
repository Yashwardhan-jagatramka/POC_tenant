package mongoDB

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"

	"project-tenant/pkg/dataAccess"
	"project-tenant/pkg/models"
)

type mongoF struct {
	tenantCollection *mongo.Collection
}

func NewMongoF() *mongoF {
	return &mongoF{
		tenantCollection: dataAccess.TenantCollection,
	}
}
func (db *mongoF) TotalDocs() (int, error) {
	ans, err := db.tenantCollection.EstimatedDocumentCount(context.Background())
	len := (int)(ans)
	return len, err
}
func (db *mongoF) InsertOne(ctx context.Context, reqBody models.Tenant) (*mongo.InsertOneResult, error) {
	return db.tenantCollection.InsertOne(ctx, reqBody)
}
func (db *mongoF) Find(ctx context.Context, filter interface{}) (*mongo.Cursor, error) {
	return dataAccess.TenantCollection.Find(ctx, filter)
}
func (db *mongoF) FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
	return dataAccess.TenantCollection.FindOne(ctx, filter)
}
func (db *mongoF) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return dataAccess.TenantCollection.UpdateOne(ctx, filter, update)
}
