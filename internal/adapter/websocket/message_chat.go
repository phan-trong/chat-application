package websocket

type IMessageChat interface {
	encode() []byte
	UnmarshalJSON(data []byte) error
}
