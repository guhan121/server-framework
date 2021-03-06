//author:liweisheng date:2015/07/07

package connector

const (
	ST_INITED = iota
	ST_CLOSED = iota
)

type Socket interface {
	ID() uint32
	Socket() interface{}
	RemoteAddress() map[string]interface{}
	Send([]byte) (int, error)
	SendBatch(...[]byte)
	Receive([]byte) (int, error)
	Disconnect()
}

type Connector interface {
	Start()
	Decode([]byte) (interface{}, error)
	Encode(string, string, map[string]interface{}) ([]byte, error)
}
