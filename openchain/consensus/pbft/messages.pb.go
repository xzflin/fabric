// Code generated by protoc-gen-go.
// source: messages.proto
// DO NOT EDIT!

/*
Package pbft is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	Request
	Phase
	RequestHashes
	Requests
	PBFT
*/
package pbft

import proto "github.com/golang/protobuf/proto"

// discarding unused import google_protobuf "google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Phase_Type int32

const (
	Phase_UNDEFINED      Phase_Type = 0
	Phase_PRE_PREPARE    Phase_Type = 1
	Phase_PREPARE        Phase_Type = 2
	Phase_COMMIT         Phase_Type = 3
	Phase_PREPARE_RESULT Phase_Type = 4
	Phase_COMMIT_RESULT  Phase_Type = 5
)

var Phase_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "PRE_PREPARE",
	2: "PREPARE",
	3: "COMMIT",
	4: "PREPARE_RESULT",
	5: "COMMIT_RESULT",
}
var Phase_Type_value = map[string]int32{
	"UNDEFINED":      0,
	"PRE_PREPARE":    1,
	"PREPARE":        2,
	"COMMIT":         3,
	"PREPARE_RESULT": 4,
	"COMMIT_RESULT":  5,
}

func (x Phase_Type) String() string {
	return proto.EnumName(Phase_Type_name, int32(x))
}

type PBFT_Type int32

const (
	PBFT_UNDEFINED      PBFT_Type = 0
	PBFT_REQUEST        PBFT_Type = 1
	PBFT_PRE_PREPARE    PBFT_Type = 2
	PBFT_PREPARE        PBFT_Type = 3
	PBFT_COMMIT         PBFT_Type = 4
	PBFT_PREPARE_RESULT PBFT_Type = 5
	PBFT_COMMIT_RESULT  PBFT_Type = 6
)

var PBFT_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "REQUEST",
	2: "PRE_PREPARE",
	3: "PREPARE",
	4: "COMMIT",
	5: "PREPARE_RESULT",
	6: "COMMIT_RESULT",
}
var PBFT_Type_value = map[string]int32{
	"UNDEFINED":      0,
	"REQUEST":        1,
	"PRE_PREPARE":    2,
	"PREPARE":        3,
	"COMMIT":         4,
	"PREPARE_RESULT": 5,
	"COMMIT_RESULT":  6,
}

func (x PBFT_Type) String() string {
	return proto.EnumName(PBFT_Type_name, int32(x))
}

// Request is the message passed by the peer to the validator.
// The peer receives a client request to deploy or invoke a chaincode, assigns
// it a unique ID, then packages into a "Request" message that is passed along
// to one of its connected validators.
type Request struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

// Phase is the message type shared among validators during the consensus phase.
// The "sequenceNumber" field is set by the leader when setting a "PRE_PREPARE".
// The "payload" field is expected to contain an array of hashes corresponding
// to received "Request" messages  when the Type is "PRE_PREPARE". When the Type
// is "PREPARE_RESULT" or "COMMIT_RESULT", "payload" should carry the candidate
// global hash.
type Phase struct {
	Type           Phase_Type `protobuf:"varint,1,opt,name=type,enum=pbft.Phase_Type" json:"type,omitempty"`
	SequenceNumber uint64     `protobuf:"varint,2,opt,name=sequenceNumber" json:"sequenceNumber,omitempty"`
	Payload        []byte     `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Phase) Reset()         { *m = Phase{} }
func (m *Phase) String() string { return proto.CompactTextString(m) }
func (*Phase) ProtoMessage()    {}

// RequestHashes contains the hashes of the "Request" messages that the leader
// will package in a consensus round. Marshalling such a message (proto.Marshal)
// should generate the "payload" of a PRE_PREPARE message (see "Phase" message).
type RequestHashes struct {
	Hashes []string `protobuf:"bytes,1,rep,name=hashes" json:"hashes,omitempty"`
}

func (m *RequestHashes) Reset()         { *m = RequestHashes{} }
func (m *RequestHashes) String() string { return proto.CompactTextString(m) }
func (*RequestHashes) ProtoMessage()    {}

// Requests will be used temporarily instead of RequestHashes. In order to use
// "RequestHashes" a datastore is needed on every validator that maps incoming
// "Request" messages to their hashes. Until this structure is implemented, the
// leader should proto.Marshal() a "Requests" message for a PRE_PREPARE
// message's "payload".
type Requests struct {
	Requests []*Requests `protobuf:"bytes,1,rep,name=requests" json:"requests,omitempty"`
}

func (m *Requests) Reset()         { *m = Requests{} }
func (m *Requests) String() string { return proto.CompactTextString(m) }
func (*Requests) ProtoMessage()    {}

func (m *Requests) GetRequests() []*Requests {
	if m != nil {
		return m.Requests
	}
	return nil
}

type PBFT struct {
	Type    PBFT_Type `protobuf:"varint,1,opt,name=type,enum=pbft.PBFT_Type" json:"type,omitempty"`
	Id      string    `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Payload []byte    `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *PBFT) Reset()         { *m = PBFT{} }
func (m *PBFT) String() string { return proto.CompactTextString(m) }
func (*PBFT) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("pbft.Phase_Type", Phase_Type_name, Phase_Type_value)
	proto.RegisterEnum("pbft.PBFT_Type", PBFT_Type_name, PBFT_Type_value)
}
