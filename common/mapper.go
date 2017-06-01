package common

import (
	"errors"

	"github.com/rudolfaraya/adapter340/structs"
)

var (
	errMapper = errors.New("Error-Mapper: Frame can't be mapped")
)

func MapperALT(commonFrameToALT structs.Msg) structs.ALT {

	ALTProcessed := structs.ALT{
		ServerTimestamp: commonFrameToALT.ServerTimestamp,
		FrameID:         commonFrameToALT.FrameID,
		Hdr:             commonFrameToALT.Hdr,
		DevID:           commonFrameToALT.DevID,
		Model:           commonFrameToALT.ContentsInt[0],
		Version:         commonFrameToALT.Version,
		Date:            commonFrameToALT.ContentsInt[1],
		Time:            commonFrameToALT.ContentsString[0],
		Cell:            commonFrameToALT.ContentsString[1],
		Latitude:        commonFrameToALT.ContentsFloat[0],
		Longitude:       commonFrameToALT.ContentsFloat[1],
		Speed:           commonFrameToALT.ContentsFloat[2],
		Heading:         commonFrameToALT.ContentsFloat[3],
		Satellites:      commonFrameToALT.ContentsInt[2],
		GPSFixed:        commonFrameToALT.ContentsBool[0],
		Distance:        commonFrameToALT.ContentsInt[3],
		PowerVolt:       commonFrameToALT.ContentsFloat[4],
		InOut:           commonFrameToALT.ContentsString[2],
		AltID:           commonFrameToALT.ContentsInt[4],
		HourMeter:       commonFrameToALT.ContentsInt[5],
		BackupVolt:      commonFrameToALT.ContentsFloat[5],
		MsgType:         commonFrameToALT.ContentsBool[1]}

	return ALTProcessed
}

func MapperEVT(commonFrameToEVT structs.Msg) structs.EVT {

	EVTProcessed := structs.EVT{
		ServerTimestamp: commonFrameToEVT.ServerTimestamp,
		FrameID:         commonFrameToEVT.FrameID,
		Hdr:             commonFrameToEVT.Hdr,
		DevID:           commonFrameToEVT.DevID,
		Model:           commonFrameToEVT.ContentsInt[0],
		Version:         commonFrameToEVT.Version,
		Date:            commonFrameToEVT.ContentsInt[1],
		Time:            commonFrameToEVT.ContentsString[0],
		Cell:            commonFrameToEVT.ContentsString[1],
		Latitude:        commonFrameToEVT.ContentsFloat[0],
		Longitude:       commonFrameToEVT.ContentsFloat[1],
		Speed:           commonFrameToEVT.ContentsFloat[2],
		Heading:         commonFrameToEVT.ContentsFloat[3],
		Satellites:      commonFrameToEVT.ContentsInt[2],
		GPSFixed:        commonFrameToEVT.ContentsBool[0],
		Distance:        commonFrameToEVT.ContentsInt[3],
		PowerVolt:       commonFrameToEVT.ContentsFloat[4],
		InOut:           commonFrameToEVT.ContentsString[2],
		EvtID:           commonFrameToEVT.ContentsInt[4],
		HourMeter:       commonFrameToEVT.ContentsInt[5],
		BackupVolt:      commonFrameToEVT.ContentsFloat[5],
		MsgType:         commonFrameToEVT.ContentsBool[1]}

	return EVTProcessed
}

