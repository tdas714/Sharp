package structure

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// A ProposalResponse is returned from an endorser to the proposal submitter.
// The idea is that this message contains the endorser's response to the
// request of a client to perform an action over a chaincode (or more
// generically on the ledger); the response might be success/error (conveyed in
// the Response field) together with a description of the action and a
// signature over it by that endorser.  If a sufficient number of distinct
// endorsers agree on the same action and produce signature to that effect, a
// transaction can be generated and sent for ordering.
type ProposalResponse struct {
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Timestamp is the time that the message
	// was created as  defined by the sender
	Timestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// A response message indicating whether the
	// endorsement of the action was successful
	Response *Response `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	// The payload of response. It is the bytes of ProposalResponsePayload
	Payload []byte `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	// The endorsement of the proposal, basically
	// the endorser's signature over the payload
	Endorsement *Endorsement `protobuf:"bytes,6,opt,name=endorsement,proto3" json:"endorsement,omitempty"`
	// The chaincode interest derived from simulating the proposal.
	Interest             *ChaincodeInterest `protobuf:"bytes,7,opt,name=interest,proto3" json:"interest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ProposalResponse) Reset()         { *m = ProposalResponse{} }
func (m *ProposalResponse) String() string { return proto.CompactTextString(m) }
func (*ProposalResponse) ProtoMessage()    {}

func (m *ProposalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalResponse.Unmarshal(m, b)
}
func (m *ProposalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalResponse.Marshal(b, m, deterministic)
}
func (dst *ProposalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalResponse.Merge(dst, src)
}
func (m *ProposalResponse) XXX_Size() int {
	return xxx_messageInfo_ProposalResponse.Size(m)
}
func (m *ProposalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalResponse proto.InternalMessageInfo

func (m *ProposalResponse) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ProposalResponse) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ProposalResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *ProposalResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ProposalResponse) GetEndorsement() *Endorsement {
	if m != nil {
		return m.Endorsement
	}
	return nil
}

func (m *ProposalResponse) GetInterest() *ChaincodeInterest {
	if m != nil {
		return m.Interest
	}
	return nil
}

// A response with a representation similar to an HTTP response that can
// be used within another message.
type Response struct {
	// A status code that should follow the HTTP status codes.
	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	// A message associated with the response code.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// A payload that can be used to include metadata with this response.
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Response) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// ProposalResponsePayload is the payload of a proposal response.  This message
// is the "bridge" between the client's request and the endorser's action in
// response to that request. Concretely, for chaincodes, it contains a hashed
// representation of the proposal (proposalHash) and a representation of the
// chaincode state changes and events inside the extension field.
type ProposalResponsePayload struct {
	// Hash of the proposal that triggered this response. The hash is used to
	// link a response with its proposal, both for bookeeping purposes on an
	// asynchronous system and for security reasons (accountability,
	// non-repudiation). The hash usually covers the entire Proposal message
	// (byte-by-byte).
	ProposalHash []byte `protobuf:"bytes,1,opt,name=proposal_hash,json=proposalHash,proto3" json:"proposal_hash,omitempty"`
	// Extension should be unmarshaled to a type-specific message. The type of
	// the extension in any proposal response depends on the type of the proposal
	// that the client selected when the proposal was initially sent out.  In
	// particular, this information is stored in the type field of a Header.  For
	// chaincode, it's a ChaincodeAction message
	Extension            []byte   `protobuf:"bytes,2,opt,name=extension,proto3" json:"extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProposalResponsePayload) Reset()         { *m = ProposalResponsePayload{} }
func (m *ProposalResponsePayload) String() string { return proto.CompactTextString(m) }
func (*ProposalResponsePayload) ProtoMessage()    {}

func (m *ProposalResponsePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalResponsePayload.Unmarshal(m, b)
}
func (m *ProposalResponsePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalResponsePayload.Marshal(b, m, deterministic)
}
func (dst *ProposalResponsePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalResponsePayload.Merge(dst, src)
}
func (m *ProposalResponsePayload) XXX_Size() int {
	return xxx_messageInfo_ProposalResponsePayload.Size(m)
}
func (m *ProposalResponsePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalResponsePayload.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalResponsePayload proto.InternalMessageInfo

func (m *ProposalResponsePayload) GetProposalHash() []byte {
	if m != nil {
		return m.ProposalHash
	}
	return nil
}

func (m *ProposalResponsePayload) GetExtension() []byte {
	if m != nil {
		return m.Extension
	}
	return nil
}

// An endorsement is a signature of an endorser over a proposal response.  By
// producing an endorsement message, an endorser implicitly "approves" that
// proposal response and the actions contained therein. When enough
// endorsements have been collected, a transaction can be generated out of a
// set of proposal responses.  Note that this message only contains an identity
// and a signature but no signed payload. This is intentional because
// endorsements are supposed to be collected in a transaction, and they are all
// expected to endorse a single proposal response/action (many endorsements
// over a single proposal response)
type Endorsement struct {
	// Identity of the endorser (e.g. its certificate)
	Endorser []byte `protobuf:"bytes,1,opt,name=endorser,proto3" json:"endorser,omitempty"`
	// Signature of the payload included in ProposalResponse concatenated with
	// the endorser's certificate; ie, sign(ProposalResponse.payload + endorser)
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endorsement) Reset()         { *m = Endorsement{} }
func (m *Endorsement) String() string { return proto.CompactTextString(m) }
func (*Endorsement) ProtoMessage()    {}

func (m *Endorsement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endorsement.Unmarshal(m, b)
}
func (m *Endorsement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endorsement.Marshal(b, m, deterministic)
}
func (dst *Endorsement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endorsement.Merge(dst, src)
}
func (m *Endorsement) XXX_Size() int {
	return xxx_messageInfo_Endorsement.Size(m)
}
func (m *Endorsement) XXX_DiscardUnknown() {
	xxx_messageInfo_Endorsement.DiscardUnknown(m)
}

var xxx_messageInfo_Endorsement proto.InternalMessageInfo

func (m *Endorsement) GetEndorser() []byte {
	if m != nil {
		return m.Endorser
	}
	return nil
}

func (m *Endorsement) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// ChaincodeInterest defines an interest about an endorsement
// for a specific single chaincode invocation.
// Multiple chaincodes indicate chaincode to chaincode invocations.
type ChaincodeInterest struct {
	Chaincodes           []*ChaincodeCall `protobuf:"bytes,1,rep,name=chaincodes,proto3" json:"chaincodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ChaincodeInterest) Reset()         { *m = ChaincodeInterest{} }
func (m *ChaincodeInterest) String() string { return proto.CompactTextString(m) }
func (*ChaincodeInterest) ProtoMessage()    {}

func (m *ChaincodeInterest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeInterest.Unmarshal(m, b)
}
func (m *ChaincodeInterest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeInterest.Marshal(b, m, deterministic)
}
func (dst *ChaincodeInterest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeInterest.Merge(dst, src)
}
func (m *ChaincodeInterest) XXX_Size() int {
	return xxx_messageInfo_ChaincodeInterest.Size(m)
}
func (m *ChaincodeInterest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeInterest.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeInterest proto.InternalMessageInfo

func (m *ChaincodeInterest) GetChaincodes() []*ChaincodeCall {
	if m != nil {
		return m.Chaincodes
	}
	return nil
}

// ChaincodeCall defines a call to a chaincode.
// It may have collections that are related to the chaincode
type ChaincodeCall struct {
	Name            string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CollectionNames []string `protobuf:"bytes,2,rep,name=collection_names,json=collectionNames,proto3" json:"collection_names,omitempty"`
	NoPrivateReads  bool     `protobuf:"varint,3,opt,name=no_private_reads,json=noPrivateReads,proto3" json:"no_private_reads,omitempty"`
	NoPublicWrites  bool     `protobuf:"varint,4,opt,name=no_public_writes,json=noPublicWrites,proto3" json:"no_public_writes,omitempty"`
	// The set of signature policies associated with states in the write-set
	// that have state-based endorsement policies.
	KeyPolicies []*common.SignaturePolicyEnvelope `protobuf:"bytes,5,rep,name=key_policies,json=keyPolicies,proto3" json:"key_policies,omitempty"`
	// Indicates we wish to ignore the namespace endorsement policy
	DisregardNamespacePolicy bool     `protobuf:"varint,6,opt,name=disregard_namespace_policy,json=disregardNamespacePolicy,proto3" json:"disregard_namespace_policy,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *ChaincodeCall) Reset()         { *m = ChaincodeCall{} }
func (m *ChaincodeCall) String() string { return proto.CompactTextString(m) }
func (*ChaincodeCall) ProtoMessage()    {}

func (m *ChaincodeCall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeCall.Unmarshal(m, b)
}
func (m *ChaincodeCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeCall.Marshal(b, m, deterministic)
}
func (m *ChaincodeCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeCall.Merge(m, src)
}
func (m *ChaincodeCall) XXX_Size() int {
	return xxx_messageInfo_ChaincodeCall.Size(m)
}
func (m *ChaincodeCall) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeCall.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeCall proto.InternalMessageInfo

func (m *ChaincodeCall) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChaincodeCall) GetCollectionNames() []string {
	if m != nil {
		return m.CollectionNames
	}
	return nil
}

func (m *ChaincodeCall) GetNoPrivateReads() bool {
	if m != nil {
		return m.NoPrivateReads
	}
	return false
}

func (m *ChaincodeCall) GetNoPublicWrites() bool {
	if m != nil {
		return m.NoPublicWrites
	}
	return false
}

func (m *ChaincodeCall) GetKeyPolicies() []*common.SignaturePolicyEnvelope {
	if m != nil {
		return m.KeyPolicies
	}
	return nil
}

func (m *ChaincodeCall) GetDisregardNamespacePolicy() bool {
	if m != nil {
		return m.DisregardNamespacePolicy
	}
	return false
}
