package dpt

import (
	"encoding/hex"
	"strings"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestPackString(t *testing.T) {
	var testString = "KNX is OK"
	var testStringByteReference = "4B4E58206973204F4B0000000000"

	// test packing
	var packedString = packString(testString)

	assert.Equal(t, testStringByteReference, strings.ToUpper(hex.EncodeToString(packedString)))
}

func TestUnpackString(t *testing.T) {
	var testString = "KNX is OK"
	var testStringByteReference = []byte("4B4E58206973204F4B0000000000")

	// test unpacking
	var unpackedString = ""
	var err = unpackString(testStringByteReference, &unpackedString)

	assert.NoError(t, err)
	assert.Equal(t, testString, unpackedString)
}
