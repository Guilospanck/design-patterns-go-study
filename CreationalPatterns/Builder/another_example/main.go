package main

import (
	"base/CreationalPatterns/Builder/another_example/domain"
	"base/CreationalPatterns/Builder/another_example/implementations"
	"fmt"
)

func main() {
	opts := domain.SensorOpts{
		Threshold:       10,
		Higher:          true,
		Lower:           false,
		TSample:         60,
		Border:          false,
		RisingEdge:      false,
		FallingEdge:     false,
		Between:         false,
		FirstThreshold:  0,
		SecondThreshold: 0,
		Type:            1,
		SocialContact:   "test@test.com",
	}

	tempSensorBuilder := implementations.NewSensorsBuilder("temp")
	rmmsSensorBuilder := implementations.NewSensorsBuilder("rmms")
	rms2SensorBuilder := implementations.NewSensorsBuilder("rms2")
	tiltSensorBuilder := implementations.NewSensorsBuilder("tilt")
	fftSensorBuilder := implementations.NewSensorsBuilder("fft")
	accrawSensorBuilder := implementations.NewSensorsBuilder("accraw")
	ntcSensorBuilder := implementations.NewSensorsBuilder("ntc")

	director := implementations.NewDirector(tempSensorBuilder)

	tempSensor := director.CreateSensor(opts)

	director.SetBuilder(rmmsSensorBuilder)
	rmmsSensor := director.CreateSensor(opts)

	director.SetBuilder(rms2SensorBuilder)
	rms2Sensor := director.CreateSensor(opts)

	director.SetBuilder(tiltSensorBuilder)
	tiltSensor := director.CreateSensor(opts)

	director.SetBuilder(fftSensorBuilder)
	fftSensor := director.CreateSensor(opts)

	director.SetBuilder(accrawSensorBuilder)
	accrawSensor := director.CreateSensor(opts)

	director.SetBuilder(ntcSensorBuilder)
	ntcSensor := director.CreateSensor(opts)

	for key, value := range tempSensor {
		fmt.Printf("%s: %+v\n===========================\n", key, value)
	}
	for key, value := range rmmsSensor {
		fmt.Printf("%s: %+v\n===========================\n", key, value)
	}
	for key, value := range rms2Sensor {
		fmt.Printf("%s: %+v\n===========================\n", key, value)
	}
	for key, value := range tiltSensor {
		fmt.Printf("%s: %+v\n===========================\n", key, value)
	}
	for key, value := range fftSensor {
		fmt.Printf("%s: %+v\n===========================\n", key, value)
	}
	for key, value := range accrawSensor {
		fmt.Printf("%s: %+v\n===========================\n", key, value)
	}
	for key, value := range ntcSensor {
		fmt.Printf("%s: %+v\n===========================\n", key, value)
	}
}
