package object_test

import (
	"fmt"
	"go-demo/task/02/object"
	"testing"
)

func TestShapeAndRectangle(t *testing.T) {
	rect := object.Rectangle{10, 14}
	circle := object.Circle{5.5}

	// 定义 Shape 类型的切片，体现多态
	shapes := []object.Shape{rect, circle}

	for _, s := range shapes {
		fmt.Printf("类型: %T\n", s)
		fmt.Printf("面积: %.2f\n", s.Area())
		fmt.Printf("周长: %.2f\n\n", s.Perimeter())
	}
}

func TestPerson(t *testing.T) {
	emp := object.Employee{
		Person:     object.Person{Name: "Amor", Age: 30},
		EmployeeID: "001",
	}

	emp.PrintInfo()
}
