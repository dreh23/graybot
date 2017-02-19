// Copyright 2017 Johannes Amorosa
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/urfave/cli"
)

var DebugFlag bool
var GraylogURL string

func registerFlags(app *cli.App) {
	app.Flags = []cli.Flag{

		cli.StringFlag{
			Name:        "g, logserver",
			Usage:       "Graylog server URL", //Schema: 192.168.11.29:13001
			Destination: &GraylogURL,
		},
		cli.BoolFlag{
			Name:  "D, daemon",
			Usage: "App runs forever in foreground ",
		},
		cli.BoolFlag{
			Name:  "X, crash",
			Usage: "App sporadically exits with a fatal error",
		},
		cli.BoolFlag{
			Name:        "d, debug",
			Usage:       "More Output to StdOut",
			Destination: &DebugFlag,
		},
	}
}
