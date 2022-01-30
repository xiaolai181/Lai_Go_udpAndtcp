package iudp

type Message struct {
	ID      uint32
	DataLen uint32
	KeyLen  uint32
	Key     []byte
	Data    []byte
}

func NewMessage(ID uint32, key []byte, Data []byte) *Message {
	return &Message{
		ID:      ID,
		DataLen: uint32(len(Data)),
		KeyLen:  uint32(len(key)),
		Key:     key,
		Data:    Data,
	}
}
func (ms *Message) GetMsgID() uint32 {
	return ms.ID
}
func (ms *Message) GetDataLen() uint32 {
	return ms.DataLen
}
func (ms *Message) GetMsgData() []byte {
	return ms.Data
}

func (ms *Message) GetKeyLen() uint32 {
	return ms.KeyLen
}
func (ms *Message) GetKey() []byte {
	return ms.Key
}
func (ms *Message) SetMsgID(id uint32) {
	ms.ID = id
}
func (ms *Message) SetDataLen(datalen uint32) {
	ms.DataLen = datalen
}

func (ms *Message) SetKeyLen(keylen uint32) {
	ms.KeyLen = keylen
}
func (ms *Message) SetKey(key []byte) {
	ms.Key = key
}
func (ms *Message) SetMsgData(data []byte) {
	ms.Data = data
}
