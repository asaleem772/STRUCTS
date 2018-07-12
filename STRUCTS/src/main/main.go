package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fileName := os.Args[1]
	fmt.Println(fileName)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	basePath := dir + "/../src/"
	fPtr, _ := os.Open(basePath + fileName)
	io.Copy(os.Stdout, fPtr)
}

//func (ftype *file)  Read(p []byte) (n int, err error)
type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return (0.5 * t.base * t.height)
}
func (s square) getArea() float64 {
	return (s.sideLength * s.sideLength)
}

func printArea(s shape) {
	fmt.Printf("Area of given shape is %.2f\n", s.getArea())
}

func main2() {

	printArea(triangle{10.5, 2.0})
	printArea(square{5.2})

}

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

//Creating Interfaces below
type bot interface {
	getGreetings() string
}

type englishBot struct{}
type spanishBot struct{}

func (eb englishBot) getGreetings() string {
	return "Hello There!"
}

func (sb spanishBot) getGreetings() string {
	return "Hola!"
}
func printGreeting(b bot) {
	fmt.Println(b.getGreetings())
}

func main1() {
	//	ali := person{"Ali", "Saleem", contactInfo{"ali.saleem@gmail.com", 12345}}
	//	aliPtr := &ali
	//	aliPtr.updateFirstName("Aly")
	//	ali.print()

	//	colors := map[string]string{
	//		"red":    "#ff0000",
	//		"Yellow": "#1200aa",
	//	}
	//	colors["blue"] = "#12aaaa"
	//	delete(colors, "red")
	//	//	fmt.Println(colors)
	//	printMap(colors)

	//	eb := englishBot{}
	//	sb := spanishBot{}
	//
	//	printGreeting(eb)
	//	printGreeting(sb)

	resp, err := http.Get("http://www.google.ie")
	if err != nil {

		fmt.Printf("Error occurred!%v", err)
		os.Exit(1)
	}

	//	bs := make([]byte, 99999)
	//	resp.Body.Read(bs)
	//	s := []string{}
	io.Copy(os.Stdout, resp.Body)
	//	fmt.Printf(s)
}
func printMap(m map[string]string) {
	for i, v := range m {
		fmt.Printf("%v , %v\n", i, v)
	}
}
func (ptrToPerson *person) updateFirstName(firstName string) {
	(ptrToPerson).firstName = firstName
	fmt.Println(*ptrToPerson)
	fmt.Println("****")
	fmt.Println(ptrToPerson)
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
