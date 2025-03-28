/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	urPath string
	client = http.Client{
		// Transport: &http.Transport{
		// 	Dial: net.Dialer{
		// 		Timeout: 30 * time.Second.Dial,
		// 	},
		// },
		Timeout: 2 * time.Second,
	}
)

func ping(domain string) (int, error) {
	url := "http://" + domain
	req, err := http.NewRequest("HEAD", url, nil)

	if err != nil {
		return 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

// net/pingCmd represents the net/ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "This ping command is used to ping a URL and print the response",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if resp, err := ping(urPath); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	pingCmd.Flags().StringVarP(&urPath, "url", "u", "", "URL to ping")

	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
	}

	NetCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// net/pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// net/pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
