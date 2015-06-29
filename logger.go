// Copyright (c) 2015 SUSE LLC. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const (
	// The default path for the logger if nothing has been specified.
	defaultLogPath = "suseconnect.log"

	// The environment variable used to specify a custom path for the logger
	// path.
	logEnv = "SUSECONNECT_LOG_FILE"
)

// Returns the output file for the logger. If the `logEnv` environment
// variable has been set, it will try to output there. Otherwise, it
// will try to output to the file as given in `defaultLogPath`. If
// everything fails, it will just output to the standard error channel.
func getLoggerFile() *os.File {
	// Determine the path to be used.
	var path string
	if env := os.Getenv(logEnv); env != "" {
		path = env
	} else {
		path = defaultLogPath
	}

	// If it's writable, use the given file, otherwise use os.Stderr.
	f, err := os.Create(path)
	if err == nil {
		return f
	}
	return os.Stderr
}

// Log the given formatted string with its parameters, and return it
// as a new error.
func loggedError(format string, params ...interface{}) error {
	str := fmt.Sprintf(format, params...)
	log.Print(str)
	return errors.New(str)
}
