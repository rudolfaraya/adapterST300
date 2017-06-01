package protoinfo

import (
	"fmt"

	"github.com/rudolfaraya/adapter340/structs"
)

func toProtoFormat(format structs.Format) FrameInfo_Format {
	switch format {
	case structs.Raw:
		return FrameInfo_Raw
	case structs.Hex:
		return FrameInfo_Hex
	default:
		panic(fmt.Sprintf("unknown frame info format %v", format))
	}
}

func fromProtoFormat(format FrameInfo_Format) structs.Format {
	switch format {
	case FrameInfo_Raw:
		return structs.Raw
	case FrameInfo_Hex:
		return structs.Hex
	default:
		panic(fmt.Sprintf("unknown proto frame info format %v", format))
	}
}

func toProtoReportType(reportType structs.ReportType) Report_ReportType {
	switch reportType {
	case structs.Unknown:
		return Report_Unknown
	case structs.PanicButton:
		return Report_PanicButton
	case structs.IgnitionOn:
		return Report_IgnitionOn
	case structs.IgnitionOff:
		return Report_IgnitionOff
	case structs.ModemOff:
		return Report_ModemOff
	case structs.TimeExpired:
		return Report_TimeExpired
	case structs.WayPoint:
		return Report_WayPoint
	case structs.BackupBattery:
		return Report_BackupBattery
	case structs.MainBattery:
		return Report_MainBattery
	case structs.OverSpeed:
		return Report_OverSpeed
	case structs.NormalSpeed:
		return Report_NormalSpeed
	case structs.HarshAcceleration:
		return Report_HarshAcceleration
	case structs.HarshBraking:
		return Report_HarshBraking
	case structs.Shocked:
		return Report_Shocked
	case structs.Collision:
		return Report_Collision
	case structs.SharpTurn:
		return Report_SharpTurn
	case structs.JammingDetected:
		return Report_JammingDetected
	case structs.ExternalData:
		return Report_ExternalData
	case structs.Tow:
		return Report_Tow
	case structs.Turn:
		return Report_Turn
	case structs.Motion:
		return Report_Motion
	case structs.BatteryAlert:
		return Report_BatteryAlert
	case structs.DisconnectedToBackupBattery:
		return Report_DisconnectedBackupBattery
	case structs.DisconnectedFromMainBattery:
		return Report_DisconnectedMainBattery
	case structs.ParkingModeNextIgnitionCut:
		return Report_ParkingModeNextIgnitionCut
	case structs.DrivingModeNextIgnitionCut:
		return Report_DrivingModeNextIgnitionCut
	case structs.DistanceModeNextIgnitionCut:
		return Report_DistanceModeNextIgnitionCut
	case structs.AngleModeNextIgnitionCut:
		return Report_AngleModeNextIgnitionCut
	case structs.EnterDeepSleepMode:
		return Report_EnterDeepSleepMode
	case structs.ExitDeepSleepMode:
		return Report_ExitDeepSleepMode
	default:
		panic(fmt.Sprintf("Unknown report type %v", reportType))
	}
}

func fromProtoReportType(protoType Report_ReportType) structs.ReportType {
	switch protoType {
	case Report_Unknown:
		return structs.Unknown
	case Report_PanicButton:
		return structs.PanicButton
	case Report_IgnitionOn:
		return structs.IgnitionOn
	case Report_IgnitionOff:
		return structs.IgnitionOff
	case Report_ModemOff:
		return structs.ModemOff
	case Report_TimeExpired:
		return structs.TimeExpired
	case Report_WayPoint:
		return structs.WayPoint
	case Report_BackupBattery:
		return structs.BackupBattery
	case Report_MainBattery:
		return structs.MainBattery
	case Report_OverSpeed:
		return structs.OverSpeed
	case Report_NormalSpeed:
		return structs.NormalSpeed
	case Report_HarshAcceleration:
		return structs.HarshAcceleration
	case Report_HarshBraking:
		return structs.HarshBraking
	case Report_Shocked:
		return structs.Shocked
	case Report_Collision:
		return structs.Collision
	case Report_SharpTurn:
		return structs.SharpTurn
	case Report_JammingDetected:
		return structs.JammingDetected
	case Report_ExternalData:
		return structs.ExternalData
	case Report_Tow:
		return structs.Tow
	case Report_Turn:
		return structs.Turn
	case Report_Motion:
		return structs.Motion
	case Report_BatteryAlert:
		return structs.BatteryAlert
	case Report_DisconnectedBackupBattery:
		return structs.DisconnectedToBackupBattery
	case Report_ParkingModeNextIgnitionCut:
		return structs.ParkingModeNextIgnitionCut
	case Report_DrivingModeNextIgnitionCut:
		return structs.DrivingModeNextIgnitionCut
	case Report_DistanceModeNextIgnitionCut:
		return structs.DistanceModeNextIgnitionCut
	case Report_AngleModeNextIgnitionCut:
		return structs.AngleModeNextIgnitionCut
	case Report_EnterDeepSleepMode:
		return structs.EnterDeepSleepMode
	case Report_ExitDeepSleepMode:
		return structs.ExitDeepSleepMode
	default:
		panic(fmt.Sprintf("Unknown proto type %v", protoType))
	}
}
