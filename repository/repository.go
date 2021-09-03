package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/forever-eight/mongo.git/ds"
)

type Repository struct {
	db *mongo.Client
}

const (
	database           = "Projects"
	projectsCollection = "Projects"
)

func New(ctx context.Context, dbUri string) (*Repository, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Println("problem")
	}

	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Println("ctx problem")
	}
	//defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Print(err)
	}

	return &Repository{
		db: client,
	}, nil
}

// Добавляет проект
func (r *Repository) AddProject(ctx context.Context, project *ds.Project) error {
	col := r.db.Database(database).Collection(projectsCollection)

	bookDoc, err := bson.Marshal(project)
	if err != nil {
		return err
	}
	_, err = col.InsertOne(ctx, bookDoc)
	if err != nil {
		return err
	}

	return nil
}

// Находит проект по названию
func (r *Repository) FindProjectByTitle(ctx context.Context, project string) ([]*ds.Project, error) {
	col := r.db.Database(database).Collection(projectsCollection)
	cursor, err := col.Find(ctx, bson.D{{"title", project}})
	if err != nil {
		return nil, err
	}

	var results []*ds.Project
	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// Находит проект по id
func (r *Repository) FindProjectByID(ctx context.Context, ID string) (*ds.Project, error) {
	col := r.db.Database(database).Collection(projectsCollection)

	oID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	got := col.FindOne(ctx, bson.D{{"_id", oID}})
	if got.Err() != nil {
		return nil, got.Err()
	}
	var doc *ds.Project
	err = got.Decode(&doc)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// Удаляет проект
func (r *Repository) DeleteProject(ctx context.Context, ID string) error {
	col := r.db.Database(database).Collection(projectsCollection)
	_, err := col.DeleteOne(ctx, bson.D{{"_id", ID}})
	if err != nil {
		return err
	}

	return nil
}
