package report

import (
	"fmt"
	"strconv"

	"github.com/rudolfaraya/adapter340/structs"
)

func ParserSTT(frameSlice []string, frameParsedCommon structs.Msg) (structs.Msg, error) {

	frameParsedSTT := structs.Msg{}

	//Parseo de Mode
	ModeParsed, err := strconv.ParseInt(frameSlice[16], 10, 32)
	if err != nil {
		fmt.Println("Hubo un error al parsear Mode")
		return frameParsedSTT, errWrongData
	}
	//Parseo de MsgNum
	MsgNumParsed, err := strconv.ParseInt(frameSlice[17], 10, 32)
	if err != nil {
		fmt.Println("Hubo un error al parsear MsgNum")
		return frameParsedSTT, errWrongData
	}
	//Parseo de HourMeter
	HourMeterParsed, err := strconv.ParseInt(frameSlice[18], 10, 32)
	if err != nil {
		fmt.Println("Hubo un error al parsear HourMeter")
		return frameParsedSTT, errWrongData
	}
	//Parseo de BackupVolt
	BackupVoltParsed, err := strconv.ParseFloat(frameSlice[19], 64)
	if err != nil {
		fmt.Println("Hubo un error al parsear BackupVolt")
		return frameParsedSTT, errWrongData
	}
	//Parseo de MsgType
	MsgTypeParsed, err := strconv.ParseBool(frameSlice[20])
	if err != nil {
		fmt.Println("Hubo un error al parsear MsgType")
		return frameParsedSTT, errWrongData
	}

	//INT FIELDS
	frameParsedCommon.ContentsInt = append(frameParsedCommon.ContentsInt, ModeParsed)
	frameParsedCommon.ContentsInt = append(frameParsedCommon.ContentsInt, MsgNumParsed)
	frameParsedCommon.ContentsInt = append(frameParsedCommon.ContentsInt, HourMeterParsed)
	//FLOAT FIELDS
	frameParsedCommon.ContentsFloat = append(frameParsedCommon.ContentsFloat, BackupVoltParsed)
	//BOOL FIELDS
	frameParsedCommon.ContentsBool = append(frameParsedCommon.ContentsBool, MsgTypeParsed)

	return frameParsedCommon, nil

}
