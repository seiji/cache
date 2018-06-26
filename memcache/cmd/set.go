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
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"

	"github.com/seiji/cache"
)

var (
	value []byte
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set object into memcachd",
	Long:  "Set object into memcachd",
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
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if value, err = ioutil.ReadAll(os.Stdin); err != nil {
			return err
		}
		if len(value) == 0 {
			return fmt.Errorf(`Read from stdin but no data found`)
		}

		key := args[0]
		c := cache.New(host, port)
		if err = c.Set(&cache.Item{
			Key:        key,
			Value:      value,
			Flags:      flags,
			Expiration: expiration,
		}); err != nil {
			return
		}
		fmt.Println("STORED")
		return
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.PersistentFlags().Uint32VarP(&flags, "flags", "f", 0, "A flags for object")
	setCmd.PersistentFlags().Int32VarP(&expiration, "expiration", "e", 0, "A expiration for object")
}
