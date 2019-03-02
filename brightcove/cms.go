package brightcove

import (
	"errors"

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
			Format:      FormatVideoCount,
			Brightcove:  GetVideoCount,
		},
		&util.Command{
			Name:        "GetVideos",
			Description: "Get videos in account",
			Format:      FormatVideos,
			Brightcove:  GetVideos,
		},
	}
}

/////////////////////////////////////////////////////////////////////

func FormatVideoCount(output *util.Table) error {
	if err := output.AddColumn("video_count"); err != nil {
		return err
	}

	// Success
	return nil
}

func GetVideoCount(client *brightcoveapi.Client, output *util.Table, args []string) error {
	if len(args) != 0 {
		return errors.New("Too many arguments")
	}
	if count, err := client.CMS.GetVideoCount(); err != nil {
		return err
	} else {
		output.Append(map[string]uint{"video_count": count})
	}

	// Success
	return nil
}

/////////////////////////////////////////////////////////////////////

func FormatVideos(output *util.Table) error {
	if err := output.AddColumn("id"); err != nil {
		return err
	}

	// Success
	return nil
}

func GetVideos(client *brightcoveapi.Client, output *util.Table, args []string) error {
	if len(args) != 0 {
		return errors.New("Too many arguments")
	}
	if videos, err := client.CMS.GetVideos(); err != nil {
		return err
	} else {
		for _, video := range videos {
			output.Append(video)
		}
	}

	// Success
	return nil
}
