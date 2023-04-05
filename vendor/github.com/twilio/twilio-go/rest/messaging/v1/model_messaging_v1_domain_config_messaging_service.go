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

// MessagingV1DomainConfigMessagingService struct for MessagingV1DomainConfigMessagingService
type MessagingV1DomainConfigMessagingService struct {
	// The unique string that we created to identify the Domain resource.
	DomainSid *string `json:"domain_sid,omitempty"`
	// The unique string that we created to identify the Domain config (prefix ZK).
	ConfigSid *string `json:"config_sid,omitempty"`
	// The unique string that identifies the messaging service
	MessagingServiceSid *string `json:"messaging_service_sid,omitempty"`
	// Any requests we receive to this domain that do not match an existing shortened message will be redirected to the fallback url. These will likely be either expired messages, random misdirected traffic, or intentional scraping.
	FallbackUrl *string `json:"fallback_url,omitempty"`
	// URL to receive click events to your webhook whenever the recipients click on the shortened links.
	CallbackUrl *string `json:"callback_url,omitempty"`
	// Date this Domain Config was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// Date that this Domain Config was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Url         *string    `json:"url,omitempty"`
}
