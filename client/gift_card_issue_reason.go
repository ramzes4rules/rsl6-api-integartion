// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/ramzes4rules/rsl6-integration-api/models"
)

// GiftCardIssueReasonService handles gift card issue reason operations
type GiftCardIssueReasonService struct {
	client *Client
}

// Create creates a new gift card issue reason
func (s *GiftCardIssueReasonService) Create(ctx context.Context, req *models.CreateReasonRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_card_issue_reasons/create", req, headers)
}

// Rename renames a gift card issue reason
func (s *GiftCardIssueReasonService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_card_issue_reasons/rename", req, headers)
}

// Delete deletes a gift card issue reason
func (s *GiftCardIssueReasonService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_card_issue_reasons/delete", req, headers)
}

// Restore restores a deleted gift card issue reason
func (s *GiftCardIssueReasonService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_card_issue_reasons/restore", req, headers)
}

// MoveToArchive moves a gift card issue reason to archive
func (s *GiftCardIssueReasonService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_card_issue_reasons/move_to_archive", req, headers)
}

// Unarchive restores a gift card issue reason from archive
func (s *GiftCardIssueReasonService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_card_issue_reasons/unarchive", req, headers)
}

// Batch executes multiple commands in a batch
func (s *GiftCardIssueReasonService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_card_issue_reasons/batch", req, headers)
}

// GetByID retrieves a gift card issue reason by ID
func (s *GiftCardIssueReasonService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.ReasonDto, error) {
	var response models.ReasonDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/gift_card_issue_reasons/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of gift card issue reasons
func (s *GiftCardIssueReasonService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.ReasonListDto, error) {
	var response models.ReasonListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/gift_card_issue_reasons/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
