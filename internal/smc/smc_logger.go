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

// filepath: /data/DEV-TERRAFORM-PROVIDER/terraform-provider-smc/internal/smc/log_context.go
package smc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	// SubsystemName is the name of the SMC API subsystem for logging
	SubsystemName = "SMC_API"
)

// LogContext holds the logging context for SMC session management operations.
// This code does not depend on terraform so, logging with terraform
// is encapsulated in this class. It allows seeing any connection issue within the ft.log file.
type LogContext struct {
	ctx         context.Context
	mgtServerId string
}

// NewLogContext creates a new LogContext with the smc subsystem
func NewLogContext(ctx context.Context, mgtServerId string) *LogContext {

	// Create a context with the smc subsystem
	logCtx := tflog.NewSubsystem(ctx, SubsystemName)
	// WARNING: when not using context.background(), the context referenced bay be canceled and generate a terraform panic

	// Add default fields
	if mgtServerId != "" {
		logCtx = tflog.SetField(logCtx, "mgt_server_id", mgtServerId)
	}

	return &LogContext{
		ctx:         logCtx,
		mgtServerId: mgtServerId,
	}
}

// Context returns the underlying context
func (lc *LogContext) Context() context.Context {
	return lc.ctx
}

// WithField adds a field to the log context and returns a new context
func (lc *LogContext) WithField(key string, value interface{}) context.Context {
	return tflog.SetField(lc.ctx, key, value)
}

// WithFields adds multiple fields to the log context and returns a new context
func (lc *LogContext) WithFields(fields map[string]interface{}) context.Context {
	ctx := lc.ctx
	for k, v := range fields {
		ctx = tflog.SetField(ctx, k, v)
	}
	return ctx
}

// Debug logs a debug message
func (lc *LogContext) Debug(msg string, additionalFields ...map[string]interface{}) {
	ctx := lc.ctx
	if len(additionalFields) > 0 {
		ctx = lc.WithFields(additionalFields[0])
	}
	tflog.SubsystemDebug(ctx, SubsystemName, msg)
}

// Info logs an info message
func (lc *LogContext) Info(msg string, additionalFields ...map[string]interface{}) {
	if lc != nil {
		ctx := lc.ctx
		if len(additionalFields) > 0 {
			ctx = lc.WithFields(additionalFields[0])
		}
		tflog.SubsystemInfo(ctx, SubsystemName, msg)
	}
}

// Warn logs a warning message
func (lc *LogContext) Warn(msg string, additionalFields ...map[string]interface{}) {
	if lc != nil {
		ctx := lc.ctx
		if len(additionalFields) > 0 {
			ctx = lc.WithFields(additionalFields[0])
		}
		tflog.SubsystemWarn(ctx, SubsystemName, msg)
	}
}

// Error logs an error message
func (lc *LogContext) Error(msg string, additionalFields ...map[string]interface{}) {
	if lc != nil {
		ctx := lc.ctx
		if len(additionalFields) > 0 {
			ctx = lc.WithFields(additionalFields[0])
		}
		tflog.SubsystemError(ctx, SubsystemName, msg)
	}
}

// Trace logs a trace message
func (lc *LogContext) Trace(msg string, additionalFields ...map[string]interface{}) {
	if lc != nil {
		ctx := lc.ctx
		if len(additionalFields) > 0 {
			ctx = lc.WithFields(additionalFields[0])
		}
		tflog.SubsystemTrace(ctx, SubsystemName, msg)
	}
}
