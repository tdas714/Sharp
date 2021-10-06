package structure

// Structure of proposal

// This structure is necessary to sign the proposal which contains the header
// and the payload. Without this structure, we would have to concatenate the
// header and the payload to verify the signature, which could be expensive
// with large payload
//
// When an endorser receives a SignedProposal message, it should verify the
// signature over the proposal bytes. This verification requires the following
// steps:
// 1. Verification of the validity of the certificate that was used to produce
//    the signature.  The certificate will be available once proposalBytes has
//    been unmarshalled to a Proposal message, and Proposal.header has been
//    unmarshalled to a Header message. While this unmarshalling-before-verifying
//    might not be ideal, it is unavoidable because i) the signature needs to also
//    protect the signing certificate; ii) it is desirable that Header is created
//    once by the client and never changed (for the sake of accountability and
//    non-repudiation). Note also that it is actually impossible to conclusively
//    verify the validity of the certificate included in a Proposal, because the
//    proposal needs to first be endorsed and ordered with respect to certificate
//    expiration transactions. Still, it is useful to pre-filter expired
//    certificates at this stage.
// 2. Verification that the certificate is trusted (signed by a trusted CA) and
//    that it is allowed to transact with us (with respect to some ACLs);
// 3. Verification that the signature on proposalBytes is valid;
// 4. Detect replay attacks;
type SignedProposal struct {
	// The bytes of Proposal
	ProposalBytes []byte `protobuf:"bytes,1,opt,name=proposal_bytes,json=proposalBytes,proto3" json:"proposal_bytes,omitempty"`
	// Signaure over proposalBytes; this signature is to be verified against
	// the creator identity contained in the header of the Proposal message
	// marshaled as proposalBytes
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedProposal) Reset() { *m = SignedProposal{} }

// Nesseceray methods of SignedProposal
func (m *SignedProposal) GetProposalBytes() []byte {
	if m != nil {
		return m.ProposalBytes
	}
	return nil
}
func (m *SignedProposal) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// A Proposal is sent to an endorser for endorsement.  The proposal contains:
// 1. A header which should be unmarshaled to a Header message.  Note that
//    Header is both the header of a Proposal and of a Transaction, in that i)
//    both headers should be unmarshaled to this message; and ii) it is used to
//    compute cryptographic hashes and signatures.  The header has fields common
//    to all proposals/transactions.  In addition it has a type field for
//    additional customization. An example of this is the ChaincodeHeaderExtension
//    message used to extend the Header for type CHAINCODE.
// 2. A payload whose type depends on the header's type field.
// 3. An extension whose type depends on the header's type field.
//
// Let us see an example. For type CHAINCODE (see the Header message),
// we have the following:
// 1. The header is a Header message whose extensions field is a
//    ChaincodeHeaderExtension message.
// 2. The payload is a ChaincodeProposalPayload message.
// 3. The extension is a ChaincodeAction that might be used to ask the
//    endorsers to endorse a specific ChaincodeAction, thus emulating the
//    submitting peer model.
type Proposal struct {
	// The header of the proposal. It is the bytes of the Header
	Header []byte `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The payload of the proposal as defined by the type in the proposal
	// header.
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// Optional extensions to the proposal. Its content depends on the Header's
	// type field.  For the type CHAINCODE, it might be the bytes of a
	// ChaincodeAction message.
	Extension            []byte   `protobuf:"bytes,3,opt,name=extension,proto3" json:"extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Proposal) Reset() { *m = Proposal{} }

// Nesseceray methods of Proposal
func (m *Proposal) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *Proposal) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}
func (m *Proposal) GetExtension() []byte {
	if m != nil {
		return m.Extension
	}
	return nil
}

// ChaincodeHeaderExtension is the Header's extentions message to be used when
// the Header's type is CHAINCODE.  This extensions is used to specify which
// chaincode to invoke and what should appear on the ledger.
type ChaincodeHeaderExtension struct {
	// The ID of the chaincode to target.
	ChaincodeId          *ChaincodeID `protobuf:"bytes,2,opt,name=chaincode_id,json=chaincodeId,proto3" json:"chaincode_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ChaincodeHeaderExtension) Reset() { *m = ChaincodeHeaderExtension{} }

// Nesseceray methods of ChaincodeHeaderExtension
func (m *ChaincodeHeaderExtension) GetChaincodeId() *ChaincodeID {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}

// ChaincodeProposalPayload is the payload of the proposal message.  It contains
// the input arguments for the chaincode invocation.  It is used by the
// endorser to invoke the chaincode.
type ChaincodeProposalPayload struct {
	// The input arguments for the chaincode invocation.  This field is
	// optional and can be omitted.
	InputArgs            [][]byte `protobuf:"bytes,1,rep,name=input_args,json=inputArgs,proto3" json:"input_args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaincodeProposalPayload) Reset() { *m = ChaincodeProposalPayload{} }

// Nesseceray methods of ChaincodeProposalPayload
func (m *ChaincodeProposalPayload) GetInputArgs() [][]byte {
	if m != nil {
		return m.InputArgs
	}
	return nil
}

// ChaincodeAction is the payload of the proposal response message.  It
// contains the result of the chaincode invocation.  It is used by the
// endorser to return the result of the chaincode invocation.
type ChaincodeAction struct {
	// This field contains the read set and the write set produced by the
	// chaincode executing this invocation.
	Results []byte `protobuf:"bytes,1,opt,name=results,proto3" json:"results,omitempty"`
	// This field contains the event generated by the chaincode.
	// Only a single marshaled ChaincodeEvent is included.
	Events []byte `protobuf:"bytes,2,opt,name=events,proto3" json:"events,omitempty"`
	// This field contains the result of executing this invocation.
	Response *Response `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	// This field contains the ChaincodeID of executing this invocation. Endorser
	// will set it with the ChaincodeID called by endorser while simulating proposal.
	// Committer will validate the version matching with latest chaincode version.
	// Adding ChaincodeID to keep version opens up the possibility of multiple
	// ChaincodeAction per transaction.
	ChaincodeId          *ChaincodeID `protobuf:"bytes,4,opt,name=chaincode_id,json=chaincodeId,proto3" json:"chaincode_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ChaincodeAction) Reset() { *m = ChaincodeAction{} }

// Nesseceray methods of ChaincodeAction
func (m *ChaincodeAction) GetResults() []byte {
	if m != nil {
		return m.Results
	}
	return nil
}
func (m *ChaincodeAction) GetEvents() []byte {
	if m != nil {
		return m.Events
	}
	return nil
}
func (m *ChaincodeAction) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}
func (m *ChaincodeAction) GetChaincodeId() *ChaincodeID {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}
