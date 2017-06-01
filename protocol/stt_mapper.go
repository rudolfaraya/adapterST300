package protocol

import (
	"fmt"

	"github.com/rudolfaraya/adapter340/structs"
)

func ModeToReportType(mode ModeType) structs.ReportType {
	switch mode {
	case ParkingMode:
		return structs.ModemOff
	case DrivingMode:
		return structs.TimeExpired
	case DistanceMode:
		return structs.WayPoint
	case AngleMode:
		return structs.Turn
	case ParkingModeNextIgnitionCut:
		return structs.ParkingModeNextIgnitionCut
	case DrivingModeNextIgnitionCut:
		return structs.DrivingModeNextIgnitionCut
	case DistanceModeNextIgnitionCut:
		return structs.DistanceModeNextIgnitionCut
	case AngleModeNextIgnitionCut:
		return structs.AngleModeNextIgnitionCut
	default:
		return structs.Unknown
	}
}

func ToMode(mode int64) (modeType ModeType, err error) {

	switch mode {
	case 1:
		modeType = ParkingMode
	case 2:
		modeType = DrivingMode
	case 4:
		modeType = DistanceMode
	case 5:
		modeType = AngleMode
	default:
		err = fmt.Errorf("invalid mode value")
	}
	return modeType, nil
}

func ToModeNextIgnitionCut(mode int64) (modeType ModeType, err error) {

	switch mode {
	case 1:
		modeType = ParkingModeNextIgnitionCut
	case 2:
		modeType = DrivingModeNextIgnitionCut
	case 4:
		modeType = DistanceModeNextIgnitionCut
	case 5:
		modeType = AngleModeNextIgnitionCut
	default:
		err = fmt.Errorf("invalid mode value")
	}
	return modeType, nil
}
