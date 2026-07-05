package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Skills struct {
	Languages  []string `json:"languages"`
	Frameworks []string `json:"frameworks"`
	Tools      []string `json:"tools"`
	//Databases  []string `json:"databases"`
	//Cloud      []string `json:"cloud"`
}

type Socials struct {
	Github    string `json:"github"`
	Linkedin  string `json:"linkedin"`
	Portfolio string `json:"portfolio"`
	Twitter   string `json:"twitter,omitempty"`
	YouTube   string `json:"youtube,omitempty"`
	DevTo     string `json:"devto,omitempty"`
}

type Project struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	CodeLink    string   `json:"codeLink"`
	TechStack   []string `json:"techStack"`
	Image       string   `json:"image,omitempty"`
	Featured    bool     `json:"featured"`
	Year        int      `json:"year"`
}

type Education struct {
	Institution string `json:"institution"`
	Field       string `json:"field"`
}

type Developer struct {
	Name         string `json:"name"`
	Title        string `json:"title"`
	Location     string `json:"location"`
	Email        string `json:"email"`
	Bio          string `json:"bio"`
	Tagline      string `json:"tagline"`
	Avatar       string `json:"avatar,omitempty"`
	CurrentRole  string `json:"currentRole"`
	Availability string `json:"availability"`
	YearsOfExp   int    `json:"yearsOfExp"`

	// Sections
	Skills     Skills       `json:"skills"`
	Socials    Socials      `json:"socials"`
	Projects   []Project    `json:"projects"`
	Education  []Education  `json:"education"`

	// Extras
	Interests []string `json:"interests"`
	FunFacts  []string `json:"funFacts"`
	Languages []string `json:"languages"`
	Hobbies   []string `json:"hobbies"`
}

type PortfolioResponse struct {
	Schema    string    `json:"$schema"`
	Developer Developer `json:"developer"`
}

func handlePortfolio(w http.ResponseWriter, r *http.Request) {
	profile := PortfolioResponse{
		Schema: "https://json-schema.org",
		Developer: Developer{
			Name:     "Raymond Nicholas",
			Title:    "Software Engineer & Open Source Enthusiast",
			Location: "Otukpo, Nigeria",
			Email:    "raymond@example.com",
			Bio:      "Passionate software engineer building scalable systems and contributing to open source.",
			Tagline:  "Code. Create. Innovate.",
			Avatar:   "https://avatar.github.com/raymondproguy.png",

			CurrentRole:  "Software Engineer",
			Availability: "Open to opportunities",
			YearsOfExp:   5,

			Skills: Skills{
				Languages:  []string{"JavaScript", "Python", "Go", "TypeScript"},
				Frameworks: []string{"React", "Node.js", "FastAPI", "Fiber", "Next.js"},
				Tools:      []string{"Git", "Docker", "VS Code", "Bash", "Postman"},
				//Databases:  []string{"PostgreSQL", "MongoDB", "Redis", "MySQL"},
				//Cloud:      []string{"AWS", "GCP", "DigitalOcean", "Azure"},
			},

			Socials: Socials{
				Github:    "https://github.com/raymondproguy",
				Linkedin:  "https://linkedin.com/in/raymondproguy",
				Portfolio: "https://raymondproguy.dev",
				Twitter:   "https://twitter.com/raymondproguy",
				DevTo:     "https://dev.to/raymondproguy",
			},

			Education: []Education{
				{
					Institution: "Benpoly",
					Field:       "Computer Science",
				},
			},

			// Projects
			Projects: []Project{
				{
					Name:        "Portfolio API",
					Description: "RESTful API for developer portfolio with Go and Fiber",
					Link:        "https://api.raymond.dev",
					CodeLink:    "https://github.com/raymond/portfolio-api",
					TechStack:   []string{"Go", "JSON"},
					Featured:    true,
					Year:        2026,
				},
				{
					Name:        "ASCII Art Generator",
					Description: "Convert text to ASCII art in your browser",
					Link:        "https://ascii.raymond.dev",
					CodeLink:    "https://github.com/raymond/ascii-art",
					TechStack:   []string{"Go", "HTML", "CSS", "JSON"},
					Featured:    false,
					Year:        2026,
				},
			},

			// Extras
			Interests: []string{
				"Building open-source tools",
				"System architecture and design",
				"Balancing ",
			},
			FunFacts: []string{
				"I drink too much coffee ☕",
				"Using resources ",
				"Built my first website at 12",
			},
			Languages: []string{"English", "French"},
			Hobbies:   []string{"Chess", "Photography", "Relating "},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(profile); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func main() {
	app := http.NewServeMux()
	app.HandleFunc("/portfolio", handlePortfolio)

	log.Println("Server running on :8000")
	log.Fatal(http.ListenAndServe(":8000", app))
}
