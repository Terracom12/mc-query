package json

type Status struct {
	Version     Version
	Description any
	Players     Players
}

type Version struct {
	Name     string
	Protocol int
}

type Players struct {
	Max    int
	Online int
	Sample []PlayerSample
}

type PlayerSample struct {
	Id   string
	Name string
}
