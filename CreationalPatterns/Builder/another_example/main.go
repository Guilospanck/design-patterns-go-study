package main

import "fmt"

func main() {
	opts := SensorOpts{
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

	tempSensorBuilder := NewSensorsBuilder("temp")
	rmmsSensorBuilder := NewSensorsBuilder("rmms")
	rms2SensorBuilder := NewSensorsBuilder("rms2")
	tiltSensorBuilder := NewSensorsBuilder("tilt")
	fftSensorBuilder := NewSensorsBuilder("fft")
	accrawSensorBuilder := NewSensorsBuilder("accraw")
	ntcSensorBuilder := NewSensorsBuilder("ntc")

	director := NewDirector(tempSensorBuilder)

	tempSensor := director.createSensor(opts)

	director.setBuilder(rmmsSensorBuilder)
	rmmsSensor := director.createSensor(opts)

	director.setBuilder(rms2SensorBuilder)
	rms2Sensor := director.createSensor(opts)

	director.setBuilder(tiltSensorBuilder)
	tiltSensor := director.createSensor(opts)

	director.setBuilder(fftSensorBuilder)
	fftSensor := director.createSensor(opts)

	director.setBuilder(accrawSensorBuilder)
	accrawSensor := director.createSensor(opts)

	director.setBuilder(ntcSensorBuilder)
	ntcSensor := director.createSensor(opts)

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
