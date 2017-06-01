package protocol

import (
	"fmt"

	"github.com/rudolfaraya/adapter340/structs"
)

func eventTypeToReportType(evtID EventType) structs.ReportType {
	switch evtID {
	case Input1GroundEvt:
		return structs.Unknown
	case Input1OpenEvt:
		return structs.Unknown
	case Input2GroundEvt:
		return structs.Unknown
	case Input2OpenEvt:
		return structs.Unknown
	case Input3GroundEvt:
		return structs.Unknown
	case Input3OpenEvt:
		return structs.Unknown
	default:
		return structs.Unknown
	}
}

func ToEvtID(evtID int64) (evtType EventType, err error) {

	switch evtID {
	case 1:
		evtType = Input1GroundEvt
	case 2:
		evtType = Input1OpenEvt
	case 3:
		evtType = Input2GroundEvt
	case 4:
		evtType = Input2OpenEvt
	case 5:
		evtType = Input3GroundEvt
	case 6:
		evtType = Input3OpenEvt
	default:
		err = fmt.Errorf("unknown evtID value")
	}
	return evtType, nil
}
