// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// GiftCardService handles gift card operations
type GiftCardService struct {
	client *Client
}

// Create creates a new gift card
func (s *GiftCardService) Create(ctx context.Context, req *models.CreateGiftCardRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_cards/create", req, headers)
}

// SetNumber sets the number of a gift card
func (s *GiftCardService) SetNumber(ctx context.Context, req *models.SetCardNumberRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_cards/set_number", req, headers)
}

// SetBarcode sets the barcode of a gift card
func (s *GiftCardService) SetBarcode(ctx context.Context, req *models.SetCardBarcodeRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_cards/set_barcode", req, headers)
}

// SetPinCode sets the pin code of a gift card
func (s *GiftCardService) SetPinCode(ctx context.Context, req *models.SetCardPinCodeRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_cards/set_pin_code", req, headers)
}

// SetExpirationDate sets the expiration date of a gift card
func (s *GiftCardService) SetExpirationDate(ctx context.Context, req *models.SetCardExpirationDateRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_cards/set_expiration_date", req, headers)
}

// SetValidityPeriod sets the validity period of a gift card
func (s *GiftCardService) SetValidityPeriod(ctx context.Context, req *models.SetGiftCardValidityPeriodRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_cards/set_validity_period", req, headers)
}

// ConfirmReadinessForIssuance confirms readiness for issuance
func (s *GiftCardService) ConfirmReadinessForIssuance(ctx context.Context, req *models.ConfirmReadinessForIssuanceRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_cards/confirm_readiness_for_issuance", req, headers)
}

// Issue issues a gift card
func (s *GiftCardService) Issue(ctx context.Context, req *models.IssueGiftCardRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_cards/issue", req, headers)
}

// CancelIssue cancels the issue of a gift card
func (s *GiftCardService) CancelIssue(ctx context.Context, req *models.CancelIssueRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_cards/cancel_issue", req, headers)
}

// Block blocks a gift card
func (s *GiftCardService) Block(ctx context.Context, req *models.BlockGiftCardRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_cards/block", req, headers)
}

// Unblock unblocks a gift card
func (s *GiftCardService) Unblock(ctx context.Context, req *models.UnblockCardRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_cards/unblock", req, headers)
}

// Subtract subtracts amount from a gift card
func (s *GiftCardService) Subtract(ctx context.Context, req *models.SubtractGiftCardRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_cards/subtract", req, headers)
}

// Refund refunds amount to a gift card
func (s *GiftCardService) Refund(ctx context.Context, req *models.RefundGiftCardRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_cards/refund", req, headers)
}

// AddGiftCardGroup adds a gift card to a group
func (s *GiftCardService) AddGiftCardGroup(ctx context.Context, req *models.AddCardGroupRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_cards/add_gift_card_group", req, headers)
}

// RemoveGiftCardGroup removes a gift card from a group
func (s *GiftCardService) RemoveGiftCardGroup(ctx context.Context, req *models.RemoveCardGroupRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_cards/remove_gift_card_group", req, headers)
}

// Batch executes multiple commands in a batch
func (s *GiftCardService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/gift_cards/batch", req, headers)
}

// GetByID retrieves a gift card by ID
func (s *GiftCardService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.GiftCardDto, error) {
	var response models.GiftCardDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/gift_cards/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of gift cards
func (s *GiftCardService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.GiftCardListDto, error) {
	var response models.GiftCardListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/gift_cards/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
