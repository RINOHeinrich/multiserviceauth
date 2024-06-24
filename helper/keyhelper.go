package helper

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type KeyManager struct {
	PublicKey  *ecdsa.PublicKey
	PrivateKey *ecdsa.PrivateKey
}

func (k *KeyManager) SavePrivateToDisk(filename string) error {
	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	// Save the private key to the file
	keyBytes, err := x509.MarshalECPrivateKey(k.PrivateKey)
	if err != nil {
		return err
	}
	err = pem.Encode(outFile, &pem.Block{Type: "PRIVATE KEY", Bytes: keyBytes})
	if err != nil {
		return err
	}
	return nil
}
func (k *KeyManager) SavePublicToDisk(filename string) error {
	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	// Save the public key to the file
	keyBytes, err := x509.MarshalPKIXPublicKey(k.PublicKey)
	if err != nil {
		return err
	}
	err = pem.Encode(outFile, &pem.Block{Type: "PUBLIC KEY", Bytes: keyBytes})
	if err != nil {
		return err
	}
	return nil
}
func (k *KeyManager) LoadPrivateKey(filename string) error {

	inFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer inFile.Close()
	// Load the private key from the file
	fileInfo, _ := inFile.Stat()
	buffer := make([]byte, fileInfo.Size())
	_, err = inFile.Read(buffer)
	if err != nil {
		return err
	}
	data, _ := pem.Decode(buffer)
	k.PrivateKey, err = x509.ParseECPrivateKey(data.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
func (k *KeyManager) LoadPublicKey(filename string) error {

	inFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer inFile.Close()
	// Load the public key from the file
	fileInfo, _ := inFile.Stat()
	buffer := make([]byte, fileInfo.Size())
	_, err = inFile.Read(buffer)
	if err != nil {
		return err
	}
	data, _ := pem.Decode(buffer)
	keys, err := x509.ParsePKIXPublicKey(data.Bytes)
	if err != nil {
		return err
	}
	k.PublicKey = keys.(*ecdsa.PublicKey)
	return nil
}
func (k *KeyManager) GenerateKeys() {
	// Generate the private key
	Keys, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	k.PrivateKey = Keys
	k.PublicKey = &Keys.PublicKey
}
func (k *KeyManager) CheckKeys(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
func (k *KeyManager) LoadConfig(conffilepath string) error {
	err := godotenv.Load(conffilepath)
	if err != nil {
		log.Default().Println(err)
	}
	privatekeypath := os.Getenv("PRIVATE_KEY")
	publickeypath := os.Getenv("PUBLIC_KEY")
	if !k.CheckKeys(privatekeypath) {
		k.GenerateKeys()
		err := k.SavePrivateToDisk(privatekeypath)
		if err != nil {
			return err
		}
		err = k.SavePublicToDisk(publickeypath)
		if err != nil {
			return err
		}
	}
	err = k.LoadPrivateKey(privatekeypath)
	if err != nil {
		return err
	}
	err = k.LoadPublicKey(publickeypath)
	if err != nil {
		return err
	}
	return nil
}
