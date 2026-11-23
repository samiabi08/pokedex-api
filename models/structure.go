package models

type Pokemon struct {
	Id       int     `json:"pokedex_id"`
	Category string  `json:"category"`
	Name     Names   `json:"name"`
	Image    Sprites `json:"sprites"`
	Types    []Types `json:"types"`
}

type Names struct {
	Fr string `json:"fr"`
}

type Sprites struct {
	Regular string `json:"regular"`
}

type Types struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
