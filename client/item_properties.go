// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/ramzes4rules/rsl6-integration-api/models"
)

// ItemPropertyService handles item property operations
type ItemPropertyService struct {
	client *Client
}

// Create creates a new item property
func (s *ItemPropertyService) Create(ctx context.Context, req *models.CreateItemPropertyRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_properties/create", req, headers)
}

// Rename renames an item property
func (s *ItemPropertyService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_properties/rename", req, headers)
}

// Delete deletes an item property
func (s *ItemPropertyService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_properties/delete", req, headers)
}

// Restore restores a deleted item property
func (s *ItemPropertyService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_properties/restore", req, headers)
}

// AddEnum adds an enum value to an item property
func (s *ItemPropertyService) AddEnum(ctx context.Context, req *models.AddEnumRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_properties/add_enum", req, headers)
}

// RenameEnum renames an enum value of an item property
func (s *ItemPropertyService) RenameEnum(ctx context.Context, req *models.RenameEnumRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_properties/rename_enum", req, headers)
}

// DeleteEnum deletes an enum value from an item property
func (s *ItemPropertyService) DeleteEnum(ctx context.Context, req *models.DeleteEnumRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_properties/delete_enum", req, headers)
}

// RestoreEnum restores a deleted enum value of an item property
func (s *ItemPropertyService) RestoreEnum(ctx context.Context, req *models.RestoreEnumRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_properties/restore_enum", req, headers)
}

// Batch executes multiple commands in a batch
func (s *ItemPropertyService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/item_properties/batch", req, headers)
}

// GetByID retrieves an item property by ID
func (s *ItemPropertyService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.PropertyDefinitionDto, error) {
	var response models.PropertyDefinitionDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/item_properties/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of item properties
func (s *ItemPropertyService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.PropertyDefinitionListDto, error) {
	var response models.PropertyDefinitionListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/item_properties/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
