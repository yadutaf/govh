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

// NichandleAccessRestrictionSotpAccount SOTP Two-Factor Authentication
type NichandleAccessRestrictionSotpAccount struct {

	// Remaining Number of remaining codes
	Remaining int64 `json:"remaining,omitempty"`

	// Status Status of this account
	Status string `json:"status,omitempty"`
}
