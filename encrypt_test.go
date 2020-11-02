package wxgameod

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestEncrypt encrypt
func TestEncrypt(t *testing.T) {

	ast := assert.New(t)
	data := "{\"kv_list\":[{\"key\":\"1\",\"value\":\"0\"}]}"
	sessionKey := "tiihtNczf5v6AKRyjwEUhQ=="
	ret:= Encrypt(sessionKey, data)
	ast.Equal("f6887a3dc76b1ff66a0c7309d0b5a11fb336d54869823c45ecc374e65a249034", ret)
}
