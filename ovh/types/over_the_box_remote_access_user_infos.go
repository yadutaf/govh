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

// OverTheBoxRemoteAccessUserInfos Infos about the remote user
type OverTheBoxRemoteAccessUserInfos struct {

	// IP IP from which the remote access will be allowed
	IP string `json:"ip,omitempty"`

	// PublicKey The public key authorized on the device (for SSH purpose)
	PublicKey string `json:"publicKey,omitempty"`

	// User The user that will access the device remotely
	User string `json:"user,omitempty"`
}
