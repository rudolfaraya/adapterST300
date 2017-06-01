package report

import (
	"fmt"
	"strconv"

	"github.com/rudolfaraya/adapter340/structs"
)

func ParserEmgEvtAlt(frameSlice []string, frameParsedCommon structs.Msg) (structs.Msg, error) {

	frameParsedEmgEvtAlt := structs.Msg{}

	//Parseo de EmgEvtAltID
	EmgEvtAltIDParsed, err := strconv.ParseInt(frameSlice[16], 10, 32)
	if err != nil {
		fmt.Println("Hubo un error al parsear EmgEvtAltID")
		return frameParsedEmgEvtAlt, errWrongData
	}
	//Parseo de HourMeter
	HourMeterParsed, err := strconv.ParseInt(frameSlice[17], 10, 32)
	if err != nil {
		fmt.Println("Hubo un error al parsear HourMeter")
		return frameParsedEmgEvtAlt, errWrongData
	}
	//Parseo de BackupVolt
	BackupVoltParsed, err := strconv.ParseFloat(frameSlice[18], 64)
	if err != nil {
		fmt.Println("Hubo un error al parsear BackupVolt")
		return frameParsedEmgEvtAlt, errWrongData
	}
	//Parseo de MsgType
	MsgTypeParsed, err := strconv.ParseBool(frameSlice[19])
	if err != nil {
		fmt.Println("Hubo un error al parsear MsgType")
		return frameParsedEmgEvtAlt, errWrongData
	}
	//INT FIELDS
	frameParsedCommon.ContentsInt = append(frameParsedCommon.ContentsInt, EmgEvtAltIDParsed)
	frameParsedCommon.ContentsInt = append(frameParsedCommon.ContentsInt, HourMeterParsed)
	//FLOAT FIELDS
	frameParsedCommon.ContentsFloat = append(frameParsedCommon.ContentsFloat, BackupVoltParsed)
	//BOOL FIELDS
	frameParsedCommon.ContentsBool = append(frameParsedCommon.ContentsBool, MsgTypeParsed)

	return frameParsedCommon, nil
}
