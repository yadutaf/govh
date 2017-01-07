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

// DomainZoneTask Tasks associated to a zone
type DomainZoneTask struct {

	// CanAccelerate Can accelerate the task
	CanAccelerate bool `json:"canAccelerate,omitempty"`

	// CanCancel Can cancel the task
	CanCancel bool `json:"canCancel,omitempty"`

	// CanRelaunch Can relaunch the task
	CanRelaunch bool `json:"canRelaunch,omitempty"`

	// Comment Comment about the task
	Comment string `json:"comment,omitempty"`

	// CreationDate Creation date of the task
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// DoneDate Done date of the task
	DoneDate *time.Time `json:"doneDate,omitempty"`

	// Function Function of the task
	Function string `json:"function,omitempty"`

	// ID Id of the task
	ID int64 `json:"id,omitempty"`

	// LastUpdate Last update date of the task
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`

	// Status Status of the task
	Status string `json:"status,omitempty"`

	// TodoDate Todo date of the task
	TodoDate *time.Time `json:"todoDate,omitempty"`
}
