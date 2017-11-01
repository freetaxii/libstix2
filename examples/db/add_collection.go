package main

import (
	"github.com/freetaxii/libstix2/datastore/sqlite3"
	"github.com/freetaxii/libstix2/resources"
)

func main() {
	databaseFilename := "freetaxii.db"
	ds := sqlite3.New(databaseFilename)

	c := objects.NewCollection()
	c.NewID()

	c.SetTitle("ZeuS IP blocklist")
	c.SetDescription("ZeuS IP blocklistn from abuse.ch (excluding hijacked sites and free hosting providers). For questions please refer to https://zeustracker.abuse.ch/blocklist.php")
	c.SetCanRead()
	c.AddMediaType("application/vnd.oasis.stix+json")

	ds.Put(c)
	ds.Close()
}
