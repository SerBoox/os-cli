package views

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/serboox/os-cli/src/models"
)

// ShowNewServer send information in stdout
func ShowNewServer(server *models.ResServers) {
	columns := []string{
		"Instance ID",
		"Status",
		"AdminPass",
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(columns)
	table.SetAutoMergeCells(false)
	table.SetRowLine(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Append([]string{
		server.Server.ID,
		"success",
		server.Server.AdminPass,
	})
	table.Render()
}
