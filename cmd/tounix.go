// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"time"

	"github.com/spf13/cobra"
)

const (
	unixtimeResTmpl = "%s -> %d\n"
)

var (
	fromFormatStr string
)

// tounixCmd represents the tounix command
var tounixCmd = &cobra.Command{
	Use:   "tounix",
	Short: "Tounix convets time string to unixtime seconds from a given format",
	Run: func(cmd *cobra.Command, args []string) {
		layout := layoutMap[fromFormatStr]
		now := time.Now()
		for i := range args {
			if args[i] == "now" {
				fmt.Printf(unixtimeResTmpl, now.Format(layout), now.Unix())
				continue
			}

			t, _ := time.Parse(layout, args[i])
			fmt.Printf(unixtimeResTmpl, args[i], t.Unix())
		}
	},
}

func init() {
	RootCmd.AddCommand(tounixCmd)

	tounixCmd.Flags().StringVarP(&fromFormatStr, "format", "f", "", "Time expression layout format")
}
