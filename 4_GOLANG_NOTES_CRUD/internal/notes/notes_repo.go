package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Repo struct {
	coll *mongo.Collection
}

func NewRepo(db *mongo.Database) *Repo {
	return &Repo{
		coll: db.Collection("notes"),
	}
}

func (r *Repo) Create(ctx context.Context, note Note) (*Note, error) {
	opCtx , cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	_, err := r.coll.InsertOne(opCtx, note)

	if err != nil {
		return &Note{}, fmt.Errorf("failed to create note: %v", err)
	}

	return &note, nil;
}

func (r *Repo) List(ctx context.Context) ([]Note, error){
	opCtx , cancel := context.WithTimeout(ctx,5*time.Second)

	defer cancel()

	// filter match all the documents in the collection
	filter := bson.M{}

	cursor, err := r.coll.Find(opCtx,filter)

	if err != nil {
		return nil, fmt.Errorf("failed to list notes: %v", err)
	}

	//cusror must be closed after use to free up resources to avoid memory leaks and to ensure that the connection to the database is properly released.
	defer cursor.Close(opCtx)

	var notes []Note

	if err := cursor.All(opCtx,&notes); err != nil {
		return nil, fmt.Errorf("Decode notes: %v",err)
	}

	return notes, nil
}

func (r *Repo) GetById(ctx context.Context, id primitive.ObjectID) (*Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	filter := bson.M{"_id": id}

	var note Note 

	err := r.coll.FindOne(opCtx, filter, options.FindOne()).Decode(&note)

	if err != nil {
		return nil, fmt.Errorf("failed to get note by id: %v", err)
	}

	return &note, nil
}

func (r *Repo) UpdateById(ctx context.Context, id primitive.ObjectID, req UpdateNoteRequest) (*Note, error){
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"title": req.Title,
			"content": req.Content,
			"pinned": req.Pinned,
			"updatedAt": time.Now().UTC(),
		},
	}

	after := options.After
	
	option := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	err := r.coll.FindOneAndUpdate(opCtx, filter, update, &option).Decode(&update)

	if err != nil {
		return nil, fmt.Errorf("failed to update note: %v", err)
	}
	return &Note{
		ID: id,
		Title: req.Title,
		Content: req.Content,
		Pinned: req.Pinned,
		UpdatedAt: time.Now().UTC(),
	}, nil
}	

func (r *Repo) DeleteById(ctx context.Context, id primitive.ObjectID) error {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	filter := bson.M{"_id": id}

	result , err := r.coll.DeleteOne(opCtx,filter)
	if err != nil {
		return fmt.Errorf("failed to delete note: %v", err)
	}
	
	if result.DeletedCount == 0 {
		return fmt.Errorf("no note deleted with id: %s", id.Hex())
	}

	return nil
}