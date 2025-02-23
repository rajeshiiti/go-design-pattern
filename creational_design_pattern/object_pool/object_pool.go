package main

import "fmt"

type iDBConnection interface {
	getID() string
}

type connection struct {
	id string
}

type connectionStatus struct {
	conn   *connection
	status bool
}

func (c *connection) getID() string {
	return c.id
}

type poolConnections struct {
	connection    []connection
	connectionMap map[string]bool // here bool is variable that shows that connection is available
	available     int
}

func (pc *poolConnections) getConnection() iDBConnection {
	if pc.available == 0 {
		fmt.Println("No connection available.")
		panic("No connection available.")
	}
	for _, val := range pc.connection {
		status := pc.connectionMap[val.id]
		if status {
			pc.connectionMap[val.id] = false
			fmt.Println("Setting false for id: ", val.id)
			pc.available = pc.available - 1
			return &val
		}
	}
	fmt.Println("No connection available.")
	panic("No connection available.")
}

func (pc *poolConnections) returnConnection(target iDBConnection) error {
	pc.connectionMap[target.getID()] = true
	pc.available = pc.available + 1
	return nil
}

func initPool(poolObjects []connection) poolConnections {
	connectionMap := make(map[string]bool)
	for _, con := range poolObjects {
		connectionMap[con.getID()] = true
	}
	pool := poolConnections{
		connection:    poolObjects,
		connectionMap: connectionMap,
		available:     len(poolObjects),
	}
	return pool

}

func main() {
	poolObjects := []connection{}

	c1 := connection{id: "1"}
	poolObjects = append(poolObjects, c1)
	c2 := connection{id: "2"}
	poolObjects = append(poolObjects, c2)
	c3 := connection{id: "3"}
	poolObjects = append(poolObjects, c3)
	c4 := connection{id: "4"}
	poolObjects = append(poolObjects, c4)
	c5 := connection{id: "5"}
	poolObjects = append(poolObjects, c5)
	fmt.Println("size: ", len(poolObjects))

	connectionPool := initPool(poolObjects)

	fmt.Println("size: ", len(connectionPool.connection))
	fmt.Println("Available connection: ", connectionPool.available)

	// case 1
	conn1 := connectionPool.getConnection()
	fmt.Println("Connection id: ", conn1.getID())

	// case 2
	conn2 := connectionPool.getConnection()
	fmt.Println("Connection id: ", conn2.getID())

	// case 3
	conn3 := connectionPool.getConnection()
	fmt.Println("Connection id: ", conn3.getID())

	// case 4
	conn4 := connectionPool.getConnection()
	fmt.Println("Connection id: ", conn4.getID())

	connectionPool.returnConnection(conn1)
	connectionPool.returnConnection(conn3)

	conn5 := connectionPool.getConnection()
	fmt.Println("Connection id: ", conn5.getID())

	conn6 := connectionPool.getConnection()
	fmt.Println("Connection id: ", conn6.getID())

	conn7 := connectionPool.getConnection()
	fmt.Println("Connection id: ", conn7.getID())

	conn8 := connectionPool.getConnection()
	fmt.Println("Connection id: ", conn8.getID())

	fmt.Println("Available connection: ", connectionPool.available)

}
