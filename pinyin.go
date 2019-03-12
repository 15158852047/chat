package main

import (
	"fmt"
	"github.com/axgle/mahonia"
	"strconv"
)

func test() string {
	str := "测试"
	str1 := strconv.QuoteToASCII(str)
	gbkstr := mahonia.NewDecoder("gbk").ConvertString(str1)
	asc := int(gbkstr[0])*256 + int(gbkstr[1]) - 65536
	fmt.Println(asc)
	if asc >= -20319 && asc <= -20284 {
		return "a"
	} else if asc >= -20283 && asc <= -19776 {
		return "b"
	} else if asc >= -19775 && asc <= -19219 {
		return "c"
	} else if asc >= -19218 && asc <= -18711 {
		return "d"
	} else if asc >= -18710 && asc <= -18527 {
		return "e"
	} else if asc >= -18526 && asc <= -18240 {
		return "f"
	} else if asc >= -18239 && asc <= -17923 {
		return "g"
	} else if asc >= -17922 && asc <= -17418 {
		return "h"
	} else if asc >= -17417 && asc <= -16475 {
		return "j"
	} else if asc >= -16474 && asc <= -16213 {
		return "k"
	} else if asc >= -16212 && asc <= -15641 {
		return "l"
	} else if asc >= -15640 && asc <= -15166 {
		return "m"
	} else if asc >= -15165 && asc <= -14923 {
		return "n"
	} else if asc >= -14922 && asc <= -14915 {
		return "o"
	} else if asc >= -14914 && asc <= -14631 {
		return "p"
	} else if asc >= -14630 && asc <= -14150 {
		return "q"
	} else if asc >= -14149 && asc <= -14091 {
		return "r"
	} else if asc >= -14090 && asc <= -13119 {
		return "s"
	} else if asc >= -13118 && asc <= -12839 {
		return "t"
	} else if asc >= -12838 && asc <= -12557 {
		return "w"
	} else if asc >= -12556 && asc <= -11848 {
		return "x"
	} else if asc >= -11847 && asc <= -11056 {
		return "y"
	} else if asc >= -11055 && asc <= -10247 {
		return "z"
	} else {
		return ""
	}
}

func main() {
	s := test()
	fmt.Println(s)
}
