package blind75

import (
	"encoding/json"
	"fmt"
)

func flatten(arr []any) {
	for _, item := range arr {
		switch v := item.(type) {
		case []any:
			flatten(v)
		default:
			fmt.Print(v, " ")
		}
	}
}

func TwentyOneFlattenDeeplyNestedArray() {
	input := `[1,2,3,[4,5,6],[7,8,[9,10,11],12],[13,14,15]]`

	var arr []any

	json.Unmarshal([]byte(input), &arr)

	fmt.Println("Before Flattening =>", arr)

	fmt.Print("After Flattening  => ")
	flatten(arr)
	fmt.Println()
}

// Done
