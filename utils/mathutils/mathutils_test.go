package mathutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormulaX3(t *testing.T) {
	assert := assert.New(t)
	args := []uint32{1, 2, 3, 4}
	assert.Equal(58, int(FormulaX3(args, 3)))
	assert.Equal(58, int(formulaX3_1Pow(args, 3)))
	assert.Equal(58, int(formulaX3_2Pow(args, 3)))

	args = []uint32{5, 6, 7, 8}
	assert.Equal(218, int(FormulaX3(args, 3)))
	assert.Equal(218, int(formulaX3_1Pow(args, 3)))
	assert.Equal(218, int(formulaX3_2Pow(args, 3)))

}

func BenchmarkFormulaX3(b *testing.B) {
	args := []uint32{1, 2, 3, 4}
	b.Run("FormulaX3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FormulaX3(args, 3)
		}
	})
	b.Run("FormulaX3-1Pow", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			formulaX3_1Pow(args, 3)
		}
	})

	b.Run("FormulaX3_2Pow", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			formulaX3_2Pow(args, 3)
		}
	})
}
