package dtos

type OrderPushInfoDto struct {
	OrderProductId string
	PushType       int // 1 api push, 2 web push
	PushUrl        string
}
