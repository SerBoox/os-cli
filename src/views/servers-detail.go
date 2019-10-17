package views

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/serboox/os-cli/src/models"
)

// ShowServersDetail send information in stdout
func ShowServersDetail(servers models.ResServersDetail) (err error) {
	columns := []string{
		"Instance Name",
		"Image ID",
		"IP Address",
		"FlavorID",
		"Status",
		"Availability Zone",
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(columns)
	table.SetAutoMergeCells(false)
	table.SetRowLine(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for i := range servers.Servers {
		for j := range servers.Servers[i].Addresses.Public {
			var imageID string
			switch image := servers.Servers[i].Image.(type) {
			case string:
				imageID = image
			case map[string]interface{}:
				if value, ok := image["id"].(string); ok {
					imageID = value
				} else {
					return fmt.Errorf("imageID not be found in server response")
				}
			default:
				return fmt.Errorf("value type ImageID not defined")
			}

			table.Append([]string{
				servers.Servers[i].Name,
				imageID,
				servers.Servers[i].Addresses.Public[j].Addr,
				servers.Servers[i].Flavor.ID,
				servers.Servers[i].Status,
				servers.Servers[i].OSExtAz,
			})
		}
	}

	table.Render()

	return
}
