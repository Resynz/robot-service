/**
 * @Author: Resynz
 * @Date: 2020/12/30 13:25
 */
package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	Mode       = "debug"
	AppPort    = 3000
	DB_CONFIGS = "./configs"
)

func getEnv(e *string, key string) {
	if c := os.Getenv(key); c != "" {
		*e = c
	}
}

func initDbConfigs() error {
	f, err := os.Stat(DB_CONFIGS)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		err = os.MkdirAll(DB_CONFIGS, os.ModePerm)
		if err != nil {
			return err
		}
		f, err = os.Stat(DB_CONFIGS)
	}
	if f != nil && !f.IsDir() {
		return fmt.Errorf("db config is not dir")
	}
	return nil
}

func CheckEnvConf() error {
	if p, err := strconv.Atoi(os.Getenv("APP_PORT")); err == nil && p > 0 {
		AppPort = p
	}
	getEnv(&Mode, "MODE")
	getEnv(&DB_CONFIGS, "DB_CONFIGS")
	if err := initDbConfigs(); err != nil {
		return err
	}
	return nil
}
