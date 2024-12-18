package helper

import "encoding/base64"

func EncodeBase64URLSafe(input []byte) string {
	return base64.URLEncoding.EncodeToString(input)
}
