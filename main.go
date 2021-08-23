package main

var eru = "Eru"

func main() {
	// Arrays and Slices
	/*
		var ages [3]int = [3]int{20, 25, 30}
		var agesShort = [3]int{20, 25, 30}
		fmt.Println(ages, agesShort)
		fmt.Println(ages, len(ages))

		names := [4]string{"Bilbo", "Frodo", "Merry", "Pippin"}
		names[1] = "Gandalf"
		fmt.Println(names, len(names))

		var scores = []int{100, 50, 60} // Slice not an array
		scores[2] = 65
		fmt.Println(scores, len(scores))
		scores = append(scores, 85) // Doesn't append it returns a new slice (we overwrite)
		fmt.Println(scores, len(scores))

		rangeOne := names[1:3] // Inclusive of initial, exclusive of second i.e. 1, 2
		fmt.Println(rangeOne)

		rangeTwo := names[2:] // Includes position 2 and everything else
		fmt.Println(rangeTwo)

		rangeThree := names[:3] // Everything up to (not including) index 3
		fmt.Println(rangeThree)

		rangeFour := append(rangeOne, "Elrond")
		fmt.Println(rangeFour)
	*/

	// String formatting and printing
	/*
		age := 3000
		gandalf := "Gandalf"
		whichBirthday := 111
		event := "birthday"
		seconds := 1.2345566

		fmt.Print("Good morning...")
		fmt.Println(gandalf, "is", age, "years old!")
		fmt.Print("No admittance except on party business!\n")

		fmt.Printf("Bilbo is %v and today is my %v\n", whichBirthday, event)
		fmt.Printf("Bilbo is %q and today is my %q\n", whichBirthday, event)
		fmt.Printf("Bilbo's age is of type %T\n", whichBirthday)
		fmt.Printf("Current second is %f\n", seconds)
		fmt.Printf("Current second is %0.2f\n", seconds)

		var str = fmt.Sprintf("Bilbo is %v and today is my %v\n", whichBirthday, event)
		fmt.Print("The saved string  is: ", str)
	*/

	// Strings
	/*
		var bilbo string = "Bilbo"
		var frodo = "Frodo"
		var gandalf string

		fmt.Println(bilbo, frodo, gandalf)

		bilbo = "Frodo's Uncle"
		gandalf = "Gandalf"

		fmt.Println(bilbo, frodo, gandalf)

		pippin := "Pippin"

		fmt.Println(pippin)
	*/

	// Integers
	/*
		var age int = 111
		var frodoAge = 30
		gandalfAge := 3000

		fmt.Println(age, frodoAge, gandalfAge)
	*/

	// Bits and Memory
	/*
		var bit8 int8 = 25
		var bit8neg int8 = -128
		var bit8un uint8 = 223

		fmt.Println(bit8, bit8neg, bit8un)
	*/

	// Floats
	/*
		var scoreOne float32 = 25.96
		var scoreTwo float64 = 34987349573459830458304.8
		scoreThree := 1.5

		fmt.Println(scoreOne, scoreTwo, scoreThree)
	*/
}
