/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package main

import (
    "os"
    "flag"
    "fmt"
    "encoding/json"
    "text/template"
)

type Settings struct {
    Branch string
    Tag string
    Hash string
    Date string
    GoVersion string
}

////////////////////////////////////////////////////////////////////////////////

func main() {
    var settings Settings

    flag.Parse()

    // Read in JSON file
    args := flag.Args()
    if len(args) != 2 {
        fmt.Fprintln(os.Stderr,"Syntax error: requires input json and template arguments")
        os.Exit(1)
    }
    file, err := os.Open(args[0])
    if err != nil {
        fmt.Fprintln(os.Stderr,err)
        os.Exit(1)
    }
    defer file.Close()
    json := json.NewDecoder(file)
    if err = json.Decode(&settings); err != nil {
        fmt.Fprintln(os.Stderr,err)
        os.Exit(1)
    }

    // Use template to output
    t,err := template.ParseFiles(args[1])
    if err != nil {
        fmt.Fprintln(os.Stderr,err)
        os.Exit(1)
    }
    t.Execute(os.Stdout,settings)
}
