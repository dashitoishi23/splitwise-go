package splits

import (
	"context"

	splitmodels "github.com/dashitoishi23/splitwise-go/pkg/splits/models"
)

type SplitService interface {
	SaveTheTransaction(ctx context.Context, transaction splitmodels.Transaction) error
	HowMuchIOwe(ctx context.Context) error
}
