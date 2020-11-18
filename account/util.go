package account

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/golang/protobuf/proto"
	"github.com/jack-koli/tron-protocol/core"
	"github.com/sasaxie/go-client-api/common/base58"

	"golang.org/x/crypto/sha3"
	"math/big"
	log "github.com/sirupsen/logrus"
	"time"
)

func GenerateKey()  string {
	key, _ := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
	return fmt.Sprintf("%x", key.D.Bytes())
}

func EcdsaFromBase64(privKey string) (key *ecdsa.PrivateKey,err error) {
	key, err = crypto.HexToECDSA(privKey)
	return
}

func AddressFromEsdaPrivKey(key *ecdsa.PrivateKey) (addr []byte) {
	// #1
	pub := append(key.X.Bytes(), key.Y.Bytes()...)

	// #2
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pub)
	hashed := hash.Sum(nil)
	last20 := hashed[len(hashed)-20:]

	// #3
	addr41 := append([]byte{0x41}, last20...)

	// #4
	hash2561 := sha256.Sum256(addr41)
	hash2562 := sha256.Sum256(hash2561[:])
	checksum := hash2562[:4]

	// #5/#6
	rawAddr := append(addr41, checksum...)
	return rawAddr
}

func AddressFromPrv(privKeyStr string) string {
	keyBytes, _ := hex.DecodeString(privKeyStr)
	key := new(ecdsa.PrivateKey)
	key.PublicKey.Curve = btcec.S256()
	key.D = new(big.Int).SetBytes(keyBytes)
	key.PublicKey.X, key.PublicKey.Y = key.PublicKey.Curve.ScalarBaseMult(keyBytes)

	// #1
	pub := append(key.X.Bytes(), key.Y.Bytes()...)

	// #2
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pub)
	hashed := hash.Sum(nil)
	last20 := hashed[len(hashed)-20:]

	// #3
	addr41 := append([]byte{0x41}, last20...)

	// #4
	hash2561 := sha256.Sum256(addr41)
	hash2562 := sha256.Sum256(hash2561[:])
	checksum := hash2562[:4]

	// #5/#6
	rawAddr := append(addr41, checksum...)
	return base58.Encode(rawAddr)
}

const (
	AddressLength = 21
	AddressPrefix = "a0"
)

func PubkeyToAddress(p ecdsa.PublicKey) []byte {
	address := crypto.PubkeyToAddress(p)

	addressTron := make([]byte, AddressLength)

	addressPrefix, err := hexutil.Decode(AddressPrefix)
	if err != nil {
		log.Error(err.Error())
	}

	addressTron = append(addressTron, addressPrefix...)
	addressTron = append(addressTron, address.Bytes()...)

	return addressTron
}

func SignTransaction(transaction *core.Transaction, key *ecdsa.PrivateKey) (err error) {
	transaction.GetRawData().Timestamp = time.Now().UnixNano() / 1000000
	rawData, err := proto.Marshal(transaction.GetRawData())

	if err != nil {
		return
	}

	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)

	contractList := transaction.GetRawData().GetContract()

	for range contractList {
		signature, err := crypto.Sign(hash, key)
		if err != nil {
			return err
		}
		transaction.Signature = append(transaction.Signature, signature)
	}

	return
}
