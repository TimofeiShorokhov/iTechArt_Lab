package main

import (
	"fmt"
	"iTechArt_Lab/task_2_Geometry/solution"
)

func main() {

	var pole solution.Pole
	var paral solution.Parallelogramm
	var cir solution.Circle

	//creating Field for our goal...
	pole.CreateArray(8)

	//making dots....
	var a = solution.Dot{X: 1, Y: 2, Color: solution.Yellow}
	var b = solution.Dot{X: 2, Y: 3, Color: solution.Brown}
	var c = solution.Dot{X: 1, Y: 6, Color: solution.Blue}
	var extraDot = solution.Dot{X: 1, Y: 0, Color: solution.Green}

	//setting three dots on field...
	pole.SetThreeDots(a, b, c)

	fmt.Println("````pole with three dots only````")
	pole.PrintPole()
	fmt.Println("-----------------------------")

	//searching fourth dot...
	paral.CreateFourthDot(a, b, c, pole)

	//making parallelogramm...
	paral.MakeParam(pole)
	fmt.Println("````pole with parallelogramm````")
	pole.PrintPole()
	fmt.Println("-----------------------------")

	//finding square of parallelogramm...
	paral.FindSquare(a, b, c)

	//finding center of parallelogramm...
	paral.FindCenter(a, c)

	//finding radius of circle...
	cir.FindRadius(paral)

	//finding length of circle...
	cir.FindLength()

	//making circle...
	cir.MakeCircle(paral, pole)

	//printing our field with parallelogram and center of circle...
	fmt.Println("````pole with center````")
	pole.PrintPole()
	fmt.Println()
	fmt.Println("-----------------------------")

	//printing info about our figures...
	solution.PrintInfo(paral, cir)
	fmt.Println()
	fmt.Println("-----------------------------")

	//changing first dot and printing changed info...
	fmt.Println("~~~changed parallelogramm~~~")
	pole.ChangeOneDot(extraDot, paral, cir)
	fmt.Println()
	fmt.Println("-----------------------------")

	//printing pole with new figure...
	fmt.Println("~~~changed pole~~~")
	pole.PrintPole()

}
