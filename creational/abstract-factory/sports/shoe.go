package sports

type IShoe interface {
	SetLogo(logo string)
	GetLogo() string
	SetSize(size string)
	GetSize() string
}

type Shoe struct {
	logo string
	size string
}

func (s *Shoe) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) GetLogo() string {
	return s.logo
}

func (s *Shoe) SetSize(size string) {
	s.size = size
}

func (s *Shoe) GetSize() string {
	return s.size
}
