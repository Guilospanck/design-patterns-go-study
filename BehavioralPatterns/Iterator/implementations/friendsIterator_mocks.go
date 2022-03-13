package implementations

type FriendsIteratorMocks struct{}

var (
	HasNextCalledFriendsIterator = false
	GetNextCalledFriendsIterator = false
)

func (iterator *FriendsIteratorMocks) HasNext() bool {
	HasNextCalledFriendsIterator = true
	return true
}

func (iterator *FriendsIteratorMocks) GetNext() string {
	GetNextCalledFriendsIterator = true
	return ""
}

func NewFriendsIteratorMocks(friends []*Friend) *FriendsIteratorMocks {
	return &FriendsIteratorMocks{}
}
