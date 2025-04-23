package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/shirou/gopsutil/disk"
	"github.com/urfave/cli/v2"
)

/*
This example shows how to use commands in urfave/cli.
-> kill command: to kill a process
-> volume command: list mounted volumes
*/

type Volume struct {
	Name       string
	Total      uint64  `json:"total"`
	Used       uint64  `json:"used"`
	Available  uint64  `json:"free"`
	UsePercent float64 `json:"usedPercent"`
	MountedOn  string
}

func main() {
	app := &cli.App{
		Name:        "Kill and Volume Commands",
		Usage:       "Kill and Volume Commands",
		Description: "A simple CLI tool demonstrating process management and volume listing",
		Commands: []*cli.Command{
			{
				Name:        "kill",
				HelpName:    "kill",
				Aliases:     []string{"k"},
				ArgsUsage:   ` `, // no arguments
				Usage:       "Kill a process by PID or name",
				Description: "This command terminate a process by PID or name",
				Flags: []cli.Flag{
					&cli.UintFlag{
						Name:    "pid",
						Aliases: []string{"p"},
						Usage:   "PID of the process to kill",
					},
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "Name of the process to kill",
					},
				},
				Action: killProcessAction,
			},
			{
				Name:        "volumes",
				HelpName:    "volumes",
				Aliases:     []string{"v"},
				ArgsUsage:   ` `, // no arguments
				Usage:       "List mounted volumes",
				Description: "This command list all mounted volumes",
				Action:      volumnAction,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func killProcessAction(c *cli.Context) error {
	// no arguments are passed
	if c.NArg() > 0 {
		return errors.New("no arguments are allowed; only flags")
	}

	// id and name are passed
	if c.IsSet("pid") && c.IsSet("name") {
		return errors.New("only one of pid or name can be set")
	}

	// if name is blank and pid is blank
	if c.String("name") == "" && c.Uint("pid") == 0 {
		return errors.New("either pid or name must be set")
	}

	// grab name if it is set or pid if it is set
	name := c.String("name")
	pid := c.Uint("pid")
	if name == "" {
		name = fmt.Sprint(pid)
	}

	if err := killProcess(c); err != nil {
		return err
	}

	fmt.Printf("process %s killed\n", name)
	return nil
}

func killProcess(c *cli.Context) error {
	if c.IsSet("pid") {
		// kill process by pid
		err := exec.Command("kill", fmt.Sprint(c.Uint("pid"))).Run()

		if err != nil {
			return err
		}
	} else if c.IsSet("name") {
		// kill process by name
		var (
			errs []string
			// found  bool
		)

		processName := c.String("name")
		// get all processes
		processes, err := exec.Command("ps", "-ax").Output()
		if err != nil {
			return err
		}
		// split processes by newline
		processesSplit := splitByNewline(string(processes))
		// iterate over processes
		for _, process := range processesSplit {
			// split process by space
			processSplit := splitBySpace(process)
			// if process name is equal to process name
			if processSplit[len(processSplit)-1] == processName {
				// kill process
				err := exec.Command("kill", processSplit[0]).Run()
				if err != nil {
					errs = append(errs, err.Error())
				}
			}
		}

		if len(errs) > 0 {
			return errors.New(strings.Join(errs, "\n"))
		}
	}

	return nil
}

func splitByNewline(s string) []string {
	return strings.Split(s, "\n")
}

func splitBySpace(s string) []string {
	// Use strings.Fields which efficiently splits a string on whitespace
	return strings.Fields(s)
}

func volumnAction(c *cli.Context) error {
	stats, err := disk.Partitions(true)

	if err != nil {
		return err
	}

	var volumes []*Volume // array of pointers to Volume struct

	for _, stat := range stats {
		usage, diskErr := disk.Usage(stat.Mountpoint)

		if diskErr != nil {
			continue
		}

		volume := &Volume{
			Name:       stat.Device,
			Total:      usage.Total,
			Used:       usage.Used,
			Available:  usage.Free,
			UsePercent: usage.UsedPercent,
			MountedOn:  stat.Mountpoint,
		}
		volumes = append(volumes, volume)
	}

	// json
	json, err := json.MarshalIndent(volumes, "", "\t")
	if err != nil {
		return err
	}

	fmt.Println(string(json))
	return nil
}
