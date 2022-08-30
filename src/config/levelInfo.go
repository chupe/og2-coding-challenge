package config

type LevelInfo struct {
	Level           int
	Production      int // per minute
	UpgradeDuration int // duration in seconds
	Cost            Cost
}

type Cost struct {
	Iron   int
	Copper int
	Gold   int
}
