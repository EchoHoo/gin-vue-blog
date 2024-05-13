package ctype

type MsgType int

const (
	TextMsg    MsgType = 1
	ImageMsg   MsgType = 2
	SystemMsg  MsgType = 3
	InRoomMsg  MsgType = 4
	OutRoomMsg MsgType = 5
)
