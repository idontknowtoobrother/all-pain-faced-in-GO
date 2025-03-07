package pain_bool

import (
	"fmt"
	"time"
)

// OH SHIT I KNOW THAT

// IT'S active or unactive right

// can we change a DeletedAt to DeactiveAt ?
// because this struct can disable or enable it
// so we can use full in one value by not handle this shit
// or just use DeletedAt and DeactiveAt so we can usefully it like show deactive date on front end, yes but it can enable it back ?

type RealShitUsefulmoreThanBool struct {
	Name       string
	DeactiveAt *time.Time // this struct is some kind like active or inactive (can active later it's just like disable or enable)

	// this field is tell this shit already deleted can't get it back ! actually can it's soft delete
	// but it's just will not show in that Management UI make UI cleaner | can recovery later another recovery UI
	DeletedAt *time.Time // Optional

}

func BestFixPainBoolRepoLayer() {
	mockUserRequestNeedUpdateDeactiveAt := time.Now()
	reqToUpdateName := RealShitUsefulmoreThanBool{
		Name:       "John",
		DeactiveAt: &mockUserRequestNeedUpdateDeactiveAt, // No update Active Or Active
	}
	bestFixIHought(&reqToUpdateName)
}

func bestFixIHought(userRequestDataFromHTTP *RealShitUsefulmoreThanBool) {

	if userRequestDataFromHTTP.Name != "" {
		fmt.Println("1.Yes User need to update name!! because name is not empty")
	}

	if userRequestDataFromHTTP.DeactiveAt != nil { // ??
		fmt.Println("2.Yes User need to update active or inactive !")
		fmt.Println("3.Yes it's a time value", userRequestDataFromHTTP.DeactiveAt)
		fmt.Println("4.Can more useful tell Front End it's show deactive date this value", userRequestDataFromHTTP.DeactiveAt.UTC().Format(time.RFC3339))
		fmt.Println()
	}
}
