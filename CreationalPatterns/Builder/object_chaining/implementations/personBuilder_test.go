package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PersonBuilder_Named(t *testing.T) {
	t.Run("should give person a name", func(t *testing.T) {
		// arrange
		sut := makePersonBuilderSut()

		// act
		result := sut.personBuilder.Named(sut.person.name)

		// assert
		assert.Equal(t, sut.person.name, sut.personBuilder.person.name)
		assert.Equal(t, sut.personBuilder, result)
	})
}

func Test_PersonBuilder_Lives(t *testing.T) {
	t.Run("should return a personAddressBuilder", func(t *testing.T) {
		// arrange
		sut := makePersonBuilderSut()

		// act
		result := sut.personBuilder.Lives()

		// assert
		assert.Equal(t, sut.personAddressBuilder, result)
	})
}

func Test_PersonAddressBuilder_At(t *testing.T) {
	t.Run("should return a personAddressBuilder and set address properly", func(t *testing.T) {
		// arrange
		sut := makePersonBuilderSut()

		// act
		result := sut.personAddressBuilder.At(sut.person.address)

		// assert
		assert.Equal(t, sut.person.address, sut.personAddressBuilder.person.address)
		assert.Equal(t, sut.personAddressBuilder, result)
	})
}

func Test_PersonAddressBuilder_WithPostalCode(t *testing.T) {
	t.Run("should return a personAddressBuilder and set pin properly", func(t *testing.T) {
		// arrange
		sut := makePersonBuilderSut()

		// act
		result := sut.personAddressBuilder.WithPostalCode(sut.person.pin)

		// assert
		assert.Equal(t, sut.person.pin, sut.personAddressBuilder.person.pin)
		assert.Equal(t, sut.personAddressBuilder, result)
	})
}

func Test_PersonBuilder_Works(t *testing.T) {
	t.Run("should return a personJobBuilder", func(t *testing.T) {
		// arrange
		sut := makePersonBuilderSut()

		// act
		result := sut.personBuilder.Works()

		// assert
		assert.Equal(t, sut.personJobBuilder, result)
	})
}

func Test_PersonJobBuilder_As(t *testing.T) {
	t.Run("should return a personJobBuilder and set position properly", func(t *testing.T) {
		// arrange
		sut := makePersonBuilderSut()

		// act
		result := sut.personJobBuilder.As(sut.person.position)

		// assert
		assert.Equal(t, sut.personJobBuilder, result)
		assert.Equal(t, sut.person.position, sut.personJobBuilder.person.position)
	})
}

func Test_PersonJobBuilder_For(t *testing.T) {
	t.Run("should return a personJobBuilder and set company properly", func(t *testing.T) {
		// arrange
		sut := makePersonBuilderSut()

		// act
		result := sut.personJobBuilder.For(sut.person.company)

		// assert
		assert.Equal(t, sut.personJobBuilder, result)
		assert.Equal(t, sut.person.company, sut.personJobBuilder.person.company)
	})
}

func Test_PersonJobBuilder_In(t *testing.T) {
	t.Run("should return a personJobBuilder and set workAddress properly", func(t *testing.T) {
		// arrange
		sut := makePersonBuilderSut()

		// act
		result := sut.personJobBuilder.In(sut.person.workAddress)

		// assert
		assert.Equal(t, sut.personJobBuilder, result)
		assert.Equal(t, sut.person.workAddress, sut.personJobBuilder.person.workAddress)
	})
}

func Test_PersonJobBuilder_WithSalary(t *testing.T) {
	t.Run("should return a personJobBuilder and set salary properly", func(t *testing.T) {
		// arrange
		sut := makePersonBuilderSut()

		// act
		result := sut.personJobBuilder.WithSalary(sut.person.salary)

		// assert
		assert.Equal(t, sut.personJobBuilder, result)
		assert.Equal(t, sut.person.salary, sut.personJobBuilder.person.salary)
	})
}

func Test_PersonBuilder_Reset(t *testing.T) {
	t.Run("should reset person", func(t *testing.T) {
		// arrange
		sut := makePersonBuilderSut()
		sut.personBuilder.person.name = sut.person.name

		// act
		assert.Equal(t, sut.person.name, sut.personBuilder.person.name)
		sut.personBuilder.Reset()

		// assert
		assert.Equal(t, sut.emptyPerson, sut.personBuilder.person)
	})
}

func Test_PersonBuilder_Build(t *testing.T) {
	t.Run("should reset person", func(t *testing.T) {
		// arrange
		sut := makePersonBuilderSut()
		sut.personBuilder.person.name = sut.person.name
		sut.personBuilder.person.address = sut.person.address

		// act
		result := sut.personBuilder.Build()

		// assert
		assert.Equal(t, sut.person.name, result.name)
		assert.Equal(t, sut.person.address, result.address)
	})
}

type personBuilderSut struct {
	personBuilder        *PersonBuilder
	personAddressBuilder *PersonAddressBuilder
	personJobBuilder     *PersonJobBuilder
	person               *Person
	emptyPerson          *Person
}

func makePersonBuilderSut() personBuilderSut {
	personBuilder := NewPersonBuilder()
	personAddressBuilder := &PersonAddressBuilder{*personBuilder}
	personJobBuilder := &PersonJobBuilder{*personBuilder}

	person := &Person{
		name:        "Guilherme",
		address:     "person address 48",
		pin:         "pin of person",
		workAddress: "person work address 22",
		company:     "Company of person",
		position:    "Dev",
		salary:      100.85,
	}

	emptyPerson := &Person{}

	return personBuilderSut{
		personBuilder,
		personAddressBuilder,
		personJobBuilder,
		person,
		emptyPerson,
	}
}
