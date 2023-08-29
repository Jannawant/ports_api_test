package transport

import (
	"context"
	"net/http"
	"portsApi/internal/common/server"
	"portsApi/internal/domain"
)

type PortService interface {
	GetPort(ctx context.Context, id string) (*domain.Port, error)
}

type HTTPServer struct {
	service PortService
}

func NewHTTPServer(service PortService) *HTTPServer {
	return &HTTPServer{
		service: service,
	}
}

func (s *HTTPServer) GetPort(w http.ResponseWriter, r *http.Request) {
	port, err := s.service.GetPort(r.Context(), r.URL.Query().Get("id"))
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}
	response := Port{
		ID:          port.ID(),
		Name:        port.Name(),
		Code:        port.Code(),
		City:        port.City(),
		Country:     port.Country(),
		Alias:       port.Alias(),
		Regions:     port.Regions(),
		Coordinates: port.Coordinates(),
		Province:    port.Province(),
		Timezone:    port.Timezone(),
		Unlocs:      port.Unlocs(),
	}

	server.RespondOK(response, w, r)
}