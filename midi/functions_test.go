package midipackage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMod(t *testing.T) {
	t.Run("mod -> a is positive, rest is 0 ", func(t *testing.T) {
		got := mod(10, 5)
		want := 0

		assert.Equal(t, want, got)
	})
	t.Run("mod -> a is negative, rest is 0 ", func(t *testing.T) {
		got := mod(10, 5)
		want := 0

		assert.Equal(t, want, got)
	})
	t.Run("mod -> a is positive, rest is not 0 ", func(t *testing.T) {
		got := mod(10, 12)
		want := 10

		assert.Equal(t, want, got)
	})
	t.Run("mod -> a is negative, rest is not 0 ", func(t *testing.T) {
		got := mod(-10, 12)
		want := 2

		assert.Equal(t, want, got)
	})

}

func TestAbs(t *testing.T) {
	t.Run("< 0", func(t *testing.T) {
		got := abs(-10)
		want := 10

		assert.Equal(t, want, got)
	})
	t.Run("> 0", func(t *testing.T) {
		got := abs(10)
		want := 10

		assert.Equal(t, want, got)
	})
	t.Run("== 0", func(t *testing.T) {
		got := abs(0)
		want := 0

		assert.Equal(t, want, got)
	})
}

func TestNatural(t *testing.T) {
	t.Run("< 0", func(t *testing.T) {
		got := natural(-10)
		want := 10

		assert.Equal(t, want, got)
	})
	t.Run("> 0", func(t *testing.T) {
		got := natural(10)
		want := 10

		assert.Equal(t, want, got)
	})
	t.Run("== 0", func(t *testing.T) {
		got := natural(0)
		want := 1

		assert.Equal(t, want, got)
	})
}

func TestPutInRange(t *testing.T) {
	t.Run("min < x < max", func(t *testing.T) {
		got := putInRange(10, 0, 100)
		want := 10

		assert.Equal(t, want, got)
	})
	t.Run("x < min < max", func(t *testing.T) {
		got := putInRange(-10, 5, 100)
		want := 5

		assert.Equal(t, want, got)
	})
	t.Run("min < max < x", func(t *testing.T) {
		got := putInRange(10, 0, 6)
		want := 6

		assert.Equal(t, want, got)
	})
	t.Run("max < x < min", func(t *testing.T) {
		got := putInRange(10, 100, 0)
		want := 10

		assert.Equal(t, want, got)
	})
	t.Run("x < max < min", func(t *testing.T) {
		got := putInRange(-10, 100, 5)
		want := 5

		assert.Equal(t, want, got)
	})
	t.Run("max < min < x", func(t *testing.T) {
		got := putInRange(10, 6, 0)
		want := 6

		assert.Equal(t, want, got)
	})
}
