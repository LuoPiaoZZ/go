package utils

import "time"

//处理日期范围，将一段时间戳按日期切割
func GetTimeRangeSlice(startTime int64, endTime int64) (startRange []int64, endRange []int64) {
	if startTime <= endTime {
		//时间戳转化为time
		start := time.Unix(startTime, 0)
		end := time.Unix(endTime, 0)
		//所选开始日期的0点0分0秒
		startDayTime := time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
		//所选结束日期的后一天凌晨
		endDayTime := time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, end.Location())
		endDayTime = endDayTime.AddDate(0, 0, 1)
		//去除不必要的时间
		//当前日期的后一天凌晨
		currentTime := time.Now()
		currentTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
		currentNextTime := currentTime.AddDate(0, 0, 1)
		if endDayTime.Unix() > currentNextTime.Unix() {
			endDayTime = currentNextTime
		}
		//添加第一天时间
		dayStartTime := startDayTime
		dayStart := dayStartTime.Unix()
		dayEndTime := dayStartTime.AddDate(0, 0, 1)
		dayEnd := dayEndTime.Unix()
		startRange = append(startRange, dayStart)
		endRange = append(endRange, dayEnd)
		//天数日期不断叠加
		for dayEnd < endDayTime.Unix() {
			dayStartTime = dayEndTime
			dayStart = dayStartTime.Unix()
			dayEndTime = dayStartTime.AddDate(0, 0, 1)
			dayEnd = dayEndTime.Unix()
			//将后续日期不断添加
			startRange = append(startRange, dayStart)
			endRange = append(endRange, dayEnd)
		}
		if len(startRange) == len(endRange) {
			return startRange, endRange
		}
	}
	return nil, nil

}
