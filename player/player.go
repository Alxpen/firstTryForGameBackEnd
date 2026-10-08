package player

import (
	"awesomeProject/chat"
	"awesomeProject/function"
)

type Player struct {
	UId        uint64
	FriendList []uint64 //朋友
	chChat     chan chat.Msg
}

func NewPlayer() *Player {
	p := &Player{
		UId:        0,
		FriendList: nil,
	}
	return p
}

func (p *Player) AddFriend(fid uint64) {
	if !function.CheckInNumberSlice(fId, p.FriendList) {
		p.FriendList = append(p.FriendList, fid)
	}
}

func (p *Player) DelFriend(fId uint64) {
	p.FriendList = function.DelEleInSlice(fId, p.FriendList)
}

func (p *Player) Run() {
	for {
		select {
		case chatMsg := <-p.chChat:
			p.ResolveChatMessage(chatMsg)
		}
	}
}

func (p *Player) ResolveChatMessage(chatMsg chat.Msg) {

}
