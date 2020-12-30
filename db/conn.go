/**
 * @Author: Resynz
 * @Date: 2020/12/30 13:43
 */
package db

import (
	"encoding/json"
	"fmt"
	"github.com/yue-best-practices/db-handler"
	"os"
	"robot-service/config"
)

var (
	Handler *db_handler.DBHandler
)

func InitDBHandler() error {
	dp := fmt.Sprintf("%s/db.json", config.DB_CONFIGS)
	rp := fmt.Sprintf("%s/redis.json", config.DB_CONFIGS)
	df, err := os.Open(dp)
	if err != nil {
		return err
	}
	defer df.Close()
	dbDec := json.NewDecoder(df)
	var dbConf db_handler.DbConfig
	err = dbDec.Decode(&dbConf)
	if err != nil {
		return err
	}
	rf, err := os.Open(rp)
	if err != nil {
		return err
	}
	defer rf.Close()
	rDec := json.NewDecoder(rf)
	var redisConf db_handler.RedisConfig
	err = rDec.Decode(&redisConf)
	if err != nil {
		return err
	}
	Handler, err = db_handler.New(&dbConf, &redisConf)
	return err
}
