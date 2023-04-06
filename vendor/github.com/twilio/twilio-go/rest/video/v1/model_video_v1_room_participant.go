/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Video
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

// VideoV1RoomParticipant struct for VideoV1RoomParticipant
type VideoV1RoomParticipant struct {
	// The unique string that we created to identify the RoomParticipant resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the participant's room.
	RoomSid *string `json:"room_sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the RoomParticipant resource.
	AccountSid *string `json:"account_sid,omitempty"`
	Status     *string `json:"status,omitempty"`
	// The application-defined string that uniquely identifies the resource's User within a Room. If a client joins with an existing Identity, the existing client is disconnected. See [access tokens](https://www.twilio.com/docs/video/tutorials/user-identity-access-tokens) and [limits](https://www.twilio.com/docs/video/programmable-video-limits) for more info.
	Identity *string `json:"identity,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The time of participant connected to the room in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format.
	StartTime *time.Time `json:"start_time,omitempty"`
	// The time when the participant disconnected from the room in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format.
	EndTime *time.Time `json:"end_time,omitempty"`
	// The duration in seconds that the participant was `connected`. Populated only after the participant is `disconnected`.
	Duration *int `json:"duration,omitempty"`
	// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
	// The URLs of related resources.
	Links *map[string]interface{} `json:"links,omitempty"`
}