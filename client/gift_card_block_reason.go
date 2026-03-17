// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// GiftCardBlockReasonService handles gift card block reason operations
type GiftCardBlockReasonService struct {
	client *Client
}

// Create creates a new gift card block reason
func (s *GiftCardBlockReasonService) Create(ctx context.Context, req *models.CreateReasonRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_card_block_reasons/create", req, headers)
}

// Rename renames a gift card block reason
func (s *GiftCardBlockReasonService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_card_block_reasons/rename", req, headers)
}

// Delete deletes a gift card block reason
func (s *GiftCardBlockReasonService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_card_block_reasons/delete", req, headers)
}

// Restore restores a deleted gift card block reason
func (s *GiftCardBlockReasonService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_card_block_reasons/restore", req, headers)
}

// MoveToArchive moves a gift card block reason to archive
func (s *GiftCardBlockReasonService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_card_block_reasons/move_to_archive", req, headers)
}

// Unarchive restores a gift card block reason from archive
func (s *GiftCardBlockReasonService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_card_block_reasons/unarchive", req, headers)
}

// Batch executes multiple commands in a batch
func (s *GiftCardBlockReasonService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_card_block_reasons/batch", req, headers)
}

// GetByID retrieves a gift card block reason by ID
func (s *GiftCardBlockReasonService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.ReasonDto, error) {
	var response models.ReasonDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/gift_card_block_reasons/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of gift card block reasons
func (s *GiftCardBlockReasonService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.ReasonListDto, error) {
	var response models.ReasonListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/gift_card_block_reasons/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
