package identicon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRand(t *testing.T) {

	t.Run("Rand", func(t *testing.T) {
		assert := assert.New(t)
		r := Rand([]byte{'1', '2', '3', '4', '5', '6', '7', '8'})
		assert.NotNil(r)
		assert.EqualValues(0, r.Intn(1))
	})
}

func TestData(t *testing.T) {

	t.Run("NewData", func(t *testing.T) {
		assert := assert.New(t)
		d := NewData([]byte{'a'})
		assert.NotNil(d)
		assert.True(len(d.hash) > 0)
	})

	t.Run("NewDataString", func(t *testing.T) {
		assert := assert.New(t)
		d := NewDataString("a")
		assert.NotNil(d)
		assert.True(len(d.hash) > 0)
	})

	t.Run("Color", func(t *testing.T) {
		assert := assert.New(t)
		d := NewDataString("a")
		assert.NotNil(d.Color())
	})
}
