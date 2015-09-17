package main

import ("testing")

func TestDistance(t *testing.T){
	var v float64
	v=distance(5,10,10,10)
	if v != 5 { t.Error("Expected 5, got ", v) }
}


func TestArea(t *testing.T) {
	var x,y,z,w float64=0,0,10,10;
	var actualRectArea,actualCirArea float64=100,0;
	r := Rectangle{x,y,z,w}
	c := Circle{x,z,x}
 	if (actualRectArea != calArea(&r)) || (actualCirArea != calArea(&c)) {
 		t.Error("Expected area of shape is not calculated.") 
 	}
 }

 func TestPerimeter(t *testing.T){
 	var x,y,z,w float64=0,0,10,10;
	var actualRectPeri,actualCirPeri float64=40,0;
	r := Rectangle{x,y,z,w}
	c := Circle{x,z,x}
 	if (actualRectPeri != calPerimeter(&r)) || (actualCirPeri != calPerimeter(&c)) {
 		t.Error("Expected Perimeter of shape is not calculated") 
 	}
 }