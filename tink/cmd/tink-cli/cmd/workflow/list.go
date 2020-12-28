package workflow

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"github.com/tinkerbell/tink/client"
	"github.com/tinkerbell/tink/protos/workflow"
)

var (
	quiet bool
	t     table.Writer

	hCreatedAt = "Created At"
	hUpdatedAt = "Updated At"
)

// listCmd represents the list subcommand for workflow command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "list all workflows",
	Example: "tink workflow list",
	Args: func(c *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("%v takes no arguments", c.UseLine())
		}
		return nil
	},
	Run: func(c *cobra.Command, args []string) {
		if quiet {
			listWorkflows()
			return
		}
		t = table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{hID, hTemplate, hDevice, hCreatedAt, hUpdatedAt})
		listWorkflows()
		t.Render()
	},
}

func listWorkflows() {
	list, err := client.WorkflowClient.ListWorkflows(context.Background(), &workflow.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	var wf *workflow.Workflow
	for wf, err = list.Recv(); err == nil && wf.Id != ""; wf, err = list.Recv() {
		printOutput(wf)
	}

	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
}

func printOutput(wf *workflow.Workflow) {
	if quiet {
		fmt.Println(wf.Id)
	} else {
		cr := wf.CreatedAt
		up := wf.UpdatedAt
		t.AppendRows([]table.Row{
			{wf.Id, wf.Template, wf.Hardware, time.Unix(cr.Seconds, 0), time.Unix(up.Seconds, 0)},
		})
	}
}

func addListFlags() {
	flags := listCmd.Flags()
	flags.BoolVarP(&quiet, "quiet", "q", false, "only display workflow IDs")
}

func init() {
	addListFlags()
	SubCommands = append(SubCommands, listCmd)
}
