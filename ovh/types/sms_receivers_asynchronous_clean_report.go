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

// SmsReceiversAsynchronousCleanReport A structure giving operation price and asynchronous task ID
type SmsReceiversAsynchronousCleanReport struct {
	TaskID int64 `json:"taskId,omitempty"`

	TotalCreditsRemoved float64 `json:"totalCreditsRemoved,omitempty"`
}