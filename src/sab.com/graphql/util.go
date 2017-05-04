package graphql

import "encoding/base64"

func ComputeBase64(input string) (encoded string) {
	encoded = base64.StdEncoding.EncodeToString([]byte(input))
	return
}
