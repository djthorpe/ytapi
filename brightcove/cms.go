package brightcove

import (
	"errors"
	"reflect"

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
		}, &util.Command{
			Name:        "GetVideoById",
			Description: "Get videos by unique id",
			Format:      FormatVideos,
			Brightcove:  GetVideoById,
		}, &util.Command{
			Name:        "GetVideoSources",
			Description: "Get video sources by unique id",
			Format:      FormatSources,
			Brightcove:  GetVideoSources,
		}, &util.Command{
			Name:        "GetAssets",
			Description: "Get video assets by unique id",
			Format:      FormatAssets,
			Brightcove:  GetAssets,
		},
	}
}

/////////////////////////////////////////////////////////////////////

func FormatVideoCount(output *util.Table) error {
	if err := output.AddColumn(&util.Field{Name: "video_count", Type: util.FIELD_UINT}); err != nil {
		return err
	}

	// Success
	return nil
}

func GetVideoCount(client *brightcoveapi.Client, output *util.Table, args []string) error {
	if len(args) != 0 {
		return errors.New("Too many arguments")
	}
	if video_count, err := client.CMS.GetVideoCount(); err != nil {
		return err
	} else if err := output.Append([][]uint{[]uint{video_count}}); err != nil {
		return err
	}

	// Success
	return nil
}

/////////////////////////////////////////////////////////////////////

func FormatVideos(output *util.Table) error {
	if err := output.AddColumnsFrom(reflect.ValueOf(brightcoveapi.Video{})); err != nil {
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
	} else if err := output.Append(videos); err != nil {
		return err
	}

	// Success
	return nil
}

func GetVideoById(client *brightcoveapi.Client, output *util.Table, args []string) error {
	if len(args) == 0 {
		return errors.New("Not enough arguments")
	}
	if videos, err := client.CMS.GetVideoById(args); err != nil {
		return err
	} else if err := output.Append(videos); err != nil {
		return err
	}

	// Success
	return nil
}

/////////////////////////////////////////////////////////////////////

func FormatSources(output *util.Table) error {
	if err := output.AddColumnsFrom(reflect.ValueOf(brightcoveapi.VideoSource{})); err != nil {
		return err
	}

	// Success
	return nil
}

func GetVideoSources(client *brightcoveapi.Client, output *util.Table, args []string) error {
	if len(args) != 1 {
		return errors.New("Missing argument")
	}
	if sources, err := client.CMS.GetVideoSources(args[0]); err != nil {
		return err
	} else if err := output.Append(sources); err != nil {
		return err
	}

	// Success
	return nil
}

/////////////////////////////////////////////////////////////////////

func FormatAssets(output *util.Table) error {
	if err := output.AddColumnsFrom(reflect.ValueOf(brightcoveapi.Asset{})); err != nil {
		return err
	}

	// Success
	return nil
}

func GetAssets(client *brightcoveapi.Client, output *util.Table, args []string) error {
	if len(args) != 1 {
		return errors.New("Missing argument")
	}
	if sources, err := client.CMS.GetAssets(args[0]); err != nil {
		return err
	} else if err := output.Append(sources); err != nil {
		return err
	}

	// Success
	return nil
}
