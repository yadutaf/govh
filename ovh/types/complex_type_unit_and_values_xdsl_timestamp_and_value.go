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

// ComplexTypeUnitAndValuesXdslTimestampAndValue A value set tagged with its unit
type ComplexTypeUnitAndValuesXdslTimestampAndValue struct {
	Unit string `json:"unit,omitempty"`

	Values []*XdslTimestampAndValue `json:"values,omitempty"`
}