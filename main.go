package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	dao "golang_wechat/dao/mysql"
	"golang_wechat/pkg"
	"log"
	"os"
	"time"
)

var mapCityFriend = map[string][]*openwechat.Friend{}

var isSend = false
var weatherMap = map[string]bool{}

func main() {
	// 日志
	logFile, err := os.OpenFile("./golang_wechat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err: ", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[golang_wechat] ")

	// 微信登录
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 登陆
	if err := bot.Login(); err != nil {
		fmt.Println(err)
		return
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取所有的好友
	friends, err := self.Friends()
	//fmt.Println(friends, err)

	// 胡冰冰
	hubingbing := friends.SearchByRemarkName(1, "冰坨子 11.22").First()
	fmt.Println(hubingbing)

	wangTianErLiJieSheng := friends.SearchByNickName(1, "王天二").First()
	fmt.Println(wangTianErLiJieSheng)

	dogRuningLijiesheng := friends.SearchByNickName(1, "motor").First()
	fmt.Println(dogRuningLijiesheng)

	friends_beijing := []*openwechat.Friend{}
	friends_wuhan := []*openwechat.Friend{}

	mapCityFriend["北京"] = append(friends_beijing, wangTianErLiJieSheng, dogRuningLijiesheng)
	mapCityFriend["武汉"] = append(friends_wuhan, hubingbing)
	//friends_wuhan := []*openwechat.Friend{}
	//mapCityFriend["武汉"] = append(friends_wuhan, hubingbing)

	for {
		//go SendMessageToFriend(mapCityFriend)
		sendMessageToFriend("黄诗美", mapCityFriend)
		time.Sleep(60 * time.Second)
	}

	//lib.SendTextToFried("asdfeiaufgadiufbviadd厉害哦dsa", wangTianErLiJieSheng)
	//lib.SendTextToFried("asdfeiaufgadiufbviadd厉害哦dsa", dogRuningLijiesheng)

	//// 获取所有的群组
	//groups, err := self.Groups()
	//fmt.Println(groups, err)

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}

//// 多个人，多个地方发送推送天气
//func SendMessageToFriend(mapCityFriend map[string][]*openwechat.Friend) {
//	// 对于同一个城市的人，只要请求一次天气，然后一个个发送就行了
//	for key, values := range mapCityFriend {
//		weather := pkg.Getweather(key)
//		for _, friend := range values {
//			sendMessageToFriend(weather, "黄诗美", friend)
//			log.Println("发送成功")
//		}
//	}
//}

// 单个人发送消息
func sendMessageToFriend(nickNameFrom string, mapCityFriend map[string][]*openwechat.Friend) {
	//messagesMorning := []string{
	//	"我放弃，不是因为我输了，而是因为我懂了。早安",
	//	"自信的女人，不一定会是漂亮的，比如说凤姐；可爱的女人不一定不雷人，比如说小月月；但是关心你的人，那一定会是我，于每日都问候，早安",
	//	"希望有这么一天你可以酷到不问任何人任何问题。早安",
	//	"那种打我一巴掌，又给我一颗糖的好，我真的不想要。早安",
	//	"我始终坚信，即使再黑暗的夜也会走到尽头，并迎来曙光。早安",
	//	"请一定要有自信，你就是一道风景，没必要在别人风景里面仰视。早安",
	//	"我多想拥抱你，可惜时光之里山南水北，可惜你我中间人来人往。早安",
	//	"人生很短，不要蜷缩在一小块阴影里。早安",
	//	"我从来没有招惹你，你为什么要来招惹我？既然招惹了，为什么半途而废？早安",
	//	"有了激情，有了梦想……。早安",
	//	"走出第一步，下一步就变得简单。想做的事，就在今天做吧，不要让未来的自己遗憾。早安",
	//	"虽然终于遇到了对的人，但是却偏偏在错误的时间里，那只能徒留遗憾，不如没有遇见。早安",
	//	"不妄想，不在其中自我沉醉。不伤害，不与自己和他人为敌。不表演，也不相信他人的表演。早安",
	//	"只是一转眼，十月就过去了；还有两个月，今年就过去了。时间和你都善变，我还是原来的我。早安",
	//	"错误和弱点，却依然认为你非常棒的人。早安",
	//	"每个人都有他的路，每条路都是正确的。人的不幸在于他们不想走自己那条路，总想走别人的路。早安",
	//	"记忆总是这个样子，在不该回忆处触碰了你的泪腺。体无完肤的爱和恨，肆虐了灵魂，感染了心肺。早安",
	//	"如果对方犯了你自己也会犯的错你就的确没有任何资格和理由去在意了。早安",
	//	"不要拿太一般的要求去要求自己，也不要太多的关心一些太一般的事。志存高远，这是一切美好和灿烂前程的开端。早安",
	//	"遗忘，从来不是件容易的事，就算你忘了，还有很多事物帮你记着那个人，比如一起听过的那首歌，比如拼过他名字的输入法……早安",
	//	"感情都有几分毒性，但过着过着你就百毒不侵了。哭啊闹啊，那都不算啥，熬久了就发现，曾经让你痛不欲生的，最终也让你骨肉相依。早安",
	//	"困难都一样，每个人都要面对无数困难。有的人将困难夸大，有的人把困难当作动力，这就是为什么有些人能够成功，而有些人注定失败。早安",
	//	"时光静好，细水长流，繁华落尽，聚散随缘，且行且珍惜。送上一份快乐，共享真挚友情；送上真挚的祝福，铭刻永恒的情谊；朋友，愿你快乐！早安",
	//	"真心祝愿你：上班偷偷傻笑，下班活泼乱跳！嘴里哼着小曲，不知不觉跑调！中午下班叽哩呱啦乱叫，晚上下班呼噜呼噜睡觉！早上醒来吓了一跳：又迟到！早安",
	//}
	// 获取当前时间
	now := time.Now()
	hour := now.Hour()
	minute := now.Minute()
	minute = minute
	today := now.Format("2006-01-02")
	var text string
	//var url string
	switch hour {

	case 6:
		//gf.SendText(messagesMorning[util.GenerateRandnum(24)])

	// 9点推送天气
	case 9:
		for key, values := range mapCityFriend {
			weather := pkg.Getweather1(key) // 获取天气
			for _, friend := range values {
				sql := "select count(*) count from wechat_message where data = ? " +
					"and type = ? and send_nick = ? and recieve_nick = ?"
				var count int
				err := dao.Db.Get(&count, sql, today, 9, nickNameFrom, friend.NickName)
				if err != nil {
					fmt.Printf("get failed, err:%v\n", err)
					return
				}
				if count == 0 {
					// 推送天气消息
					text = fmt.Sprintf("城市: %s \n"+
						"日期: %s %s\n"+
						"天气: %s %s %s\n"+
						"最高温度:%s\n"+
						"最低温度:%s\n"+
						"当前温度:%s\n",
						weather.City, weather.Date, weather.Week, weather.Wea, weather.Win, weather.WinSpeed,
						weather.TemDay, weather.TemNight, weather.Tem)
					fmt.Printf("%+v\n", weather)
					fmt.Println(text)
					friend.SendText(text)

					// 插入到数据库中
					insertSql := `insert into wechat_message(send_nick, recieve_nick, data, type, content) values (?,?,?,?,?)`
					ret, err := dao.Db.Exec(insertSql, nickNameFrom, friend.NickName, today, 9, text)
					theID, err := ret.LastInsertId() // 新插入数据的id
					if err != nil {
						fmt.Printf("get lastinsert ID failed, err:%v\n", err)
						return
					}
					log.Printf("insert success, the id is %d.\n", theID)
				}

			}
		}

		//url = fmt.Sprintf("./image/%d.png", util.GenerateRandnum(6))
		//fmt.Println("url==>", url)
		//
		//open, err := os.Open(url)
		//if err != nil {
		//	fmt.Println("图片打开错误, err==>", err)
		//	return
		//}
		//gf.SendImage(open)

		// 推送湖北疫情

		// 推送湖北高风险地区
	//case 9:
	// 1、删除历史数据
	//pkg.DelAllData()
	// 2、同步新的数据
	//pkg.GetAllData()
	// 3、查询湖北数据
	//if data, err := pkg.GetDataByProvence("武汉"); err != nil {
	//
	//}
	// 推送一首歌
	case 12:
		//gf.SendText("该午休了")
	case 13:
		//gf.SendText("学习吧")
	}
}
