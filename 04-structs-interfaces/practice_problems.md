# Practice Problems - Structs and Interfaces

These practice problems reinforce struct and interface concepts. Focus on proper interface design and method implementation.

## Problem 1: Shape Calculator with Interfaces

**Description:** Create a shape calculator that can compute areas and perimeters for different geometric shapes using interfaces.

**Requirements:**
- Define a `Shape` interface with `Area()` and `Perimeter()` methods
- Implement `Circle` and `Rectangle` structs that satisfy the interface
- Create a function that takes a slice of shapes and computes total area and perimeter

**Examples:**
- Circle with radius 5: Area = 78.54, Perimeter = 31.42
- Rectangle 4x6: Area = 24.00, Perimeter = 20.00

**Constraints:**
- Use math.Pi for π calculations
- All dimensions are positive floats

### Solution

```go
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func calculateTotals(shapes []Shape) (totalArea, totalPerimeter float64) {
	for _, shape := range shapes {
		totalArea += shape.Area()
		totalPerimeter += shape.Perimeter()
	}
	return
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
		Circle{Radius: 3},
		Rectangle{Width: 2, Height: 8},
	}

	fmt.Println("Individual shapes:")
	for i, shape := range shapes {
		fmt.Printf("Shape %d: Area = %.2f, Perimeter = %.2f\n",
			i+1, shape.Area(), shape.Perimeter())
	}

	totalArea, totalPerimeter := calculateTotals(shapes)
	fmt.Printf("\nTotals: Area = %.2f, Perimeter = %.2f\n", totalArea, totalPerimeter)
}
```

**Explanation:**
1. Define Shape interface with Area() and Perimeter() methods
2. Implement interface for Circle and Rectangle structs
3. Use value receivers since methods don't modify the structs
4. Create polymorphic function that works with any Shape
5. Demonstrate interface satisfaction through usage

## Problem 2: Employee Management System

**Description:** Create an employee management system using structs and interfaces.

**Requirements:**
- Define `Employee` interface with `GetSalary()` and `GetDetails()` methods
- Implement `Manager` and `Developer` structs
- Managers get base salary + bonus per subordinate
- Developers get base salary + experience bonus
- Create a function to calculate total payroll

**Examples:**
- Manager with 3 subordinates: Salary = 50000 + 3*2000 = 56000
- Developer with 5 years experience: Salary = 40000 + 5*1000 = 45000

### Solution

```go
package main

import "fmt"

type Employee interface {
	GetSalary() float64
	GetDetails() string
}

type Manager struct {
	Name         string
	BaseSalary   float64
	BonusPerSub  float64
	Subordinates int
}

func (m Manager) GetSalary() float64 {
	return m.BaseSalary + float64(m.Subordinates)*m.BonusPerSub
}

func (m Manager) GetDetails() string {
	return fmt.Sprintf("Manager: %s (Base: $%.0f, Subs: %d)",
		m.Name, m.BaseSalary, m.Subordinates)
}

type Developer struct {
	Name            string
	BaseSalary      float64
	ExperienceBonus float64
	YearsExperience int
}

func (d Developer) GetSalary() float64 {
	return d.BaseSalary + float64(d.YearsExperience)*d.ExperienceBonus
}

func (d Developer) GetDetails() string {
	return fmt.Sprintf("Developer: %s (Base: $%.0f, Exp: %d years)",
		d.Name, d.BaseSalary, d.YearsExperience)
}

func calculatePayroll(employees []Employee) float64 {
	total := 0.0
	for _, emp := range employees {
		total += emp.GetSalary()
	}
	return total
}

func main() {
	employees := []Employee{
		Manager{Name: "Alice", BaseSalary: 50000, BonusPerSub: 2000, Subordinates: 3},
		Developer{Name: "Bob", BaseSalary: 40000, ExperienceBonus: 1000, YearsExperience: 5},
		Manager{Name: "Charlie", BaseSalary: 55000, BonusPerSub: 2500, Subordinates: 2},
		Developer{Name: "Diana", BaseSalary: 35000, ExperienceBonus: 1200, YearsExperience: 3},
	}

	fmt.Println("Employee Details:")
	for _, emp := range employees {
		fmt.Printf("%s - Salary: $%.0f\n", emp.GetDetails(), emp.GetSalary())
	}

	totalPayroll := calculatePayroll(employees)
	fmt.Printf("\nTotal Payroll: $%.0f\n", totalPayroll)
}
```

**Explanation:**
1. Define Employee interface with salary and details methods
2. Implement different salary calculation logic for Manager and Developer
3. Use value receivers since methods don't modify struct fields
4. Demonstrate polymorphism with calculatePayroll function
5. Show different bonus structures for different employee types