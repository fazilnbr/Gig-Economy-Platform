/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Chat
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
	"strings"
)

// Optional parameters for the method 'UpdateChannel'
type UpdateChannelParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	//
	Type *string `json:"Type,omitempty"`
	// The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this channel belongs to.
	MessagingServiceSid *string `json:"MessagingServiceSid,omitempty"`
}

func (params *UpdateChannelParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *UpdateChannelParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *UpdateChannelParams) SetType(Type string) *UpdateChannelParams {
	params.Type = &Type
	return params
}
func (params *UpdateChannelParams) SetMessagingServiceSid(MessagingServiceSid string) *UpdateChannelParams {
	params.MessagingServiceSid = &MessagingServiceSid
	return params
}

// Update a specific Channel.
func (c *ApiService) UpdateChannel(ServiceSid string, Sid string, params *UpdateChannelParams) (*ChatV3Channel, error) {
	path := "/v3/Services/{ServiceSid}/Channels/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Type != nil {
		data.Set("Type", *params.Type)
	}
	if params != nil && params.MessagingServiceSid != nil {
		data.Set("MessagingServiceSid", *params.MessagingServiceSid)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV3Channel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
