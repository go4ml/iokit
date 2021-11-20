package iokit

import (
	"crypto/rand"
	"sudachen.xyz/pkg/fu"
	"gotest.tools/assert"
	"testing"
)

func Test_Crypto(t *testing.T) {
	buf := make([]byte, 4096)
	rand.Read(buf)
	dat, err := fu.Encrypt("helloworld", buf)
	assert.NilError(t, err)
	dat, err = fu.Decrypt("helloworld", dat)
	assert.NilError(t, err)
	assert.DeepEqual(t, dat, buf)
}
