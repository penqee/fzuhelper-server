/*
Copyright 2024 The west2-online Authors.

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

package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap/zapcore"

	"github.com/west2-online/fzuhelper-server/pkg/constants"
)

var (
	targetDir = "."
	service   = "svc"
)

func Test_updateLogger(t *testing.T) {
	now := time.Now()
	date := now.Format("2006-01-02")

	PatchConvey("Test updateLogger", t, func() {
		Mock(getCurrentDirectory).Return(targetDir, nil).Build()
		Mock(time.Now).Return(now).Build()
		control.updateLogger(service)
		infoMsg := "info"
		errorMsg := "error"
		Info(infoMsg)
		Error(errorMsg)

		PatchConvey("Try open file and read all", func() {
			logPath := fmt.Sprintf("%s/%s/%s/%s.log", targetDir, constants.LogFilePath, date, service)
			stderrPath := fmt.Sprintf("%s/%s/%s/%s_stderr.log", targetDir, constants.LogFilePath, date, service)
			logFile, err := os.Open(logPath)
			So(err, ShouldBeNil)
			errFile, err := os.Open(stderrPath)
			So(err, ShouldBeNil)

			infoLogB, err := io.ReadAll(logFile)
			So(err, ShouldBeNil)
			errorLogB, err := io.ReadAll(errFile)
			So(err, ShouldBeNil)

			PatchConvey("Check result", func() {
				infoLog := string(infoLogB)
				errorLog := string(errorLogB)
				infoFields := strings.Split(infoLog, "\t")
				errorFields := strings.Split(errorLog, "\t")

				So(len(infoFields), ShouldEqual, 4)
				So(len(infoFields), ShouldEqual, len(errorFields))
				So(len(infoFields[0]), ShouldEqual, len(errorFields[0]))
				So(infoFields[1], ShouldEqual, strings.ToUpper(infoMsg))
				So(errorFields[1], ShouldEqual, strings.ToUpper(errorMsg))
				So(len(infoFields[2]), ShouldEqual, len(infoFields[2]))
				So(infoFields[3], ShouldEqual, infoMsg+"\n")
				So(errorFields[3], ShouldEqual, errorMsg+"\n")
			})

			PatchConvey("Release resource", func() {
				So(logFile.Close(), ShouldBeNil)
				So(errFile.Close(), ShouldBeNil)
				So(os.Remove(logPath), ShouldBeNil)
				So(os.Remove(stderrPath), ShouldBeNil)
				So(os.Remove(fmt.Sprintf("%s/%s/%s", targetDir, constants.LogFilePath, date)), ShouldBeNil)
				_ = os.Remove(fmt.Sprintf("%s/%s", targetDir, constants.LogFilePath))
			})
		})
	})
}

func Test_scheduleUpdateLogger(t *testing.T) {
	now := time.Now().Truncate(24 * time.Hour).Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	date := now.Format("2006-01-02")
	logPath := fmt.Sprintf("%s/%s/%s/%s.log", targetDir, constants.LogFilePath, date, service)
	stderrPath := fmt.Sprintf("%s/%s/%s/%s_stderr.log", targetDir, constants.LogFilePath, date, service)

	PatchConvey("Test scheduleUpdateLogger", t, func() {
		Mock(getCurrentDirectory).Return(targetDir, nil).Build()
		Mock(time.Now).Return(now).Build()
		control.scheduleUpdateLogger("svc")

		time.Sleep(3 * time.Second) // waiting for update logger

		So(logFileHandler.Load(), ShouldNotBeNil)
		So(stdErrFileHandler.Load(), ShouldNotBeNil)

		PatchConvey("Release resource", func() {
			So(os.Remove(logPath), ShouldBeNil)
			So(os.Remove(stderrPath), ShouldBeNil)
			So(os.Remove(fmt.Sprintf("%s/%s/%s", targetDir, constants.LogFilePath, date)), ShouldBeNil)
			_ = os.Remove(fmt.Sprintf("%s/%s", targetDir, constants.LogFilePath))
		})
	})
}

func Test_parseLevel(t *testing.T) {
	Convey("Test parseLevel", t, func() {
		So(parseLevel("Debug"), ShouldEqual, zapcore.DebugLevel)
		So(parseLevel("DEBUG"), ShouldEqual, zapcore.DebugLevel)
		So(parseLevel("Info"), ShouldEqual, zapcore.InfoLevel)
		So(parseLevel("Info"), ShouldEqual, zapcore.InfoLevel)
		So(parseLevel(""), ShouldEqual, zapcore.InfoLevel)
	})
}