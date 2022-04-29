package service

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

func TestSha1Encrypt(t *testing.T) {
	at := assert.New(t)
	signature := "45c4b51cb74856496d9fef687120f702e028ffd3"
	timestamp := "12345678"
	nonce := "12345678"
	//echostr := "2sdsdsfdd"
	tmpArray := []string{timestamp, nonce, "lax4832"}
	sort.Strings(tmpArray)
	h := sha1.New()
	h.Write([]byte(strings.Join(tmpArray, "")))
	calcVal := hex.EncodeToString(h.Sum(nil))
	at.Equal(calcVal, signature)
}
