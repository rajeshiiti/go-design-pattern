package main

import "fmt"

type house struct {
	wall   string
	window string
	door   string
	floor  int
}

type houseBuilder struct {
	wall   string
	window string
	door   string
	floor  int
}

type director struct {
	builder houseBuilder
}

func getHouseBuilder() houseBuilder {
	return houseBuilder{}
}

func (h *houseBuilder) buildWindow(windowType string) *houseBuilder {
	h.window = windowType
	return h
}

func (h *houseBuilder) buildWall(wall string) *houseBuilder {
	h.wall = wall
	return h
}

func (h *houseBuilder) buildDoor(door string) *houseBuilder {
	h.door = door
	return h
}

func (h *houseBuilder) buildFloor(floor int) *houseBuilder {
	h.floor = floor
	return h
}

func (h *houseBuilder) buildHouse() house {
	return house{
		door:   h.door,
		wall:   h.wall,
		floor:  h.floor,
		window: h.window,
	}
}

func newDirector(h houseBuilder) *director {
	return &director{
		builder: h,
	}
}

func (d *director) buildHouse() house {
	d.builder.buildDoor("Steel door")
	d.builder.buildFloor(2)
	d.builder.buildWall("Paint wall")
	d.builder.buildWindow("Glass window")

	return d.builder.buildHouse()

}

func main() {
	b := getHouseBuilder()
	d := newDirector(b)
	myHouse := d.buildHouse()
	fmt.Println("Wall: ", myHouse.wall)
	fmt.Println("Door: ", myHouse.door)
	fmt.Println("Floor: ", myHouse.floor)
	fmt.Println("Window: ", myHouse.window)
}
