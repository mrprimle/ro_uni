package main

import ("fmt")

type Station struct {
	freeSpots []int
	visitedSpots []int
}

type Car struct {
	waitTime int
	isParked bool
}

func main()  {

	station := Station {freeSpots: []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, visitedSpots: []int {}}

	cars1 := []Car {}
	cars2 := []Car {}

	for car := 0; car < 10; car += 1 {
		carID := Car {waitTime: 0, isParked: false}
		cars1 = append( cars1, carID)
		cars2 = append( cars2, carID)
	}



	fmt.Println(station)
	fmt.Println(cars1)
	fmt.Println(cars2)

	for _, car1 := range cars1 {
		go CarThread (car1, station)
	}

	for _, car2 := range cars2 {
		car, station := CarThread (car2, station)
		fmt.Println(car, station)
	}
	
}

func CarThread (car Car, station Station) (Car, Station) {
	if len(station.freeSpots) != 0 {
		for spot := range station.freeSpots {
			station.visitedSpots = append(station.visitedSpots, spot)
		}
		car.isParked = true
	} else {
		car.waitTime++
	}

	if car.waitTime == 5 {
		fmt.Println("I'm leaving")
	}


	return car, station
}