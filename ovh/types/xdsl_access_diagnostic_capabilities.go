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

// XdslAccessDiagnosticCapabilities Describe the capabilities of the access diagnostic
type XdslAccessDiagnosticCapabilities struct {
	Incident bool `json:"incident,omitempty"`

	IsActiveOnLns bool `json:"isActiveOnLns,omitempty"`

	IsModemConnected bool `json:"isModemConnected,omitempty"`

	LineTest bool `json:"lineTest,omitempty"`

	Ping bool `json:"ping,omitempty"`

	ProposedProfileID bool `json:"proposedProfileId,omitempty"`

	Sync bool `json:"sync,omitempty"`
}