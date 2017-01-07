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

// SSLGatewayTask SSL Gateway tasks
type SSLGatewayTask struct {

	// Action The action made
	Action string `json:"action,omitempty"`

	// CreationDate Creation date of your task
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// ID Id of the task
	ID int64 `json:"id,omitempty"`

	// Progress Task progress percentage
	Progress int64 `json:"progress,omitempty"`

	// Status Current status of your task
	Status string `json:"status,omitempty"`
}
