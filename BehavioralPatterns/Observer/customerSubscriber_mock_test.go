package main

type CustomerSubscriberMock struct {
	id           string
	updateCalled bool
}

func (sub *CustomerSubscriberMock) Update(data string) {
	sub.updateCalled = true
}

func (sub *CustomerSubscriberMock) GetID() string {
	return sub.id
}

func NewCustomerSubscriberMock(id string) *CustomerSubscriberMock {
	return &CustomerSubscriberMock{
		id:           id,
		updateCalled: false,
	}
}
