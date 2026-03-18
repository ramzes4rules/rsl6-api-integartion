// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/ramzes4rules/rsl6-integration-api/models"
)

// ItemService handles item operations
type ItemService struct {
	client *Client
}

// Create creates a new item
func (s *ItemService) Create(ctx context.Context, req *models.CreateItemRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/create", req, headers)
}

// Rename renames an item
func (s *ItemService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/rename", req, headers)
}

// SetDescription sets the description of an item
func (s *ItemService) SetDescription(ctx context.Context, req *models.SetDescriptionRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/set_description", req, headers)
}

// SetPublicName sets the public name of an item
func (s *ItemService) SetPublicName(ctx context.Context, req *models.SetPublicNameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/set_public_name", req, headers)
}

// SetArticle sets the article of an item
func (s *ItemService) SetArticle(ctx context.Context, req *models.SetArticleRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/set_article", req, headers)
}

// SetExternalID sets the external ID of an item
func (s *ItemService) SetExternalID(ctx context.Context, req *models.SetExternalIDRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/set_external_id", req, headers)
}

// SetItemGroupID sets the item group ID
func (s *ItemService) SetItemGroupID(ctx context.Context, req *models.SetItemGroupIDRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/set_item_group_id", req, headers)
}

// SetRestriction sets the restriction of an item
func (s *ItemService) SetRestriction(ctx context.Context, req *models.SetRestrictionRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/set_restriction", req, headers)
}

// SetBooleanPropertyValue sets a boolean property value
func (s *ItemService) SetBooleanPropertyValue(ctx context.Context, req *models.SetBooleanPropertyValueRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/set_boolean_property_value", req, headers)
}

// SetIntPropertyValue sets an integer property value
func (s *ItemService) SetIntPropertyValue(ctx context.Context, req *models.SetIntPropertyValueRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/set_int_property_value", req, headers)
}

// SetStringPropertyValue sets a string property value
func (s *ItemService) SetStringPropertyValue(ctx context.Context, req *models.SetStringPropertyValueRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/set_string_property_value", req, headers)
}

// SetDatePropertyValue sets a date property value
func (s *ItemService) SetDatePropertyValue(ctx context.Context, req *models.SetDatePropertyValueRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/set_date_property_value", req, headers)
}

// SetEnumPropertyValue sets an enum property value
func (s *ItemService) SetEnumPropertyValue(ctx context.Context, req *models.SetEnumPropertyValueRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/set_enum_property_value", req, headers)
}

// SetEnumsPropertyValue sets multiple enum property values
func (s *ItemService) SetEnumsPropertyValue(ctx context.Context, req *models.SetEnumsPropertyValueRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/set_enums_property_value", req, headers)
}

// AddMediaContent adds media content to an item
func (s *ItemService) AddMediaContent(ctx context.Context, req *models.AddMediaContentRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/add_media_content", req, headers)
}

// RemoveMediaContent removes media content from an item
func (s *ItemService) RemoveMediaContent(ctx context.Context, req *models.RemoveMediaContentRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/remove_media_content", req, headers)
}

// RenameMediaContent renames media content
func (s *ItemService) RenameMediaContent(ctx context.Context, req *models.RenameMediaContentRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/rename_media_content", req, headers)
}

// SetMediaContentSource sets the source of media content
func (s *ItemService) SetMediaContentSource(ctx context.Context, req *models.SetMediaContentSourceRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/set_media_content_source", req, headers)
}

// SetMediaContentType sets the type of media content
func (s *ItemService) SetMediaContentType(ctx context.Context, req *models.SetMediaContentTypeRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/set_media_content_type", req, headers)
}

// SetDefaultImage sets the default image of an item
func (s *ItemService) SetDefaultImage(ctx context.Context, req *models.SetDefaultImageRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/set_default_image", req, headers)
}

// ClearDefaultImage clears the default image of an item
func (s *ItemService) ClearDefaultImage(ctx context.Context, req *models.ClearDefaultImageRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/clear_default_image", req, headers)
}

// AddSaleItem adds a sale item
func (s *ItemService) AddSaleItem(ctx context.Context, req *models.AddSaleItemRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/add_sale_item", req, headers)
}

// DeleteSaleItem deletes a sale item
func (s *ItemService) DeleteSaleItem(ctx context.Context, req *models.DeleteSaleItemRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/delete_sale_item", req, headers)
}

// RestoreSaleItem restores a sale item
func (s *ItemService) RestoreSaleItem(ctx context.Context, req *models.RestoreSaleItemRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/restore_sale_item", req, headers)
}

// SetSaleItemExternalID sets the external ID of a sale item
func (s *ItemService) SetSaleItemExternalID(ctx context.Context, req *models.SetSaleItemExternalIDRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/set_sale_item_external_id", req, headers)
}

// SetSaleItemPropertyValue sets the property value of a sale item
func (s *ItemService) SetSaleItemPropertyValue(ctx context.Context, req *models.SetSaleItemPropertyValueRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/set_sale_item_property_value", req, headers)
}

// Delete deletes an item
func (s *ItemService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/delete", req, headers)
}

// Restore restores a deleted item
func (s *ItemService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/restore", req, headers)
}

// Batch executes multiple commands in a batch
func (s *ItemService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/items/batch", req, headers)
}

// GetByID retrieves an item by ID
func (s *ItemService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.ItemDto, error) {
	var response models.ItemDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/items/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of items
func (s *ItemService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.ItemListDto, error) {
	var response models.ItemListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/items/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
