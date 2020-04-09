module github.com/atc0005/send2teams

// Use local copy of library package (instead of fetching remote content)
replace github.com/atc0005/go-teams-notify => ../go-teams-notify

require (
	//gopkg.in/dasrick/go-teams-notify.v1 v1.2.0

	// temporarily use our fork while developing changes for potential
	// inclusion in the upstream project
	//
	// Note: Due to `replace` directive and `v0.0.0` here, we use the current
	// state of this library package instead of fetching remote content
	github.com/atc0005/go-teams-notify v0.0.0
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

go 1.13
