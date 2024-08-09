package conf

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	Init struct {
		Port     int    `yaml:"port"`
		SavePath string `yaml:"save_dir"`
		WorkMax  int    `yaml:"work_max"`
	} `yaml:"init"`
	Log struct {
		Path  string `yaml:"path"`
		Level string `yaml:"level"`
		LogNu int    `yaml:"log_Nu"`
	} `yaml:"log"`
}

func NewConfig() *Config {
	c := new(Config)
	c.Init.Port = 2045
	c.Init.SavePath = "./video"
	c.Init.WorkMax = 1
	c.Log.Level = "debug"
	c.Log.Path = "./log"
	c.Log.LogNu = 10
	return c
}

var ConfMap Config

func ConfInit() {
	useDefaultConfig := true
	// 读取YAML配置文件内容
	yamlFile, err := os.ReadFile("./conf.yaml")
	if err != nil {
		log.Println("无法读取YAML文件：%v", err)
	} else {
		// 解析YAML配置文件
		configTemp := NewConfig()
		err = yaml.Unmarshal(yamlFile, &configTemp)
		if err != nil {
			log.Println("无法解析YAML文件：%v", err)
		} else {
			useDefaultConfig = false
			ConfMap = *configTemp
		}
	}
	if useDefaultConfig {
		ConfMap = *NewConfig()
	}
	savePath := ConfMap.Init.SavePath
	_, err = os.Stat(savePath)
	if err != nil && os.IsNotExist(err) {
		err = os.Mkdir(savePath, 0755)
		if err != nil {
			panic(err)
		}
	}
	if !CheckWritePermission(savePath) {
		panic("save_dir: " + strconv.Quote(savePath) + " can not write")
	}
	// 打印配置项的值
	confjson, _ := json.Marshal(ConfMap)
	fmt.Println("conf:", string(confjson))
}

// 检查文件夹是否可写
func CheckWritePermission(dirPath string) bool {
	dirPath, _ = filepath.Abs(dirPath)
	tmpFile, err := os.CreateTemp(dirPath, "test*")
	if err != nil {
		return false
	}
	defer os.Remove(tmpFile.Name()) // 确保临时文件被删除
	defer tmpFile.Close()
	return true
}
