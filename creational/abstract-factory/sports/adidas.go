package sports

type AdidasSports struct {
}

type AdidasShoe struct {
	Shoe
}

type AdidasShirt struct {
	Shirt
}

func (a *AdidasSports) CreateShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: "42",
		},
	}
}

func (a *AdidasSports) CreateShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			color:    "red",
			material: "cotton",
		},
	}
}
