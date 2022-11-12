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
	db *mongo.Collection
}

func NewSplitService(client *mongo.Client) SplitService {
	return &splitService{
		db: client.Database("ExpenseSharing").Collection("Transaction"),
	}
}

func (s *splitService) SaveTheTransaction(ctx context.Context, transaction splitmodels.Transaction) error {
	_, err := s.db.InsertOne(ctx, transaction)

	if err != nil {
		return err
	}

	return nil
}

func (s *splitService) HowMuchIOwe(ctx context.Context, MobileNumber string) ([]splitmodels.Debt, error) {
	var debts []splitmodels.Debt
	cursor, err := s.db.Find(ctx, bson.M{"spentBy": bson.M{"mobile": bson.M{"$ne": MobileNumber},
		"split": MobileNumber}})

	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &debts); err != nil {
		return nil, err
	}

	return debts, nil

}
