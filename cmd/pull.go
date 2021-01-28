/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"io"
	"net/http"
	url2 "net/url"
	"os"

	"github.com/spf13/cobra"
)

const (
	url = "http://containeros.cn:8000/"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull image from the remote repository",
	Long: `Usage: unikernel pull <image name>`,
	Run: func(cmd *cobra.Command, args []string) {
		//f, err := os.OpenFile(args[0], os.O_RDWR|os.O_CREATE, 0666)
		//if err != nil {
		//	panic(err)
		//}
		var req http.Request
		var err error
		req.Method = "GET"
		req.Close = true
		req.URL, err = url2.Parse(url + args[0])
		if err != nil {
			panic(err)
		}
		header := http.Header{}
		req.Header = header
		resp, err := http.DefaultClient.Do(&req)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode == 404 {
			println("Can not find image in the remote repository: ", args[0])
		} else {
			f, err := os.OpenFile(args[0], os.O_RDWR | os.O_CREATE, 0666)
			f.Seek(0, 0)
			_, err = io.Copy(f, resp.Body)
			if err != nil {
				panic(err)
			}
			println("Image has been pulled from the remote repository: ", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
