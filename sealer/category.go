package sealer

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

type Category struct {
	TubeDiameter       string
	Success            string
	Preheat            string
	PreheatCurrent     string
	PreheatCoilVoltage string
	MaxPreheatFreq     string
	WeldReport         string
	WeldCurrent        string
	WeldCoilVoltage    string
	MaxWeldFreq        string
	WeldTimeIDS        string
	WeldTimeTSM        string
	HeatLevel          string
	CoilNoise          string
	ISDSupplyVoltage   string
	ISDMinVoltage      string
}

func ParseCategory(in string) (Category, error) {
	category := Category{}
	matcher := regexp.MustCompile(`[>a-zA-Z :]+(?P<td>[\d.]+)[>a-zA-Z ,]+: (?P<report>[a-zA-Z|\d .]+)`)
	matches := matcher.FindStringSubmatch(in)
	if len(matches) != 3 {
		return category, errors.New("failed to match on regex")
	}

	category.TubeDiameter = matches[1]

	report := strings.Split(matches[2], "|")
	if len(report) != 15 {
		return category, errors.New("not enough fields in heater report")
	}
	category.Success = report[0]
	category.Preheat = report[1]
	category.PreheatCurrent = report[2]
	category.PreheatCoilVoltage = report[3]
	category.MaxPreheatFreq = report[4]
	category.WeldReport = report[5]
	category.WeldCurrent = report[6]
	category.WeldCoilVoltage = report[7]
	category.MaxWeldFreq = report[8]
	category.WeldTimeIDS = report[9]
	category.WeldTimeTSM = report[10]
	category.HeatLevel = report[11]
	category.CoilNoise = report[12]
	category.ISDSupplyVoltage = report[13]
	category.ISDMinVoltage = report[14]
	return category, nil
}
