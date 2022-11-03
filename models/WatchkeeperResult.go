package models

import (
	"time"
)

type WatchkeeperResult struct {
	Id             int       `json:"id"`
	StaffID        int       `json:"staff_id"`
	StaffName      string    `json:"staff_name"`
	StaffPhone     string    `json:"staff_phone"`
	StaffEmail     string    `json:"staff_email"`
	Title          string    `json:"title"`
	Labels         []string  `json:"labels"`
	Role           string    `json:"role"`
	Classification string    `json:"classification"`
	Position       string    `json:"position"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	CreateTime     time.Time `json:"create_time"`
	ModifyTime     time.Time `json:"modify_time"`
}

type WatchKeeperInfo struct {
	AppPrimary []*WatchkeeperResult `json:"app_primary"`
	Platform   []*WatchkeeperResult `json:"platform"`
	TeamLeader []*WatchkeeperResult `json:"team_leader"`
}

func NewWatchkeeperResult() *WatchkeeperResult {
	return &WatchkeeperResult{}
}

func (wkr *WatchkeeperResult) ToWatchkeeperResult(wk *Watchkeeper) *WatchkeeperResult {
	result := new(WatchkeeperResult)
	result.Id = wk.Id
	result.StaffID = wk.Staff.MemberId
	if len(wk.Staff.RealName) == 0 {
		result.StaffName = wk.Staff.Account
	} else {
		result.StaffName = wk.Staff.RealName
	}

	result.StaffPhone = wk.Staff.Phone
	result.StaffEmail = wk.Staff.Email
	switch wk.Role {
	case WKROperation:
		result.Role = "Operation"
		result.Title = "运行值班"
		result.Labels = []string{"一线"}
	case WKRManager:
		result.Role = "Manager"
		result.Title = "一线值班"
		result.Labels = []string{"一线"}
	case WKRLeader:
		result.Role = "Leader"
		result.Title = "值班团队长"
	}
	switch wk.Classification {
	case WKCApplication:
		result.Classification = "Application"
		result.Labels = append(result.Labels, "应用")
	case WKCPlatformTenant:
		result.Classification = "Tenant"
		result.Title = "租户"
		result.Labels = append(result.Labels, "平台")
	case WKCPlatformNetwork:
		result.Classification = "Network"
		result.Title = "网络"
		result.Labels = append(result.Labels, "平台")
	case WKCPlatformOS:
		result.Classification = "OS"
		result.Title = "操作系统"
		result.Labels = append(result.Labels, "平台")
	case WKCPlatformDatabase:
		result.Classification = "Database"
		result.Title = "数据库"
		result.Labels = append(result.Labels, "平台")
	case WKCPlatformMidware:
		result.Classification = "Midware"
		result.Title = "中间件"
		result.Labels = append(result.Labels, "平台")
	}
	if wk.IsRemote {
		result.Position = "非现场"
	} else {
		result.Position = "南湖园区"
	}

	if wk.StartTime.Year() == 1 {
		currentTime := time.Now().Local()
		todayTime1 := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 8, 30, 0, 0, time.Local)
		if currentTime.Before(todayTime1) {
			result.Position = "非现场"
			result.StartTime = todayTime1.Add(-12 * time.Hour)
		} else {
			todayTime2 := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 20, 30, 0, 0, time.Local)
			if currentTime.Before(todayTime2) {
				result.StartTime = todayTime1
			} else {
				result.Position = "非现场"
				result.StartTime = todayTime2
			}
		}

	} else {
		result.StartTime = wk.StartTime.Local()
	}
	if wk.EndTime.Year() == 1 {
		currentTime := time.Now().Local()
		todayTime1 := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 8, 30, 0, 0, time.Local)
		if currentTime.Before(todayTime1) {
			result.EndTime = todayTime1
		} else {
			todayTime2 := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 20, 30, 0, 0, time.Local)
			if currentTime.Before(todayTime2) {
				result.EndTime = todayTime2
			} else {
				result.EndTime = todayTime2.Add(12 * time.Hour)
			}
		}
	} else {
		result.EndTime = wk.EndTime.Local()
	}

	result.CreateTime = wk.CreateTime.Local()
	result.ModifyTime = wk.ModifyTime.Local()
	return result
}

// func FindWatchkeeperByTime(t time.Time) {

// }

func (wkr *WatchkeeperResult) FindByTime(t time.Time) (wkInfo *WatchKeeperInfo) {
	var wk Watchkeeper

	wkInfo = &WatchKeeperInfo{
		AppPrimary: []*WatchkeeperResult{},
		Platform:   []*WatchkeeperResult{},
		TeamLeader: []*WatchkeeperResult{},
	}
	ap, err := wk.findWKByTime(t, WKROperation, WKCApplication)
	if err != nil {
		return nil
	}
	for _, w := range ap {
		wkInfo.AppPrimary = append(wkInfo.AppPrimary, wkr.ToWatchkeeperResult(w))
	}
	apm, err := wk.findWKByTime(t, WKRManager, WKCApplication)
	if err != nil {
		return nil
	}
	for _, w := range apm {
		wkInfo.AppPrimary = append(wkInfo.AppPrimary, wkr.ToWatchkeeperResult(w))
	}

	pt, err := wk.findWKByTime(t, WKRManager, WKCPlatformTenant)
	if err == nil {
		for _, w := range pt {
			wkInfo.Platform = append(wkInfo.Platform, wkr.ToWatchkeeperResult(w))
		}
	}
	pos, err := wk.findWKByTime(t, WKRManager, WKCPlatformOS)
	if err == nil {
		for _, w := range pos {
			wkInfo.Platform = append(wkInfo.Platform, wkr.ToWatchkeeperResult(w))
		}
	}
	pn, err := wk.findWKByTime(t, WKRManager, WKCPlatformNetwork)
	if err == nil {
		for _, w := range pn {
			wkInfo.Platform = append(wkInfo.Platform, wkr.ToWatchkeeperResult(w))
		}
	}
	pm, err := wk.findWKByTime(t, WKRManager, WKCPlatformMidware)
	if err == nil {
		for _, w := range pm {
			wkInfo.Platform = append(wkInfo.Platform, wkr.ToWatchkeeperResult(w))
		}
	}
	pd, err := wk.findWKByTime(t, WKRManager, WKCPlatformDatabase)
	if err == nil {
		for _, w := range pd {
			wkInfo.Platform = append(wkInfo.Platform, wkr.ToWatchkeeperResult(w))
		}
	}

	tm, err := wk.findWKByTime(t, WKRLeader, WKCApplication)
	if err == nil {
		for _, w := range tm {
			wkInfo.TeamLeader = append(wkInfo.TeamLeader, wkr.ToWatchkeeperResult(w))
		}
	}

	return
}
