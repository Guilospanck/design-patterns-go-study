package implementations

type Person struct {
	// Personal details
	name, address, pin string
	// Job details
	workAddress, company, position string
	salary                         float64
}

type PersonBuilder struct {
	person *Person
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

// PersonBuilder chain methods
func (pb *PersonBuilder) Named(name string) *PersonBuilder {
	pb.person.name = name
	return pb
}

func (pb *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*pb}
}

// methods for PersonAddressBuilder
func (pab *PersonAddressBuilder) At(address string) *PersonAddressBuilder {
	pab.person.address = address
	return pab
}

func (pab *PersonAddressBuilder) WithPostalCode(pin string) *PersonAddressBuilder {
	pab.person.pin = pin
	return pab
}

func (pb *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*pb}
}

// methods for PersonJobBuilder
func (pjb *PersonJobBuilder) As(position string) *PersonJobBuilder {
	pjb.person.position = position
	return pjb
}

func (pjb *PersonJobBuilder) For(company string) *PersonJobBuilder {
	pjb.person.company = company
	return pjb
}

func (pjb *PersonJobBuilder) In(companyAddress string) *PersonJobBuilder {
	pjb.person.workAddress = companyAddress
	return pjb
}

func (pjb *PersonJobBuilder) WithSalary(salary float64) *PersonJobBuilder {
	pjb.person.salary = salary
	return pjb
}

// Reset person
func (pb *PersonBuilder) Reset() {
	pb.person = &Person{}
}

// Buikd a person from PersonBuilder
func (pb *PersonBuilder) Build() *Person {
	return pb.person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
}
