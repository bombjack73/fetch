// Copyright © 2017 Marc Vandenbosch
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type GitHubRepository struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	HtmlUrl     string `json:"html_url"`
	Fork        bool   `json:"fork"`
}

type GitHubSearchResult struct {
	Items []GitHubRepository `json"items"`
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		for _, arg := range args {
			fmt.Println("search called: " + arg)
			//default: in:name,description
			//user:username in case of username/repo
			//+fork:true
			resp, err := http.Get("https://api.github.com/search/repositories?q=" + arg + "+language:go&order=desc")
			//todo see err handling in go source
			if err != nil {
				fmt.Println("error") //todo explain
				continue
			}
			defer resp.Body.Close()

			res := new(GitHubSearchResult)
			json.NewDecoder(resp.Body).Decode(res)
			fmt.Println(len(res.Items))
			for i := 0; i < len(res.Items); i++ {
				fmt.Printf("%+v\n", res.Items[i])
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
