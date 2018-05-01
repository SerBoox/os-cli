package views

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/serboox/os-cli/models"
)

//ShowServersDetail send information in stdout
func ShowServersDetail(servers models.ResServersDetail) {

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
	for _, server := range servers.Servers {
		for _, instances := range server.Addresses.Public {

			var imageID string
			switch server.Image.(type) {
			case map[string]interface{}:
				imageMap := server.Image.(map[string]interface{})
				value, ok := imageMap["id"]
				if ok {
					imageID = value.(string)
				}
			case string:
				imageID = server.Image.(string)
			}

			table.Append([]string{
				server.Name,
				imageID,
				instances.Addr,
				server.Flavor.ID,
				server.Status,
				server.OSExtAz,
			})
		}
	}
	table.Render()
}
