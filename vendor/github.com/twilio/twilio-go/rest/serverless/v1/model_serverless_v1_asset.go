/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Serverless
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

// ServerlessV1Asset struct for ServerlessV1Asset
type ServerlessV1Asset struct {
	// The unique string that we created to identify the Asset resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Asset resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Service that the Asset resource is associated with.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The string that you assigned to describe the Asset resource. It can be a maximum of 255 characters.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The date and time in GMT when the Asset resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the Asset resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the Asset resource.
	Url *string `json:"url,omitempty"`
	// The URLs of the Asset resource's nested resources.
	Links *map[string]interface{} `json:"links,omitempty"`
}
