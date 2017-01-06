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

// EmailExchangeExternalContact External contact for this exchange service
type EmailExchangeExternalContact struct {

	// CreationDate Creation date
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// DisplayName Contact display name
	DisplayName string `json:"displayName,omitempty"`

	// ExternalEmailAddress Contact email
	ExternalEmailAddress string `json:"externalEmailAddress,omitempty"`

	// FirstName Contact first name
	FirstName string `json:"firstName,omitempty"`

	// HiddenFromGAL Hide the contact in Global Address List
	HiddenFromGAL bool `json:"hiddenFromGAL,omitempty"`

	// ID Contact id
	ID int64 `json:"id,omitempty"`

	// Initials Contact initals
	Initials string `json:"initials,omitempty"`

	// LastName Contact last name
	LastName string `json:"lastName,omitempty"`

	// Organization2010 If specified, indicates to which organization this external contact belongs (Exchange 2010 only)
	Organization2010 string `json:"organization2010,omitempty"`

	// State Contact state
	State string `json:"state,omitempty"`

	// TaskPendingID Task pending id
	TaskPendingID int64 `json:"taskPendingId,omitempty"`
}