package splits

import (
	"context"

	splitmodels "github.com/dashitoishi23/splitwise-go/pkg/splits/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SplitService interface {
	SaveTheTransaction(ctx context.Context, transaction splitmodels.Transaction) error
	HowMuchIOwe(ctx context.Context, MobileNumer string) ([]splitmodels.Debt, error)
	HowMuchOthersOweToMe(ctx context.Context, MobileNumber string) ([]splitmodels.Transaction, error)
	ChangePaymentStatus(ctx context.Context, MobileNumber string) (bool, error)
}

type splitService struct {
	db *mongo.Database
}

func NewSplitService(client *mongo.Client) SplitService {
	return &splitService{
		db: client.Database("ExpenseSharing"),
	}
}

func (s *splitService) SaveTheTransaction(ctx context.Context, transaction splitmodels.Transaction) error {
	_, err := s.db.Collection("Transactions").InsertOne(ctx, transaction)

	if err != nil {
		return err
	}

	return nil
}

func (s *splitService) HowMuchIOwe(ctx context.Context, MobileNumber string) ([]splitmodels.Debt, error) {
	cursor, err := s.db.Collection("Transactions").Find(ctx, bson.M{"spentBy": bson.M{"mobile": bson.M{"$ne": MobileNumber},
		"split": MobileNumber}})

	if err != nil {
		return nil, err
	}

	var debts []splitmodels.Debt

	for cursor.Next(context.TODO()) {
		var result bson.D

		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
	}
}
