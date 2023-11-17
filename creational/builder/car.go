package builder

// Car represents the complex object being built.
type Car struct {
	Color         string
	EngineType    string
	HasRoof       bool
	HasNavigation bool
}

// CarBuilder provides an interface for constructing the parts of the car.
type CarBuilder interface {
	SetColor(color string) CarBuilder
	SetEngineType(engineType string) CarBuilder
	SetSunroof(hasSunroof bool) CarBuilder
	SetNavigation(hasNavigation bool) CarBuilder
	Build() *Car
}

// NewCarBuilder creates a new CarBuilder.
func NewCarBuilder() CarBuilder {
	return &carBuilder{
		car: &Car{}, // Initialize the car attribute
	}
}

// carBuilder implements the CarBuilder interface. It's a concrete builder
type carBuilder struct {
	car *Car
}

func (cb *carBuilder) SetColor(color string) CarBuilder {
	cb.car.Color = color
	return cb
}

func (cb *carBuilder) SetEngineType(engineType string) CarBuilder {
	cb.car.EngineType = engineType
	return cb
}

func (cb *carBuilder) SetSunroof(hasRoof bool) CarBuilder {
	cb.car.HasRoof = hasRoof
	return cb
}

func (cb *carBuilder) SetNavigation(hasNavigation bool) CarBuilder {
	cb.car.HasNavigation = hasNavigation
	return cb
}

func (cb *carBuilder) Build() *Car {
	return cb.car
}

// Director provides an interface to build cars.
type Director struct {
	builder CarBuilder
}

func (d *Director) ConstructCar(color, engineType string, hasSunroof, hasNavigation bool) *Car {
	d.builder.SetColor(color).
		SetEngineType(engineType).
		SetSunroof(hasSunroof).
		SetNavigation(hasNavigation)

	return d.builder.Build()
}
