package splitendpoints

import (
	"context"

	splitservice "github.com/dashitoishi23/splitwise-go/pkg/splits/service"
	"github.com/go-kit/kit/endpoint"
)

type Set struct {
	SaveTheTransactionEndpoint endpoint.Endpoint
}

func New(svc splitservice.SplitService) Set {
	var saveTheTransactionEndpoint endpoint.Endpoint
	{
		saveTheTransactionEndpoint = SaveTheTransactionEndpoint(svc)
	}

	return Set{
		SaveTheTransactionEndpoint: saveTheTransactionEndpoint,
	}
}

func SaveTheTransactionEndpoint(svc splitservice.SplitService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SaveTheTransactionRequest)
		req.Transaction.InitFields()

		res := svc.SaveTheTransaction(context.TODO(), req.Transaction)

		return SaveTheTransactionResponse{res}, res
	}
}
