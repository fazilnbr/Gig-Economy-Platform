/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Conversations
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

// ConversationsV1ServiceConversationScopedWebhook struct for ConversationsV1ServiceConversationScopedWebhook
type ConversationsV1ServiceConversationScopedWebhook struct {
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`
	// The unique ID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this conversation.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
	ChatServiceSid *string `json:"chat_service_sid,omitempty"`
	// The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
	ConversationSid *string `json:"conversation_sid,omitempty"`
	// The target of this webhook: `webhook`, `studio`, `trigger`
	Target *string `json:"target,omitempty"`
	// An absolute API resource URL for this webhook.
	Url *string `json:"url,omitempty"`
	// The configuration of this webhook. Is defined based on target.
	Configuration *interface{} `json:"configuration,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
}
