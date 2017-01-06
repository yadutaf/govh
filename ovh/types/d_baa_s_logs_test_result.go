/*
 * OVH API - EU
 *
 * Build your own OVH world.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: api-subscribe@ml.ovh.net
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package types

import (
	"time"
)

// DBaaSLogsTestResult Config test results
type DBaaSLogsTestResult struct {

	// Stderr Standard error
	Stderr string `json:"stderr,omitempty"`

	// Stdout Standard output
	Stdout string `json:"stdout,omitempty"`

	// UpdatedAt Last config test update
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}