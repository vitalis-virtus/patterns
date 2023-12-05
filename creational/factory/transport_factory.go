package factory

import "errors"

func GetTransport(transportType string) (Transporter, error) {
	switch transportType {
	case "car":
		return NewCar(), nil
	case "bicycle":
		return NewBicycle(), nil
	default:
		return nil, errors.New("wrong trasport type")
	}
}
