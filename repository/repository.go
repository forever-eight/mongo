package repository

import (
	"context"
	"fmt"
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

func NewRepos(ctx context.Context, dbUri string) (*Repository, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Println("problem")
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Println("ctx problem")
	}

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
func (r *Repository) FindProjectByID(ctx context.Context, ID primitive.ObjectID) (*ds.Project, error) {
	col := r.db.Database(database).Collection(projectsCollection)

	/*oID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}*/

	got := col.FindOne(ctx, bson.D{{"_id", ID}})
	if got.Err() != nil {
		return nil, got.Err()
	}
	var doc *ds.Project
	err := got.Decode(&doc)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// Изменяет проект
func (r *Repository) ChangeProject(ctx context.Context, ID primitive.ObjectID, project *ds.Project) error {
	col := r.db.Database(database).Collection(projectsCollection)

	rID := bson.M{"_id": ID}

	upd := bson.D{{Key: "$set", Value: bson.M{"title": project.Title, "channels": project.Channels}}}
	fmt.Println(upd)
	_, err := col.UpdateOne(ctx, rID, upd)
	if err != nil {
		log.Fatal("Error on updating one Hero", err)
	}

	return nil

}

// Удаляет проект
func (r *Repository) DeleteProject(ctx context.Context, ID primitive.ObjectID) error {
	col := r.db.Database(database).Collection(projectsCollection)
	/*oID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}*/
	_, err := col.DeleteOne(ctx, bson.D{{"_id", ID}})
	if err != nil {
		return err
	}

	return nil
}
