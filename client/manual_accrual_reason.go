// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// ManualAccrualReasonService handles manual accrual reason operations
type ManualAccrualReasonService struct {
	client *Client
}

// Create creates a new manual accrual reason
func (s *ManualAccrualReasonService) Create(ctx context.Context, req *models.CreateReasonRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/manual_accrual_reasons/create", req, headers)
}

// Rename renames a manual accrual reason
func (s *ManualAccrualReasonService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/manual_accrual_reasons/rename", req, headers)
}

// Delete deletes a manual accrual reason
func (s *ManualAccrualReasonService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/manual_accrual_reasons/delete", req, headers)
}

// Restore restores a deleted manual accrual reason
func (s *ManualAccrualReasonService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/manual_accrual_reasons/restore", req, headers)
}

// MoveToArchive moves a manual accrual reason to archive
func (s *ManualAccrualReasonService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/manual_accrual_reasons/move_to_archive", req, headers)
}

// Unarchive restores a manual accrual reason from archive
func (s *ManualAccrualReasonService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/manual_accrual_reasons/unarchive", req, headers)
}

// Batch executes multiple commands in a batch
func (s *ManualAccrualReasonService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/manual_accrual_reasons/batch", req, headers)
}

// GetByID retrieves a manual accrual reason by ID
func (s *ManualAccrualReasonService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.ReasonDto, error) {
	var response models.ReasonDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/manual_accrual_reasons/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of manual accrual reasons
func (s *ManualAccrualReasonService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.ReasonListDto, error) {
	var response models.ReasonListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/manual_accrual_reasons/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
