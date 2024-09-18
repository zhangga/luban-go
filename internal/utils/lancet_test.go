package utils_test

import (
	"github.com/duke-git/lancet/v2/stream"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestLancet_Stream(t *testing.T) {
	result := stream.Of(1, 2, 3, 4, 5).Filter(func(i int) bool {
		return i%2 == 0
	}).Map(func(i int) int {
		return i * 2
	}).ToSlice()

	assert.Equal(t, result, []int{4, 8})
}
