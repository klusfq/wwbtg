package frame

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
)

/**
 * 实现一个加载配置文件的基础库
 */

type Config map[string]string

type ConfigStore struct {
	merge map[string]Config
	app   map[string]Config
}

func LoadFrameConf(fnm string) {
	aFnm := "./" + fnm + ".json"

	_, fnx, _, ok := runtime.Caller(0)
	if ok {
	}
	fmt.Println("org -- ", fnx)

	content, err := os.ReadFile(aFnm)
	if err != nil {
		log.Panic(err)
	}

	var tmpConf map[string]any
	err = json.Unmarshal(content, &tmpConf)
	if err != nil {
		log.Panic(err)
	}

	for key, value := range tmpConf {
		fmt.Printf("%s -- %v", key, value)
	}
}

func LoadAppConf(app, fnm string) {
}
