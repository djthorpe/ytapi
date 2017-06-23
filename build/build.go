/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type Settings struct {
	Branch         string `json:"branch"`
	Tag            string `json:"tag"`
	Hash           string `json:"hash"`
	Date           string `json:"date"`
	GoVersion      string `json:"goversion"`
	ClientSecret   string `json:"client_secret"`
	ServiceAccount string `json:"service_account"`
}

////////////////////////////////////////////////////////////////////////////////

func decodefile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func main() {
	var settings Settings

	flag.Parse()

	// Read in JSON file
	args := flag.Args()
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Syntax error: requires input json and template arguments")
		os.Exit(1)
	}
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()
	json := json.NewDecoder(file)
	if err = json.Decode(&settings); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Convert Client Secret and Service Account
	if settings.ClientSecret != "" {
		if settings.ClientSecret, err = decodefile(settings.ClientSecret); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	if settings.ServiceAccount != "" {
		if settings.ServiceAccount, err = decodefile(settings.ServiceAccount); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	// Use template to output
	t, err := template.ParseFiles(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	t.Execute(os.Stdout, settings)
}
