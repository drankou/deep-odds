package storage

import (
	"betsapi_scrapper/types"
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"os"
	"reflect"
	"sync"
	"time"
)

type MongoWrapper struct {
	client   *mongo.Client
	database *mongo.Database
	lock     sync.Mutex
}

const DefaultMongoDatabase = "betsapi"

func (m *MongoWrapper) Connect(connectionString string, database string) {
	log.Infof("Mongo - connecting to %s", connectionString)

	// connect is non blocking(at least not completely) - it returns error if something is wrong with conn string
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	opt := options.Client().ApplyURI(connectionString)
	var poolSize = uint64(100)
	opt.MaxPoolSize = &poolSize
	var socketTimeout = 300 * time.Second
	opt.SocketTimeout = &socketTimeout
	client, connectionErr := mongo.Connect(ctx, opt)
	if connectionErr != nil {
		log.Fatalf("mongo - cannot connect: %v", connectionErr)
	}

	// since the above note, we need to check the connection with ping functionality, after this point the db state
	// could be described as ready
	pingCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	pingErr := client.Ping(pingCtx, nil)
	if pingErr != nil {
		log.Fatalf("mongo - cannot connect(2): %v", pingErr)
	}

	log.Infof("Mongo - connected")
	m.client = client
	m.database = m.client.Database(database)
}

func (m *MongoWrapper) Insert(tableName string, data interface{}) (interface{}, error) {
	collection := m.getCollection(tableName, false)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	res, insertErr := collection.InsertOne(ctx, data)
	if insertErr != nil {
		return "", insertErr
	} else {
		return res.InsertedID, nil
	}
}

func (m *MongoWrapper) CreateCollection(tableName string) *mongo.Collection {
	collection := m.getCollection(tableName, false)

	return collection
}

func (m *MongoWrapper) UpdateById(tableName string, id string, updateMap map[string]interface{}) (int64, error) {
	filter := make(map[string]interface{})
	filter["event.id"] = id

	return m.Update(tableName, filter, updateMap, false, false)
}

