package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
	"koazee-examples-v2/model"
	"time"
)

func main() {
	fmt.Println("Print the duration for a given album.")
	album := model.Albums[0]
	d := koazee.StreamOf(album.Songs).
		Reduce(func(duration time.Duration, s *model.Song) time.Duration {
			return duration + s.Duration
		}).Val()
	fmt.Printf("- The duration for album %s by %s is %v\n", album.Title, album.Artist, d)

}
