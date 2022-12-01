package lib

import "github.com/eatmoreapple/openwechat"

// 单个好友发送文本消息
func SendTextToFried(text string, gf *openwechat.Friend) {
	gf.SendText(text) // 发送文本
}

// 单个好友发送图片消息
func SendImageToFriend(gf *openwechat.Friend) {
	//gf.SendImage("")  // 发送图片
}

func SendFileToFriend(gf *openwechat.Friend) {
	//gf.SendFile()  // 发送文件
}

func SendVideoToFriend(gf *openwechat.Friend) {
	//gf.SendVideo()   // 发送视频
}
