package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 获取天气预报接口
// http://wjhsh.net/Jimc-p-10250861.html
// https://blog.csdn.net/u012140251/article/details/89529540

// 免费天气API接口（http://www.tianqiapi.com）
//1.接口没有调用次数、频率和IP限制, 永久免费。
//
//2.JSON接口调用文档：http://doc.tianqiapi.com/603579
//
//3.IFrame调用：http://www.tianqiapi.com/

// 这里我使用的是  http://www.tianqiapi.com/index/doc
// 1、注册开发账号可以获得
/*
	测试豆：2000个
	appid：94853111
	appsecret：5x1bFpJg
*/
// 2、

var (
	WeatherVersion = "v1"
)

type Weather struct {
	Cityid        string `json:"cityid"`
	Date          string `json:"date"`
	Week          string `json:"week"`
	UpdateTime    string `json:"update_time"`
	City          string `json:"city"`
	CityEn        string `json:"cityEn"`
	Country       string `json:"country"`
	CountryEn     string `json:"countryEn"`
	Wea           string `json:"wea"`
	WeaImg        string `json:"wea_img"`
	Tem           string `json:"tem"`
	Tem1          string `json:"tem1"`
	Tem2          string `json:"tem2"`
	Win           string `json:"win"`
	WinSpeed      string `json:"win_speed"`
	WinMeter      string `json:"win_meter"`
	Humidity      string `json:"humidity"`
	Visibility    string `json:"visibility"`
	Pressure      string `json:"pressure"`
	Air           string `json:"air"`
	AirPm25       string `json:"air_pm25"`
	AirLevel      string `json:"air_level"`
	AirTips       string `json:"air_tips"`
	Alarm         Alarm  `json:"alarm"`
	WinSpeedDay   string `json:"win_speed_day"`
	WinSpeedNight string `json:"win_speed_night"`
	Aqi           Aqi    `json:"aqi"`
}

type Alarm struct {
	AlarmType    string `json:"alarm_type"`
	AlarmLevel   string `json:"alarm_level"`
	AlarmContent string `json:"alarm_content"`
}

type Aqi struct {
	UpdateTime string `json:"update_time"`
	Cityid     string `json:"cityid"`
	City       string `json:"city"`
	CityEn     string `json:"cityEn"`
	Country    string `json:"country"`
	CountryEn  string `json:"countryEn"`
	Air        string `json:"air"`
	AirLevel   string `json:"air_level"`
	AirTips    string `json:"air_tips"`
	Pm25       string `json:"pm25"`
	Pm25Desc   string `json:"pm25_desc"`
	Pm10       string `json:"pm10"`
	Pm10Desc   string `json:"pm10_desc"`
	O3         string `json:"o3"`
	O3Desc     string `json:"o3_desc"`
	No2        string `json:"no2"`
	No2Desc    string `json:"no2_desc"`
	So2        string `json:"so2"`
	So2Desc    string `json:"so2_desc"`
	Co         string `json:"co"`
	CoDesc     string `json:"co_desc"`
	Kouzhao    string `json:"kouzhao"`
	Yundong    string `json:"yundong"`
	Waichu     string `json:"waichu"`
	Kaichuang  string `json:"kaichuang"`
	Jinghuaqi  string `json:"jinghuaqi"`
}

type ResWeather struct {
	City     string `json:"city"`
	Date     string `json:"date"`
	Week     string `json:"week"`
	Wea      string `json:"wea"`
	Win      string `json:"win"`
	WinSpeed string `json:"win_speed"`
	AirTips  string `json:"air_tips"`
	MaxTem   string `json:"max_tem"`  // 最高温度
	MinTem   string `json:"max_tem"`  // 最低温度
	CurrTem  string `json:"curr_tem"` // 当前温度
	Alarm    *Alarm `json:"alarm"`
}

// 根据城市名获取天气
func Getweather(cityName string) *ResWeather {
	url := fmt.Sprintf("https://v0.yiketianqi.com/api?unescape=1&version=v61&appid=94853111&appsecret=5x1bFpJg&city=%s", cityName)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("获取天气失败,", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取内容失败", err)
		return nil
	}

	weather := &Weather{}
	err = json.Unmarshal(body, weather)
	if err != nil {
		fmt.Println("转换出错")
		return nil
	}
	return &ResWeather{
		City:     weather.City,
		Date:     weather.Date,
		Wea:      weather.Wea,
		Week:     weather.Week,
		Win:      weather.Win,
		WinSpeed: weather.WinSpeed,
		AirTips:  weather.AirTips,
		MaxTem:   weather.Tem1,
		MinTem:   weather.Tem2,
		CurrTem:  weather.Tem,
		Alarm:    &weather.Alarm,
	}
}
