package initial

import (
	"fmt"
	"math/rand"
	"time"
)

func GetNowStr() string {
	t := time.Now().Hour()
	if t < 11 && t > 6 {
		return "上午"
	} else if t >= 11 && t < 13 {
		return "中午"
	} else if t >= 13 && t < 17 {
		return "下午"
	} else if t >= 17 && t < 20 {
		return "傍晚"
	} else {
		return "晚上"
	}
}

func GetMsgTime(t time.Time) string {
	s := t.Day()
	n := time.Now().Day()

	if s == n {
		str := fmt.Sprintf("%d:%d", t.Hour(), t.Minute())
		return str
	} else if s-n == 1 {
		return "昨天"
	} else if s-n == 2 {
		return "前天"
	} else {
		str := fmt.Sprintf("%d-%d", t.Month(), t.Day())
		return str
	}
}

func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min)
	randNum = randNum + min
	return randNum
}
