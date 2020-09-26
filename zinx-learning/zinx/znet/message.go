package znet

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

// 创建一个Message消息包
func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

func (m *Message) GetMsgId() uint32 {
	return m.Id
}
func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}
func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgId(uid uint32) {
	m.Id = uid
}
func (m *Message) SetDataLen(len uint32) {
	m.DataLen = len
}
func (m *Message) SetMsgData(data []byte) {
	m.Data = data
}
