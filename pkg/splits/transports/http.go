package splittransports

import (
	commonmodels "github.com/dashitoishi23/splitwise-go/models"
	splitendpoints "github.com/dashitoishi23/splitwise-go/pkg/splits/endpoints"
	"github.com/dashitoishi23/splitwise-go/util"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHttpHandler(endpoints splitendpoints.Set) []commonmodels.HttpServerConfig {
	var splitServers []commonmodels.HttpServerConfig

	serverOptions := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(util.ErrorEncoder),
	}

	saveTransactionHandler := *httptransport.NewServer(
		endpoints.SaveTheTransactionEndpoint,
		util.DecodeHTTPGenericRequest[splitendpoints.SaveTheTransactionRequest],
		util.EncodeHTTPGenericResponse,
		serverOptions...,
	)

	howMuchIOweHandler := *httptransport.NewServer(
		endpoints.HowMuchIOweEndpoint,
		util.DecodeHTTPGenericRequest[splitendpoints.HowMuchIOweRequest],
		util.EncodeHTTPGenericResponse,
		serverOptions...,
	)

	splitServers = append(splitServers, commonmodels.HttpServerConfig{
		Server:  &saveTransactionHandler,
		Route:   "/save_the_transaction",
		Methods: []string{"POST"},
	}, commonmodels.HttpServerConfig{
		Server:  &howMuchIOweHandler,
		Route:   "/how_much_i_owe",
		Methods: []string{"POST"},
	})

	return splitServers
}
