import "fmt"

func changeValue(str *string) {
	*str = "Changed Value!!"

}

func main() {

	toChange := "hello"
	var Pointer *string = &toChange
	fmt.Println(*Pointer)
	// fmt.Println(toChange)
	// changeValue(&toChange)
	// fmt.Println(toChange)

}
