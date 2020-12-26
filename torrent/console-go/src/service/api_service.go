package service

import (
	"context"
	"net/http"
	"os/exec"

	"github.com/jsawatzky/server/torrent/console-go/openapi"
)

// ApiService is a service that implents the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type ApiService struct {
	filebotChan chan *exec.Cmd
}

// NewApiService creates a default api service
func NewApiService(filebotChan chan *exec.Cmd) openapi.DefaultApiServicer {
	return &ApiService{filebotChan}
}

// DownloadNotifyPost - Notify server of a completed download
func (s *ApiService) DownloadNotifyPost(ctx context.Context, mediaType string) (openapi.ImplResponse, error) {

	if mediaType != "movies" && mediaType != "tvshows" {
		return openapi.Response(http.StatusBadRequest, "media_type must be one of movie or tvshow"), nil
	}

	cmd := exec.Command("/usr/bin/run-filebot", mediaType)
	s.filebotChan <- cmd

	return openapi.Response(http.StatusOK, nil), nil
}
