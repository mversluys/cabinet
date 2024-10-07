package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

type mame struct {
	Build    string    `xml:"mame"`
	Debug    string    `xml:"debug,attr"`
	Mamecfg  int       `xml:"mameconfig,attr"`
	Machines []machine `xml:"machine"`
}

type machine struct {
	Name         string `xml:"name,attr"`
	Sourcefile   string `xml:"sourcefile,attr"`
	Sampleof     string `xml:"sampleof,attr"`
	Description  string `xml:"description"`
	Year         string `xml:"year"`
	Manufacturer string `xml:"manufacturer"`
	Input        input  `xml:"input"`
}

type input struct {
	Players  string    `xml:"players,attr"`
	Coins    string    `xml:"coins,attr"`
	Service  string    `xml:"service,attr"`
	Controls []control `xml:"control"`
}

type control struct {
	Type    string `xml:"type,attr"`
	Player  string `xml:"player,attr"`
	Buttons string `xml:"buttons,attr"`
	Ways    string `xml:"ways,attr"`
}

const MAME_EXECUTABLE = "mame"

func mameGetMachines(appPath string, romset ...string) (map[string]machine, error) {
	result := &mame{}
	command := filepath.Join(appPath, MAME_EXECUTABLE)

	args := append([]string{"-listxml"}, romset...)

	log.Println("reading mame information", romset)
	cmd := exec.Command(command, args...)
	cmd.Dir = appPath

	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error getting output: %w", err)
	}

	err = xml.Unmarshal(stdout, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling XML: %w", err)
	}

	machinesByName := make(map[string]machine)

	for _, m := range result.Machines {
		machinesByName[m.Name] = m
	}

	return machinesByName, nil
}

func mamePlay(ctx context.Context, appPath string, rom string) error {
	log.Println("playing", rom)

	command := filepath.Join(appPath, MAME_EXECUTABLE)
	cmd := exec.Command(command, rom)
	cmd.Dir = appPath

	return cmd.Run()
}
