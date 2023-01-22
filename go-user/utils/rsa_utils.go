package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

type RSAUtils struct {
}

// bits = 4096 or sth...
func GenerateKey(bits int) (*rsa.PublicKey, *rsa.PrivateKey) {

	// The GenerateKey method takes in a reader that returns random bits, and the number of bits
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Fatalln("[FATAL]: RSA generate Key - *s", err)
		panic(err)
	}

	// Get public key...
	// The public key is a part of the *rsa.PrivateKey struct
	publicKey := privateKey.PublicKey

	return &publicKey, privateKey
}

func Encryption(pubKey *rsa.PublicKey, plainText string) string {

	cipherText, err := rsa.EncryptOAEP(
		sha512.New(),
		rand.Reader,
		pubKey,
		[]byte(plainText),
		nil)
	if err != nil {
		log.Fatalln("[FATAL]: RSA encryption failed - *s", err)
		panic(err)
	}

	return string(cipherText)
}

func Decryption(priKey *rsa.PrivateKey, cipher []byte) string {
	// The first argument is an optional random data generator (the rand.Reader we used before)
	// we can set this value as nil
	// The OEAPOptions in the end signify that we encrypted the data using OEAP, and that we used
	// SHA256 to hash the input.
	plainText, err := priKey.Decrypt(nil, cipher, &rsa.OAEPOptions{Hash: crypto.SHA512})
	if err != nil {
		log.Fatalln("[FATAL]: RSA decryption failed - *s", err)
		panic(err)
	}
	return string(plainText)
}

/**
 * 	SaveRSAPubKeyToFile
 */
func SaveRSAPubKeyToFile(path_with_name string, publicKey *rsa.PublicKey) {
	/**
	 * 	Check Mandatory Parameters
	 */

	if publicKey == nil {
		log.Println("publicKey is missing!")
		panic("publicKey is missing!")
	}

	if !(len(path_with_name)>0) {
		log.Println("file path is missing!")
		panic("file path is missing!")
	}

	// Serialize public key to ASN.1(DER) by x509 std.
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		log.Fatalln("[FATAL]: Failed to serialize pubKey to X509 std. - *s", err)
		panic(err)
	}

	// Create file in this path
	keyFile, err := os.Create(path_with_name)
	if err != nil {
		log.Fatalln("[FATAL]: Failed to create file to store pubKey. - *s", err)
		panic(err)
	}

	defer keyFile.Close()
	// Create an struct obj named pem.Block
	privateBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	// Save keys to files
	err = pem.Encode(keyFile, &privateBlock)

	if err != nil {
		log.Fatalln("[FATAL]: Save Public File Failed - *s", err)
		panic(err)
	}
}

/**
 * 	SaveRSAPubKeyToFile
 */
func SaveRSAPriKeyToFile(path_with_name string, privateKey *rsa.PrivateKey) {
	/**
	 * 	Check Mandatory Parameters
	 */

	if privateKey == nil {
		log.Println("privateKey is missing!")
		panic("privateKey is missing!")
	}

	if !(len(path_with_name)>0) {
		log.Println("file path is missing!")
		panic("file path is missing!")
	}

	// Serialize private key to ASN.1(DER) by x509 std.
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)

	// Create file in this path
	keyFile, err := os.Create(path_with_name)
	if err != nil {
		log.Fatalln("[FATAL]: Failed to create file to store priKey. - *s", err)
		panic(err)
	}

	defer keyFile.Close()
	// Create an struct obj named pem.Block
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	// Save keys to files
	err = pem.Encode(keyFile, &privateBlock)

	if err != nil {
		log.Fatalln("[FATAL]: Save PrivateKey File Failed - *s", err)
		panic(err)
	}
}

func GetRSAPubKeyFromFile(path_with_name string) rsa.PublicKey {

	if !(len(path_with_name)>0) {
		log.Println("file path is missing!")
		panic("file path is missing!")
	}
	// Open file
	file, err := os.Open(path_with_name)
	if err != nil {
		log.Fatalln("[FATAL]: Failed to open public key - *s", err)
		panic(err)
	}
	defer file.Close()
	// Read file content
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	// decode pem
	block, _ := pem.Decode(buf)
	// deserialize x509 key to rsa.PublicKey
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatalln("[FATAL]: Failed to deserialize pubKey to X509 std. - *s", err)
		panic(err)
	}
	key := publicKeyInterface.(*rsa.PublicKey)
	return *key
}

func GetRSAPriKeyFromFile(path_with_name string) rsa.PrivateKey {

	if !(len(path_with_name)>0) {
		log.Println("file path is missing!")
		panic("file path is missing!")
	}
	// Open file
	file, err := os.Open(path_with_name)
	if err != nil {
		log.Fatalln("[FATAL]: Failed to open private key - *s", err)
		panic(err)
	}
	defer file.Close()
	// Read file content
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	// decode pem
	block, _ := pem.Decode(buf)
	// deserialize x509 key to rsa.PublicKey
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalln("[FATAL]: Failed to deserialize priKey to X509 std. - *s", err)
		panic(err)
	}
	return *privateKey
}

func Sign(content *string, privateKey *rsa.PrivateKey) string {

	contentHash := sha256.New()
	_, err := contentHash.Write([]byte(*content))
	if err != nil {
		log.Fatalln("[FATAL]: Failed to deserialize pubKey to X509 std. - *s", err)
		panic(err)
	}
	contentHashSum := contentHash.Sum(nil)
	// In order to generate the signature, SignPSS needs a random number generator for salt,
	// private key, the hashing algorithm used, and the hash sum
	// of the content
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA512, contentHashSum, nil)
	if err != nil {
		log.Fatalln(err)
	}

	return string(signature)
}

/**
 * content here is the one to be verify with out signature attribute
 */
func Verify(publicKey *rsa.PublicKey, signature string, content *string) bool {

	contentHash := sha256.New()
	_, err := contentHash.Write([]byte(*content))
	if err != nil {
		log.Fatalln("[FATAL]: Failed to deserialize pubKey to X509 std. - *s", err)
		panic(err)
	}
	contentHashSum := contentHash.Sum(nil)

	// To verify the signature, VerifyPSS needs the public key, the hashing algorithm
	// the hash sum of content and the signature
	// there is an optional "options" parameter which can omit for now
	err_verify := rsa.VerifyPSS(publicKey, crypto.SHA512, contentHashSum, []byte(signature), nil)
	if err_verify != nil {
		log.Println("[FAILED]: Invalid Signature - *s", err)
		return false
	}
	return true
}
