package splits

import (
	"context"
	"errors"

	splitmodels "github.com/dashitoishi23/splitwise-go/pkg/splits/models"
	"github.com/dashitoishi23/splitwise-go/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/exp/slices"
)

type SplitService interface {
	SaveTheTransaction(ctx context.Context, transaction splitmodels.Transaction) error
	HowMuchIOwe(ctx context.Context, MobileNumer string) ([]splitmodels.Debt, error)
	HowMuchOthersOweToMe(ctx context.Context, MobileNumber string) ([]splitmodels.Transaction, error)
	// ChangePaymentStatus(ctx context.Context, MobileNumber string) (bool, error)
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
	for idx := range transaction.Split {
		transaction.Split[idx].PaymentStatus = splitmodels.Pending
	}

	_, err := s.db.InsertOne(ctx, transaction)

	if err != nil {
		return err
	}

	return nil
}

func (s *splitService) HowMuchIOwe(ctx context.Context, MobileNumber string) ([]splitmodels.Debt, error) {
	var debts []splitmodels.Transaction
	var owedTransactions []splitmodels.Debt
	cursor, err := s.db.Find(ctx, bson.M{"spentBy.mobile": bson.M{"$ne": MobileNumber},
		"split.mobile": MobileNumber, "split.paymentstatus": splitmodels.Pending})

	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &debts); err != nil {
		return nil, err
	}

	if len(debts) == 0 {
		return nil, errors.New(util.DEBT_NOT_FOUND)
	}

	for _, debt := range debts {
		idx := slices.IndexFunc(debt.Split, func(t splitmodels.Split) bool { return t.Mobile == MobileNumber })

		if idx == -1 {
			return nil, errors.New(util.DEBT_NOT_FOUND)
		}
		owedTransactions = append(owedTransactions, splitmodels.Debt{
			TotalAmount:   debt.TotalAmount,
			Place:         debt.Place,
			Date:          debt.Date,
			SpentBy:       debt.SpentBy,
			Npeople:       debt.NPeople,
			MyShare:       debt.Split[idx].ShareAmount,
			PaymentStatus: debt.Split[idx].PaymentStatus.String(),
		})
	}

	return owedTransactions, nil

}

func (s *splitService) HowMuchOthersOweToMe(ctx context.Context, MobileNumber string) ([]splitmodels.Transaction, error) {
	cursor, err := s.db.Find(ctx, bson.M{"spentBy": bson.M{"mobile": MobileNumber}})

	if err != nil {
		return nil, err
	}

	var owedTransactions []splitmodels.Transaction

	if err := cursor.All(ctx, &owedTransactions); err != nil {
		return nil, err
	}

	return owedTransactions, nil
}
