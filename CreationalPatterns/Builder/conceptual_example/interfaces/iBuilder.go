package interfaces

import (
	"base/CreationalPatterns/Builder/conceptual_example/domain"
)

type IBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() domain.House
}
