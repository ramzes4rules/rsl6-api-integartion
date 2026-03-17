// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// LoyaltyCardBlockReasonService handles loyalty card block reason operations
type LoyaltyCardBlockReasonService struct {
	client *Client
}

// Create creates a new loyalty card block reason
func (s *LoyaltyCardBlockReasonService) Create(ctx context.Context, req *models.CreateReasonRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_block_reasons/create", req, headers)
}

// Rename renames a loyalty card block reason
func (s *LoyaltyCardBlockReasonService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_block_reasons/rename", req, headers)
}

// Delete deletes a loyalty card block reason
func (s *LoyaltyCardBlockReasonService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_block_reasons/delete", req, headers)
}

// Restore restores a deleted loyalty card block reason
func (s *LoyaltyCardBlockReasonService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_block_reasons/restore", req, headers)
}

// MoveToArchive moves a loyalty card block reason to archive
func (s *LoyaltyCardBlockReasonService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_block_reasons/move_to_archive", req, headers)
}

// Unarchive restores a loyalty card block reason from archive
func (s *LoyaltyCardBlockReasonService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_block_reasons/unarchive", req, headers)
}

// Batch executes multiple commands in a batch
func (s *LoyaltyCardBlockReasonService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_block_reasons/batch", req, headers)
}

// GetByID retrieves a loyalty card block reason by ID
func (s *LoyaltyCardBlockReasonService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.ReasonDto, error) {
	var response models.ReasonDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/loyalty_card_block_reasons/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of loyalty card block reasons
func (s *LoyaltyCardBlockReasonService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.ReasonListDto, error) {
	var response models.ReasonListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/loyalty_card_block_reasons/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
