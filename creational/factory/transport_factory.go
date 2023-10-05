package factory

import "errors"

func GetTransport(transportType string) (ITransport, error) {
	switch transportType {
	case "car":
		return NewCar(), nil
	case "bicycle":
		return NewBicycle(), nil
	default:
		return nil, errors.New("wrong trasport type")
	}
}
