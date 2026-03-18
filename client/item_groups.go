// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/ramzes4rules/rsl6-integration-api/models"
)

// ItemGroupService handles item group operations
type ItemGroupService struct {
	client *Client
}

// Create creates a new item group
func (s *ItemGroupService) Create(ctx context.Context, req *models.CreateItemGroupRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_groups/create", req, headers)
}

// Rename renames an item group
func (s *ItemGroupService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_groups/rename", req, headers)
}

// SetParent sets the parent of an item group
func (s *ItemGroupService) SetParent(ctx context.Context, req *models.SetParentRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_groups/set_parent", req, headers)
}

// Delete deletes an item group
func (s *ItemGroupService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_groups/delete", req, headers)
}

// Restore restores a deleted item group
func (s *ItemGroupService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_groups/restore", req, headers)
}

// Batch executes multiple commands in a batch
func (s *ItemGroupService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_groups/batch", req, headers)
}

// GetByID retrieves an item group by ID
func (s *ItemGroupService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.ItemGroupDto, error) {
	var response models.ItemGroupDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/item_groups/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of item groups
func (s *ItemGroupService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.ItemGroupListDto, error) {
	var response models.ItemGroupListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/item_groups/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
