package runner

import (
	"github.com/khulnasoft-lab/gologger"
	updateutils "github.com/khulnasoft-lab/utils/update"
)

const banner = `
    __    __  __       _  __
   / /_  / /_/ /_____ | |/ /
  / __ \/ __/ __/ __ \|   /
 / / / / /_/ /_/ /_/ /   |
/_/ /_/\__/\__/ .___/_/|_|
             /_/
`

// Version is the current version of httpx
const version = `v1.3.6`

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\tkhulnasoft-lab.io\n\n")
}

// GetUpdateCallback returns a callback function that updates httpx
func GetUpdateCallback() func() {
	return func() {
		showBanner()
		updateutils.GetUpdateToolCallback("httpx", version)()
	}
}
