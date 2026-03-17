// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"

	"github.com/rsl6/rsloyalty/models"
)

// CustomerSegmentService handles customer segment operations
type CustomerSegmentService struct {
	client *Client
}

// AddToStaticSegment adds a customer to a static segment
func (s *CustomerSegmentService) AddToStaticSegment(ctx context.Context, req *models.AddToStaticSegmentRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customer_segments/add_to_static_segment", req, headers)
}

// RemoveFromStaticSegment removes a customer from a static segment
func (s *CustomerSegmentService) RemoveFromStaticSegment(ctx context.Context, req *models.RemoveFromStaticSegmentRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customer_segments/remove_from_static_segment", req, headers)
}

// Batch executes multiple customer segment commands in a batch
func (s *CustomerSegmentService) Batch(ctx context.Context, req *models.CustomerSegmentBatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customer_segments/batch", req, headers)
}
