// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/ramzes4rules/rsl6-integration-api/models"
)

// OpeningHoursService handles opening hours operations
type OpeningHoursService struct {
	client *Client
}

// Create creates a new opening hours entry
func (s *OpeningHoursService) Create(ctx context.Context, req *models.CreateOpeningHoursRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/opening_hours/create", req, headers)
}

// Rename renames an opening hours entry
func (s *OpeningHoursService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/opening_hours/rename", req, headers)
}

// SetWorkPeriod sets the work period of an opening hours entry
func (s *OpeningHoursService) SetWorkPeriod(ctx context.Context, req *models.SetWorkPeriodRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/opening_hours/set_work_period", req, headers)
}

// Delete deletes an opening hours entry
func (s *OpeningHoursService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/opening_hours/delete", req, headers)
}

// Restore restores a deleted opening hours entry
func (s *OpeningHoursService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/opening_hours/restore", req, headers)
}

// MoveToArchive moves an opening hours entry to archive
func (s *OpeningHoursService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/opening_hours/move_to_archive", req, headers)
}

// Unarchive restores an opening hours entry from archive
func (s *OpeningHoursService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/opening_hours/unarchive", req, headers)
}

// Batch executes multiple commands in a batch
func (s *OpeningHoursService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/opening_hours/batch", req, headers)
}

// GetByID retrieves an opening hours entry by ID
func (s *OpeningHoursService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.OpeningHoursDto, error) {
	var response models.OpeningHoursDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/opening_hours/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of opening hours entries
func (s *OpeningHoursService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.OpeningHoursListDto, error) {
	var response models.OpeningHoursListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/opening_hours/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
