package splits

import (
	"context"
	"errors"
	"fmt"

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
	ChangePaymentStatus(ctx context.Context, MobileNumber string, TransactionIdentifier string) (bool, error)
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
		transaction.Split[idx].PaymentStatus = splitmodels.Pending.String()
	}

	transaction.OverallPaymentStatus = splitmodels.Pending.String()

	_, err := s.db.InsertOne(ctx, transaction)

	if err != nil {
		return err
	}

	return nil
}

func (s *splitService) HowMuchIOwe(ctx context.Context, MobileNumber string) ([]splitmodels.Debt, error) {
	var debts []splitmodels.Transaction
	var owedTransactions []splitmodels.Debt
	cursor, err := s.db.Find(ctx, bson.M{"$and": []bson.M{
		{"spentby.mobile": bson.M{"$ne": MobileNumber}},
		{"split.mobile": MobileNumber},
		{"split.paymentstatus": splitmodels.Pending.String()},
	}})

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
			PaymentStatus: debt.Split[idx].PaymentStatus,
		})
	}

	return owedTransactions, nil

}

func (s *splitService) HowMuchOthersOweToMe(ctx context.Context, MobileNumber string) ([]splitmodels.Transaction, error) {
	var debts []splitmodels.Transaction
	cursor, err := s.db.Find(ctx, bson.M{"spentby.mobile": MobileNumber, "overallpaymentstatus": splitmodels.Pending.String()})

	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &debts); err != nil {
		return nil, err
	}

	if len(debts) == 0 {
		return nil, errors.New(util.DEBT_NOT_FOUND)
	}

	return debts, nil
}

func (s *splitService) ChangePaymentStatus(ctx context.Context, MobileNumber string, TransactionIdentifier string) (bool, error) {
	cursor := s.db.FindOne(ctx, bson.M{"identifier": TransactionIdentifier})

	if cursor.Err() != nil {
		if cursor.Err().Error() == mongo.ErrNoDocuments.Error() {
			return false, errors.New(util.TRANSACTION_NOT_FOUND)
		}

		return false, cursor.Err()
	}

	var existingTransaction splitmodels.Transaction

	if err := cursor.Decode(&existingTransaction); err != nil {
		return false, err
	}

	idx := slices.IndexFunc(existingTransaction.Split, func(s splitmodels.Split) bool { return s.Mobile == MobileNumber })

	if idx == -1 {
		return false, errors.New(util.DEBT_NOT_FOUND)
	}

	existingTransaction.Split[idx].PaymentStatus = splitmodels.Paid.String()

	index := slices.IndexFunc(existingTransaction.Split, func(s splitmodels.Split) bool { return s.PaymentStatus == splitmodels.Pending.String() })

	if index == -1 {
		existingTransaction.OverallPaymentStatus = splitmodels.Paid.String()
	}

	updRes, err := s.db.UpdateOne(ctx, bson.M{"_id": existingTransaction.Id}, bson.M{"$set": existingTransaction})

	if err != nil {
		return false, err
	}

	fmt.Print(updRes)

	return true, nil
}
