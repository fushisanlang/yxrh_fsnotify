package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/fsnotify/fsnotify"
	"github.com/larspensjo/config"
	"github.com/wonderivan/logger"
)

type Watch struct {
	watch *fsnotify.Watcher
}

var (
	msg_oper   string
	msg_file   string
	ipaddr     string
	corpid     string
	corpsecret string
	appid      string
)

var (
	configFile = flag.String("configfile", "/usr/local/yxrh_fsnotify/conf/config.ini", "General configuration file")
)

var TOPIC = make(map[string]string)
var parameter string

func readconf(parameter string) string {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}
	if cfg.HasSection("topicArr") {
		section, err := cfg.SectionOptions("topicArr")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("topicArr", v)
				if err == nil {
					TOPIC[v] = options
				}
			}
		}
	}
	return TOPIC[parameter]
}

func (w *Watch) watchDir(dir string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			path, err := filepath.Abs(path)

			if err != nil {
				return err
			}
			err = w.watch.Add(path)
			if err != nil {
				return err
			}

			logger.Info("监控 : ", path)
		}
		return nil
	})
	go func() {
		for {
			select {
			case ev := <-w.watch.Events:
				{
					if searchinmap(ev.Name) == true {
						break
					}
					if ev.Op&fsnotify.Create == fsnotify.Create {
						logger.Alert("创建文件 : ", ev.Name)
						fi, err := os.Stat(ev.Name)
						if err == nil && fi.IsDir() {
							w.watch.Add(ev.Name)
							logger.Alert("添加监控 : ", ev.Name)
						}
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						logger.Alert("写入文件 : ", ev.Name)
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						logger.Alert("删除文件 : ", ev.Name)
						fi, err := os.Stat(ev.Name)
						if err == nil && fi.IsDir() {
							w.watch.Remove(ev.Name)
							logger.Alert("删除监控 : ", ev.Name)
						}
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						logger.Alert("重命名文件 : ", ev.Name)
						w.watch.Remove(ev.Name)
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						logger.Alert("修改权限 : ", ev.Name)
					}
				}
			case err := <-w.watch.Errors:
				{
					logger.Error("error : ", err)
					return
				}
			}
		}
	}()
}

var (
	dirs   [0]string
	dirs_s []string = dirs[:]
)

func getnamelist(num int) string {
	dirnamenum := "dirname" + strconv.Itoa(num)
	dirname := readconf(dirnamenum)
	return dirname
}

func getskipnamelist(num int) string {
	skipnamenum := "skipname" + strconv.Itoa(num)
	skipname := readconf(skipnamenum)
	return skipname
}
func searchinmap(path string) bool {

	skipnamemap := make(map[int]string)
	for index := 1; ; index++ {
		if getskipnamelist(index) == "" {
			break
		} else {
			skipnamemap[index] = getskipnamelist(index)
		}
	}

	result := map[int]bool{1: false}
	for _, value := range skipnamemap {
		ok, _ := filepath.Match(value+`/*`, path)
		if ok {
			result[1] = true
			break
		}
	}
	return result[1]
}

func main() {
	logger.SetLogger(`{"Console":{"level":"INFO"},"File": {"filename":"/usr/local/yxrh_fsnotify/log/fsnotify.log","level": "ALRT","maxlines": 1000000,"maxsize": 1,"maxdays": -1,"append": true,"permit": "0664"}}`)

	watch, _ := fsnotify.NewWatcher()
	w := Watch{
		watch: watch}
	for index := 1; ; index++ {
		dir := getnamelist(index)
		if dir == "" {
			break
		}
		w.watchDir(dir)
	}
	select {}
}
