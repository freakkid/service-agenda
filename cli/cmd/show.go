// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	service "github.com/freakkid/service-agenda/cli/service"
	"github.com/spf13/cobra"
	tools	"github.com/freakkid/service-agenda/cli/tools"
	"os"
)

// showCmd represents the show command
var ushowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show user account",
	Long: `Use this command to show every user's information.`,
	Run: func(cmd *cobra.Command, args []string) {
		limit, _ := cmd.Flags().GetString("limit")

		// validate	
		ok, message := tools.ValidateLimit(limit)
		if !ok {
			fmt.Fprintln(os.Stderr, message)
			os.Exit(1)
		}

		var Item []service.RetJson
		Item = service.ListAllUsers(limit)
		fmt.Printf("%-5s%-15s%-25s%-25s\n", "Id", "Username", "Phone number", "E-mail")
		for _, user := range Item{
			fmt.Printf("%-5d%-15s%-25s%-25s\n", user.Id, user.Username, user.Phone, user.Email)
		}
		fmt.Printf("\nTotal number is %d\n", len(Item))
	},
}

func init() {
	userCmd.AddCommand(ushowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	ushowCmd.Flags().StringP("limit", "l", "2", "limit length of the result")
}
