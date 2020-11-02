package wxgameod

import (
	"wxgameod/utils"
)

// Encrypt encrypt
func Encrypt(sessionKey, data string) string {
	return utils.GenerateSha256(sessionKey, data)
}
