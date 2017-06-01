package common

import (
	"strconv"
	"strings"
	"time"

	"github.com/rudolfaraya/adapter340/protocol"
	"github.com/rudolfaraya/adapter340/protoinfo"
	"github.com/rudolfaraya/adapter340/structs"
)

func ToFormatPlatform(msg structs.Msg, rawFrame string) *protoinfo.FrameInfo {
	msgFormated := &structs.FrameInfo{
		FrameID:         strconv.Itoa(msg.FrameID),
		ServerTimestamp: time.Now(),
		Adapter:         "ST340",
		Received:        true,
		Data:            []byte(rawFrame),
		DataFormat:      "",
		DevID:           strconv.FormatInt(msg.DevID, 10)}

	satellites := int(msg.ContentsInt[2])

	//-------------------Time Parsing----------------------
	tsLayout := "20060102;15:04:05;"

	date := strconv.Itoa(int(msg.ContentsInt[1]))
	timeDev := msg.ContentsString[0]

	t, _ := time.Parse(tsLayout, date+timeDev)
	//-------------------Time Parsing----------------------

	var hourMeter int64
	reportType := protocol.AlertTypeToReportType(0)

	if msg.Hdr == "ST300EVT" || msg.Hdr == "ST300ALT" {
		typeAlert, _ := protocol.ToAltID(msg.ContentsInt[4])
		reportType = protocol.AlertTypeToReportType(typeAlert)
	} else if msg.Hdr == "ST300EMG" {
		typeEmergency, _ := protocol.ToEmgID(msg.ContentsInt[4])
		reportType = protocol.EmgTypeToReportType(typeEmergency)
	} else if msg.Hdr == "ST300STT" {
		//Asignar HourMeter, campo que sólo se encuentra en trama STT
		hourMeter = msg.ContentsInt[6]

		//para comprobar salida de dispositivo en campo InOut de una trama STT (conexión a cortacorriente)
		if strings.Compare(msg.ContentsString[2], "100000") == 0 || strings.Compare(msg.ContentsString[2], "000000") == 0 {
			typeMode, _ := protocol.ToModeNextIgnitionCut(msg.ContentsInt[4])
			reportType = protocol.ModeToReportType(typeMode)
		} else if strings.Compare(msg.ContentsString[2], "100010") == 0 || strings.Compare(msg.ContentsString[2], "000010") == 0 {
			typeMode, _ := protocol.ToMode(msg.ContentsInt[4])
			reportType = protocol.ModeToReportType(typeMode)
		} else {
			reportType = structs.Unknown
		}
	} else {
		reportType = structs.Unknown
	}

	newReport := structs.Report{
		ReportID:        UniqueID(),
		InstallationID:  0,
		AssetID:         0,
		DevID:           strconv.FormatInt(msg.DevID, 10),
		ServerTimestamp: time.Now(),
		FrameID:         strconv.Itoa(msg.FrameID),
		Type:            reportType,
		Timestamp:       t,
		Latitude:        msg.ContentsFloat[0],
		Longitude:       msg.ContentsFloat[1],
		Speed:           msg.ContentsFloat[2],
		Heading:         msg.ContentsFloat[3],
		GPSFixed:        msg.ContentsBool[0],
		HourMeter:       hourMeter,
		Satellites:      &satellites,
		PowerVolt:       &msg.ContentsFloat[4],
		BackupVolt:      &msg.ContentsFloat[5]}

	//msgFormated.Reports[0].Timestamp = msg.ContentsString[0]
	//msgFormated.Reports[0].Satellites = msg.ContentsInt[2]
	//msgFormated.Reports[0].PowerVolt = msg.ContentsFloat[4]
	//msgFormated.Reports[0].BackupVolt = msg.ContentsFloat[5]
	msgFormated.Reports = append(msgFormated.Reports, newReport)

	pb := protoinfo.ToProto(msgFormated)

	return pb
}
