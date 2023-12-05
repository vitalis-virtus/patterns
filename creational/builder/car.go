package builder

// Car represents the complex object being built.
type Car struct {
	Color         string
	EngineType    string
	HasRoof       bool
	HasNavigation bool
}

// Builder provides an interface for constructing the parts of the car.
type Builder interface {
	SetColor(color string) Builder
	SetEngineType(engineType string) Builder
	SetSunroof(hasSunroof bool) Builder
	SetNavigation(hasNavigation bool) Builder
	Build() *Car
}

// NewCarBuilder creates a new CarBuilder.
func NewCarBuilder() Builder {
	return &carBuilder{
		car: &Car{}, // Initialize the car attribute
	}
}

// carBuilder implements the CarBuilder interface. It's a concrete builder
type carBuilder struct {
	car *Car
}

func (cb *carBuilder) SetColor(color string) Builder {
	cb.car.Color = color
	return cb
}

func (cb *carBuilder) SetEngineType(engineType string) Builder {
	cb.car.EngineType = engineType
	return cb
}

func (cb *carBuilder) SetSunroof(hasRoof bool) Builder {
	cb.car.HasRoof = hasRoof
	return cb
}

func (cb *carBuilder) SetNavigation(hasNavigation bool) Builder {
	cb.car.HasNavigation = hasNavigation
	return cb
}

func (cb *carBuilder) Build() *Car {
	return cb.car
}

// director provides an interface to build cars.
type director struct {
	builder Builder
}

type DirectorBuilder interface {
	ConstructCar(color, engineType string, hasSunroof, hasNavigation bool) *Car
}

func NewDirector(b Builder) DirectorBuilder {
	return &director{
		builder: b,
	}
}

func (d *director) ConstructCar(color, engineType string, hasSunroof, hasNavigation bool) *Car {
	d.builder.
		SetColor(color).
		SetEngineType(engineType).
		SetSunroof(hasSunroof).
		SetNavigation(hasNavigation)

	return d.builder.Build()
}
