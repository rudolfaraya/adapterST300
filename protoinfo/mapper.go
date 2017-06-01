package protoinfo

import (
	"errors"
	"time"

	"github.com/rudolfaraya/adapter340/structs"
)

// ToProto maps a *frameinfo.FrameInfo to its proto type
func ToProto(frameInfo *structs.FrameInfo) *FrameInfo {
	fi := &FrameInfo{
		FrameId:         frameInfo.FrameID,
		ServerTimestamp: frameInfo.ServerTimestamp.Unix(),
		Adapter:         frameInfo.Adapter,
		Received:        frameInfo.Received,
		Data:            frameInfo.Data,
		DataFormat:      toProtoFormat(frameInfo.DataFormat),
		DevId:           frameInfo.DevID,
		Reports:         toProtoReports(frameInfo),
		Tags:            frameInfo.Tags,
	}
	if frameInfo.ParsingError != nil {
		fi.ParsingError = frameInfo.ParsingError.Error()
	}
	return fi
}

func ToProtoReport(r *structs.Report) *Report {
	report := &Report{
		ReportId:        r.ReportID,
		InstallationId:  int32(r.InstallationID),
		AssetId:         int32(r.AssetID),
		DevId:           r.DevID,
		ServerTimestamp: r.ServerTimestamp.Unix(),
		FrameId:         r.FrameID,
		Type:            toProtoReportType(r.Type),
		Timestamp:       r.Timestamp.Unix(),
		Latitude:        r.Latitude,
		Longitude:       r.Longitude,
		Speed:           r.Speed,
		Heading:         r.Heading,
		GpsFixed:        r.GPSFixed,
		HourMeter:       int64(r.HourMeter),
		Tags:            r.Tags,
	}
	if r.Satellites != nil {
		report.Satellites = &Int32Value{int32(*r.Satellites)}
	}
	if r.PowerVolt != nil {
		report.PowerVolt = &DoubleValue{*r.PowerVolt}
	}
	if r.BackupVolt != nil {
		report.BackupVolt = &DoubleValue{*r.BackupVolt}
	}
	return report
}

func toProtoReports(frameInfo *structs.FrameInfo) []*Report {
	reports := make([]*Report, 0, len(frameInfo.Reports))
	for _, r := range frameInfo.Reports {
		report := ToProtoReport(&r)
		reports = append(reports, report)
	}
	return reports
}

func FromProto(pb *FrameInfo) *structs.FrameInfo {
	fi := &structs.FrameInfo{
		FrameID:         pb.FrameId,
		ServerTimestamp: time.Unix(pb.ServerTimestamp, 0).UTC(),
		Adapter:         pb.Adapter,
		Received:        pb.Received,
		Data:            pb.Data,
		DataFormat:      fromProtoFormat(pb.DataFormat),
		DevID:           pb.DevId,
		Reports:         fromProtoReports(pb),
		Tags:            pb.Tags,
	}
	if pb.ParsingError != "" {
		fi.ParsingError = errors.New(pb.ParsingError)
	}
	return fi
}

func FromProtoReport(pb *Report) structs.Report {
	report := structs.Report{
		ReportID:        pb.ReportId,
		InstallationID:  int(pb.InstallationId),
		AssetID:         int(pb.AssetId),
		DevID:           pb.DevId,
		ServerTimestamp: time.Unix(pb.ServerTimestamp, 0).UTC(),
		FrameID:         pb.FrameId,
		Type:            fromProtoReportType(pb.Type),
		Timestamp:       time.Unix(pb.Timestamp, 0).UTC(),
		Latitude:        pb.Latitude,
		Longitude:       pb.Longitude,
		Speed:           pb.Speed,
		Heading:         pb.Heading,
		GPSFixed:        pb.GpsFixed,
		HourMeter:       int64(pb.HourMeter),
		Tags:            pb.Tags,
	}
	if pb.Satellites != nil {
		s := int(pb.Satellites.Value)
		report.Satellites = &s
	}
	if pb.PowerVolt != nil {
		pv := pb.PowerVolt.Value
		report.PowerVolt = &pv
	}
	if pb.BackupVolt != nil {
		bv := pb.BackupVolt.Value
		report.BackupVolt = &bv
	}
	return report
}

func fromProtoReports(frameInfo *FrameInfo) []structs.Report {
	reports := make([]structs.Report, 0, len(frameInfo.Reports))
	for _, r := range frameInfo.Reports {
		report := FromProtoReport(r)
		reports = append(reports, report)
	}
	return reports
}