func (m *MongoWrapper) Update(tableName string, filter map[string]interface{}, updateData map[string]interface{}, upsert bool, addToSet bool) (int64, error) {
	collection := m.getCollection(tableName, true)

	dataWrapper := make(map[string]map[string]interface{})
	if addToSet {
		dataWrapper["$addToSet"] = updateData
	} else {
		dataWrapper["$set"] = updateData
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, insertErr := collection.UpdateOne(ctx, filter, dataWrapper, &options.UpdateOptions{Upsert: &upsert})
	if insertErr == mongo.ErrUnacknowledgedWrite {
		return 0, nil
	} else if insertErr != nil {
		return 0, insertErr
	} else {
		return res.UpsertedCount, nil
	}
}

func (m *MongoWrapper) UpdateMany(tableName string, filter map[string]interface{}, updateData map[string]interface{}, upsert bool) (int64, error) {
	collection := m.getCollection(tableName, false)

	dataWrapper := make(map[string]map[string]interface{})
	dataWrapper["$set"] = updateData

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, insertErr := collection.UpdateMany(ctx, filter, dataWrapper, &options.UpdateOptions{Upsert: &upsert})

	if insertErr == mongo.ErrUnacknowledgedWrite {
		return 0, nil
	} else if insertErr != nil {
		return 0, insertErr
	} else {
		return res.UpsertedCount, nil
	}
}

func (m *MongoWrapper) ReadOne(tableName string, entryType reflect.Type, where interface{}, opts ...*options.FindOneOptions) (interface{}, error) {
	collection := m.getCollection(tableName, false)

	ctx := context.Background()
	finalEntry := reflect.New(entryType).Interface()
	err := collection.FindOne(ctx, where, opts...).Decode(finalEntry)
	if err != nil {
		return nil, err
	}

	return finalEntry, nil
}

func (m *MongoWrapper) ReadOneDoc(tableName string, where primitive.M, opts ...*options.FindOneOptions) (map[string]interface{}, error) {
	collection := m.getCollection(tableName, false)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var found primitive.M
	err := collection.FindOne(ctx, where, opts...).Decode(&found)
	if err == nil {
		return found, nil
	} else {
		return nil, err
	}
}

func (m *MongoWrapper) StreamAll(tableName string, entryType reflect.Type, where interface{}, opts ...*options.FindOptions) (chan interface{}, error) {
	collection := m.getCollection(tableName, false)
	ctx := context.Background()
	if where == nil {
		where = bson.D{{}}
	}
	cursor, err := collection.Find(ctx, where, opts...)
	if err != nil {
		return nil, err
	}

	out := make(chan interface{})
	go func() {
		for {
			if cursor.Next(ctx) == false {
				cursorErr := cursor.Err()
				if cursorErr != nil {
					log.Error(cursorErr)
				}
				close(out)
				break
			}
			finalEntry := reflect.New(entryType).Interface()
			decodeErr := cursor.Decode(finalEntry)
			if decodeErr != nil {
				log.Errorf("Cannot decode mongodb entry: %v", decodeErr)
				continue
			}

			out <- finalEntry.(interface{})
		}
	}()
	return out, nil
}

func (m *MongoWrapper) StreamCollection(tableName string, where map[string]interface{}, opts ...*options.FindOptions) (chan map[string]interface{}, error) {
	collection := m.getCollection(tableName, false)
	ctx := context.Background()
	cursor, err := collection.Find(ctx, where, opts...)
	if err != nil {
		return nil, err
	}

	var out = make(chan map[string]interface{})
	go func() {
		for {
			if cursor.Next(ctx) == false {
				close(out)
				break
			}
			bsonedEntry := make(map[string]interface{})
			decodeErr := cursor.Decode(&bsonedEntry)
			if decodeErr != nil {
				log.Errorf("Cannot decode mongodb entry: %v", decodeErr)
				continue
			}
			out <- bsonedEntry
		}
	}()
	return out, nil
}

func (m *MongoWrapper) ReadAll(tableName string, entryType reflect.Type, where map[string]interface{}, opts ...*options.FindOptions) ([]interface{}, error) {
	collection := m.getCollection(tableName, false)
	var out []interface{}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, where, opts...)
	if err != nil {
		return nil, err
	}
	for {
		if cursor.Next(ctx) == false {
			break
		}
		finalEntry := reflect.New(entryType).Interface()
		decodeErr := cursor.Decode(finalEntry)
		if decodeErr != nil {
			log.Errorf("Cannot decode mongodb entry: %v", decodeErr)
			continue
		}

		out = append(out, finalEntry)
	}
	return out, nil
}

func (m *MongoWrapper) DeleteMany(tableName string, filter interface{}) (int64, error) {
	collection := m.getCollection(tableName, false)

	if filter == nil {
		return -1, errors.New(fmt.Sprintf("DeleteMany called with empty filter - this will effectively truncate the %s collection. Aborting", tableName))
	}
	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		return -1, err
	}
	return result.DeletedCount, nil
}

func (m *MongoWrapper) getCollection(tableName string, ignoreAcknowledgement bool) *mongo.Collection {
	if os.Getenv("ENVIRONMENT") != "prod" {
		tableName = "dev-" + tableName
	}

	if ignoreAcknowledgement {
		return m.database.Collection(tableName, &options.CollectionOptions{WriteConcern: writeconcern.New(writeconcern.W(0), writeconcern.J(false))})
	} else {
		return m.database.Collection(tableName)
	}
}

func (m *MongoWrapper) GetFootballEvents(filter map[string]interface{}) ([]*types.FootballEvent, error) {
	entries, err := m.ReadAll("football_event", reflect.TypeOf(types.FootballEvent{}), filter)
	if err != nil {
		return nil, err
	}

	events := entriesToEvents(entries)

	return events, err
}

func entriesToEvents(entries []interface{}) []*types.FootballEvent {
	var events []*types.FootballEvent
	for _, entry := range entries {
		event, ok := entry.(*types.FootballEvent)
		if !ok {
			log.Error("wrong type assertion")
			continue
		}
		events = append(events, event)
	}

	return events
}

var mongoInstance *MongoWrapper
var getMongoInstanceOnce sync.Once

func GetMongoWrapper() *MongoWrapper {
	getMongoInstanceOnce.Do(func() {
		mongoInstance = &MongoWrapper{}
		mongoInstance.Connect(os.Getenv("MONGO_CONNECTION_STRING"), DefaultMongoDatabase)
	})
	return mongoInstance
}
