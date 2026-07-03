package main

import (
	"log"
  "net/http"
)

type ProGuy struct {
	Name     string   `json:"name"`
	Bio      string   `json:"bio"`
	Stack    []string `json:"stack"`
	Location string   `json:"location"`
}

type Person struct{}

type Skills struct {
	Languages  []string `json:"languages"`
	Frameworks []string `json:"frameworks"`
	Tools      []string `json:"tools"`
}

type Socials struct {
	Github    string `json:"github"`
	Linkedin  string `json:"linkedin"`
	Portfolio string `json:"portfolio"`
}

type Developer struct {
	Name         string   `json:"name"`
	Location     string   `json:"location"`
	CurrentRole  string   `json:"currentRole"`
	Availability string   `json:"availability"`
	Skills       Skills   `json:"skills"`
	Interests    []string `json:"interests"`
	FunFacts     []string `json:"funFacts"`
	Socials      Socials  `json:"socials"`
}

type PortfolioResponse struct {
	Schema    string    `json:"$schema"`
	Developer Developer `json:"developer"`
}

func main() {
	app := http.NewServeMux()

	app.HandleFunc("/portfolio", handlePortfolio)

	err := http.ListenAndServe(":8000", app)

	log.Fatal(err)
}

func handlePortfolio(w http.ResponseWriter, r *http.Request) {
	profile := PortfolioResponse{
		Schema: "https://json-schema.org",
		Developer: Developer{
			Name:         "Raymond Nicholas",
			Location:     "Otukpo, Nigeria",
			CurrentRole:  "Software Engineer",
			Availability: "Open to opportunities",
			Skills: Skills{
				Languages:  []string{"JavaScript", "Python", "Go"},
				Frameworks: []string{"React", "Node.js", "FastAPI", "Fiber"},
				Tools:      []string{"Git", "Docker", "VS Code", "Bash"},
			},
			Interests: []string{
				"Building open-source tools",
				"System architecture and design",
				"Ascii Art Web Generation",
			},
			FunFacts: []string{
				"I drink too much coffee",
				"I can debug code in my sleep",
			},
			Socials: Socials{
				Github:    "https://github.com",
				Linkedin:  "https://linkedin.com",
				Portfolio: "https://your-portfolio.com",
			},
		},
	}

	return c.Status(fiber.StatusOK).JSONMarshalIndent(profile, "", "  ")
}
