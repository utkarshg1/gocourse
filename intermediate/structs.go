package intermediate

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	address   Address
}

type Address struct{
	City string
	Country string
}

func (p Person) fullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) incrementAge() {
	p.Age++
}
func main() {
	p := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		address: Address{
			City: "Pune",
			Country: "India",
		},
	}
	fmt.Println(p)
	fmt.Println(p.FirstName)
	fmt.Println(p.LastName)
	fmt.Println(p.Age)
	fmt.Println(p.address)
	fmt.Println(p.fullName())
	p.incrementAge()
	fmt.Println(p.Age)
	p.incrementAge()
	fmt.Println(p.Age)
	user := struct {
		Username string
		Email string
	}{
		Username: "Meowdy",
		Email: "meowdy@example.com",
	}
	fmt.Println(user)
}

