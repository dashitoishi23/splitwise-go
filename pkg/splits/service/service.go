package splits

import (
	"context"

	splitmodels "github.com/dashitoishi23/splitwise-go/pkg/splits/models"
)

type SplitService interface {
	SaveTheTransaction(ctx context.Context, transaction splitmodels.Transaction) error
	HowMuchIOwe(ctx context.Context, MobileNumer string) ([]splitmodels.Debt, error)
	HowMuchOthersOweToMe(ctx context.Context, MobileNumber string) ([]splitmodels.Transaction, error)
	ChangePaymentStatus(ctx context.Context, MobileNumber string) (bool, error)
}

type splitService struct {
}

func NewSplitService() SplitService {
	return &splitService{}
}

func (s *splitService) SaveTheTransaction(ctx context.Context, transaction splitmodels.Transaction) error {
	return ctx.Err()
}
