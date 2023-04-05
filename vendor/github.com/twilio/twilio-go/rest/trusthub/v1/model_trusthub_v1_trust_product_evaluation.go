/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Trusthub
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// TrusthubV1TrustProductEvaluation struct for TrusthubV1TrustProductEvaluation
type TrusthubV1TrustProductEvaluation struct {
	// The unique string that identifies the Evaluation resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the trust_product resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique string of a policy that is associated to the trust_product resource.
	PolicySid *string `json:"policy_sid,omitempty"`
	// The unique string that we created to identify the trust_product resource.
	TrustProductSid *string `json:"trust_product_sid,omitempty"`
	Status          *string `json:"status,omitempty"`
	// The results of the Evaluation which includes the valid and invalid attributes.
	Results     *[]interface{} `json:"results,omitempty"`
	DateCreated *time.Time     `json:"date_created,omitempty"`
	Url         *string        `json:"url,omitempty"`
}
