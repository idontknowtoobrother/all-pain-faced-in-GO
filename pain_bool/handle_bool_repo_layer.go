package pain_bool

import "fmt"

type UserRequestDataFromHTTP struct {
	Name string
	Shit bool
}

func PainBoolRepoLayer() {
	reqToUpdateName := UserRequestDataFromHTTP{
		Name: "John",
		Shit: false, // user doesn't input but it's default val...
	}
	updateShitOnRepositoryLayer(&reqToUpdateName)
}

func updateShitOnRepositoryLayer(userRequestDataFromHTTP *UserRequestDataFromHTTP) {

	if userRequestDataFromHTTP.Name != "" {
		fmt.Println("1.Yes User need to update name!! because name is not empty")
	}

	if !userRequestDataFromHTTP.Shit { // ??
		fmt.Println("2.How u fcking know it's user input ?")
		fmt.Println("3.Because if they didn't it's a fcking 'FALSE'")
		fmt.Println()
	}
}
