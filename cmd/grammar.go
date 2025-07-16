package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/eyewritecode/jlpt-crawler/internal/crawler"
	"github.com/eyewritecode/jlpt-crawler/internal/utils"
	"github.com/spf13/cobra"
)

var grammarCmd = &cobra.Command{
	Use:   "grammar",
	Short: "Used to download grammar list cards for the specified JLPT level",
	Run: func(cmd *cobra.Command, args []string) {
		level, err := cmd.Flags().GetString("level")
		if err != nil {
			fmt.Println("Error retrieving --level flag:", err)
			os.Exit(1)
		}
		url, err := getGrammarListUrl(level)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		items, err := crawler.FetchAllGrammarItems(url)
		if err != nil {
			fmt.Println("Failed to fetch grammar items:", err)
			os.Exit(1)
		}
		fmt.Printf("Found %d grammar points for JLPT %s:\n", len(items), level)
		for i, item := range items {
			fmt.Printf("%d. %s: %s\n", i+1, item.Word, item.DetailLink)
		}
	},
}

func init() {
	rootCmd.AddCommand(grammarCmd)
}

func getGrammarListUrl(level string) (string, error) {
	url, exists := utils.GRAMMAR_URL[strings.ToLower(level)]
	if !exists {
		return "", fmt.Errorf("JLPT Level Not Found")
	}
	return url, nil
}
