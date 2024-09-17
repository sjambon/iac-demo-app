package handlers

import "github.com/Azure/go-autorest/autorest/date"

type event struct {
	// ID - An unique identifier for the event.
	ID *string `json:"id,omitempty"`
	// Topic - The resource path of the event source.
	Topic *string `json:"topic,omitempty"`
	// Subject - A resource path relative to the topic path.
	Subject *string `json:"subject,omitempty"`
	// Data - Event data specific to the event type.
	Data map[string]string `json:"data,omitempty"`
	// EventType - The type of the event that occurred.
	EventType *string `json:"eventType,omitempty"`
	// EventTime - The time (in UTC) the event was generated.
	EventTime *date.Time `json:"eventTime,omitempty"`
	// MetadataVersion - READ-ONLY; The schema version of the event metadata.
	MetadataVersion *string `json:"metadataVersion,omitempty"`
	// DataVersion - The schema version of the data object.
	DataVersion *string `json:"dataVersion,omitempty"`
}
