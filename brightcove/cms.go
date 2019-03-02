package brightcove

import (
	"errors"
	"fmt"

	// Frameworks
	"github.com/djthorpe/ytapi/brightcoveapi"
	"github.com/djthorpe/ytapi/util"
)

////////////////////////////////////////////////////////////////////////////////
// Register CMS commands

func RegisterCMSCommands() []*util.Command {
	return []*util.Command{
		&util.Command{
			Name:        "GetVideoCount",
			Description: "Get count of videos in account",
			Brightcove:  GetVideoCount,
		},
		&util.Command{
			Name:        "GetVideos",
			Description: "Get videos in account",
			Brightcove:  GetVideos,
		},
	}
}

func GetVideoCount(client *brightcoveapi.Client, args []string) error {
	if len(args) != 0 {
		return errors.New("Too many arguments")
	}
	if count, err := client.CMS.GetVideoCount(); err != nil {
		return err
	} else {
		fmt.Println(count)
	}

	// Success
	return nil
}

func GetVideos(client *brightcoveapi.Client, args []string) error {
	if len(args) != 0 {
		return errors.New("Too many arguments")
	}
	if videos, err := client.CMS.GetVideos(); err != nil {
		return err
	} else {
		fmt.Println(videos)
	}

	// Success
	return nil
}
