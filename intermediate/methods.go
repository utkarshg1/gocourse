package intermediate

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r *Rectangle) scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

func main() {
	r1 := Rectangle{
		width:  12,
		height: 5,
	}
	fmt.Println(r1)
	r1.scale(1.5)
	fmt.Println(r1)
	fmt.Println(r1.perimeter())
	fmt.Println(r1.area())
}