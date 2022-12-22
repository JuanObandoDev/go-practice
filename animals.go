package main

type animal interface {
	move() string
}

type dog struct{}
type fish struct{}
type bird struct{}

func (d dog) move() string {
	return "dog is walking"
}

func (f fish) move() string {
	return "fish is swimming"
}

func (b bird) move() string {
	return "bird is flying"
}

func moveAnimal(a animal) string {
	return a.move()
}

func main() {
	d := dog{}
	f := fish{}
	b := bird{}

	println(moveAnimal(d))
	println(moveAnimal(f))
	println(moveAnimal(b))
}
