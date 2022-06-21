package main

// type Doctor struct {
// 	Number     int
// 	ActorName  string
// 	Episodes   []string
// 	Companions []string
// }

// type Animal struct {
// 	Name   string `required max:"100`
// 	Origin string
// }

// type Bird struct {
// 	Animal
// 	SpeedKPH float32
// 	CanFly   bool
// }

func main() {
	// nameAge := make(map[string]int)
	// nameAge = map[string]int{
	// 	"Alex": 21,
	// 	"Ro":   26,
	// }
	// nameAge["Ash"] = 23
	// peep, ok := nameAge["Row"]

	// fmt.Println(peep, ok)

	// aDoctor := Doctor{
	// 	Number:    3,
	// 	ActorName: "Jon Pertwee",
	// 	Episodes:  []string{},
	// 	Companions: []string{
	// 		"Liz Shaw",
	// 		"Jo Grant",
	// 		"Sarah Jane Smith",
	// 	},
	// }

	// fmt.Println(aDoctor.Companions[2])

	// aDoctor := struct{ name string }{name: "John Pertwee"}
	// anotherDoctor := aDoctor
	// // anotherDoctor := &aDoctor
	// anotherDoctor.name = "Tom Baker"

	// fmt.Println(aDoctor)
	// fmt.Println(anotherDoctor)

	// b := Bird{}
	// b.Name = "Emu"
	// b.Origin = "Australia"
	// b.SpeedKPH = 48
	// b.CanFly = false
	// fmt.Println(b.Name)

	// b := Bird{
	// 	Animal:   Animal{Name: "Emu", Origin: "Australia"},
	// 	SpeedKPH: 48,
	// 	CanFly:   false,
	// }
	// fmt.Println(b.Name)

	// t := reflect.TypeOf(Animal{})
	// field, _ := t.FieldByName("Name")
	// fmt.Println(field.Tag)
}

// Maps:
// * Collections of value types that are accessed via keys
// * Created via literals or via make function
// * Members accessed via [key] syntax
// 	*-myMap["key"] = "value"
// * Check for presence with "value, ok" form of result
// * Multiple assignmnets refer to same underlying data

// Structs:
// * Collections of disparate data types that describe a single concept
// * Keyed by named fields
// * Normally created as types, but anonymous structs are allowed
// * Structs are value types
// * No inheritance, but can use composition via embedding
// * Tags can be added to struct fields to describe field
