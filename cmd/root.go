package cmd

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sando",
	Short: "Display progress bars from brandonsanderson.com",
	Long:  "Display progress bars from brandonsanderson.com",

	Run: func(cmd *cobra.Command, args []string) {
		c := colly.NewCollector(
			colly.AllowedDomains("www.brandonsanderson.com"),
		)

		c.OnHTML("small", func(e *colly.HTMLElement) {
			fmt.Println(e.Text)
		})

		error := c.Visit("https://www.brandonsanderson.com/")
		if error != nil {
			fmt.Println(error)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
