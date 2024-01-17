/*
(c) Copyright 2018, Gemalto. All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"log"

	helmscan "scan/pkg/helmscan"

	"github.com/spf13/cobra"
)

var (
	globalUsage = `
Scans helm charts for vulnerebility checks, hardcoded secrets, keys, passwords etc
`
)

var version = "SNAPSHOT"

func NewRootCmd() *cobra.Command {

	// s := &helmscan.Scan{}

	cmd := &cobra.Command{
		Use:          "scan",
		Short:        fmt.Sprintf("(helm-scan %s)", version),
		Long:         globalUsage,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {

			log.Println("args", args)

			if len(args) == 0 {
				return errors.New("this command needs at least 1 argument: chart name")
			} else if len(args) > 1 {
				return errors.New("this command accepts only 1 argument: chart name")
			}

			chartDirectoryPath := args[0]

			cmdResult := helmscan.Scan(chartDirectoryPath)
			log.Println("CMD RES", cmdResult)
			return cmdResult
		},
	}

	return cmd
}
