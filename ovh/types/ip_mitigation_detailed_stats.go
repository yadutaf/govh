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

// IPMitigationDetailedStats Detailed statistics about your traffic
type IPMitigationDetailedStats struct {

	// DestPort Traffic dest port
	DestPort int64 `json:"destPort,omitempty"`

	Fragments bool `json:"fragments,omitempty"`

	// IcmpCode ICMP protocol code
	IcmpCode int64 `json:"icmpCode,omitempty"`

	// IcmpType ICMP protocol type
	IcmpType int64 `json:"icmpType,omitempty"`

	In *IPMitigationTraffic `json:"in,omitempty"`

	Out *IPMitigationTraffic `json:"out,omitempty"`

	// Protocol Used protocol. See RFC5237 for more information
	Protocol int64 `json:"protocol,omitempty"`

	// SrcPort Traffic source port
	SrcPort int64 `json:"srcPort,omitempty"`

	Syn bool `json:"syn,omitempty"`
}
