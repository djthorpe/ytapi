package main

import (
	"os"
	"flag"
	"fmt"
	"os/user"
	"path/filepath"
)

var (
	flagCredentials         = flag.String("credentials",".credentials","Folder containing credentials")
	flagClientSecret        = flag.String("clientsecret","client_secret.json","Client secret filename")
)

////////////////////////////////////////////////////////////////////////////////

func userDir() string {
	currentUser, _ := user.Current()
	return currentUser.HomeDir
}

// TODO: Make these relative to home directory, not absolute

func GetCredentialsPath() string {
	return filepath.Join(userDir(),*flagCredentials)
}

func GetClientSecretPath() string {
	return filepath.Join(GetCredentialsPath(),*flagClientSecret)
}

////////////////////////////////////////////////////////////////////////////////

func main() {
	flag.Parse()
	fmt.Println(GetClientSecretPath())
	os.Exit(0)
}

