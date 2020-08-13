package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct {}

func(f *ManufacturingDirector) Construct() {

}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {

}

type VehicleProduct struct {
	Wheels int
	Seats int
	Structure string
}
type CarBuilder struct{}

func (c *CarBuilder) SetWheels() BuildProcess {
	return nil
}

func (c *CarBuilder) SetSeats() BuildProcess {
	return nil
}

func(c *CarBuilder) SetStruture() BuildProcess {
	return nil
}

func(c *CarBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}

type BikeBuilder struct{}

func(b *BikeBuilder) SetWheels() BuildProcess {
	return nil
}

func(b *BikeBuilder) SetSeats() BuildProcess {
	return nil
}

func(b *BikeBuilder) SetStructure() BuildProcess {
	return nil
}

func (b *BikeBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}