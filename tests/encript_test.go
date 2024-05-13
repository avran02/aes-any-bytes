package encription_test

import (
	"testing"

	aes "github.com/avran02/aes-any-bytes"

	"github.com/stretchr/testify/assert"
)

var testsEncript = []struct {
	data     []byte
	key      []byte
	excepted []byte
}{
	{
		data:     []byte{204, 43, 81, 3, 72, 70, 89, 231, 250, 161, 123, 139, 141, 140, 189, 70},
		key:      []byte{49, 50, 51, 52, 53, 54, 55, 56, 57, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 48, 49, 50},
		excepted: []byte{126, 99, 209, 219, 100, 157, 185, 164, 220, 234, 160, 129, 127, 56, 87, 38, 65, 152, 171, 64, 129, 72, 124, 33, 58, 28, 130, 188, 64, 74, 215, 200},
	},
	{
		data:     []byte{48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48},
		key:      []byte{110, 99, 111, 100, 57, 50, 55, 209, 128, 208, 178, 208, 189, 49, 54, 208, 178, 208, 190, 209, 131, 208, 179, 208, 186, 208, 187, 55, 121, 104, 102, 54},
		excepted: []byte{249, 140, 62, 147, 15, 128, 32, 152, 223, 255, 109, 247, 97, 166, 89, 194, 249, 140, 62, 147, 15, 128, 32, 152, 223, 255, 109, 247, 97, 166, 89, 194, 65, 56, 217, 220, 215, 120, 194, 108, 182, 26, 10, 150, 93, 153, 201, 113},
	},
}

func TestEncript(t *testing.T) {
	_, err := aes.Encript([]byte(""), []byte("12345678901234567890123456789012"))
	assert.Error(t, err)
	for _, test := range testsEncript {
		t.Run("", func(t *testing.T) {
			assert := assert.New(t)
			encriptedData, err := aes.Encript(test.data, test.key)
			assert.NoError(err)
			assert.Equal(test.excepted, encriptedData)
		})
	}
}
