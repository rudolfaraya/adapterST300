package protocol

import (
	"fmt"

	"github.com/rudolfaraya/adapter340/structs"
)

func AlertTypeToReportType(altID AlertType) structs.ReportType {
	switch altID {
	case StartOverSpeedAlt:
		return structs.OverSpeed
	case StopOverSpeedAlt:
		return structs.NormalSpeed
	case DisconnectedGPSAntennaAlt, ReconnectedGPSAntennaAlt, ExitedGeoFenceAlt, EnteredGeoFenceAlt,
		ShortedGPSAntennaAlt, BackupBatteryErrorAlt:
		return structs.Unknown
	case BatteryLowLevelAlt:
		return structs.BatteryAlert
	case ShockedAlt:
		return structs.Shocked
	case CollisionAlt:
		return structs.Collision
	case DeviatedFromRouteAlt, EnteredIntoRouteAlt:
		return structs.Unknown
	case EngineExceedSpeedAlt, EngineVehicleSpeedAlt, EngineCoolantTempAlt, EngineOilPressureAlt,
		EngineRPMAlt, EngineHardBrakeAlt, EngineErrCodeAlt:
		return structs.Unknown
	case IgnitionOnAlt:
		return structs.IgnitionOn
	case IgnitionOffAlt:
		return structs.IgnitionOff
	case ConnectedToMainPowerAlt:
		return structs.MainBattery
	case DisconnectedFromMainPowerAlt:
		return structs.DisconnectedFromMainBattery
	case ConnectedToBackupBatteryAlt:
		return structs.BackupBattery
	case DisconnectedToBackupBatteryAlt:
		return structs.DisconnectedToBackupBattery
	case FastAccelerationFromDPAAlt:
		return structs.HarshAcceleration
	case FastBrakingFromDPAAlt:
		return structs.HarshBraking
	case SharpTurnFromDPAAlt:
		return structs.SharpTurn
	case OverSpeedFromDPAAlt:
		return structs.Unknown
	case JammingDetectedAlt:
		return structs.JammingDetected
	case InsertedIButtonAlt, RemovedIButtonAlt, DriveLessThanPredefinedTimeAlt,
		StoppedMoreThanPredefinedTimeAlt, DeadCenterAlt, OverRPMAlt, CompletedAutoRPMCalibrationAlt,
		CompletedAutoOdometerCalibrationAlt, CompletedAutoOdometerCalibrationDualGearSystemAlt:
		return structs.Unknown
	case StopLimitAtIgnitionONAlt, MovingAfterStopLimitAtIgnitionONAlt:
		return structs.Unknown
	case EnterDeepSleepModeAlt:
		return structs.EnterDeepSleepMode
	case ExitDeepSleepModeAlt:
		return structs.ExitDeepSleepMode
	default:
		return structs.Unknown
	}
}

func ToAltID(altID int64) (altType AlertType, err error) {

	switch altID {
	case 0:
		altType = Unknown
	case 1:
		altType = StartOverSpeedAlt
	case 2:
		altType = StopOverSpeedAlt
	case 3:
		altType = DisconnectedGPSAntennaAlt
	case 4:
		altType = ReconnectedGPSAntennaAlt
	case 5:
		altType = ExitedGeoFenceAlt
	case 6:
		altType = EnteredGeoFenceAlt
	case 8:
		altType = ShortedGPSAntennaAlt
	case 9:
		altType = EnterDeepSleepModeAlt
	case 10:
		altType = ExitDeepSleepModeAlt
	case 13:
		altType = BackupBatteryErrorAlt
	case 14:
		altType = BatteryLowLevelAlt
	case 15:
		altType = ShockedAlt
	case 16:
		altType = CollisionAlt
	case 18:
		altType = DeviatedFromRouteAlt
	case 19:
		altType = EnteredIntoRouteAlt
	case 22:
		altType = EngineExceedSpeedAlt
	case 23:
		altType = EngineVehicleSpeedAlt
	case 24:
		altType = EngineCoolantTempAlt
	case 25:
		altType = EngineOilPressureAlt
	case 26:
		altType = EngineRPMAlt
	case 27:
		altType = EngineHardBrakeAlt
	case 28:
		altType = EngineErrCodeAlt
	case 33:
		altType = IgnitionOnAlt
	case 34:
		altType = IgnitionOffAlt
	case 40:
		altType = ConnectedToMainPowerAlt
	case 41:
		altType = DisconnectedFromMainPowerAlt
	case 44:
		altType = ConnectedToBackupBatteryAlt
	case 45:
		altType = DisconnectedToBackupBatteryAlt
	case 46:
		altType = FastAccelerationFromDPAAlt
	case 47:
		altType = FastBrakingFromDPAAlt
	case 48:
		altType = SharpTurnFromDPAAlt
	case 49:
		altType = OverSpeedFromDPAAlt
	case 50:
		altType = JammingDetectedAlt
	case 59:
		altType = InsertedIButtonAlt
	case 60:
		altType = RemovedIButtonAlt
	case 61:
		altType = DriveLessThanPredefinedTimeAlt
	case 62:
		altType = StoppedMoreThanPredefinedTimeAlt
	case 63:
		altType = DeadCenterAlt
	case 64:
		altType = OverRPMAlt
	case 65:
		altType = CompletedAutoRPMCalibrationAlt
	case 66:
		altType = CompletedAutoOdometerCalibrationAlt
	case 67:
		altType = CompletedAutoOdometerCalibrationDualGearSystemAlt
	case 68:
		altType = StopLimitAtIgnitionONAlt
	case 69:
		altType = MovingAfterStopLimitAtIgnitionONAlt
	case 73:
		altType = RapidFuelReductionAlt
	default:
		err = fmt.Errorf("unknown AltID value")
	}
	return altType, nil
}
