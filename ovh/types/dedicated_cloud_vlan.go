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

// DedicatedCloudVlan Dedicated Cloud Vlan
type DedicatedCloudVlan struct {
	Name string `json:"name,omitempty"`

	// RoutingRateLimit Speed in Mbps
	RoutingRateLimit string `json:"routingRateLimit,omitempty"`

	State string `json:"state,omitempty"`

	TType string `json:"type,omitempty"`

	VlanID int64 `json:"vlanId,omitempty"`

	VlanNumber int64 `json:"vlanNumber,omitempty"`
}