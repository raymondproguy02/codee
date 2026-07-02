package main

type Skills struct {
	Languages  string `json:"languages"`
	Frameworks string `json:"frameworks"`
	Tools      string `json:"tools"`
}

type Socials struct {
	Github    string `json:"github"`
	Linkedin  string `json:"linkedin"`
	Instagram string `json:"instagram"`
	Discord   string `json:"discord"`
	X         string `json:"x"`
}

type Developer struct {
	Name string `json:"name"`
}
