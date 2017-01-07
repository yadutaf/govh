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

// OrderHostingWebNewPost ...
type OrderHostingWebNewPost struct {
	DNSZone string `json:"dnsZone,omitempty"`

	Domain string `json:"domain,omitempty"`

	Module string `json:"module,omitempty"`

	Offer string `json:"offer,omitempty"`
}
