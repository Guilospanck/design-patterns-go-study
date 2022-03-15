package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Single_Instance(t *testing.T) {
	t.Run("should create only one instance", func(t *testing.T) {
		// arrange
		sut := makeSut()

		// act
		result := GetInstance()
		result2 := GetInstance()
		result3 := GetInstance()

		// assert
		assert.Equal(t, sut.instance, result)
		assert.Equal(t, sut.instance, result2)
		assert.Equal(t, sut.instance, result3)
	})
}

type singleSut struct {
	instance *single
}

func makeSut() singleSut {
	instance := GetInstance()

	return singleSut{instance}
}
