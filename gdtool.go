package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli"

	"github.com/taryk/gdtool/connect"
	"github.com/taryk/gdtool/command/diff"
	"github.com/taryk/gdtool/command/dups"
	"github.com/taryk/gdtool/command/cache"
	"github.com/taryk/gdtool/local"
	"github.com/taryk/gdtool/remote"
	. "github.com/taryk/gdtool/core"
)

type optionalParams struct {
	recursive, cached bool
}

func getFileTree(where, what string, params *optionalParams) FileList {
	var fileList FileList
	if params.cached {
		fileList, err := cache.LoadFileTreeFromCache(where, what)
		if err == nil {
			return *fileList
		}
	}
	switch where {
	case "local":
		fileList, _ = local.GetFileList(what, params.recursive)
	case "remote":
		remote.Init()
		googleDrive := connect.Connect()
		fileList, _ = remote.GetFileList(googleDrive, what, params.recursive)
	}
	return fileList
}

func cmdList(where, what string, params *optionalParams) error {
	fileList := getFileTree(where, what, params)
	fmt.Print(FileTreeStr("", fileList))
	return nil
}

func cmdDiff(remoteId, localPath string, params *optionalParams) error {
	remoteFiles := getFileTree("remote", remoteId, params)
	localFiles := getFileTree("local", localPath, params)
	uniqueRemote, uniqueLocal := diff.Compare(remoteFiles, localFiles)
	if len(uniqueRemote) == 0 && len(uniqueLocal) == 0 {
		color.Green("The directories are equal\n")
		return nil
	}
	if len(uniqueRemote) > 0 {
		color.Yellow("The following files are only on remote\n")
		fmt.Print(FileTreeStr("", uniqueRemote))
	}
	if len(uniqueLocal) > 0 {
		color.Yellow("The following files are only on local\n")
		fmt.Print(FileTreeStr("", uniqueLocal))
	}
	return nil
}

func cmdCache(where, what string, params *optionalParams) error {
	fileList := getFileTree(where, what, params)
	cache.CacheFileList(fileList, where, what)
	return nil
}

func cmdDups(where, what string, params *optionalParams) error {
	fileList := getFileTree(where, what, params)
	groupedFileList := dups.GroupByName(fileList)
	dups.PrintDuplicates(groupedFileList)
	return nil
}

func processParams(c *cli.Context) (string, string, *optionalParams, error) {
	var where, what string
	flag:
	for _, where = range []string{"remote", "local"} {
		if c.IsSet(where) {
			what = c.String(where)
			break flag
		}
	}
	if len(what) == 0 {
		return "", "", nil,
			cli.NewExitError("--remote|--local is missing", 2)
	}
	params := prepareParams(c)
	return where, what, params, nil
}

func prepareParams(c *cli.Context) *optionalParams {
	params := &optionalParams{
		recursive: c.Bool("recursive"),
		cached: c.Bool("cached"),
	}
	return params
}

func main() {
	InitLoggers("debug", "warn", "error")
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Usage = "Manage your Google Drive data from a command-line interface"
	locationFlags := []cli.Flag {
		cli.StringFlag{
			Name: "remote",
			Usage: "ID of a folder on Google Drive",
		},
		cli.StringFlag{
			Name: "local",
			Usage: "Path on a local file system",
		},
	}
	optionalFlags := []cli.Flag {
		cli.BoolFlag{
			Name: "recursive, r",
			Usage: "Go through subfolders recursively",
		},
		cli.BoolTFlag{
			Name: "cached",
			Usage: "Use cached data rather than getting a real list of files",
		},
	}
	singleLocationFlags := append(locationFlags, optionalFlags...)
	app.Commands = []cli.Command{
		{
			Name: "list",
			Usage: "Get a list of files in a specific directory",
			Flags: singleLocationFlags,
			Action: func(c *cli.Context) error {
				where, what, params, err := processParams(c)
				if err != nil {
					return err
				}
				return cmdList(where, what, params)
			},
		},
		{
			Name: "dups",
			Usage: "Show duplicates in a specific location",
			Flags: singleLocationFlags,
			Action: func(c *cli.Context) error {
				where, what, params, err := processParams(c)
				if err != nil {
					return err
				}
				return cmdDups(where, what, params)
			},
		},
		{
			Name: "diff",
			Usage: "Show a difference between a remote and a local locations",
			Flags: optionalFlags,
			ArgsUsage: "REMOTE_FOLDER_ID LOCAL_PATH",
			Action: func(c *cli.Context) error {
				remote := c.Args().Get(0)
				local := c.Args().Get(1)
				var missing_params []string
				if len(remote) == 0 {
					missing_params[0] = "REMOTE_FOLDER_ID"
				}
				if len(local) == 0 {
					missing_params = append(missing_params, "LOCAL_PATH")
				}
				if len(missing_params) > 0 {
					return cli.NewExitError(
						"You have to specify " +
							strings.Join(missing_params, " and "), 2)
				}
				return cmdDiff(remote, local, prepareParams(c))
			},
		},
		{
			Name: "cache",
			Usage: "Cache a specific remote or local location",
			Flags: singleLocationFlags,
			Action: func(c *cli.Context) error {
				where, what, params, err := processParams(c)
				if err != nil {
					return err
				}
				return cmdCache(where, what, params)
			},
		},
	}
	app.Run(os.Args)
}
