/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

type joke struct {
	// Werd          string `json:"word"`
	// Definition    string
	// Pronunciation string
	Words string
}

type word struct {
	Werd          string `json:"word"`
	Definition    string
	Pronunciation string
}

//const url = "https://api.chucknorris.io/jokes/random"
const url = "https://geek-jokes.sameerkumar.website/api"

//const url = "https://random-words-api.vercel.app/word"

// jokesCmd represents the jokes command
var jokesCmd = &cobra.Command{
	Use:   "jokes",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("jokes called")
		fmt.Println("url = ", url)
		//joke := GetJoke(url)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}

		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		//fmt.Println("body = ", body)
		var joke string
		if err := json.Unmarshal(body, &joke); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(url)
		fmt.Println("joke in jokesCmd = ", joke)

	},
}

func init() {
	rootCmd.AddCommand(jokesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jokesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jokesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	fmt.Println("jokes func init()")
}

func GetJoke(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println("body = ", body)
	var joke string
	if err := json.Unmarshal(body, &joke); err != nil {
		fmt.Println(err)
		return err.Error()
	}

	fmt.Println("joke from Get Joke = ", joke)
	return joke

}
