// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// PosTypeService handles POS type operations
type PosTypeService struct {
	client *Client
}

// Create creates a new POS type
func (s *PosTypeService) Create(ctx context.Context, req *models.CreatePosTypeRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/pos_types/create", req, headers)
}

// Rename renames a POS type
func (s *PosTypeService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/pos_types/rename", req, headers)
}

// Delete deletes a POS type
func (s *PosTypeService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/pos_types/delete", req, headers)
}

// Restore restores a deleted POS type
func (s *PosTypeService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/pos_types/restore", req, headers)
}

// MoveToArchive moves a POS type to archive
func (s *PosTypeService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/pos_types/move_to_archive", req, headers)
}

// Unarchive restores a POS type from archive
func (s *PosTypeService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/pos_types/unarchive", req, headers)
}

// Batch executes multiple commands in a batch
func (s *PosTypeService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/pos_types/batch", req, headers)
}

// GetByID retrieves a POS type by ID
func (s *PosTypeService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.PosTypeDto, error) {
	var response models.PosTypeDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/pos_types/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of POS types
func (s *PosTypeService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.PosTypeListDto, error) {
	var response models.PosTypeListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/pos_types/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
