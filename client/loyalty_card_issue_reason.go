// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// LoyaltyCardIssueReasonService handles loyalty card issue reason operations
type LoyaltyCardIssueReasonService struct {
	client *Client
}

// Create creates a new loyalty card issue reason
func (s *LoyaltyCardIssueReasonService) Create(ctx context.Context, req *models.CreateReasonRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_issue_reasons/create", req, headers)
}

// Rename renames a loyalty card issue reason
func (s *LoyaltyCardIssueReasonService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_issue_reasons/rename", req, headers)
}

// Delete deletes a loyalty card issue reason
func (s *LoyaltyCardIssueReasonService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_issue_reasons/delete", req, headers)
}

// Restore restores a deleted loyalty card issue reason
func (s *LoyaltyCardIssueReasonService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_issue_reasons/restore", req, headers)
}

// MoveToArchive moves a loyalty card issue reason to archive
func (s *LoyaltyCardIssueReasonService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_issue_reasons/move_to_archive", req, headers)
}

// Unarchive restores a loyalty card issue reason from archive
func (s *LoyaltyCardIssueReasonService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_issue_reasons/unarchive", req, headers)
}

// Batch executes multiple commands in a batch
func (s *LoyaltyCardIssueReasonService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/loyalty_card_issue_reasons/batch", req, headers)
}

// GetByID retrieves a loyalty card issue reason by ID
func (s *LoyaltyCardIssueReasonService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.ReasonDto, error) {
	var response models.ReasonDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/loyalty_card_issue_reasons/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of loyalty card issue reasons
func (s *LoyaltyCardIssueReasonService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.ReasonListDto, error) {
	var response models.ReasonListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/loyalty_card_issue_reasons/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
