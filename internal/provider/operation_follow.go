// Copyright 2026 Forcepoint LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code not auto-generated
package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
)

// OperationStatus maps to the SMC API Operation_Status schema returned by
// async operations (upload, refresh) as a 202 response body and on each
// GET to the follower URL.
type OperationStatus struct {
	InProgress  bool    `json:"in_progress"`
	Success     bool    `json:"success"`
	Progress    float64 `json:"progress"`
	LastMessage string  `json:"last_message"`
	Follower    string  `json:"follower"`
	ErrorCode   int     `json:"error_code"`
}

// renderProgressBar returns a fixed-width ASCII progress bar string.
// When color is true, the bar is wrapped in ANSI green (requires a color-capable terminal).
func renderProgressBar(progress float64, message string, color bool) string {
	const barWidth = 20
	filled := int(progress / 100.0 * float64(barWidth))
	if filled > barWidth {
		filled = barWidth
	}
	bar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)
	if color {
		return fmt.Sprintf("\033[32m[%s]\033[0m %.0f%% - %s", bar, progress, message)
	}
	return fmt.Sprintf("[%s] %.0f%% - %s", bar, progress, message)
}

func parseOperationStatus(body []byte) (OperationStatus, error) {
	var status OperationStatus
	if err := json.Unmarshal(body, &status); err != nil {
		return status, fmt.Errorf("failed to parse operation status: %w", err)
	}
	return status, nil
}

// followOperation polls the follower URL every 3 seconds until the operation
// completes or the timeout elapses. progressFn is called on each poll tick
// with the current progress percentage and last message.
func followOperation(
	ctx context.Context,
	client *smc.SmcClient,
	status OperationStatus,
	timeout time.Duration,
	progressFn func(progress float64, message string),
) error {
	if !status.InProgress {
		if !status.Success {
			return fmt.Errorf("operation failed (error_code=%d): %s", status.ErrorCode, status.LastMessage)
		}
		return nil
	}

	if status.Follower == "" {
		return fmt.Errorf("operation is in progress but no follower URL was provided")
	}

	timer := time.NewTimer(timeout)
	ticker := time.NewTicker(3 * time.Second)
	defer timer.Stop()
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context cancelled while waiting for operation: %w", ctx.Err())
		case <-timer.C:
			return fmt.Errorf("operation timed out after %v, last message: %s", timeout, status.LastMessage)
		case <-ticker.C:
			respData, err := client.DoRequest(smc.Options{
				Method:  "GET",
				URL:     status.Follower,
				Context: ctx,
			})
			if err != nil {
				return fmt.Errorf("error polling operation status: %w", err)
			}
			if respData.StatusCode != 200 {
				return fmt.Errorf("unexpected status %d polling operation follower: %s", respData.StatusCode, string(respData.Body))
			}
			status, err = parseOperationStatus(respData.Body)
			if err != nil {
				return err
			}
			progressFn(status.Progress, status.LastMessage)
			if !status.InProgress {
				if !status.Success {
					return fmt.Errorf("operation failed (error_code=%d): %s", status.ErrorCode, status.LastMessage)
				}
				return nil
			}
		}
	}
}
