package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
	"koazee-examples-v2/model"
)

// Print the albums sorted by the number of tracks

type AlbumTracks struct {
	Title  string
	Tracks int
}

func main() {
	fmt.Println("Print the albums sorted by the number of tracks:")
	koazee.StreamOf(model.Albums).
	Map(func(a *model.Album) *AlbumTracks {
		return &AlbumTracks{a.Title, len(a.Songs)}
	}).Sort(func(first, second *AlbumTracks) int {
		if first.Tracks > second.Tracks {
			return -1
		}
		return 1
	}).ForEach(func(a *AlbumTracks) {
		fmt.Printf("- '%s' has %d tracks\n", a.Title, a.Tracks)
	}).Do()
}
