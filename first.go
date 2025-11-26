package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func printMe(str1 string) {
	fmt.Println(str1)
}

func div(x int, y int) (int, int, error) {
	var er error
	if y == 0 {
		er = errors.New("division by zero")
		// throw
		return 0, 0, er
	}
	var z int = x / y
	y = x % y
	return z, y, nil
}
func main() {

	var intArr = [3]int32{1, 2, 3}   // Fixed length, same type, indexable starting at 0
	_intArr := [3]int32{11, 22, 333} // Fixed length, same type, indexable starting at 0
	// contiguous in memory so compiler does not need to store each location just first one
	fmt.Println(intArr)
	fmt.Println(_intArr)

	var slice = []int32{1, 2, 3, 55, 66}
	fmt.Println("slice", slice, cap(slice))
	slice = append(slice, 978)
	fmt.Println("slice", slice, cap(slice))

	// make slice with a specific capacity
	var slice2 []int32 = make([]int32, 5, 12)
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
	slice2 = append(slice2, slice...) // spread operator
	// specifing capacity save memory, if we know what capacity we need

	//Map  { k : v}
	var map_ = make(map[string]map[string]int)

	// Initialize inner maps before assigning
	map_["alice"] = map[string]int{"age": 12}
	map_["bob"] = map[string]int{"age": 14}

	// Define a custom type for the inner map
	type UserInfo map[string]interface{}
	// Define a custom type for the outer map
	type Users map[string]UserInfo
	users := Users{
		"alice": UserInfo{"age": 12, "name": "Alice Smith"},
		"bob":   UserInfo{"age": 14, "name": "Bob Martin"},
	}

	// key value type uint
	type MapKeyUint8 map[string]uint8
	unitMap := make(MapKeyUint8)
	unitMap["Math"] = 10
	unitMap["English"] = 100
	println("unitMap Scores", unitMap["Math"], unitMap["English"], unitMap["Arabic"])
	var score, exists = unitMap["Math"]
	println("unitMap Scores", score, exists)
	delete(unitMap, "Math")
	score, exists = unitMap["Math"]
	println("unitMap Scores", score, exists)
	fmt.Println(users)
	const str1 string = "Hello World🥁🥁🥁🥁"
	printMe(str1)
	const a = 10
	const b = 3
	var x, y, er = div(a, b)
	if er != nil {
		fmt.Println(er.Error())
	} else if y == 0 {
		fmt.Println("no reminder!")
	}

	switch y {
	case 1:
		fmt.Println("a")
	case 0:
		fmt.Println("b")
	default:
		fmt.Println("default")
	}
	fmt.Printf("div(%v, %v): %v, %v \n", a, b, x, y)
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	var num int = 32767
	var num32 int32 = 6476700
	var numu uint = 32767
	var _float32 float32 = 32767.32767
	var _float64 float64 = 32767.32767
	fmt.Println(num, numu, num32, _float32, _float64, _float32/float32(_float64))
	const str string = "hello world السلام عليكم <UNK> <UNK>"
	fmt.Println("str byte len", len(str))
	fmt.Println("str char len", utf8.RuneCountInString(str))
	for i := 1; i <= 5; i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		fmt.Println("i =", 100/i)
	}

	var map3 = map[string]int{"a": 1, "b": 2, "c": 3}
	var map4 = map[string]string{"a": "1", "b": "2", "c": "3"}
	fmt.Println(map3, map4, map4["X"])
	var emp Employee = Employee{
		name:   "Alice Smith",
		age:    12,
		salary: 2500.55,
		role:   "manager",
		dept:   Dept{name: "Sales"},
		task:   task{name: "sell", duration: "6H"},
	}
	printMe(emp.toString())
	fmt.Println(emp.dept.name)
	fmt.Println(emp.task.name)
	fmt.Println(emp.task.duration)
	fmt.Println(emp.duration) // since task is sub strauct we can call its attribute directly from employee

	for k, v := range map3 {
		fmt.Println("print by KV", k, v)
	}

	for i := 0; i < 50; i += 10 {
		fmt.Println("print i", i)
	}

	var loopString = "Résumé 🥁🥁  ﷺ" // by defualt treated as array of bytes not array of chaters
	for i, v := range loopString {
		fmt.Println("print str", i, v, string(v))
	}
	// using index directly may not be correct if there is non asci charters!

	var strCat = "a"
	strCat += "b" // each concate is a new string and re-assign opeation
	strCat += "c"
	// str is immutable so we can not change a or b or c
	// else we can  use StringBuilder package

	var strBuilder = strings.Builder{}
	strBuilder.WriteString("append 1\n")
	strBuilder.WriteString("append 2\n")
	for i := 3; i <= 5; i++ {
		strBuilder.WriteString(
			fmt.Sprintf("append %d\n", i)) // build the string
	}
	var finalStr = strBuilder.String() // convert to string
	/**
	append 1
	append 2
	append 3
	append 4
	append 5
	*/
	fmt.Println(finalStr)

}

type Employee struct {
	name       string
	age        int
	role       string
	salary     float32
	dateJoined string
	dept       Dept
	task       // the name is task of type task
}

type Dept struct {
	name string
}
type task struct {
	name     string
	duration string
}

func (em Employee) toString() string {
	return fmt.Sprintf("Name=%s, Age=%d,  role=%s, salary=%.2f\n",
		em.name, em.age, em.role, em.salary)
}
