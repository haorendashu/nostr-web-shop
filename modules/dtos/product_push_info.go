package dtos

type ProductPushInfoDto struct {
	Id           int64
	Pid          int64
	Status       int
	NoticePubkey string
	PushAddress  string
	PushKey      string
	PushType     int // 1 api push, 2 web push

	Name string
}
