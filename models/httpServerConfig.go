package commonmodels

import httptransport "github.com/go-kit/kit/transport/http"

type HttpServerConfig struct {
	Server  *httptransport.Server
	Route   string
	Methods []string
}
