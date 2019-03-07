package initial

import "time"

func GetNowStr() string {
	t := time.Now().Hour()
	if t < 11 && t > 6 {
		return "上午"
	} else if t>=11 && t<13 {
		return "中午"
	} else if t>=13 && t < 17 {
		return "下午"
	} else if t>=17 && t<20{
		return "傍晚"
	} else {
		return "晚上"
	}
}
