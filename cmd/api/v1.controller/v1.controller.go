package v1_controller

type IControllerV1 interface {
	IUserControllerV1
	IVideoControllerV1
	IShareControllerV1
}

type ControllerV1 struct{}

func NewControllerV1() IControllerV1 {
	return &ControllerV1{}
}
