package main

type ProGuy struct {
	Name     string   `json:"name"`
	Bio      string   `json:"bio"`
	Stack    []string `json:"stack"`
	Location string   `json:"location"`
}

type Person struct{}


package main

import (
	"log"
	"://github.com"
)

// 1. Map out your nested JSON structure into strict Go structs
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
	app := fiber.New()

	// Serve the portfolio endpoint
	app.Get("/portfolio", handlePortfolio)

	log.Fatal(app.Listen(":3000"))
}

func handlePortfolio(c *fiber.Ctx) error {
	// 2. Populate your portfolio with your actual information
	profile := PortfolioResponse{
		Schema: "https://json-schema.org",
		Developer: Developer{
			Name:         "Raymond", // Change to your full name
			Location:     "Otukpo, Nigeria",
			CurrentRole:  "Software Engineer",
			Availability: "Open to opportunities",
			Skills: Skills{
				Languages:  []string{"JavaScript", "Python", "Go"},
				Frameworks: []string{"React", "Node.js", "Django", "Fiber"},
				Tools:      []string{"Git", "Docker", "VS Code"},
			},
			Interests: []string{
				"Building open-source tools",
				"System architecture",
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

	// 3. Output beautiful, indent-formatted JSON by default for readability
	return c.Status(fiber.StatusOK).JSONMarshalIndent(profile, "", "  ")
}
