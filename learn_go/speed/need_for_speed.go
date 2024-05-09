package main

import "fmt"

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		distance:     0,
		battery:      100,
	}
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery < car.batteryDrain {
		return car
	} else {
		car.distance += car.speed
		car.battery -= car.batteryDrain
		return car
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	factor := track.distance / car.speed
	batteryDrain := car.batteryDrain * factor
	currentBattery := car.battery - batteryDrain
	if currentBattery >= 0 {
		return true
	} else {
		return false
	}
}

func main() {
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)
	car = Drive(car)

	fmt.Println(car)
}
