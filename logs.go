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
	logrus "github.com/Sirupsen/logrus"
	"gopkg.in/gemnasium/logrus-graylog-hook.v2"
)

var log *logrus.Logger

func initlogging() {
	log = logrus.New()

	if DebugFlag {
		logrus.SetLevel(logrus.DebugLevel)
		log.Debug("Debug Mode")
	}

	if len(GraylogURL) > 0 {

		logrus.SetLevel(logrus.WarnLevel)
		log.Infof("Log messages are now sent to Graylog (udp://%s)", GraylogURL)

		hook := graylog.NewAsyncGraylogHook(GraylogURL, map[string]interface{}{
			"AppName":    "Graybot",
			"Version":    appVersion,
			"ReleaseTag": appReleaseTag,
			"CommitId":   appShortCommitID,
		})
		defer hook.Flush()
		log.Hooks.Add(hook)
	}
}
