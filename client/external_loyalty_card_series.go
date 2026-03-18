// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/ramzes4rules/rsl6-integration-api/models"
)

// ExternalLoyaltyCardSeriesService handles external loyalty card series operations
type ExternalLoyaltyCardSeriesService struct {
	client *Client
}

// Create creates a new external loyalty card series
func (s *ExternalLoyaltyCardSeriesService) Create(ctx context.Context, req *models.CreateExternalLoyaltyCardSeriesRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/external_loyalty_card_series/create", req, headers)
}

// Rename renames an external loyalty card series
func (s *ExternalLoyaltyCardSeriesService) Rename(ctx context.Context, req *models.RenameExternalCardSeriesRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/external_loyalty_card_series/rename", req, headers)
}

// AddCirculation adds a circulation to external loyalty card series
func (s *ExternalLoyaltyCardSeriesService) AddCirculation(ctx context.Context, req *models.AddCirculationRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/external_loyalty_card_series/add_circulation", req, headers)
}

// RenameCirculation renames a circulation in external loyalty card series
func (s *ExternalLoyaltyCardSeriesService) RenameCirculation(ctx context.Context, req *models.RenameCirculationRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/external_loyalty_card_series/rename_circulation", req, headers)
}

// Batch executes multiple commands in a batch
func (s *ExternalLoyaltyCardSeriesService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/external_loyalty_card_series/batch", req, headers)
}

// GetByID retrieves an external loyalty card series by ID
func (s *ExternalLoyaltyCardSeriesService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.ExternalCardSeriesDto, error) {
	var response models.ExternalCardSeriesDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/external_loyalty_card_series/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of external loyalty card series
func (s *ExternalLoyaltyCardSeriesService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.ExternalCardSeriesListDto, error) {
	var response models.ExternalCardSeriesListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/external_loyalty_card_series/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
