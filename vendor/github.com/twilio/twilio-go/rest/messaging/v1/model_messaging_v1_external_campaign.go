/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Messaging
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

// MessagingV1ExternalCampaign struct for MessagingV1ExternalCampaign
type MessagingV1ExternalCampaign struct {
	// The unique string that identifies a US A2P Compliance resource `QE2c6890da8086d771620e9b13fadeba0b`.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that the Campaign belongs to.
	AccountSid *string `json:"account_sid,omitempty"`
	// ID of the preregistered campaign.
	CampaignId *string `json:"campaign_id,omitempty"`
	// The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) that the resource is associated with.
	MessagingServiceSid *string `json:"messaging_service_sid,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
}
