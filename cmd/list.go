package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/stikypiston/jots/internal/storage"
	"github.com/stikypiston/jots/internal/ui"
)

var listDate string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List journal entries",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Default to today
		var t time.Time
		var err error

		if listDate != "" {
			t, err = time.Parse("2006-01-02", listDate)
			if err != nil {
				return fmt.Errorf("invalid date format, use YYYY-MM-DD")
			}
		} else {
			t = time.Now()
		}

		df, err := storage.LoadDay(t)
		if err != nil {
			return err
		}

		if len(df.Entries) == 0 {
			fmt.Println("No entries for this day.")
			return nil
		}

		for _, e := range df.Entries {
			fmt.Println(ui.FormatEntry(e))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&listDate, "date", "d", "", "Specify date to list entries for")
}
