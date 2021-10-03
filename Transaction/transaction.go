package transaction

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/gob"
	"encoding/hex"
	"log"
	"math/big"
	"time"
)

// Create Transaction structure for blockchain
type Transaction struct {
	Header []byte
	Post   []byte
}

// Method of transaction to getHeader
func (tx *Transaction) GetHeader() []byte {
	return tx.Header
}

// Method of transaction to getPost
func (tx *Transaction) GetPost() []byte {
	return tx.Post
}

// Method of transaction to setHeader
func (tx *Transaction) SetHeader(header []byte) {
	tx.Header = header
}

// Method of transaction to setPost
func (tx *Transaction) SetPost(post []byte) {
	tx.Post = post
}

// Method of transaction to Serialize using gob
func (tx *Transaction) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}

// Method of transaction to Deserialize using gob
func DeserializeTransaction(data []byte) Transaction {
	var transaction Transaction

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&transaction)
	if err != nil {
		log.Panic(err)
	}

	return transaction
}

// Structure of transaction Header
type Header struct {
	Type      []byte
	Timestamp time.Time
	UserType  []byte
	Pubkey    string
}

// Method of Header to get Heade fields

func (h *Header) GetType() []byte {
	return h.Type
}

func (h *Header) GetTimestamp() time.Time {
	return h.Timestamp
}

func (h *Header) GetUserType() []byte {
	return h.UserType
}

// // Method of Header to get PubKey
func (h *Header) GetPubkey() string {
	return h.Pubkey
}

// Method of Header to set Header fields

func (h *Header) SetType(type1 []byte) {
	h.Type = type1
}

func (h *Header) SetTimestamp(timestamp time.Time) {
	h.Timestamp = timestamp
}

func (h *Header) SetUserType(userType []byte) {
	h.UserType = userType
}

// Method of Header to Serialize using gob
func (h *Header) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(h)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}

// Method of Header to
func DeserializeHeader(data []byte) Header {
	var header Header

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&header)
	if err != nil {
		log.Panic(err)
	}

	return header
}

// Structure of transaction Post
type Post struct {
	PostID      []byte
	PostData    []byte
	Signture    string
	PrivateData []byte
}

// Method of Post to get Post fields

func (p *Post) GetPostID() []byte {
	return p.PostID
}

func (p *Post) GetPostData() []byte {
	return p.PostData
}

func (p *Post) GetSignture() string {
	return p.Signture
}

func (p *Post) GetPrivateData() []byte {
	return p.PrivateData
}

// Method of Post to set Post fields

func (p *Post) SetPostID(postID []byte) {
	p.PostID = postID
}

func (p *Post) SetPostData(postData []byte) {
	p.PostData = postData
}

func (p *Post) SetSignture(signture string) {
	p.Signture = signture
}

func (p *Post) SetPrivateData(privateData []byte) {
	p.PrivateData = privateData
}

// Method of Post to Serialize using gob
func (p *Post) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(p)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}

// Method of Post to Deserialize using gob
func DeserializePost(data []byte) Post {
	var post Post

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&post)
	if err != nil {
		log.Panic(err)
	}

	return post
}

// Method of post that takes ecdsa privateKey, signs postData and returns signature, ecdsa pubkey and postID
func (p *Post) SignPost(privateKey *ecdsa.PrivateKey) (string, []byte, []byte) {
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, p.GetPostData())
	if err != nil {
		log.Panic(err)
	}

	signature := append(r.Bytes(), s.Bytes()...)

	pubKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

	postID := append(p.GetPostID(), pubKey...)

	return hex.EncodeToString(signature), pubKey, postID
}

// Method of post that verifies the signature
func (p *Post) VerifyPost(signature string, pubKey []byte, postID []byte) bool {
	r := big.Int{}
	s := big.Int{}

	sigBytes := []byte(signature)

	r.SetBytes(sigBytes[:len(sigBytes)/2])
	s.SetBytes(sigBytes[len(sigBytes)/2:])

	x := big.Int{}
	y := big.Int{}
	keyLen := len(pubKey)
	x.SetBytes(pubKey[:keyLen/2])
	y.SetBytes(pubKey[keyLen/2:])

	rawPubKey := ecdsa.PublicKey{Curve: elliptic.P256(), X: &x, Y: &y}
	if !ecdsa.Verify(&rawPubKey, p.GetPostData(), &r, &s) {
		return false
	}

	if !bytes.Equal(postID, p.GetPostID()) {
		return false
	}

	return true
}
