package main

type ISensorsBuilder interface {
	reset()
	setThreshold(int)
	setHigher(bool)
	setLower(bool)
	setTSample(uint)
	setBorder(bool)
	setRisingEdge(bool)
	setFallingEdge(bool)
	setBetween(bool)
	setFirstThreshold(int)
	setSecondThreshold(int)
	setType(uint)
	setSocialContact(string)
	getSensor() Sensor
}
