package main

import "fmt"

type AutomobileSpecification struct {
	Name               string
	NDoors             int
	RimCenterBore      float32
	RimBoltPatter      string
	Petrol             string
	EngineVolumeLiters float32
	EngineName         string
	HorsePower         float32
	WeightKilogram     int
	HasSunroof         bool
	HasRoof            bool
	NSportFeatures     int
}

// 1) receive by reference
func isSportCar(spec *AutomobileSpecification) bool {
	// 2) init NSportFeatures field
	spec.NSportFeatures = 0

	if spec.NDoors < 4 {
		spec.NSportFeatures += 1
	}

	if spec.HorsePower > 150 {
		spec.NSportFeatures += 1
	}

	if spec.WeightKilogram < 1050 {
		spec.NSportFeatures += 1
	}

	return spec.NSportFeatures > 2
}

func main() {
	car1 := AutomobileSpecification{
		Name:               "ВАЗ2106",
		NDoors:             4,
		RimCenterBore:      58.5,
		RimBoltPatter:      "4x98",
		Petrol:             "A92",
		EngineVolumeLiters: 1.57,
		EngineName:         "2106",
		HorsePower:         71.5,
		WeightKilogram:     1035,
		HasSunroof:         false,
		HasRoof:            true,
	}

	// 3) pass by reference
	fmt.Printf("is %s a sport cat? %v\n", car1.Name, isSportCar(&car1))
	fmt.Printf("is %s a sport cat? %v\n", car1.Name, isSportCar(&car1))
	fmt.Printf("is %s a sport cat? %v\n", car1.Name, isSportCar(&car1))
}
