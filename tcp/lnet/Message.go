package lnet

type Message struct {
	Datalen uint32
	ID      uint32
	Data    []byte
}

func NewMsgPackage(ID uint32, data []byte) *Message {
	return &Message{
		ID:      ID,
		Datalen: uint32(len(data)),
		Data:    data,
	}
}
func (msg *Message) GetDatalen() uint32 {
	return msg.Datalen
}

//GetMsgID 获取消息ID
func (msg *Message) GetMsgID() uint32 {
	return msg.ID
}

//GetData 获取消息内容
func (msg *Message) GetData() []byte {
	return msg.Data
}

//SetDataLen 设置消息数据段长度
func (msg *Message) SetDatalen(len uint32) {
	msg.Datalen = len
}

//SetMsgID 设计消息ID
func (msg *Message) SetMsgID(msgID uint32) {
	msg.ID = msgID
}

//SetData 设计消息内容
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
