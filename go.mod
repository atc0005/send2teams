module github.com/atc0005/send2teams

require (
	//gopkg.in/dasrick/go-teams-notify.v1 v1.2.0

	// temporarily use our fork while developing changes for potential
	// inclusion in the upstream project
	github.com/atc0005/go-teams-notify v1.3.1-0.20200327233154-4876da39881a
	github.com/stretchr/testify v1.4.0 // indirect
	gopkg.in/yaml.v2 v2.2.4 // indirect
)

go 1.13
