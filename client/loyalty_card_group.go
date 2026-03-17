// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// LoyaltyCardGroupService handles loyalty card group operations
type LoyaltyCardGroupService struct {
	client *Client
}

// Create creates a new loyalty card group
func (s *LoyaltyCardGroupService) Create(ctx context.Context, req *models.CreateGroupRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_groups/create", req, headers)
}

// Rename renames a loyalty card group
func (s *LoyaltyCardGroupService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_groups/rename", req, headers)
}

// SetColor sets the color of a loyalty card group
func (s *LoyaltyCardGroupService) SetColor(ctx context.Context, req *models.SetColorRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_groups/set_color", req, headers)
}

// Delete deletes a loyalty card group
func (s *LoyaltyCardGroupService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_groups/delete", req, headers)
}

// Restore restores a deleted loyalty card group
func (s *LoyaltyCardGroupService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_groups/restore", req, headers)
}

// MoveToArchive moves a loyalty card group to archive
func (s *LoyaltyCardGroupService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_groups/move_to_archive", req, headers)
}

// Unarchive restores a loyalty card group from archive
func (s *LoyaltyCardGroupService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_groups/unarchive", req, headers)
}

// Batch executes multiple commands in a batch
func (s *LoyaltyCardGroupService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_groups/batch", req, headers)
}

// GetByID retrieves a loyalty card group by ID
func (s *LoyaltyCardGroupService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.GroupDto, error) {
	var response models.GroupDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/loyalty_card_groups/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of loyalty card groups
func (s *LoyaltyCardGroupService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.GroupListDto, error) {
	var response models.GroupListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/loyalty_card_groups/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
