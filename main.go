package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	cmd := cobra.Command{
		Use:                        "use to show hello world",
		Aliases:                    nil,
		SuggestFor:                 nil,
		Short:                      "hw",
		GroupID:                    "",
		Long:                       "",
		Example:                    "show hello world",
		ValidArgs:                  nil,
		ValidArgsFunction:          nil,
		Args:                       nil,
		ArgAliases:                 nil,
		BashCompletionFunction:     "",
		Deprecated:                 "",
		Annotations:                nil,
		Version:                    "",
		PersistentPreRun:           nil,
		PersistentPreRunE:          nil,
		PreRun:                     nil,
		PreRunE:                    nil,
		Run:                        run,
		RunE:                       nil,
		PostRun:                    nil,
		PostRunE:                   nil,
		PersistentPostRun:          nil,
		PersistentPostRunE:         nil,
		FParseErrWhitelist:         cobra.FParseErrWhitelist{},
		CompletionOptions:          cobra.CompletionOptions{},
		TraverseChildren:           false,
		Hidden:                     false,
		SilenceErrors:              false,
		SilenceUsage:               false,
		DisableFlagParsing:         false,
		DisableAutoGenTag:          false,
		DisableFlagsInUseLine:      false,
		DisableSuggestions:         false,
		SuggestionsMinimumDistance: 0,
	}
	if err := cmd.Execute(); err != nil {
		log.Fatalf("cobra err=%v\n", err)
	}

	fmt.Printf("successful")
}

func run(cmd *cobra.Command, args []string) {
	fmt.Printf("cobra-command=%s, args=%v\n", cmd.Use, args)

}

func what() {
	fmt.Printf("hello world")
}
