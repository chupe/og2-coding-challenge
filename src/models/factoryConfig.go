package models

// FactoryInfo contains information about each factory configuration presets
// It should be loaded from a config file once on app startup
type FactoryInfo struct {
	Type FactoryType
	Info []LevelInfo
}

type LevelInfo struct {
	Level           int
	Production      int // per minute
	UpgradeDuration int // duration in seconds
	Cost            Ores
}

type Ores struct {
	Iron   int
	Copper int
	Gold   int
}

var IronConfig = FactoryInfo{
	Type: "iron",
	Info: ironLevelInfo,
}

var ironLevelInfo = []LevelInfo{
	{
		Level:           1,
		Production:      10 * 60,
		UpgradeDuration: 15,
		Cost: Ores{
			Iron:   300,
			Copper: 100,
			Gold:   1,
		},
	},
	{
		Level:           2,
		Production:      20 * 60,
		UpgradeDuration: 30,
		Cost: Ores{
			Iron:   800,
			Copper: 250,
			Gold:   2,
		},
	},
}

var CopperConfig = FactoryInfo{
	Type: "Copper",
	Info: copperLevelInfo,
}

var copperLevelInfo = []LevelInfo{
	{
		Level:           1,
		Production:      3 * 60,
		UpgradeDuration: 15,
		Cost: Ores{
			Iron:   200,
			Copper: 70,
			Gold:   0,
		},
	},
	{
		Level:           2,
		Production:      7 * 60,
		UpgradeDuration: 30,
		Cost: Ores{
			Iron:   400,
			Copper: 150,
			Gold:   0,
		},
	},
}

var GoldConfig = FactoryInfo{
	Type: "Gold",
	Info: goldLevelInfo,
}

var goldLevelInfo = []LevelInfo{
	{
		Level:           1,
		Production:      2,
		UpgradeDuration: 15,
		Cost: Ores{
			Iron:   0,
			Copper: 100,
			Gold:   2,
		},
	},
	{
		Level:           2,
		Production:      3,
		UpgradeDuration: 30,
		Cost: Ores{
			Iron:   0,
			Copper: 200,
			Gold:   4,
		},
	},
}
