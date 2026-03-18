// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/ramzes4rules/rsl6-integration-api/models"
)

// CustomerPropertyService handles customer property operations
type CustomerPropertyService struct {
	client *Client
}

// Create creates a new customer property
func (s *CustomerPropertyService) Create(ctx context.Context, req *models.CreateCustomerPropertyRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customer_properties/create", req, headers)
}

// Rename renames a customer property
func (s *CustomerPropertyService) Rename(ctx context.Context, req *models.RenameCustomerPropertyRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customer_properties/rename", req, headers)
}

// Delete deletes a customer property
func (s *CustomerPropertyService) Delete(ctx context.Context, req *models.DeleteCustomerPropertyRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customer_properties/delete", req, headers)
}

// Restore restores a deleted customer property
func (s *CustomerPropertyService) Restore(ctx context.Context, req *models.RestoreCustomerPropertyRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customer_properties/restore", req, headers)
}

// AddEnum adds an enum value to a customer property
func (s *CustomerPropertyService) AddEnum(ctx context.Context, req *models.AddEnumCustomerPropertyRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customer_properties/add_enum", req, headers)
}

// RenameEnum renames an enum value of a customer property
func (s *CustomerPropertyService) RenameEnum(ctx context.Context, req *models.RenameEnumCustomerPropertyRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customer_properties/rename_enum", req, headers)
}

// DeleteEnum deletes an enum value from a customer property
func (s *CustomerPropertyService) DeleteEnum(ctx context.Context, req *models.DeleteEnumCustomerPropertyRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customer_properties/delete_enum", req, headers)
}

// RestoreEnum restores a deleted enum value of a customer property
func (s *CustomerPropertyService) RestoreEnum(ctx context.Context, req *models.RestoreEnumCustomerPropertyRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customer_properties/restore_enum", req, headers)
}

// Batch executes multiple customer property commands in a batch
func (s *CustomerPropertyService) Batch(ctx context.Context, req *models.CustomerPropertyBatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customer_properties/batch", req, headers)
}

// GetByID retrieves a customer property by ID
func (s *CustomerPropertyService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.PropertyDefinitionDto, error) {
	var response models.PropertyDefinitionDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/customer_properties/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of customer properties
func (s *CustomerPropertyService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.PropertyDefinitionListDto, error) {
	var response models.PropertyDefinitionListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/customer_properties/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
