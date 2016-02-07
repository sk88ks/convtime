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
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

const (
	timeResTmpl = "%s -> %s\n"
)

var (
	locationStr string
	toFormatStr string
)

func convertToTime(unixtimeStr string, location *time.Location, layout string) string {
	ut, err := strconv.ParseInt(unixtimeStr, 10, 64)
	if err != nil {
		return ""
	}

	l := layoutMap[layout]

	t := time.Unix(ut, 0).In(location)
	if l != "" {
		return t.Format(l)
	}
	return t.String()
}

// unixtoCmd represents the unixto command
var unixtoCmd = &cobra.Command{
	Use:   "unixto",
	Short: "Unitxto convets unixtime seconds to formatted time expression",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		l, err := time.LoadLocation(locationStr)
		if err != nil {
			fmt.Println("Not found location")
			return
		}

		for i := range args {
			fmt.Printf(timeResTmpl, args[i], convertToTime(args[i], l, toFormatStr))
		}

	},
}

func init() {
	RootCmd.AddCommand(unixtoCmd)

	//unixtoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	unixtoCmd.Flags().StringVarP(&locationStr, "location", "l", "", "Location setting")
	unixtoCmd.Flags().StringVarP(&toFormatStr, "format", "f", "", "Time expression layout format")
}
