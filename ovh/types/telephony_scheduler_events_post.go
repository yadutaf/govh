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

// TelephonySchedulerEventsPost ...
type TelephonySchedulerEventsPost struct {
	Category string `json:"category,omitempty"`

	DateEnd *time.Time `json:"dateEnd,omitempty"`

	DateStart *time.Time `json:"dateStart,omitempty"`

	Description string `json:"description,omitempty"`

	Title string `json:"title,omitempty"`

	UID string `json:"uid,omitempty"`
}
