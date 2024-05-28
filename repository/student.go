package repository

import (
	"context"
	"errors"
	"fmt"
	"go_eduhub_nosql/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id string) (*model.Student, error)
	Store(s *model.Student) error
	Update(id string, s *model.Student) error
}

type studentRepoImpl struct {
	ctx context.Context
	db  *mongo.Collection
}

func NewStudentRepo(ctx context.Context, db *mongo.Client) *studentRepoImpl {
	return &studentRepoImpl{
		ctx: ctx,
		db:  db.Database("camp").Collection("students"),
	}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	var listStudents []model.Student

	cursor, err := s.db.Find(s.ctx, options.Find())
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &listStudents); err != nil {
		log.Fatal(err)
	}

	return listStudents, nil
}

func (s *studentRepoImpl) FetchByID(id string) (*model.Student, error) {
	var student model.Student

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	err = s.db.FindOne(s.ctx, bson.M{"_id": objectId}).Decode(&student)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("data not found")
		}
		return nil, err
	}

	return &student, nil
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	doc := bson.D{
		{"name", student.Name},
		{"email", student.Email},
		{"address", student.Address},
	}

	_, err := s.db.InsertOne(s.ctx, doc)
	if err != nil {
		return err
	}

	return nil
}

func (s *studentRepoImpl) Update(id string, student *model.Student) error {
	opts := options.Update().SetUpsert(true)
	fmt.Println("student: ", student)
	fmt.Println("id: ", id)
	idPri, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", idPri}}
	fmt.Println("idPri: ", idPri)

	update := bson.D{
		{"$set", bson.D{{"name", student.Name}}},
		{"$set", bson.D{{"email", student.Email}}},
		{"$set", bson.D{{"address", student.Address}}},
	}
	result, err := s.db.UpdateOne(s.ctx, filter, update, opts)
	if err != nil {
		return err
	}

	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
		return nil
	}

	if result.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}

	return nil
}
