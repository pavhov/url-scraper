package config_lib

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"service/src/internal/builtin_lib"
	"strconv"
	"strings"
	"sync"

	"service/src/internal/json_lib"
)

var Config *Conf

type Interface interface {
	Init() (err error)
	Get(k string) any

	defineConfig(path string) (err error)
	prepareConfigs(data *map[string]any, key string)

	setFloat(skey, swap string)
	setBoll(skey, swap string)
	setArr(skey, swap string)
}

type Conf struct {
	sync.Mutex
	config map[string]any
}

func NewConf() *Conf {
	var instance Conf
	instance.config = make(map[string]any)
	return &instance
}

func (c *Conf) Get(k string) any {
	c.Lock()
	defer c.Unlock()
	return c.config[k]
}

func (c *Conf) Init() (err error) {
	defer builtin_lib.Recovery()

	path, _ := os.Getwd()

	if err = c.defineConfig(path + "/src/configs/"); err != nil {
		return
	}
	Config = c
	return
}

func (c *Conf) defineConfig(path string) (err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() == true {
			if err := c.defineConfig(path + f.Name() + "/"); err != nil {
				return err
			}
		} else {
			file, err := ioutil.ReadFile(path + f.Name())
			if err != nil {
				return err
			}
			var data = json_lib.Decode[map[string]any](string(file))
			if data == nil {
				return err
			}
			c.prepareConfigs(&data, strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())))
		}
	}
	return
}

func (c *Conf) prepareConfigs(data *map[string]any, key string) {
	for k, v := range *data {
		skey := key + "_" + k
		swap := os.Getenv(skey)
		c.config[skey] = v
		switch v.(type) {
		case float64:
			c.setFloat(skey, swap)
			break
		case bool:
			c.setBoll(skey, swap)
			break
		case map[string]any:
			var rr = v.(map[string]any)
			if swap != "" {
				rr = json_lib.Decode[map[string]any](swap)
				c.config[skey] = rr
			}
			c.prepareConfigs(&rr, skey)
			break
		case []any:
			c.config[skey] = v.([]any)
			c.setArr(skey, swap)
			break
		default:
			if swap != "" {
				c.config[skey] = swap
			}
			break
		}
	}
}

func (c *Conf) setFloat(key, val string) {
	if val != "" {
		val, _ := strconv.Atoi(val)
		c.config[key] = float64(val)
	}
}

func (c *Conf) setBoll(key, val string) {
	if val != "" {
		val, _ := strconv.ParseBool(val)
		c.config[key] = val
	}
}

func (c *Conf) setArr(key, val string) {
	if val != "" {
		c.config[key] = json_lib.Decode[[]any](val)
	}
}
