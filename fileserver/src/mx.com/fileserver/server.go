package fileserver

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"labix.org/v2/mgo"

	"mx.com/log"
)

type Config struct {
	BindHost   string
	FilePath   string
	MgoHost    string
	MgoDbName  string
	MgoColName string
}

type Server struct {
	MgoCol   *mgo.Collection
	FilePath string
}

func (s *Server) file(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
	}
	if r.Method == "POST" {
		s.putFile(r)
	}
}

func (s *Server) getFile(fileName string) (fileContent string, err error) {
	//TODO
	return
}

func (s *Server) putFile(r *http.Request) (fileName string, err error) {

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("<fileServer> putFile failed :", err)
		return
	}

	md5Sum := md5.Sum(buf)
	fileName = fmt.Sprintf("%x", md5Sum)

	file, err := os.Create(s.FilePath + "/" + fileName)
	defer file.Close()

	if err != nil {
		log.Error("<fileServer> create file failed: ", err)
		return
	}
	_, err = file.Write(buf)
	if err != nil {
		log.Error("<fileServer> write file failed:", err)
	}
	return
}

func RunServer(conf *Config) (err error) {
	session, err := mgo.Dial(conf.MgoHost)
	if err != nil {
		log.Fatal("mgodb cannot be connected!")
	}
	MgoCol := session.DB(conf.MgoDbName).C(conf.MgoColName)

	s := &Server{
		FilePath: conf.FilePath,
		MgoCol:   MgoCol,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/file", s.file)
	err = http.ListenAndServe(conf.BindHost, mux)
	return
}
