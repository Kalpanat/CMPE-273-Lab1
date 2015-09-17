package main
import ("fmt"; "math")

func distance(x1, y1, x2, y2 float64) float64 { 
	a := x2 - x1
	b := y2 - y1 
	return math.Sqrt(a*a + b*b) }

type Rectangle struct { x1, y1, x2, y2 float64 }
func (r Rectangle) area() float64 { 
	l := distance(r.x1, r.y1, r.x1, r.y2) 
	w := distance(r.x1, r.y1, r.x2, r.y1) 
	return l * w }
func (r Rectangle) perimeter() float64{
	l := distance(r.x1, r.y1, r.x1, r.y2) 
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2*(l+w) }

type Circle struct { x, y, r float64 }
func (c Circle) area() float64 { return math.Pi * c.r*c.r }
func (c Circle) perimeter() float64 { return 2 * math.Pi * c.r }

type Shape interface { 
	area() float64 
	perimeter() float64
}
func calArea(shapes ...Shape) float64 { 
	var area float64 
	for _, s := range 
	shapes { area = s.area() } 
	return area }
func calPerimeter(shapes ...Shape) float64 { 
	var perimeter float64 
	for _, s := range 
	shapes { perimeter = s.perimeter() } 
	return perimeter }

func main() {
	r := Rectangle{0, 0, 10, 10} 
	fmt.Println("Area Of rectangle: ",calArea(&r))
	fmt.Println("Perimeter Of rectangle: ",calPerimeter(&r))
	c := Circle{0, 0, 5} 
	fmt.Println("Area Of circle: ",calArea(&c))
	fmt.Println("Perimeter Of circle: ",calPerimeter(&c))
}