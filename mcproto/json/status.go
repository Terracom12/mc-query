package json

type Status struct {
	Version     Version
	Description string
	Players     Players
}

type Version struct {
	Name     string
	Protocol int
}
type Players struct {
	Max    int
	Online int
}
