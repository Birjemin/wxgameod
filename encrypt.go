package wxgameod

import (
	"github.com/birjemin/wxgameod/utils"
)

// Encrypt encrypt
func Encrypt(sessionKey, data string) string {
	return utils.GenerateSha256(sessionKey, data)
}
