package model

import (
	"github.com/wesovilabs/koazee"
	"time"
)

type Song struct {
	Title    string
	Duration time.Duration
}

type Album struct {
	Title  string
	Artist string
	Label  string
	Year   int
	Songs  []*Song
}

var albums = []*Album{
	{Title: "These Days", Artist: "Bon Jovi", Year: 1995, Label: "Mercury Records", Songs: []*Song{
		{Title: "Hey God", Duration: 6*time.Minute + 12*time.Second},
		{Title: "Something For The Pain", Duration: 4*time.Minute + 46*time.Second},
		{Title: "This Ain't a Love Song", Duration: 5*time.Minute + 1*time.Second},
		{Title: "These Days", Duration: 6*time.Minute + 23*time.Second},
		{Title: "Lie To Me", Duration: 5*time.Minute + 36*time.Second},
		{Title: "Damned", Duration: 4*time.Minute + 34*time.Second},
		{Title: "My Guitar Lies Bleeding In My Arms", Duration: 5*time.Minute + 42*time.Second},
		{Title: "Letting You Go", Duration: 5*time.Minute + 50*time.Second},
		{Title: "Hearts Breaking Even", Duration: 5*time.Minute + 5*time.Second},
		{Title: "Something To Believe In", Duration: 5*time.Minute + 26*time.Second},
		{Title: "If That's What It Takes", Duration: 5*time.Minute + 15*time.Second},
		{Title: "Diamond Ring", Duration: 3*time.Minute + 46*time.Second},
	}},
	{Title: "Born in the U.S.A.", Artist: "Bruce Springsteen", Year: 1984, Label: "Columbia Records", Songs: []*Song{
		{Title: "Born in the U.S.A.", Duration: 4*time.Minute + 42*time.Second},
		{Title: "Cover Me", Duration: 3*time.Minute + 23*time.Second},
		{Title: "Darlington County", Duration: 4*time.Minute + 15*time.Second},
		{Title: "Downbound Train", Duration: 3*time.Minute + 40*time.Second},
		{Title: "I'm on Fire", Duration: 2*time.Minute + 42*time.Second},
		{Title: "No Surrender", Duration: 4*time.Minute + 3*time.Second},
		{Title: "Bobby Jean", Duration: 3*time.Minute + 50*time.Second},
		{Title: "I'm Goin' Down", Duration: 3*time.Minute + 31*time.Second},
		{Title: "Glory Days", Duration: 4*time.Minute + 20*time.Second},
		{Title: "Dancing in the Dark", Duration: 4*time.Minute + 4*time.Second},
		{Title: "My Hometown", Duration: 4*time.Minute + 32*time.Second},
	}},
	{Title: "Thriller", Artist: "Michael Jackson", Year: 1982, Label: "Epic Records", Songs: []*Song{
		{Title: "Wanna Be Startin' Somethin'", Duration: 6*time.Minute + 3*time.Second},
		{Title: "Baby Be Mine", Duration: 4*time.Minute + 20*time.Second},
		{Title: "The Girl Is Mine", Duration: 3*time.Minute + 42*time.Second},
		{Title: "Thriller", Duration: 5*time.Minute + 58*time.Second},
		{Title: "Beat It", Duration: 4*time.Minute + 19*time.Second},
		{Title: "Billie Jean", Duration: 4*time.Minute + 54*time.Second},
		{Title: "Human Nature", Duration: 4*time.Minute + 5*time.Second},
		{Title: "P.Y.T. (Pretty Young Thing)", Duration: 3*time.Minute + 58*time.Second},
		{Title: "The Lady in My Life", Duration: 4*time.Minute + 59*time.Second},
	}},
	{Title: "A Night at the Opera", Artist: "Queen", Year: 1975, Label: "EMI", Songs: []*Song{
		{Title: "Death on Two Legs (Dedicated to...)", Duration: 3*time.Minute + 43*time.Second},
		{Title: "Lazing on a Sunday Afternoon", Duration: 1*time.Minute + 3*time.Second},
		{Title: "I'm in Love with My Car", Duration: 3*time.Minute + 5*time.Second},
		{Title: "You're My Best Friend", Duration: 2*time.Minute + 50*time.Second},
		{Title: "39", Duration: 3*time.Minute + 30*time.Second},
		{Title: "Sweet Lady", Duration: 4*time.Minute + 1*time.Second},
		{Title: "Seaside Rendezvous", Duration: 2*time.Minute + 13*time.Second},
		{Title: "Love of My Life", Duration: 3*time.Minute + 38*time.Second},
		{Title: "Good Company", Duration: 3*time.Minute + 26*time.Second},
	}},
	{Title: "21", Artist: "Adele", Year: 1975, Label: "Columbia Recors", Songs: []*Song{
		{Title: "Rolling in the Deep", Duration: 3*time.Minute + 49*time.Second},
		{Title: "Rumour Has It", Duration: 3*time.Minute + 43*time.Second},
		{Title: "Turning Tables", Duration: 4*time.Minute + 10*time.Second},
		{Title: "Don't You Remember", Duration: 4*time.Minute + 03*time.Second},
		{Title: "Set Fire to the Rain", Duration: 4*time.Minute + 1*time.Second},
		{Title: "He Won't Go", Duration: 4*time.Minute + 37*time.Second},
		{Title: "Take It All", Duration: 3*time.Minute + 48*time.Second},
		{Title: "I'll Be Waiting", Duration: 4*time.Minute + 1*time.Second},
		{Title: "One and Only", Duration: 5*time.Minute + 48*time.Second},
		{Title: "Lovesong", Duration: 5*time.Minute + 16*time.Second},
		{Title: "Someone like You", Duration: 4*time.Minute + 45*time.Second},
	}},
}

var stream = koazee.StreamOf(albums)
