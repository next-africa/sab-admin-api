package util

import "encoding/base64"

func ComputeBase64(input string) (encoded string) {
	encoded = base64.StdEncoding.EncodeToString([]byte(input))
	return
}

func DecodeBase64(input string) (decoded string, err error) {
	var decodedByte []byte
	decodedByte, err = base64.StdEncoding.DecodeString(input)
	decoded = string(decodedByte)
	return

}
