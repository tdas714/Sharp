package transaction

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"
)

type TxType int32

const (
	Transaction_POST   TxType = 0
	Transaction_CRITIC TxType = 1
)

// Structure of Transaction , Header, PostID, POstData, PostHash, Signature
type Transaction struct {
	Header          *TransactionHeader
	PostData        *TransactionPostData
	PostHash        string
	Signature       string
	TransactionHash string
}

// Structure of TransactionHeader
type TransactionHeader struct {
	TransactionType  *TxType
	TransactionTime  time.Time
	TransactionFrom  string
	TransactionTo    string
	CriticizedTxHash string
}

// Structure of TransactionPostData
type TransactionPostData struct {
	PostID   string
	PostData []byte
}

type Signature struct {
	R *big.Int
	S *big.Int
}

// Method to calculate hash of transactionpost data using crypto library
func (tp *TransactionPostData) HashData(time *time.Time) string {
	sum := sha512.Sum512(append(tp.PostData, []byte(time.String())...))
	s := fmt.Sprintf("%x", sum)
	tp.PostID = fmt.Sprintf("%x", sha256.Sum224(append([]byte(s), []byte(time.String())...)))
	return s[:]
}

// Method to calculate hash of transaction using crypto library
func (t *Transaction) Hash() string {
	sum := sha512.Sum512(append(t.Header.TransactionType.Serialize(), []byte(t.Header.TransactionTime.String())...))
	s := fmt.Sprintf("%x", sum)
	t.TransactionHash = fmt.Sprintf("%x", sha256.Sum224(append([]byte(s), []byte(t.PostHash)...)))
	return s[:]
}

// Function that signs Postdata using ecdsa library
func (t *Transaction) Sign(privateKey *ecdsa.PrivateKey) {
	r, s, _ := ecdsa.Sign(rand.Reader, privateKey, []byte(t.PostHash))
	// prepare a signature structure to marshal into json
	signature := &Signature{
		R: r,
		S: s,
	}
	// marshal to json
	signature_json, err := json.Marshal(signature)
	if err != nil {
		log.Fatal(err)
	}
	// encode to hex
	signature_hex := hex.EncodeToString(signature_json)
	t.Signature = signature_hex

}

// Function that verifies Postdata using ecdsa library
func (t *Transaction) Verify() bool {
	// hex decode signature and then demarshal
	signature_json, err := hex.DecodeString(t.Signature)
	if err != nil {
		log.Fatal(err)
	}
	var signature Signature
	err = json.Unmarshal(signature_json, &signature)
	if err != nil {
		log.Fatal(err)
	}
	pubKey := StringToKey(t.Header.TransactionFrom)
	// verify signature
	return ecdsa.Verify(&pubKey, []byte(t.PostHash), signature.R, signature.S)
}

// /==================================
// Method fo TxType to serialize it to byte array
func (t *TxType) Serialize() []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, t)
	return buf.Bytes()
}

// Method to deserialize byte array to TxType
func DeserializeTxType(b []byte) TxType {
	var txType TxType
	buf := bytes.NewReader(b)
	binary.Read(buf, binary.BigEndian, &txType)
	return txType
}

// Method to convert string to ecdsa.PublicKey
func StringToKey(s string) ecdsa.PublicKey {
	x, _ := new(big.Int).SetString(s[:64], 16)
	y, _ := new(big.Int).SetString(s[64:], 16)
	return ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}
}

// Method to convert ecdsa.Publickey to string
func KeyToString(key ecdsa.PublicKey) string {
	return fmt.Sprintf("%x%x", key.X, key.Y)
}

// /=====================================
// Method of creating new Transaction
func NewTransaction(privateKey *ecdsa.PrivateKey, from ecdsa.PublicKey, to ecdsa.PublicKey, postData []byte, txtype TxType, ctxhash string) *Transaction {
	var transactionTo string
	var criticizedTx string
	if txtype == Transaction_POST {
		transactionTo = ""
		criticizedTx = ""
	} else if txtype == Transaction_CRITIC {
		transactionTo = KeyToString(to)
		criticizedTx = ctxhash
	}
	// Create Transaction Header
	header := &TransactionHeader{
		TransactionType:  &txtype,
		TransactionTime:  time.Now(),
		TransactionFrom:  KeyToString(from),
		TransactionTo:    transactionTo,
		CriticizedTxHash: criticizedTx,
	}
	// Create Transaction PostData
	txPostData := &TransactionPostData{
		PostID:   "",
		PostData: postData,
	}
	// Calculate PostHash
	posthash := txPostData.HashData(&header.TransactionTime)
	// Create Transaction
	transaction := &Transaction{
		Header:          header,
		PostData:        txPostData,
		PostHash:        posthash,
		Signature:       "",
		TransactionHash: "",
	}
	transaction.Hash()
	transaction.Sign(privateKey)
	return transaction
}

// /===========================================
