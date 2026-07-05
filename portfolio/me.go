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
	Databases  []string `json:"databases"`
	Cloud      []string `json:"cloud"`
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

type Experience struct {
	Company     string   `json:"company"`
	Role        string   `json:"role"`
	Location    string   `json:"location"`
	Current     bool     `json:"current"`
	Description []string `json:"description"`
	TechUsed    []string `json:"techUsed"`
}

type Education struct {
	Institution string `json:"institution"`
	Field       string `json:"field"`
}

type Developer struct {
	Name         string   `json:"name"`
	Title        string   `json:"title"`
	Location     string   `json:"location"`
	Email        string   `json:"email"`
	Bio          string   `json:"bio"`
	Tagline      string   `json:"tagline"`
	Avatar       string   `json:"avatar,omitempty"`
	CurrentRole  string   `json:"currentRole"`
	Availability string   `json:"availability"`
	YearsOfExp   int      `json:"yearsOfExp"`
	
	// Sections
	Skills        Skills         `json:"skills"`
	Socials       Socials        `json:"socials"`
	Projects      []Project      `json:"projects"`
	Experience    []Experience   `json:"experience"`
	Education     []Education    `json:"education"`
	Certifications []Certification `json:"certifications"`
	Blog          []BlogPost     `json:"blog"`
	Testimonials  []Testimonial  `json:"testimonials"`
	
	// Extras
	Interests     []string       `json:"interests"`
	FunFacts      []string       `json:"funFacts"`
	Languages     []string       `json:"languages"` // Spoken: "English", "French", etc.
	Hobbies       []string       `json:"hobbies"`
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
			Avatar:   "https://avatar.example.com/raymond.png",

			CurrentRole:  "Software Engineer",
			Availability: "Open to opportunities",
			YearsOfExp:   5,

			Skills: Skills{
				Languages:  []string{"JavaScript", "Python", "Go", "TypeScript"},
				Frameworks: []string{"React", "Node.js", "FastAPI", "Fiber", "Next.js"},
				Tools:      []string{"Git", "Docker", "VS Code", "Bash", "Postman"},
				Databases:  []string{"PostgreSQL", "MongoDB", "Redis", "MySQL"},
				Cloud:      []string{"AWS", "GCP", "DigitalOcean"},
			},

			Socials: Socials{
				Github:    "https://github.com/raymond",
				Linkedin:  "https://linkedin.com/in/raymond",
				Portfolio: "https://raymondproguy.dev",
				Twitter:   "https://twitter.com/raymond",
				DevTo:     "https://dev.to/raymond",
			},

			// Experience
			Experience: []Experience{
				{
					Company:   "TechCorp Inc.",
					Role:      "Senior Software Engineer",
					Location:  "Remote",
					StartDate: "2022-01",
					EndDate:   "Present",
					Current:   true,
					Description: []string{
						"Built microservices handling 10k+ requests/sec",
						"Led a team of 5 engineers on cloud migration",
						"Reduced API response time by 40%",
					},
					TechUsed: []string{"Go", "React", "AWS", "Docker"},
				},
				{
					Company:   "StartupHub",
					Role:      "Full Stack Developer",
					Location:  "Lagos, Nigeria",
					StartDate: "2019-06",
					EndDate:   "2021-12",
					Current:   false,
					Description: []string{
						"Developed a fintech platform used by 50k+ users",
						"Integrated payment gateways and banking APIs",
					},
					TechUsed: []string{"Python", "FastAPI", "React", "PostgreSQL"},
				},
			},

			// Education
			Education: []Education{
				{
					Institution: "University of Nigeria",
					Degree:      "Bachelor of Science",
					Field:       "Computer Science",
					StartDate:   "2015-09",
					EndDate:     "2019-06",
					Grade:       "First Class Honours",
				},
			},

			// Certifications
			Certifications: []Certification{
				{
					Name:        "AWS Certified Developer",
					Issuer:      "Amazon Web Services",
					Date:        "2023-03",
					Link:        "https://aws.amazon.com/verify",
					CredentialID: "ABC123XYZ",
				},
			},

			// Projects
			Projects: []Project{
				{
					Name:        "Portfolio API",
					Description: "RESTful API for developer portfolio with Go and Fiber",
					Link:        "https://api.raymond.dev",
					CodeLink:    "https://github.com/raymond/portfolio-api",
					TechStack:   []string{"Go", "Fiber", "MongoDB"},
					Featured:    true,
					Year:        2024,
				},
				{
					Name:        "ASCII Art Generator",
					Description: "Convert images to ASCII art in your browser",
					Link:        "https://ascii.raymond.dev",
					CodeLink:    "https://github.com/raymond/ascii-art",
					TechStack:   []string{"JavaScript", "Canvas", "CSS"},
					Featured:    false,
					Year:        2023,
				},
			},

			// Blog
			Blog: []BlogPost{
				{
					Title:       "Building Scalable APIs with Go Fiber",
					Slug:        "building-scalable-apis-go-fiber",
					Description: "Learn how to build production-ready APIs with Go Fiber",
					Date:        "2024-05-15",
					Tags:        []string{"Go", "API", "Backend"},
					Link:        "https://blog.raymond.dev/go-fiber",
					ReadTime:    8,
				},
			},

			// Testimonials
			Testimonials: []Testimonial{
				{
					Name:    "Jane Doe",
					Role:    "CTO",
					Company: "TechCorp",
					Text:    "Raymond is one of the best engineers I've worked with. He delivered beyond expectations.",
					Rating:  5,
				},
			},

			// Extras
			Interests: []string{
				"Building open-source tools",
				"System architecture and design",
				"Ascii Art Web Generation",
			},
			FunFacts: []string{
				"I drink too much coffee ☕",
				"I can debug code in my sleep",
				"Built my first website at 12",
			},
			Languages: []string{"English (Fluent)", "French (Intermediate)"},
			Hobbies:   []string{"Chess", "Photography", "Gaming"},
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
