package structs

import "time"

// ReportType is the type of a report.
//go:generate stringer -type=ReportType
type ReportType int

const (
	// Unknown is used as type of unmapped events
	Unknown ReportType = iota

	// 1: PanicButton occurs when the panic button is pressed
	PanicButton

	// 2: IgnitionOn occurs when the ignition is on.
	IgnitionOn

	// 3: IgnitionOff occurs when the ignition is off.
	IgnitionOff

	// 4: ModemOff occurs when the asset stands still, the ignition is off and a report by time is sent.
	ModemOff

	// 5: TimeExpired occurs when the asset stands still, the ignition is on and a report by time is sent.
	TimeExpired

	// 6: WayPoint occurs when the asset is moving and the position is reported (this could be by time, by distance
	// or using another criteria).
	WayPoint

	// 7.BackupBattery occurs when the backup battery starts working.
	BackupBattery

	// 8: MainBattery occurs when the main battery starts working.
	MainBattery

	// 9: OverSpeed occurs when the asset exceeds a certain speed (configured in the device).
	OverSpeed

	// 10: NormalSpeed occurs when the asset goes back to a certain speed, considered normal (configured in the device).
	NormalSpeed

	// 11: HarshAcceleration occurs when when the driver applies more force than normal to the asset's accelerator.
	HarshAcceleration

	// 12: HarshBraking occurs when when the driver applies more force than normal to the asset's brake.
	HarshBraking

	// 13: Shocked occurs when a shock impact is given to the asset.
	Shocked

	// 14: Collision occurs when a collision happened to the asset.
	Collision

	// 15: SharpTurn occurs when the driver changes the direction of the asset suddenly.
	SharpTurn

	// 16: JammingDetected occurs when the device detects a Jamming signal.
	JammingDetected

	// 17: ExternalData occurs when the device reports external data (from another device, usually CAN data)
	ExternalData

	// 18: Tow occurs when the asset is being towed
	Tow

	// 19: Turn occurs when the asset turns more than "x" degrees
	Turn

	// 20: Motion occurs when the asset is in motion and the ignition is off but isn't being towed (or if the ignition status is unknown)
	Motion

	// 21: BatteryAlert occurs when the asset battery voltage is in a specific range, usually a low one and the user
	// should be notified.
	BatteryAlert

	// 22: Occurs when the dispositive is disconnected from backup battery
	DisconnectedToBackupBattery

	// 23: Occurs when the dispositive is disconnected from main battery.
	DisconnectedFromMainBattery

	// 24: Ocurrs when the asset stands still, the ignition is off, a report by time is sent and the SwitchStart is on
	ParkingModeNextIgnitionCut

	// 25: Ocurrs when the asset stands still, the ignition is on, a report by time is sent and the SwitchStart is on
	DrivingModeNextIgnitionCut

	// 26: Occurs when the asset is moving, the position is reported and the SwitchStart is on
	DistanceModeNextIgnitionCut

	// 27: Occurs when the asset turns more than "x" degrees and the SwitchStart is on
	AngleModeNextIgnitionCut

	// 28: Occurs when the dispositive is in Deep Sleep Mode
	EnterDeepSleepMode

	// 29: Occurs when the dispositive exit form Deep Sleep Mode
	ExitDeepSleepMode
)

// Report is a measurement from a modem device, like ST600, MT4000, etc. A report has a type, and is originated from a
// single frame that was measured in a specific place and time (so it has position and it's timestamped).
type Report struct {
	// ReportID is the ID of the report, a non-empty string, that identifies uniquely this report.
	ReportID string

	// InstallationID is the ID of the related installation. Could be zero in early stages of the pipeline.
	InstallationID int

	// AssetID is the ID of the asset, that identifies uniquely the asset. Could be zero in early stages of the
	// pipeline, when the asset of the report is still unknown.
	AssetID int

	// DevID is the device ID, a non-empty string, with a maximum length of 20, that uniquely identifies a device.
	DevID string

	// ServerTimestamp is the report creation timestamp in the server.
	// This value can diverge significantly from the real timestamp if the frame was buffered in the device prior
	// to its transmission to the server.
	ServerTimestamp time.Time

	// FrameID is the ID of the frame, a non-empty string, source of the report.
	FrameID string

	// Type is the type of the report.
	Type ReportType

	// Timestamp is the timestamp when the measurement was made (in the device).
	// This value can diverge significantly from the server timestamp if the frame was buffered in the device prior
	// to its transmission to the server.
	Timestamp time.Time

	// Latitude is the latitude in decimal degrees.
	Latitude float64

	// Longitude is the longitude in decimal degrees.
	Longitude float64

	// Speed is the speed of the asset in km per hour.
	Speed float64

	// Heading is the orientation of the asset in sexagesimal degrees.
	Heading float64

	// GPSFixed indicates if the GPS is fixed (check https://en.wikipedia.org/wiki/Fix_(position))
	GPSFixed bool

	// Satellites is the number of visible satellites. Optional field.
	Satellites *int

	// PowerVolt is the voltage of the main power. Optional field.
	PowerVolt *float64

	// BackupVolt is the voltage of the backup power. Optional field.
	BackupVolt *float64

	HourMeter int64

	// FIXME: move location fields to a separate struct
	// LocationEstimated indicates if the fields related to the location are an estimation (based in a previous
	// measurement). The location is compromised by the latitude, longitude, heading and street address
	LocationEstimated bool

	// Tags is a optional group of free form tags, useful to mark special messages (for example, a "smoke-test"
	// tag indicates a message used to verify that the infrastructure is in place and operating correctly,
	// so should be ignored by consumers)
	Tags []string
}
