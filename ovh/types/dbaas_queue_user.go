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

// DBaasQueueUser User
type DBaasQueueUser struct {

	// ID User ID
	ID string `json:"id,omitempty"`

	// Name User name
	Name string `json:"name,omitempty"`

	// Roles List of roles this user belongs to
	Roles []string `json:"roles,omitempty"`
}
