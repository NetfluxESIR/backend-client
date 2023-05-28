package internal

import (
	"context"
	"fmt"
	"github.com/NetfluxESIR/backend/pkg/api/gen"
	"github.com/google/uuid"
)

func (c *Client) SetSteps(ctx context.Context, videoId string, step string, status string, logs string) error {
	id, err := uuid.Parse(videoId)
	if err != nil {
		return err
	}
	s := gen.UpdateProcessingStepJSONBodyStatus(status)
	response, err := c.apiClient.UpdateProcessingStepWithResponse(ctx, id, gen.UpdateProcessingStepJSONRequestBody{
		Step:   gen.UpdateProcessingStepJSONBodyStep(step),
		Status: &s,
		Logs:   func() *string { return &logs }(),
	})
	if err != nil {
		return err
	}
	if response.StatusCode() != 200 {
		return fmt.Errorf("unexpected status code: %s", string(response.Body))
	}
	return nil
}
