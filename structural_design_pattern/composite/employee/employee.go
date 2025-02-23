package main

import "fmt"

type EmployeeComponents interface {
	ShowDetails(indent string)
}

type Empolyee struct {
	Name string
	Role string
}

func (e *Empolyee) ShowDetails(indent string) {
	fmt.Println(indent + "-" + e.Name + "-" + e.Role)
}

type Manager struct {
	Name      string
	Role      string
	Empolyees []EmployeeComponents
}

func (e *Manager) AddEmployee(emp EmployeeComponents) {
	e.Empolyees = append(e.Empolyees, emp)
}

func (e *Manager) ShowDetails(indent string) {
	fmt.Println(indent + "-" + e.Name + "-" + e.Role)
	for _, emp := range e.Empolyees {
		emp.ShowDetails(indent + " ")
	}
}

func main() {
	dev1 := &Empolyee{Name: "Ram", Role: "SDE1"}
	dev2 := &Empolyee{Name: "Shyam", Role: "SDE2"}
	dev3 := &Empolyee{Name: "Krishna", Role: "SDE3"}

	techLead := &Manager{Name: "Vishnu", Role: "TechLead"}
	techLead.AddEmployee(dev1)
	techLead.AddEmployee(dev2)

	manager := &Manager{Name: "Shiva", Role: "Manager"}
	manager.AddEmployee(techLead)
	manager.AddEmployee(dev3)

	manager.ShowDetails("")
}
