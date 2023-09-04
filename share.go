package doppler

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/pbkdf2"
)

const (
	algorithm     = "aes-256-gcm"
	keyLength     = 256 / 8
	ivLength      = 12
	saltLength    = 16
	saltRounds    = 100000
	authTagLength = 16
)

type PlainTextArgs struct {
	Secret      string `json:"secret"`
	ExpireViews int32  `json:"expire_views"`
	ExpireDays  int32  `json:"expire_days"`
}
type PlainTextResp struct {
	Url              string `json:"url"`
	AuthenticatedUrl string `json:"authenticated_url"`
	Password         string `json:"password"`
	Success          bool   `json:"success"`
}

/*
Generate a Doppler Share link by sending a plain text secret. This endpoint is not end-to-end encrypted as you are sending the secret in plain text. At no point do we store the plain text secret or the password in our systems. The receive flow the user goes through will be end-to-end encrypted where the encrypted secret will be decrypted on the browser.
*/
func (dp *doppler) SharePlainTextSecret(args PlainTextArgs) (data *PlainTextResp, err error) {

	payload, err := json.Marshal(args)
	if err != nil {
		return
	}

	req, _ := http.NewRequest(
		http.MethodPost,
		"/v1/share/secrets/plain",
		bytes.NewReader(payload))

	body, err := dp.makeApiRequest(req)
	if err != nil {
		return
	}

	data = new(PlainTextResp)
	err = json.Unmarshal(body, &data)

	return
}

type EncryptedSecretArgs struct {
	Secret      string `json:"secret"`
	ExpireViews int32  `json:"expire_views"`
	ExpireDays  int32  `json:"expire_days"`
}

/*
Generate a Doppler Share link by sending an encrypted secret. The receive flow the user goes through will be end-to-end encrypted where the encrypted secret will be decrypted on the browser.
*/
func (dp *doppler) ShareEncryptedSecret(args interface{}) (data interface{}, err error) {
	payload, err := json.Marshal(args)
	if err != nil {
		return
	}

	req, _ := http.NewRequest(
		http.MethodPost,
		"/v1/share/secrets/encrypted",
		bytes.NewReader(payload))

	body, err := dp.makeApiRequest(req)
	if err != nil {
		return
	}

	data = new(interface{})
	err = json.Unmarshal(body, &data)

	return
}

/*
Go implementation of the Encryption Example here: https://docs.doppler.com/reference/share-secret-encrypted
*/
func EncryptSecret(secret string) (map[string]interface{}, error) {
	password := make([]byte, 32)
	_, err := rand.Read(password)
	if err != nil {
		return nil, err
	}
	passwordHex := fmt.Sprintf("%x", password)
	hashedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(passwordHex)))

	salt := make([]byte, saltLength)
	_, err = rand.Read(salt)
	if err != nil {
		return nil, err
	}
	key := pbkdf2.Key([]byte(passwordHex), salt, saltRounds, keyLength, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, ivLength)
	_, err = rand.Read(iv)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	encryptedData := aesgcm.Seal(nil, iv, []byte(secret), nil)
	encrypted_secret := append(append(salt, iv...), encryptedData...)

	return map[string]interface{}{
			"encrypted_secret":       base64.StdEncoding.EncodeToString(encrypted_secret),
			"password":               passwordHex,
			"hashed_password":        hashedPassword,
			"expire_views":           1,
			"expire_days":            1,
			"encryption_kdf":         "pbkdf2",
			"encryption_salt_rounds": 1000000},
		nil
}
