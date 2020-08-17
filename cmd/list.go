package cmd

import (
	"fmt"
	"os"

	"github.com/sprinkelcell/Sticky-Notes/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show list of all notes",

	Run: func(cmd *cobra.Command, args []string) {
		notes, err := db.GetNoteList()
		if err != nil {
			fmt.Printf("Something went wrong. Please try after some time")
			os.Exit(1)
		}
		if len(notes) == 0 {
			fmt.Printf("List is empty")
			return
		}
		fmt.Println("List of Notes is as follows")
		for index, note := range notes {
			fmt.Printf("%d. %s\n", index+1, note.Value.Short)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

}
