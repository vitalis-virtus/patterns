package factory

type Bicycle struct {
	Transport
}

func NewBicycle() ITransport {
	return &Bicycle{
		Transport{
			name:   "bicycle",
			wheels: 2,
		},
	}
}
