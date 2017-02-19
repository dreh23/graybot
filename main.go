//& Copyright 2017 Johannes Amorosa
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

// This Program logs things to a Greylog server. It is just for testing and has
// no production relevance

package main

import (
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/dreh23/nicename"
	"github.com/urfave/cli"
)

var appVersion, appReleaseTag, appShortCommitID string

func main() {
	app := cli.NewApp()
	app.Name = "Celluloid Filesystemlinker"
	app.Usage = "This Application links files"
	app.Author = "Johannes Amorosa"
	app.Email = "johannesa@celluloid-vfx.com"
	app.Version = "XXX" + "\n" + mainVersion()

	registerFlags(app)

	app.Action = func(c *cli.Context) bool {

		initlogging()

		if c.Bool("daemon") {

			rand.Seed(time.Now().UnixNano())
			for {
				w := rand.Intn(100)
				x := rand.Int31n(3 + 1)
				y := rand.Intn(4)
				z := rand.Intn(100)

				// Error
				if w == 1 {
					log.Error("What a mess!")
				}

				// Jackpot!
				if c.Bool("crash") {
					if z == 1 {
						log.Error("Oh noes! I hope init scrapes me from the floor.")
					}
				}

				// Warn and Info
				if y == 1 {
					a := strings.ToLower(nicename.First[z])
					ind := "a"
					if string(a[0]) == "a" ||
						string(a[0]) == "e" ||
						string(a[0]) == "i" ||
						string(a[0]) == "o" ||
						string(a[0]) == "u" {
						ind = "an"
					}

					log.Warnf("This is %s %s warning", ind, a)
				} else {
					log.Infof("Graybot is sleeping for %v seconds - doing nothing", x)
				}

				time.Sleep(time.Duration(x) * time.Second)
			}
		}
		log.Warn("This Service runs just in Daemon mode try the commandline option -D")
		return false
	} //app.Action
	_ = app.Run(os.Args)

}

func mainVersion() string {
	s := ""
	s = s + "Version: " + appVersion + "\n"
	s = s + "Release-Tag: " + appReleaseTag + "\n"
	s = s + "Commit-ID: " + appShortCommitID + "\n"
	return s
}
