// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// ExternalGiftCardSeriesService handles external gift card series operations
type ExternalGiftCardSeriesService struct {
	client *Client
}

// Create creates a new external gift card series
func (s *ExternalGiftCardSeriesService) Create(ctx context.Context, req *models.CreateExternalGiftCardSeriesRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/external_gift_card_series/create", req, headers)
}

// Rename renames an external gift card series
func (s *ExternalGiftCardSeriesService) Rename(ctx context.Context, req *models.RenameExternalCardSeriesRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/external_gift_card_series/rename", req, headers)
}

// AddCirculation adds a circulation to external gift card series
func (s *ExternalGiftCardSeriesService) AddCirculation(ctx context.Context, req *models.AddCirculationRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/external_gift_card_series/add_circulation", req, headers)
}

// RenameCirculation renames a circulation in external gift card series
func (s *ExternalGiftCardSeriesService) RenameCirculation(ctx context.Context, req *models.RenameCirculationRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/external_gift_card_series/rename_circulation", req, headers)
}

// Batch executes multiple commands in a batch
func (s *ExternalGiftCardSeriesService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/external_gift_card_series/batch", req, headers)
}

// GetByID retrieves an external gift card series by ID
func (s *ExternalGiftCardSeriesService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.ExternalCardSeriesDto, error) {
	var response models.ExternalCardSeriesDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/external_gift_card_series/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of external gift card series
func (s *ExternalGiftCardSeriesService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.ExternalCardSeriesListDto, error) {
	var response models.ExternalCardSeriesListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/external_gift_card_series/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
