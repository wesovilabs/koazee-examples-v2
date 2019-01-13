package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
	"koazee-examples-v2/model"
)

func main() {
	fmt.Println("Print the title of albums provided by 'Columbia Records'")
	koazee.StreamOf(model.Albums).
		Filter(func(a *model.Album) bool {
			return a.Label == "Columbia Records"
		}).ForEach(func(a *model.Album) {
		fmt.Printf(" - %s\n", a.Title)
	}).Do()
}
