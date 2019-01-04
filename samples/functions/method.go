package main

import "fmt"

type VideoCourse struct {
	Minute   int
	Language string
}

func (v *VideoCourse) Description() string {
	return fmt.Sprintf("Video Course for %s with a duration of %d minutes\n", v.Language, v.Minute)
}

func main() {
	course := VideoCourse{Minute: 60, Language: "Go"}
	fmt.Printf("%+v\n", course)

	description := course.Description()
	fmt.Println(description)
}
