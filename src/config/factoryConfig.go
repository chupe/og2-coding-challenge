package config

import "sync"

var once sync.Once
var facConfig *FactoryConfig

type FactoryConfig map[string][]LevelInfo

func NewFactoryConfig() *FactoryConfig {
	once.Do(func() {
		facConfig = loadFactoryConfig()
	})

	return facConfig
}

func loadFactoryConfig() *FactoryConfig {
	config := make(FactoryConfig)
	config["iron"] = []LevelInfo{
		{
			Level:           1,
			Production:      10 * 60,
			UpgradeDuration: 15,
			Cost: Cost{
				Iron:   300,
				Copper: 100,
				Gold:   1,
			},
		},
		{
			Level:           2,
			Production:      20 * 60,
			UpgradeDuration: 30,
			Cost: Cost{
				Iron:   800,
				Copper: 250,
				Gold:   2,
			},
		},
	}

	config["copper"] = []LevelInfo{
		{
			Level:           1,
			Production:      3 * 60,
			UpgradeDuration: 15,
			Cost: Cost{
				Iron:   200,
				Copper: 70,
				Gold:   0,
			},
		},
		{
			Level:           2,
			Production:      7 * 60,
			UpgradeDuration: 30,
			Cost: Cost{
				Iron:   400,
				Copper: 150,
				Gold:   0,
			},
		},
	}

	config["gold"] = []LevelInfo{
		{
			Level:           1,
			Production:      2,
			UpgradeDuration: 15,
			Cost: Cost{
				Iron:   0,
				Copper: 100,
				Gold:   2,
			},
		},
		{
			Level:           2,
			Production:      3,
			UpgradeDuration: 30,
			Cost: Cost{
				Iron:   0,
				Copper: 200,
				Gold:   4,
			},
		},
	}

	return &config
}
