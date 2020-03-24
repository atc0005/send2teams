module github.com/atc0005/send2teams

require (
	//gopkg.in/dasrick/go-teams-notify.v1 v1.2.0

	// temporarily use our fork until upstream webhook URL FQDN validation
	// changes can be made
	github.com/atc0005/go-teams-notify v1.2.1-0.20200324114153-d5bf5e1bebf3
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	gopkg.in/yaml.v2 v2.2.4 // indirect
)

go 1.13
