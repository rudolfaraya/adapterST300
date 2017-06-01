package common

func MapperAlt(altID int64) string {

	switch altID {
	case 1:
		message := "Start driving faster than SPEED_LIMIT"
		return message
	case 2:
		message := "Ended over-speed condition"
		return message
	case 3:
		message := "Disconnected GPS antenna"
		return message
	case 4:
		message := "Reconnected GPS antenna after disconnected"
		return message
	case 5:
		message := "The vehicle exited the geo-fenced area that has the following ID"
		return message
	case 6:
		message := "The vehicle entered the geo-fenced area that has the following ID"
		return message
	case 8:
		message := "Shorted GPS antenna"
		return message
	case 9:
		message := "Enter the Deep Sleep Mode"
		return message
	case 10:
		message := "Exit from the Deep Sleep Mode"
		return message
	case 13:
		message := "Backup battery error"
		return message
	case 14: //Analizado en measurement
		message := "Vehicle battery goes down to a very low level"
		return message
	case 15: //Analizado en measurement
		message := "Shocked"
		return message
	case 16: //Analizado en measurement
		message := "Collision has occurred to the vehicle"
		return message
	case 18:
		message := "Deviated from the predefined route"
		return message
	case 19:
		message := "Entered into the predefined route"
		return message
	case 22: //Analizado en measurement
		message := "Engine :Exceed Speed"
		return message
	case 23:
		message := "Engine :Vehicle Speed Violation"
		return message
	case 24:
		message := "Engine :Coolant Temperature Violation"
		return message
	case 25:
		message := "Engine :Oil Pressure Violation"
		return message
	case 26:
		message := "Engine :RPM Violation"
		return message
	case 27: //Analizado en measurement
		message := "Engine :Sudden Hard Brake Violation"
		return message
	case 28:
		message := "Engine :Error Code(DTC)"
		return message
	case 33: //Analizado en measurement
		message := "Ignition ON"
		return message
	case 34: //Analizado en measurement
		message := "Ignition OFF"
		return message
	case 40: //Analizado en measurement
		message := "Connected to the Main Power source"
		return message
	case 41:
		message := "Disconnected to the Main Power source"
		return message
	case 44: //Analizado en measurement
		message := "Connected to the Back-up Battery"
		return message
	case 45:
		message := "Disconnected to the Back-up Battery"
		return message
	case 46: //Analizado en measurement
		message := "Alert of fast accelaration from Driver Pattern Analysis"
		return message
	case 47: //Analizado en measurement
		message := "Alert of fast accelaration from Driver Pattern Analysis"
		return message
	case 48: //Analizado en measurement
		message := "Alert of sharp turn from Driver Pattern Analysis"
		return message
	case 49:
		message := "Alert of over speed from Driver Pattern Analysis"
		return message
	case 50: //Analizado en measurement
		message := "Jamming detected"
		return message
	case 59:
		message := "Inserted I-Button"
		return message
	case 60:
		message := "Removed I-Button"
		return message
	case 61: //Analizado en measurement
		message := "The vehicle turns on but doesn't drive predefined time"
		return message
	case 62: //Analizado en measurement
		message := "Stopped more than predefined time"
		return message
	case 63:
		message := "Dead center"
		return message
	case 64:
		message := "Over RPM"
		return message
	case 65:
		message := "Completed automatic RPM calibration"
		return message
	case 66:
		message := "Completed automatic odometer calibration"
		return message
	case 67:
		message := "Completed automatic odometer calibration as another type in dual gear system"
		return message
	case 68:
		message := "Alert of Stop limit at Ignition ON"
		return message
	case 69:
		message := "Alert of Moving after Alert 68"
		return message
	case 73:
		message := "Alert of rapid reduction of the Fuel"
		return message
	case 75:
		message := "The tempeture is higher than the preset temperature"
		return message
	case 76:
		message := "The tempeture is lower than the preset temperature"
		return message
	default:
		message := "NO ALERT"
		return message
	}
}
