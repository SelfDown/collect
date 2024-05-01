package collect

import "time"

func UnixTime(delay int, unit string) int64 {
	if delay == 0 {
		return time.Now().Unix()
	}
	if unit == "second" {
		tmp := time.Second * time.Duration(delay)
		return time.Now().Add(tmp).Unix()
	} else if unit == "minute" {
		tmp := time.Minute * time.Duration(delay)
		return time.Now().Add(tmp).Unix()
	} else if unit == "hour" { // 小时
		tmp := time.Hour * time.Duration(delay)
		return time.Now().Add(tmp).Unix()
	} else if unit == "day" { // 天
		tmp := time.Hour * 24 * time.Duration(delay)
		return time.Now().Add(tmp).Unix()
	} else if unit == "mon" { // 月
		tmp := time.Hour * 30 * 24 * time.Duration(delay)
		return time.Now().Add(tmp).Unix()
	} else {
		return time.Now().Unix()
	}

}
