// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// PosService handles POS (point of sale) operations
type PosService struct {
	client *Client
}

// Create creates a new POS
func (s *PosService) Create(ctx context.Context, req *models.CreatePosRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/poses/create", req, headers)
}

// Rename renames a POS
func (s *PosService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/poses/rename", req, headers)
}

// SetPosType sets the type of a POS
func (s *PosService) SetPosType(ctx context.Context, req *models.SetPosTypeRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/poses/set_pos_type", req, headers)
}

// SetAuthorization sets the authorization of a POS
func (s *PosService) SetAuthorization(ctx context.Context, req *models.SetPosAuthorizationRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/poses/set_authorization", req, headers)
}

// MoveToStore moves a POS to a store
func (s *PosService) MoveToStore(ctx context.Context, req *models.MoveToStoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/poses/move_to_store", req, headers)
}

// Block blocks a POS
func (s *PosService) Block(ctx context.Context, req *models.BlockPosRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/poses/block", req, headers)
}

// Unblock unblocks a POS
func (s *PosService) Unblock(ctx context.Context, req *models.UnblockPosRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/poses/unblock", req, headers)
}

// Delete deletes a POS
func (s *PosService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/poses/delete", req, headers)
}

// Restore restores a deleted POS
func (s *PosService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/poses/restore", req, headers)
}

// MoveToArchive moves a POS to archive
func (s *PosService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/poses/move_to_archive", req, headers)
}

// Unarchive restores a POS from archive
func (s *PosService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/poses/unarchive", req, headers)
}

// Batch executes multiple commands in a batch
func (s *PosService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/poses/batch", req, headers)
}

// GetByID retrieves a POS by ID
func (s *PosService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.PosDto, error) {
	var response models.PosDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/poses/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of POSes
func (s *PosService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.PosListDto, error) {
	var response models.PosListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/poses/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
