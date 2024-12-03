// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/freetaxii/libstix2/datastore"
	"github.com/freetaxii/libstix2/datastore/sqlite3"
	"github.com/freetaxii/libstix2/objects/indicator"
)

// These global variables hold build information. The Build variable will be
// populated by the Makefile and uses the Git Head hash as its identifier.
// These variables are used in the console output for --version and --help.
var (
	Version = "0.0.1"
	Build   string
)

// These global variables are for dealing with command line options
var (
	sOptDatabaseFilename = flag.String("filename", "Database Filename. uses freetaxii.db if empty", "string")
	bOptHelp             = flag.Bool("help", false, "Help")
	bOptVer              = flag.Bool("version", false, "Version")
)

func main() {
	processCommandLineFlags()

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))

	var ds datastore.Datastorer
	ds = sqlite3.New(logger, *sOptDatabaseFilename, nil)
	defer ds.Close()

	i := indicator.New()
	i.SetName("Malware C2 Indicator 2016")
	i.AddLabels("BadStuff")
	i.AddTypes("compromised")

	// Set modified time to be one hour from now
	//modifiedTime := time.Now().Add(time.Hour)
	//i.SetModified(modifiedTime)

	i.SetValidFrom(time.Now())
	i.CreateKillChainPhase("lockheed-martin-cyber-kill-chain", "delivery")

	data, _ := i.EncodeToString()
	fmt.Println(data)

	// Add to datastore
	err := ds.AddObject(i)
	if err != nil {
		fmt.Println(err)
	}
}

// --------------------------------------------------
// Private functions
// --------------------------------------------------

// processCommandLineFlags - This function will process the command line flags
// and will print the version or help information as needed.
func processCommandLineFlags() {
	if sOptDatabaseFilename == nil {
		*sOptDatabaseFilename = "freetaxii.db"
	}

	// Lets check to see if the version command line flag was given. If it is
	// lets print out the version infomration and exit.
	if *bOptVer {
		printOutputHeader()
		os.Exit(0)
	}

	// Lets check to see if the help command line flag was given. If it is lets
	// print out the help information and exit.
	if *bOptHelp {
		printOutputHeader()
		os.Exit(0)
	}
}

// printOutputHeader - This function will print a header for all console output
func printOutputHeader() {
	fmt.Println("")
	fmt.Println("FreeTAXII - TAXII Table Creator")
	fmt.Println("Copyright: Bret Jordan")
	fmt.Println("Version:", Version)
	if Build != "" {
		fmt.Println("Build:", Build)
	}
	fmt.Println("")
}
