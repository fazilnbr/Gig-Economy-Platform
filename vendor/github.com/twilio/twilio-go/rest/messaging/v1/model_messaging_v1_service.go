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

// MessagingV1Service struct for MessagingV1Service
type MessagingV1Service struct {
	// The unique string that we created to identify the Service resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Service resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The URL we call using `inbound_method` when a message is received by any phone number or short code in the Service. When this property is `null`, receiving inbound messages is disabled. All messages sent to the Twilio phone number or short code will not be logged and received on the Account. If the `use_inbound_webhook_on_number` field is enabled then the webhook url defined on the phone number will override the `inbound_request_url` defined for the Messaging Service.
	InboundRequestUrl *string `json:"inbound_request_url,omitempty"`
	// The HTTP method we use to call `inbound_request_url`. Can be `GET` or `POST`.
	InboundMethod *string `json:"inbound_method,omitempty"`
	// The URL that we call using `fallback_method` if an error occurs while retrieving or executing the TwiML from the Inbound Request URL. If the `use_inbound_webhook_on_number` field is enabled then the webhook url defined on the phone number will override the `fallback_url` defined for the Messaging Service.
	FallbackUrl *string `json:"fallback_url,omitempty"`
	// The HTTP method we use to call `fallback_url`. Can be: `GET` or `POST`.
	FallbackMethod *string `json:"fallback_method,omitempty"`
	// The URL we call to [pass status updates](https://www.twilio.com/docs/sms/api/message-resource#message-status-values) about message delivery.
	StatusCallback *string `json:"status_callback,omitempty"`
	// Whether to enable [Sticky Sender](https://www.twilio.com/docs/sms/services#sticky-sender) on the Service instance.
	StickySender *bool `json:"sticky_sender,omitempty"`
	// Whether to enable the [MMS Converter](https://www.twilio.com/docs/sms/services#mms-converter) for messages sent through the Service instance.
	MmsConverter *bool `json:"mms_converter,omitempty"`
	// Whether to enable [Smart Encoding](https://www.twilio.com/docs/sms/services#smart-encoding) for messages sent through the Service instance.
	SmartEncoding      *bool   `json:"smart_encoding,omitempty"`
	ScanMessageContent *string `json:"scan_message_content,omitempty"`
	// Whether to enable [Fallback to Long Code](https://www.twilio.com/docs/sms/services#fallback-to-long-code) for messages sent through the Service instance.
	FallbackToLongCode *bool `json:"fallback_to_long_code,omitempty"`
	// Whether to enable [Area Code Geomatch](https://www.twilio.com/docs/sms/services#area-code-geomatch) on the Service Instance.
	AreaCodeGeomatch *bool `json:"area_code_geomatch,omitempty"`
	// Reserved.
	SynchronousValidation *bool `json:"synchronous_validation,omitempty"`
	// How long, in seconds, messages sent from the Service are valid. Can be an integer from `1` to `14,400`.
	ValidityPeriod *int `json:"validity_period,omitempty"`
	// The absolute URL of the Service resource.
	Url *string `json:"url,omitempty"`
	// The absolute URLs of related resources.
	Links *map[string]interface{} `json:"links,omitempty"`
	// A string that describes the scenario in which the Messaging Service will be used. Examples: [notification, marketing, verification, poll ..]
	Usecase *string `json:"usecase,omitempty"`
	// Whether US A2P campaign is registered for this Service.
	UsAppToPersonRegistered *bool `json:"us_app_to_person_registered,omitempty"`
	// A boolean value that indicates either the webhook url configured on the phone number will be used or `inbound_request_url`/`fallback_url` url will be called when a message is received from the phone number. If this field is enabled then the webhook url defined on the phone number will override the `inbound_request_url`/`fallback_url` defined for the Messaging Service.
	UseInboundWebhookOnNumber *bool `json:"use_inbound_webhook_on_number,omitempty"`
}
