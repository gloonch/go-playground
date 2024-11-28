package main

import "fmt"

// product
type house struct {
	window string
	door   string
	wall   string
	floor  int
}

// builder
type houseBuilder struct {
	window string
	door   string
	wall   string
	floor  int
}

func newHouseBuilder() houseBuilder {
	return houseBuilder{}
}

func (b *houseBuilder) buildWindow(value string) *houseBuilder {
	b.window = value
	return b
}

func (b *houseBuilder) buildDoor(value string) *houseBuilder {
	b.door = value
	return b
}

func (b *houseBuilder) buildWall(value string) *houseBuilder {
	b.wall = value
	return b
}

func (b *houseBuilder) buildFloor(value int) *houseBuilder {
	b.floor = value
	return b
}

// Director will not be able to return the instance of builder
func (b *houseBuilder) buildHouse() house {
	return house{
		window: b.window,
		door:   b.door,
		wall:   b.wall,
		floor:  b.floor,
	}
}

// director
type director struct {
	builder houseBuilder
}

func newDirector(b houseBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) buildHouse() house {

	d.builder.buildWindow("Wooden window").
		buildDoor("Steel door").
		buildFloor(3).
		buildWall("Stone Wall")

	return d.builder.buildHouse()
}

func main() {
	houseBuilder := newHouseBuilder()

	director := newDirector(houseBuilder)
	normalHouse := director.buildHouse()

	fmt.Println("Normal house wall is :", normalHouse.wall)
	fmt.Println("Normal house window is :", normalHouse.window)
	fmt.Println("Normal house door is :", normalHouse.door)
	fmt.Println("Normal house floor is :", normalHouse.floor)

}
