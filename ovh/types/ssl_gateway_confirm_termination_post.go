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

// SSLGatewayConfirmTerminationPost ...
type SSLGatewayConfirmTerminationPost struct {
	Commentary string `json:"commentary,omitempty"`

	Reason string `json:"reason,omitempty"`

	Token string `json:"token,omitempty"`
}
