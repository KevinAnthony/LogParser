package sealer

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func Parse() *cli.Command {
	var outfile, infile string
	return &cli.Command{
		Name:  "sealer",
		Usage: "parse sealer log file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "output file name",
				Destination: &outfile,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "input",
				Aliases:     []string{"i"},
				Usage:       "input file name",
				Destination: &infile,
				Required:    true,
			},
		},
		Action: func(ctx *cli.Context) error {
			data, err := ioutil.ReadFile(infile)
			if err != nil {
				return errors.Wrap(err, "failure to open input log")
			}
			var event EventLog
			if err := xml.Unmarshal(data, &event); err != nil {
				return errors.Wrap(err, "failure to parse xml logs")
			}
			for i := range event.Entries {
				cat, err := ParseCategory(event.Entries[i].CategoryStr)
				if err != nil {
					return errors.Wrap(err, "failure to parse category string")
				}
				event.Entries[i].Category = cat
			}

			return output(outfile, event)
		},
	}
}

func output(outname string, event EventLog) error {
	file, err := os.Create(outname)
	if err != nil {
		return errors.Wrap(err, "failed to open output for writing")
	}
	defer file.Close()

	//nolint: lll
	if _, err := fmt.Fprintln(file, "Computer Name, Software Ver, Log Ver, Component Name, Process, Process ID, Creation DateTime, TimeZone Offset In Minutes, Tread ID, System Date Time, Success, Preheat, Preheat Current, Preheat Coil Voltage, Max Preheat Frequency, Weld Report, Weld Coil Voltage, Max Weld Frequency, Weld Time IDS, Weld Time TSM, Heat Level, Coil Noise, ISD Supply Voltage, ISD Minimum Voltage"); err != nil {
		return errors.Wrap(err, "failed to print line to file")
	}

	for _, ent := range event.Entries {
		if _, err := fmt.Fprintf(file, "%s,%s,%s,%s,%s,%s,%s,%s,%d,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s",
			event.ComputerName,
			event.SoftwareVersion,
			event.LogVersion,
			event.ComponentName,
			event.Process,
			event.ProcessID,
			event.CreationDateTime,
			event.TimeZoneOffsetMinutes,
			ent.ThreadID,
			ent.SysDateTime,
			ent.Category.Success,
			ent.Category.Preheat,
			ent.Category.PreheatCurrent,
			ent.Category.PreheatCoilVoltage,
			ent.Category.MaxPreheatFreq,
			ent.Category.WeldReport,
			ent.Category.WeldCoilVoltage,
			ent.Category.MaxWeldFreq,
			ent.Category.WeldTimeIDS,
			ent.Category.WeldTimeTSM,
			ent.Category.HeatLevel,
			ent.Category.CoilNoise,
			ent.Category.ISDSupplyVoltage,
			ent.Category.ISDMinVoltage,
		); err != nil {
			return errors.Wrap(err, "failed to print line to file")
		}

		// this is for cross-platform compatibility
		if _, err := fmt.Fprintln(file, ""); err != nil {
			return errors.Wrap(err, "failed to print line to file")
		}
	}
	return nil
}
