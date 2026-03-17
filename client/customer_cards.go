// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"

	"github.com/rsl6/rsloyalty/models"
)

// CustomerCardsService handles customer cards operations
type CustomerCardsService struct {
	client *Client
}

// BindLoyaltyCard binds a loyalty card to a customer
func (s *CustomerCardsService) BindLoyaltyCard(ctx context.Context, req *models.BindLoyaltyCardRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customer_cards/bind", req, headers)
}

// UnbindLoyaltyCard unbinds a loyalty card from a customer
func (s *CustomerCardsService) UnbindLoyaltyCard(ctx context.Context, req *models.UnbindLoyaltyCardRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customer_cards/unbind", req, headers)
}

// Batch executes multiple customer cards commands in a batch
func (s *CustomerCardsService) Batch(ctx context.Context, req *models.CustomerCardsBatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customer_cards/batch", req, headers)
}
