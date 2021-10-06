package structure

import "github.com/golang/protobuf/proto"

// ChaincodeEvent is used for events and registrations that are specific to chaincode
//string type - "chaincode"
type ChaincodeEvent struct {
	ChaincodeId          string   `protobuf:"bytes,1,opt,name=chaincode_id,json=chaincodeId,proto3" json:"chaincode_id,omitempty"`
	TxId                 string   `protobuf:"bytes,2,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	EventName            string   `protobuf:"bytes,3,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaincodeEvent) Reset()         { *m = ChaincodeEvent{} }
func (m *ChaincodeEvent) String() string { return proto.CompactTextString(m) }
func (*ChaincodeEvent) ProtoMessage()    {}

func (m *ChaincodeEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeEvent.Unmarshal(m, b)
}
func (m *ChaincodeEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeEvent.Marshal(b, m, deterministic)
}
func (dst *ChaincodeEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeEvent.Merge(dst, src)
}
func (m *ChaincodeEvent) XXX_Size() int {
	return xxx_messageInfo_ChaincodeEvent.Size(m)
}
func (m *ChaincodeEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeEvent proto.InternalMessageInfo

func (m *ChaincodeEvent) GetChaincodeId() string {
	if m != nil {
		return m.ChaincodeId
	}
	return ""
}

func (m *ChaincodeEvent) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *ChaincodeEvent) GetEventName() string {
	if m != nil {
		return m.EventName
	}
	return ""
}

func (m *ChaincodeEvent) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}
