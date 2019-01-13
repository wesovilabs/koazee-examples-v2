package mapReduce

import "time"

type JobExperience struct {
	From         time.Time
	To           time.Time
	Technologies []string
}

type Candidate struct {
	Fullname   string
	Age        int
	Male       bool
	Experience []JobExperience
}
