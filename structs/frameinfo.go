package structs

import (
	"time"
)

// Format is the type of format of the Data field of a FrameInfo
type Format string

const (
	// Raw represents a Data field with the content as-is was read from the socket.
	Raw Format = ""

	// Hex represents a Data field with the content as an hex string.
	Hex Format = "hex"
)

// FrameInfo contains the information relative to a single frame received and processed by an adapter.
type FrameInfo struct {
	// FrameID is the ID of the frame, a non-empty string who identifies uniquely the associated raw frame.
	FrameID string

	// ServerTimestamp is the timestamp of the reception of the frame in the server.
	ServerTimestamp time.Time

	// Adapter identifies the adapter that received(or sent) the frame.
	Adapter string

	// Received indicates if the frame was received (true) from a device, or was sent (false) to a device.
	Received bool

	// Data is the raw frame.
	Data []byte

	// DataFormat is the format of the Data field.
	DataFormat Format

	// DevID is the device ID, a string with a maximum length of 20, that uniquely identifies a device.
	// This field can be empty, the parser couldn't be capable to extract the device ID in case of a
	// parsing error.
	DevID string

	// ParsingError contains a posible parsing error, optional.
	ParsingError error

	// Reports contains the extracted reports from the frame, generally is only one. Optional, this field
	// could be nil or empty if the frame isn't recognized or if there was a parsing error.
	Reports []Report

	// Tags is a optional group of free form tags, useful to mark special messages (for example, a "smoke-test"
	// tag indicates a message used to verify that the infrastructure is in place and operating correctly,
	// so should be ignored by consumers)
	Tags []string
}
