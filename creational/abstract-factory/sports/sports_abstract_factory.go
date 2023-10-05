package sports

import "errors"

type Sports interface {
	CreateShoe() IShoe
	CreateShirt() IShirt
}

func GetSports(brand string) (Sports, error) {
	switch brand {
	case "adidas":
		return &AdidasSports{}, nil
	case "nike":
		return &NikeSports{}, nil
	default:
		return nil, errors.New("no such brand")
	}
}
