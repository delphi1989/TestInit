package inheritance

import "fmt"
import "reflect"

type Animal struct {
	Name string `required max:"100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly bool
}


func CreateAnonymousStruct() {
	b := struct{
		name string
	} {
		"Delphi",
	}

	anotherB := b
	anotherB.name = "Kuo"

	refB := &b
	refB.name = "Changed"

	fmt.Println(b)
	fmt.Println(anotherB)
	fmt.Println(refB)
	fmt.Println(*refB)
}

func CreateInheritance() {
	// 這樣宣告不行
	//bird := Bird{
	//	Name: "Delphi",
	//	Origin: "USA",
	//	SpeedKPH: 32.0,
	//	CanFly: true,
	//}

	// 這樣宣告可以
	bird := Bird{
		SpeedKPH: 32.0,
		CanFly: true,
	}
	bird.Name = "Delphi"
	bird.Origin = "USA"

	// 這樣也行
	birdB := Bird {
		Animal: Animal {
			Name: "Delphi",
			Origin: "USA",
		},
		SpeedKPH: 32.0,
		CanFly: true,
	}

	fmt.Println(bird)
	fmt.Println(birdB)
}

func TestTag() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}