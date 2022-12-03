package model

type WechatMessage struct {
	ID           int    `db:"id"`
	SendNick     string `db:"send_nick"`
	ReceieveNick string `db:"recieve_nick"`
	Data         string `db:"data"`
	Type         int    `db:"type"`
}
