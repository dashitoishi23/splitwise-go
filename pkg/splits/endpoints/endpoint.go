package splitendpoints

import (
	"context"

	splitservice "github.com/dashitoishi23/splitwise-go/pkg/splits/service"
	"github.com/go-kit/kit/endpoint"
)

type Set struct {
	SaveTheTransactionEndpoint   endpoint.Endpoint
	HowMuchIOweEndpoint          endpoint.Endpoint
	HowMuchOthersOweToMeEndpoint endpoint.Endpoint
	ChangePaymentStatusEndpoint  endpoint.Endpoint
}

func New(svc splitservice.SplitService) Set {
	var saveTheTransactionEndpoint endpoint.Endpoint
	{
		saveTheTransactionEndpoint = SaveTheTransactionEndpoint(svc)
	}

	var howMuchIOweEndpoint endpoint.Endpoint
	{
		howMuchIOweEndpoint = HowMuchIOweEndpoint(svc)
	}

	var howMuchOthersOweToMeEndpoint endpoint.Endpoint
	{
		howMuchOthersOweToMeEndpoint = HowMuchOthersOweToMeEndpoint(svc)
	}

	var changePaymentStatusEndpoint endpoint.Endpoint
	{
		changePaymentStatusEndpoint = ChangePaymentStatusEndpoint(svc)
	}

	return Set{
		SaveTheTransactionEndpoint:   saveTheTransactionEndpoint,
		HowMuchIOweEndpoint:          howMuchIOweEndpoint,
		HowMuchOthersOweToMeEndpoint: howMuchOthersOweToMeEndpoint,
		ChangePaymentStatusEndpoint:  changePaymentStatusEndpoint,
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

func HowMuchIOweEndpoint(svc splitservice.SplitService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(HowMuchIOweRequest)

		res, err := svc.HowMuchIOwe(context.TODO(), req.MobileNumber)

		return HowMuchIOweResponse{res, err}, err
	}
}

func HowMuchOthersOweToMeEndpoint(svc splitservice.SplitService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(HowMuchIOweRequest)

		res, err := svc.HowMuchOthersOweToMe(context.TODO(), req.MobileNumber)

		return HowMuchOthersOweToMeResponse{res, err}, err
	}
}

func ChangePaymentStatusEndpoint(svc splitservice.SplitService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ChangePaymentStatusRequest)

		res, err := svc.ChangePaymentStatus(context.TODO(), req.MobileNumber, req.TransactionIdentifier)

		return ChangePaymentStatusResponse{res, err}, err
	}
}
