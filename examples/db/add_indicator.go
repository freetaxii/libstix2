package main

import (
	"github.com/freetaxii/libstix2/datastore/sqlite3"
	"github.com/freetaxii/libstix2/objects"
	"time"
)

func main() {
	databaseFilename := "freestix.db"
	ds := sqlite3.New(databaseFilename)

	i := objects.NewIndicator("2.0")

	i.SetName("Malware C2 Indicator 2016")
	i.SetDescription("Some more information about this")
	i.AddLabel("BadStuff")

	i.SetValidFrom(time.Now())
	i.AddKillChainPhase("lockheed-martin-cyber-kill-chain", "delivery")

	ds.Put(i)
	ds.Close()
}
