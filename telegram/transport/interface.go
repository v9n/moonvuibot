package transport

type Transport interface {
	Send(interface{}) interface{}
	Set(interface{}) interface{}
	Forward(interface{}) interface{}
}
