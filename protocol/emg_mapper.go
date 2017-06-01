package protocol

import (
	"fmt"

	"github.com/rudolfaraya/adapter340/structs"
)

func EmgTypeToReportType(emgType EmergencyType) structs.ReportType {
	switch emgType {
	case PanicButtonEmg:
		return structs.PanicButton
	case ParkingLockEmg:
		return structs.Unknown
	case RemovingMainPowerEmg:
		return structs.Unknown
	case AntiTheftEmg:
		return structs.Unknown
	case AntiTheftDoorEmg:
		return structs.Unknown
	case MotionEmg:
		return structs.Unknown
	case AntiTheftShockEmg:
		return structs.Unknown
	default:
		return structs.Unknown
	}
}

func ToEmgID(emgID int64) (emgType EmergencyType, err error) {

	switch emgID {
	case 1:
		emgType = PanicButtonEmg
	case 2:
		emgType = ParkingLockEmg
	case 3:
		emgType = RemovingMainPowerEmg
	case 5:
		emgType = AntiTheftEmg
	case 6:
		emgType = AntiTheftDoorEmg
	case 7:
		emgType = MotionEmg
	case 8:
		emgType = AntiTheftShockEmg
	default:
		err = fmt.Errorf("unknown EmgID value")
	}
	return emgType, nil
}
