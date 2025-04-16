package main

import (
	"errors"
	"fmt"
)

type TargetSystem int8

var (
	UnknownSystem = errors.New("UnknownSystem")
)

type TargetError struct {
	Err   error
	Value string
}

func (te *TargetError) Error() (err string) {
	err = ""

	switch te.Err {
	case UnknownSystem:
		err = fmt.Sprintf(
			"%v: got = %v", te.Err, te.Value,
		)
	}
	return
}

const (
	Android TargetSystem = iota
	Linux
	Windows
)

var supportedSystems map[string]TargetSystem = map[string]TargetSystem{
	"android": Android,
	"linux":   Linux,
	"windows": Windows,
}

func getTarget(targetName string) (ts TargetSystem, err error) {
	ts = -1
	for k, v := range supportedSystems {
		if targetName == k {
			ts = v
			break
		}
	}

	if ts == -1 {
		err = &TargetError{
			Err:   UnknownSystem,
			Value: targetName,
		}
	}
	return
}
