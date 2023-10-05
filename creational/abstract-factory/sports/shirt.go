package sports

type IShirt interface {
	SetColor(color string)
	GetColor() string
	SetMaterial(material string)
	GetMaterial() string
}

type Shirt struct {
	color    string
	material string
}

func (s *Shirt) SetColor(color string) {
	s.color = color
}

func (s *Shirt) GetColor() string {
	return s.color
}

func (s *Shirt) SetMaterial(material string) {
	s.material = material
}

func (s *Shirt) GetMaterial() string {
	return s.material
}
