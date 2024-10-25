package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"slices"
)

const MIN = 1
const MAX = 5

func rF64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

type Shape3D interface {
	Vol() float64
}

type Cube struct {
	x float64
}

type Cuboid struct {
	x float64
	y float64
	z float64
}

type Sphere struct {
	r float64
}

func (c Cube) Vol() float64 {
	return c.x * c.x * c.x
}

func (c Cuboid) Vol() float64 {
	return c.x * c.y * c.z
}

func (c Sphere) Vol() float64 {
	return 4 / 3 * math.Pi * c.r * c.r * c.r
}

type shapes []Shape3D

func compareShapes(a, b Shape3D) int {
	aVol := a.Vol()
	bVol := b.Vol()

	switch {
	case aVol > bVol:
		return 1
	case aVol == bVol:
		return 0
	case aVol < bVol:
		return -1
	default:
		fmt.Println("Error!")
		os.Exit(1)
	}

	return 0
}

func PrintShapes(a shapes) {
	for _, v := range a {
		switch v.(type) {
		case Cube:
			fmt.Printf("Cube: volume %.2f\n", v.Vol())
		case Cuboid:
			fmt.Printf("Cuboid: volume: %.2f\n", v.Vol())
		case Sphere:
			fmt.Printf("Sphere: volume: %.2f\n", v.Vol())
		default:
			fmt.Println("Unknown data type!")
		}
	}
	fmt.Println()
}

func main() {
	data := shapes{}

	for i := 0; i < 3; i++ {
		cube := Cube{rF64(MIN, MAX)}
		cuboid := Cuboid{rF64(MIN, MAX), rF64(MIN, MAX), rF64(MIN, MAX)}
		sphere := Sphere{rF64(MIN, MAX)}
		data = append(data, cube)
		data = append(data, cuboid)
		data = append(data, sphere)
	}
	PrintShapes(data)

	slices.SortFunc(data, compareShapes)
	PrintShapes(data)

	slices.SortFunc(data, func(a, b Shape3D) int { return compareShapes(b, a) })
	PrintShapes(data)
}
