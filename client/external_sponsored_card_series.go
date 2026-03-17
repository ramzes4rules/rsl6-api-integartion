// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// ExternalSponsoredCardSeriesService handles external sponsored card series operations
type ExternalSponsoredCardSeriesService struct {
	client *Client
}

// Create creates a new external sponsored card series
func (s *ExternalSponsoredCardSeriesService) Create(ctx context.Context, req *models.CreateExternalSponsoredCardSeriesRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/external_sponsored_card_series/create", req, headers)
}

// Rename renames an external sponsored card series
func (s *ExternalSponsoredCardSeriesService) Rename(ctx context.Context, req *models.RenameExternalCardSeriesRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/external_sponsored_card_series/rename", req, headers)
}

// AddCirculation adds a circulation to external sponsored card series
func (s *ExternalSponsoredCardSeriesService) AddCirculation(ctx context.Context, req *models.AddCirculationRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/external_sponsored_card_series/add_circulation", req, headers)
}

// RenameCirculation renames a circulation in external sponsored card series
func (s *ExternalSponsoredCardSeriesService) RenameCirculation(ctx context.Context, req *models.RenameCirculationRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/external_sponsored_card_series/rename_circulation", req, headers)
}

// Batch executes multiple commands in a batch
func (s *ExternalSponsoredCardSeriesService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/external_sponsored_card_series/batch", req, headers)
}

// GetByID retrieves an external sponsored card series by ID
func (s *ExternalSponsoredCardSeriesService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.ExternalCardSeriesDto, error) {
	var response models.ExternalCardSeriesDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/external_sponsored_card_series/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of external sponsored card series
func (s *ExternalSponsoredCardSeriesService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.ExternalCardSeriesListDto, error) {
	var response models.ExternalCardSeriesListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/external_sponsored_card_series/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
