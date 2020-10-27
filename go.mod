module github.com/freetaxii/libstix2

go 1.13

require (
	github.com/gologme/log v1.2.0
	github.com/mattn/go-sqlite3 v1.13.0
	github.com/pborman/getopt v0.0.0-20190409184431-ee0cd42419d3 // indirect
	github.com/pborman/uuid v1.2.0
)

replace (
	github.com/freetaxii/libstix2 latest =>  github.com/wxj95/libstix2 latest
)