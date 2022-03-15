package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Single_Instance(t *testing.T) {
	t.Run("should get only one instance", func(t *testing.T) {
		// arrange
		instance := makeSut()

		// act
		result := GetInstance()
		result2 := GetInstance()
		result3 := GetInstance()

		// assert
		assert.Equal(t, result, instance.instance)
		assert.Equal(t, result2, instance.instance)
		assert.Equal(t, result3, instance.instance)
	})
}

type singleSut struct {
	instance *single
}

func makeSut() singleSut {
	instance := GetInstance()

	return singleSut{instance}
}
