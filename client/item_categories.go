// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// ItemCategoryService handles item category operations
type ItemCategoryService struct {
	client *Client
}

// Create creates a new item category
func (s *ItemCategoryService) Create(ctx context.Context, req *models.CreateItemCategoryRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_categories/create", req, headers)
}

// Rename renames an item category
func (s *ItemCategoryService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_categories/rename", req, headers)
}

// SetDescription sets the description of an item category
func (s *ItemCategoryService) SetDescription(ctx context.Context, req *models.SetDescriptionRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_categories/set_description", req, headers)
}

// SetColor sets the color of an item category
func (s *ItemCategoryService) SetColor(ctx context.Context, req *models.SetColorRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_categories/set_color", req, headers)
}

// AddItem adds an item to the category
func (s *ItemCategoryService) AddItem(ctx context.Context, req *models.AddItemToCategoryRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_categories/add_item", req, headers)
}

// RemoveItem removes an item from the category
func (s *ItemCategoryService) RemoveItem(ctx context.Context, req *models.RemoveItemFromCategoryRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_categories/remove_item", req, headers)
}

// AddItemGroup adds an item group to the category
func (s *ItemCategoryService) AddItemGroup(ctx context.Context, req *models.AddItemGroupToCategoryRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_categories/add_item_group", req, headers)
}

// RemoveItemGroup removes an item group from the category
func (s *ItemCategoryService) RemoveItemGroup(ctx context.Context, req *models.RemoveItemGroupFromCategoryRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_categories/remove_item_group", req, headers)
}

// Delete deletes an item category
func (s *ItemCategoryService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_categories/delete", req, headers)
}

// Restore restores a deleted item category
func (s *ItemCategoryService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_categories/restore", req, headers)
}

// MoveToArchive moves an item category to archive
func (s *ItemCategoryService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_categories/move_to_archive", req, headers)
}

// Unarchive restores an item category from archive
func (s *ItemCategoryService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_categories/unarchive", req, headers)
}

// Batch executes multiple commands in a batch
func (s *ItemCategoryService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_categories/batch", req, headers)
}

// GetByID retrieves an item category by ID
func (s *ItemCategoryService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.ItemCategoryDto, error) {
	var response models.ItemCategoryDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/item_categories/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of item categories
func (s *ItemCategoryService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.ItemCategoryListDto, error) {
	var response models.ItemCategoryListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/item_categories/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
