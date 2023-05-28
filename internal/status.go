package internal

import (
	"context"
	"fmt"
	"github.com/NetfluxESIR/backend/pkg/api/gen"
	"github.com/google/uuid"
	"net/http"
)

func (c *Client) SetStatus(ctx context.Context, videoId string, status string) error {
	id, err := uuid.Parse(videoId)
	if err != nil {
		return err
	}
	response, err := c.apiClient.UpdateProcessingStatusWithResponse(ctx, id, gen.UpdateProcessingStatusJSONRequestBody{
		Status: gen.UpdateProcessingStatusJSONBodyStatus(status),
	}, func(ctx context.Context, req *http.Request) error {
		// Add the token to the request
		req.Header.Set("Authorization", "Bearer "+c.token)
		return nil
	})
	if err != nil {
		return err
	}
	if response.StatusCode() != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", string(response.Body))
	}
	return nil
}
