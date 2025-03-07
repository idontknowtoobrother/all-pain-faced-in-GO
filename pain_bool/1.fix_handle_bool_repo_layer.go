package pain_bool

import "fmt"

// helper shit
func BoolPtr(b bool) *bool {
	return &b
}

func BoolVal(b *bool) (val bool, isNil bool) {
	if b == nil {
		return false, true
	}
	return *b, false
}

// end helper shit

type FixUserRequestDataFromHTTP struct {
	Name string
	Shit *bool // use ptr
}

func FixPainBoolRepoLayer() {
	mockUserRequestNeedUpdateShit := BoolPtr(true)
	reqToUpdateName := FixUserRequestDataFromHTTP{
		Name: "John",
		Shit: mockUserRequestNeedUpdateShit, // user doesn't input but it's a pointer so value is "NIL"
	}
	fixUpdateShitOnRepositoryLayer(&reqToUpdateName)
}

func fixUpdateShitOnRepositoryLayer(userRequestDataFromHTTP *FixUserRequestDataFromHTTP) {

	if userRequestDataFromHTTP.Name != "" {
		fmt.Println("1.Yes User need to update name!! because name is not empty")
	}

	if shitVal, isNil := BoolVal(userRequestDataFromHTTP.Shit); !isNil { // ??
		fmt.Println("2.Yes User need to update shit!! because it's not nil")
		fmt.Println("3.Yes it's a bool value", shitVal)
		fmt.Println()
		userRequestDataFromHTTP.Shit = BoolPtr(shitVal) // use helper bool ptr take it ..
	}
}
