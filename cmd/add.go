package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sprinkelcell/Sticky-Notes/db"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add note in your sticky notes",

	Run: func(cmd *cobra.Command, args []string) {
		short := strings.Join(args, " ")
		
		fmt.Print("Enter Description:-")
		reader := bufio.NewReader(os.Stdin)

		long,err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("Something went wrong. Please try after some time")
			os.Exit(1)
		}
		_, err = db.CreateNote(db.Description{Short: short, Long: long})
		if err != nil {
			fmt.Printf("Something went wrong. Please try after some time")
			os.Exit(1)
		}
		fmt.Println("note added successfully.")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)

}
