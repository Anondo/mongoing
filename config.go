package main

type mongoConfig struct {
	uri string
	db  string
}

var mCfg mongoConfig

func loadConfig() {
	mCfg = mongoConfig{
		uri: "mongodb://localhost:27017",
		db:  "hello",
	}
}

func mongodb() mongoConfig {
	return mCfg
}
