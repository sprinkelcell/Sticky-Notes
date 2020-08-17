package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/sprinkelcell/Sticky-Notes/db"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a description of a particular note",

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
			note, err := db.GetNoteByKey(notes[id-1].Key)
			if err != nil {
				fmt.Printf("Something went wrong. Please try after some time")
				os.Exit(1)
			}
			fmt.Println("-------------------------------")
			fmt.Println(note.Value.Short)
			fmt.Printf("Description -> %s\n", note.Value.Long)
		}
	},
}

func init() {
	RootCmd.AddCommand(showCmd)
}
