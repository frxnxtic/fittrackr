package database

import (
	"Fittrackr/pkg/model"
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var client *firestore.Client

func InitFirestore() error {
	ctx := context.Background()
	var err error
	client, err = firestore.NewClient(ctx, "fittrackr-dkozlov", option.WithCredentialsFile("cmd/key.json"))
	if err != nil {
		return err
	}
	return nil
}

func PostModel[T model.Model](ctx context.Context, model T, modelName string) error {
	_, _, err := client.Collection(modelName).Add(ctx, model)
	return err
}

func GetAllModels[T model.Model](ctx context.Context, modelName string) ([]T, error) {
	iter := client.Collection(modelName).OrderBy("id", firestore.Asc).Documents(ctx)
	var models []T
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var model T
		err = doc.DataTo(&model)
		if err != nil {
			return nil, err
		}

		models = append(models, model)
	}
	return models, nil
}

func GetModel[T model.Model](ctx context.Context, modelId string, modelName string) (*T, error) {
	doc, err := client.Collection(modelName).Where("id", "==", modelId).Documents(ctx).Next()
	if err != nil {
		return nil, err
	}

	var model T
	err = doc.DataTo(&model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func UpdateModel[T model.Model](ctx context.Context, modelId string, model T, modelName string) error {
	doc, err := client.Collection(modelName).Where("id", "==", modelId).Documents(ctx).Next()
	doc.Ref.Set(ctx, model)
	return err
}

func DeleteModel(ctx context.Context, modelId string, modelName string) error {
	doc, err := client.Collection(modelName).Where("id", "==", modelId).Documents(ctx).Next()
	doc.Ref.Delete(ctx)
	return err
}
