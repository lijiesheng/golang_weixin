package pkg

import (
	"fmt"
	"testing"
)

func TestGetweather(t *testing.T) {
	weather := Getweather("武汉")
	var text string
	if weather.Alarm == nil {
		text = fmt.Sprintf("城市: %s \n"+
			"日期: %s : %s\n "+
			"最高温度 : %s, 最低温度 : %s, 当前温度 : %s\n"+
			"小黄提示您：%s",
			weather.City, weather.Date, weather.Week, weather.MaxTem, weather.MinTem, weather.CurrTem, weather.AirTips)
	} else {
		text = fmt.Sprintf("城市: %s \n"+
			"日期: %s : %s\n"+
			"最高温度 : %s, 最低温度 : %s, 当前温度 : %s\n"+
			"小黄提示您：%s",
			weather.City, weather.Date, weather.Week, weather.MaxTem, weather.MinTem, weather.CurrTem, weather.AirTips)
	}
	fmt.Println(text)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Printf("%+v\n", weather)
}

func TestGetweather1(t *testing.T) {
	getweather1 := Getweather1("江夏")
	fmt.Printf("%+v\n", getweather1)
}
