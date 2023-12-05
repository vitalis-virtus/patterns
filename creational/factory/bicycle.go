package factory

type Bicycle struct {
	Transport
}

func NewBicycle() Transporter {
	return &Bicycle{
		Transport{
			name:   "bicycle",
			wheels: 2,
		},
	}
}
