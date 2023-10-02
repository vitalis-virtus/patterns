package factory

type Car struct {
	Transport
}

func NewCar() ITransport {
	return &Car{
		Transport{
			name:   "car",
			wheels: 4,
		},
	}
}
