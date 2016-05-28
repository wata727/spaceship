package main

import (
	"github.com/mitchellh/cli"
	"github.com/wata727/spaceship/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &command.InitCommand{
				Meta: *meta,
			}, nil
		},
		"build": func() (cli.Command, error) {
			return &command.BuildCommand{
				Meta: *meta,
			}, nil
		},
		"apply": func() (cli.Command, error) {
			return &command.ApplyCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
