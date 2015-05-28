package main

import (
	"flag"
	"runtime"

	"mx.com/config"
	"mx.com/fileserver"
	"mx.com/log"
)

type FileServerConf struct {
	MaxProcs   int    `json:"max_procs`
	BindHost   string `json:"bind_host"`
	FilePath   string `json:"file_path"`
	MgoHost    string `json:"mgo_host"`
	MgoDbName  string `json:"mgo_database"`
	MgoColName string `json:"mgo_collection"`
}

func main() {
	var conf FileServerConf
	confName := flag.String("f", "fileserver.conf", "the config file")
	if !flag.Parsed() {
		flag.Parse()
	}
	err := config.LoadEx(&conf, *confName)
	if err != nil {
		log.Error("<fileserver> load conf failed:", err)
		return
	}
	runtime.GOMAXPROCS(conf.MaxProcs)

	serverConf := &fileserver.Config{
		BindHost:   conf.BindHost,
		FilePath:   conf.FilePath,
		MgoHost:    conf.MgoHost,
		MgoDbName:  conf.MgoDbName,
		MgoColName: conf.MgoColName,
	}
	err = fileserver.RunServer(serverConf)
	if err != nil {
		log.Fatal("start file failed")
	}
}
