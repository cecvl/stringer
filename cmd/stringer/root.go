package stringer

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
    Use:  "stringer",
	Version: version,
    Short: "stringer - a simple CLI to transform and inspect strings",
    Long: `stringer is a super fancy CLI (initial cli attempt)
   
One can use stringer to modify or inspect strings straight from the terminal`,
    Run: func(cmd *cobra.Command, args []string) {
        // This function is currently empty because:
        // 1. It might be a placeholder for future implementation.
        // 2. The functionality might be defined elsewhere and this function serves as a stub.
        // 3. It could be intentionally left empty for testing purposes.
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}