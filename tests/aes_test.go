package encription_test

import (
	"testing"

	aes "github.com/avran02/aes-any-bytes"

	"github.com/stretchr/testify/assert"
)

var testsEncriptAndDecript = []struct {
	key  []byte
	data []byte
}{
	// 256 bit key
	{
		data: []byte("Some tandom text just to check if encription and decription works and some symbols $%@!^&*()[]{}"),
		key:  []byte("12345678901234567890123456789012"),
	}, {
		key:  []byte("12345678901234567890123456789012"),
		data: []byte("test"),
	}, {
		key:  []byte("12345678901234567890123456789012"),
		data: []byte("1234567890123456"),
	}, {
		key:  []byte("12345678901234567890123456789012"),
		data: []byte("фср8934а8рофцвгшапй"),
	}, {
		key:  []byte("12345678901234567890123456789012"),
		data: []byte("dui12hf7hqw8einfu91243gf9hqwiehjfu9gqh9q8whefu7g123"),
	},
	// 192 bit key
	{
		data: []byte("Some tandom text just to check if encription and decription works and some symbols $%@!^&*()[]{}"),
		key:  []byte("123456789012345678901234"),
	},
	{
		data: []byte("s"),
		key:  []byte("123456789012345678901234"),
	},
	{
		data: []byte("Какой-то текст на русском языке"),
		key:  []byte("123456789012345678901234"),
	},
	{
		data: []byte("s"),
		key:  []byte("123456789012345678901234"),
	},
	// 128 bit key
	{
		data: []byte("Some tandom text just to check if encription and decription works and some symbols $%@!^&*()[]{}"),
		key:  []byte("1234567890123456"),
	},
	{
		data: []byte("s"),
		key:  []byte("1234567890123456"),
	},
	{
		data: []byte("Какой-то текст на русском языке"),
		key:  []byte("1234567890123456"),
	},
}

func TestEncriptAndDecript(t *testing.T) {
	_, err := aes.Encript([]byte(""), []byte("12345678901234567890123456789012"))
	assert.Error(t, err)
	for _, test := range testsEncriptAndDecript {
		t.Run("", func(t *testing.T) {
			assert := assert.New(t)
			encriptedData, err := aes.Encript(test.data, test.key)
			assert.NoError(err)
			decriptedData, err := aes.Decript(encriptedData, test.key)
			assert.NoError(err)
			assert.Equal(test.data, decriptedData)
		})
	}
}
