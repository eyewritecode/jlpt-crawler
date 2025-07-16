package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/eyewritecode/jlpt-crawler/internal/crawler"
	"github.com/eyewritecode/jlpt-crawler/internal/utils"
	"github.com/schollz/progressbar/v3"
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
		bar := progressbar.Default(int64(len(items)), "Downloading...")
		for _, item := range items {
			err := crawler.DownloadGrammarCard(item.DetailLink, item.Word, "images")
			if err != nil {
				fmt.Printf("\nFailed to download %s: %v\n", item.Word, err)
			}
			bar.Add(1)
		}
		fmt.Println("\nDownload complete.")
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
