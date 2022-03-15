package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Builder_GetBuilder(t *testing.T) {
	t.Run("Should return proper normal builder", func(t *testing.T) {
		// arrange
		sut := makeBuilderSut()

		// act
		result := GetBuilder("normal")

		// assert
		assert.Equal(t, sut.normalBuilder, result)
	})

	t.Run("Should return proper igloo builder", func(t *testing.T) {
		// arrange
		sut := makeBuilderSut()

		// act
		result := GetBuilder("igloo")

		// assert
		assert.Equal(t, sut.iglooBuilder, result)
	})
}

type builderSut struct {
	iglooBuilder  *iglooBuilder
	normalBuilder *normalBuilder
}

func makeBuilderSut() builderSut {
	iglooBuilder := NewIglooBuilder()
	normalBuilder := NewNormalBuilder()

	return builderSut{iglooBuilder, normalBuilder}
}
