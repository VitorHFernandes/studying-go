package auxiliar

import (
	"fmt"

	"github.com/badoux/checkmail"
)

// ! - GO are not OOP language, so for declare a public function, use caseUp in first character on name function
func Writer() {
	fmt.Println("Writing on Writer")
	writer()

	err := checkmail.ValidateFormat("vitorfernandesdeutsch@gmail.com")
	fmt.Println(err)
}
