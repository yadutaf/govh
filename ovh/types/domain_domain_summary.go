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

// DomainSummary Values of number account (account, mailing list, redirection and responder)
type DomainSummary struct {

	// Account Number of mailboxes
	Account int64 `json:"account,omitempty"`

	// MailingList Number of mailing lists
	MailingList int64 `json:"mailingList,omitempty"`

	// Redirection Number of redirections
	Redirection int64 `json:"redirection,omitempty"`

	// Responder Number of responders
	Responder int64 `json:"responder,omitempty"`
}
