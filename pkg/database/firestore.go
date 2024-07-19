package database

import (
	"Fittrackr/pkg/models"
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var client *firestore.Client

func InitFirestore() error {
	ctx := context.Background()
	var err error
	client, err = firestore.NewClient(ctx, "fittrackr-dkozlov", option.WithCredentialsFile("cmd/fittrackr-dkozlov-firebase-adminsdk-sovt7-a54294dc41.json"))
	if err != nil {
		return err
	}
	return nil
}

func PostExercise(ctx context.Context, exerciseID string, exercise models.Exercise) error {
	_, err := client.Collection("exercises").Doc(exerciseID).Set(ctx, exercise)
	return err
}

func GetExercise(ctx context.Context, exerciseID string) (*models.Exercise, error) {
	doc, err := client.Collection("exercises").Doc(exerciseID).Get(ctx)
	if err != nil {
		return nil, err
	}

	var exercise models.Exercise
	err = doc.DataTo(&exercise)
	if err != nil {
		return nil, err
	}

	return &exercise, nil
}
