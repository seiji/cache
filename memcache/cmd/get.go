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
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/seiji/cache"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get object from memcached",
	Long:  "Get object from memcached",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Specify key")
		}
		key := args[0]
		if err := checkKey(key); err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]
		c := cache.New(host, port)
		var item *cache.Item
		var err error
		if item, err = c.Get(key); err != nil && !cache.ResumableErr(err) {
			return err
		}
		if err == nil {
			fmt.Fprint(os.Stdout, string(item.Value[:]))
		} else {
			fmt.Println(err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
