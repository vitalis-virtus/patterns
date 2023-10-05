package sports

type NikeSports struct {
}

type NikeShoe struct {
	Shoe
}

type NikeShirt struct {
	Shirt
}

func (n *NikeSports) CreateShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: "41",
		},
	}
}

func (n *NikeSports) CreateShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			color:    "green",
			material: "silk",
		},
	}
}