func MapperEMG(commonFrameToEMG structs.Msg) structs.EMG {

	EMGProcessed := structs.EMG{
		ServerTimestamp: commonFrameToEMG.ServerTimestamp,
		FrameID:         commonFrameToEMG.FrameID,
		Hdr:             commonFrameToEMG.Hdr,
		DevID:           commonFrameToEMG.DevID,
		Model:           commonFrameToEMG.ContentsInt[0],
		Version:         commonFrameToEMG.Version,
		Date:            commonFrameToEMG.ContentsInt[1],
		Time:            commonFrameToEMG.ContentsString[0],
		Cell:            commonFrameToEMG.ContentsString[1],
		Latitude:        commonFrameToEMG.ContentsFloat[0],
		Longitude:       commonFrameToEMG.ContentsFloat[1],
		Speed:           commonFrameToEMG.ContentsFloat[2],
		Heading:         commonFrameToEMG.ContentsFloat[3],
		Satellites:      commonFrameToEMG.ContentsInt[2],
		GPSFixed:        commonFrameToEMG.ContentsBool[0],
		Distance:        commonFrameToEMG.ContentsInt[3],
		PowerVolt:       commonFrameToEMG.ContentsFloat[4],
		InOut:           commonFrameToEMG.ContentsString[2],
		EmgID:           commonFrameToEMG.ContentsInt[4],
		HourMeter:       commonFrameToEMG.ContentsInt[5],
		BackupVolt:      commonFrameToEMG.ContentsFloat[5],
		MsgType:         commonFrameToEMG.ContentsBool[1]}

	return EMGProcessed
}

func MapperSTT(commonFrameToSTT structs.Msg) structs.STT {

	STTProcessed := structs.STT{
		ServerTimestamp: commonFrameToSTT.ServerTimestamp,
		FrameID:         commonFrameToSTT.FrameID,
		Hdr:             commonFrameToSTT.Hdr,
		DevID:           commonFrameToSTT.DevID,
		Model:           commonFrameToSTT.ContentsInt[0],
		Version:         commonFrameToSTT.Version,
		Date:            commonFrameToSTT.ContentsInt[1],
		Time:            commonFrameToSTT.ContentsString[0],
		Cell:            commonFrameToSTT.ContentsString[1],
		Latitude:        commonFrameToSTT.ContentsFloat[0],
		Longitude:       commonFrameToSTT.ContentsFloat[1],
		Speed:           commonFrameToSTT.ContentsFloat[2],
		Heading:         commonFrameToSTT.ContentsFloat[3],
		Satellites:      commonFrameToSTT.ContentsInt[2],
		GPSFixed:        commonFrameToSTT.ContentsBool[0],
		Distance:        commonFrameToSTT.ContentsInt[3],
		PowerVolt:       commonFrameToSTT.ContentsFloat[4],
		InOut:           commonFrameToSTT.ContentsString[2],
		Mode:            commonFrameToSTT.ContentsInt[4],
		MsgNum:          commonFrameToSTT.ContentsInt[5],
		HourMeter:       commonFrameToSTT.ContentsInt[6],
		BackupVolt:      commonFrameToSTT.ContentsFloat[5],
		MsgType:         commonFrameToSTT.ContentsBool[1]}

	return STTProcessed
}

func Mapper(commonFrame structs.Msg) (structs.Msg, structs.STT, structs.EMG, structs.EVT, structs.ALT, error) {

	emptyFrame := structs.Msg{}
	emptySTT := structs.STT{}
	emptyEMG := structs.EMG{}
	emptyEVT := structs.EVT{}
	emptyALT := structs.ALT{}

	if commonFrame.Hdr == "ST300STT" {
		STTProcessed := MapperSTT(commonFrame)

		return emptyFrame, STTProcessed, emptyEMG, emptyEVT, emptyALT, nil
	} else if commonFrame.Hdr == "ST300EMG" {
		EMGProcessed := MapperEMG(commonFrame)

		return emptyFrame, emptySTT, EMGProcessed, emptyEVT, emptyALT, nil
	} else if commonFrame.Hdr == "ST300EVT" {
		EVTProcessed := MapperEVT(commonFrame)

		return emptyFrame, emptySTT, emptyEMG, EVTProcessed, emptyALT, nil
	} else if commonFrame.Hdr == "ST300ALT" {
		ALTProcessed := MapperALT(commonFrame)

		return emptyFrame, emptySTT, emptyEMG, emptyEVT, ALTProcessed, nil
	}
	return emptyFrame, emptySTT, emptyEMG, emptyEVT, emptyALT, errMapper
}
