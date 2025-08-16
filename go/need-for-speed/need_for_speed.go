package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery, batteryDrain, speed, distance int
}

type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}
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
	var remainingBattery = car.battery - car.batteryDrain

	if remainingBattery >= 0 {
		car.battery = remainingBattery
		car.distance += car.speed
	} else {
		//no changes
	}
	return car

}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	//battery /discharging * speed = distanza percorribile

	var distanceCanRun = (car.battery / car.batteryDrain) * car.speed
	if distanceCanRun >= track.distance {
		return true
	} else {
		return false
	}

}
