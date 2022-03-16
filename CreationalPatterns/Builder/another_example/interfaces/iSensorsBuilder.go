package interfaces

import "base/CreationalPatterns/Builder/another_example/domain"

type ISensorsBuilder interface {
	Reset()
	SetThreshold(int)
	SetHigher(bool)
	SetLower(bool)
	SetTSample(uint)
	SetBorder(bool)
	SetRisingEdge(bool)
	SetFallingEdge(bool)
	SetBetween(bool)
	SetFirstThreshold(int)
	SetSecondThreshold(int)
	SetType(uint)
	SetSocialContact(string)
	GetSensor() domain.Sensor
}
