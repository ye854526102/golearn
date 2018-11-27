package conf

import (
	"runtime"
	"flag"
	"log"
	"fmt"
	"github.com/larspensjo/config"
)

var (
	configFile = flag.String("configfile", "E:\\wamp\\wwwroot\\istore2_dev\\trunk\\app\\configs\\appConfigs.ini", "General configuration file")
)

//topic list

type dbConfig struct {
	dbHost string
	dbPort int
	dbUser string
	dbName string
}

var Conf = make(map[string]string)
var db = new(dbConfig)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	//set config file std
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}
	//set config file std End

	//Initialized topic from the configuration
	if cfg.HasSection("production") {
		section, err := cfg.SectionOptions("production")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("production", v)
				if err == nil {
					Conf[v] = options
				}
			}
		}
	}
	//Initialized topic from the configuration END

	fmt.Println(Conf["db.host"])
	fmt.Println(Conf["db.port"])
	fmt.Println(Conf["db.user"])
	fmt.Println(Conf["db.dbname"])
}
func main() {

}

func (*dbConfig) parser() {

}
