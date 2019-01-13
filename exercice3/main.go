package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
	"koazee-examples-v2/model"
)

func main() {

	album := model.Albums[0]
	tracks := koazee.StreamOf(album.Songs).
		Sort(func(l, r *model.Song) int {
			if l.Title < r.Title {
				return -1
			}
			return 1
		}).Do().Out().Val().([]*model.Song)

	fmt.Println("Print the tracks of a given album sorted alphabetically asc")
	koazee.StreamOf(tracks).
		ForEach(func(s *model.Song) {
			fmt.Printf(" - %s\n", s.Title)
		}).Do()
	fmt.Println("\nPrint the tracks of a given album sorted alphabetically desc")
	koazee.StreamOf(tracks).
		Reverse().
		ForEach(func(s *model.Song) {
			fmt.Printf(" - %s\n", s.Title)
		}).Do()

}
