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

// HostingWebCron Hosting crons
type HostingWebCron struct {

	// Command Command to execute
	Command string `json:"command,omitempty"`

	// Description Description field for you
	Description string `json:"description,omitempty"`

	// Email Email used to receive error log ( stderr )
	Email string `json:"email,omitempty"`

	// Frequency Frequency ( crontab format ) defined for the script ( minutes are ignored )
	Frequency string `json:"frequency,omitempty"`

	// ID Cron's id
	ID int64 `json:"id,omitempty"`

	// Language Cron language
	Language string `json:"language,omitempty"`

	// Status Cron status
	Status string `json:"status,omitempty"`
}
