// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

	"github.com/spf13/cobra"

	"github.com/seiji/cache"
)

var deleteAllCmd = &cobra.Command{
	Use:   "delete_all",
	Short: "Delete all object from memcached",
	Long:  "Delete all object from memcached",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := cache.New(host, port)
		var err error
		if err = c.DeleteAll(); err != nil && !cache.ResumableErr(err) {
			return err
		}

		if err == nil {
			fmt.Println("OK")
		} else {
			fmt.Println(err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteAllCmd)
}
