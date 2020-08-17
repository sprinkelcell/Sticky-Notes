package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sprinkelcell/Sticky-Notes/db"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a note from list",

	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {

			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("'%s' this id is not correct\n", arg)
			} else {
				ids = append(ids, id)
			}
		}
		notes, err := db.GetNoteList()
		if err != nil {
			fmt.Printf("Something went wrong. Please try after some time")
			os.Exit(1)
		}
		for _, id := range ids {
			if id <= 0 || id > len(notes) {
				fmt.Printf("This id '%d' is incorrect", id)
				continue
			}
			err := db.RemoveNote(notes[id-1].Key)
			if err != nil {
				fmt.Printf("Something went wrong. Please try after some time")
				os.Exit(1)
			}
			fmt.Printf("id no '%d' removed from list\n", id)
		}
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)

}
