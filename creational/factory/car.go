package factory

type Car struct {
	Transport
}

func NewCar() Transporter {
	return &Car{
		Transport{
			name:   "car",
			wheels: 4,
		},
	}
}
