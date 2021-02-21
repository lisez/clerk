package providers

import (
	"clerk/jobs"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/xeipuuv/gojsonschema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongodbProvider ...
type MongodbProvider struct {
	Config     jobs.ClerkRemoteConfig
	Timeout    time.Duration
	Datebase   string
	Collection string
}

// Check ...
func (p *MongodbProvider) Check() {
	if p.Datebase == "" {
		log.Fatal("no database")
	}
	if p.Collection == "" {
		log.Fatal("no collection")
	}
}

func (p *MongodbProvider) isFatalError(err error) {
	if err != nil {
		log.Fatalf("mongo error: %v", err)
	}
}

// DocValidator ...
func (p *MongodbProvider) DocValidator() func(jdoc map[string]interface{}) {
	schemaLoader := gojsonschema.NewGoLoader(p.Config.Schema)
	schema, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		log.Fatalf("load json schema failed: %s", err)
	}

	return func(jdoc map[string]interface{}) {
		docLoader := gojsonschema.NewGoLoader(jdoc)
		result, err := schema.Validate(docLoader)
		if err != nil {
			log.Fatalf("load json file failed: %s", err)
		}

		docID := jdoc["_id"]
		if result.Valid() {
			fmt.Printf("%s: pass\n", docID)
		} else {
			fmt.Printf("%s: invalid JSON, reasons: %s\n", docID, result.Errors())
		}
	}
}

// Start ...
func (p *MongodbProvider) Start() {
	client, err := mongo.NewClient(options.Client().ApplyURI(p.Config.URI))
	p.isFatalError(err)

	ctx, cancel := context.WithTimeout(context.Background(), p.Timeout)
	defer cancel()
	err = client.Connect(ctx)
	p.isFatalError(err)
	defer client.Disconnect(ctx)

	database := client.Database(p.Datebase)
	collection := database.Collection(p.Collection)

	cursor, err := collection.Find(ctx, bson.M{})
	p.isFatalError(err)
	defer cursor.Close(ctx)

	validator := p.DocValidator()
	for cursor.Next(ctx) {
		var bdoc bson.M
		err = cursor.Decode(&bdoc)
		p.isFatalError(err)
		validator(bdoc)
	}
}
