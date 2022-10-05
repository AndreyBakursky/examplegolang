package example_for_review

import (
	"context"
	"fmt"
	"time"
)

type Client struct {
	api     *client.SchoolGoExternal
	timeout time.Duration
	metrics *school_go_external.Metrics
}

func NewSchoolGoExternalClient(
	api *client.SchoolGoExternal,
	timeout time.Duration,
	metrics *school_go_external.Metrics,
) *Client {
	return &Client{
		api:     api,
		timeout: timeout,
		metrics: metrics,
	}
}

func (c *Client) SendOrder(ctx context.Context, request models.Order, orderId int64) (*models.Success, error) {
	response, err := interaction.
		CreateCall(ctx, serviceName, c.timeout).
		WithMetrics(c.metrics.GetSendOrder()).
		Invoke(
			gen.NewSendOrderInteractionRequest(
				c.api.Operations,
				operations.NewSendOrderParamsWithContext(ctx).
					WithBody(&request).WithID(orderId),
			),
		)

	if err != nil {
		return nil, fmt.Errorf("can`t send request to breadwinners: %w", err)
	}

	if response, ok := response.(*operations.SendOrderOK); ok {
		return response.Payload, nil
	}

	return nil, fmt.Errorf("invalid response model in breadwinners ChainSearch method")
}
