package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Factory struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type FactoryType        `json:"type" bson:"type" validate:"required,alpha" example:"iron"`
	// The idea is to use this information to calculate the ores information on each user load by calculating how much time did each factory spend on a level
	// Time to upgrade is calculated as current time - map value with the highest level. If the value is negative there is an upgrade pending.
	// Similarly level is calculated as maximum map value that has time in the past.
	// Using a map is favoring readability vs an array that favors efficiency
	UpgradeData []time.Time `json:"upgradeData" bson:"upgradeData" validate:"required" example:"2021-05-25T00:00:00.0Z"`
}

func (f *Factory) GetLevel() int {
	level := len(f.UpgradeData)
	if f.UnderConstruction() {
		return level - 1
	}

	return level
}

func (f *Factory) GetRate() int {
	return f.GetConfig()[f.GetLevel()-1].Production
}

func (f *Factory) TimeToUpgrade() time.Time {
	var timeToUpgrade time.Time
	for _, v := range f.UpgradeData {
		if v.Before(time.Now().UTC()) {
			timeToUpgrade = v
		}
	}

	return timeToUpgrade
}

func (f *Factory) UnderConstruction() bool {
	underConstruction := false
	for _, v := range f.UpgradeData {
		if v.After(time.Now().UTC()) {
			underConstruction = true
		}
	}

	return underConstruction
}

func (f *Factory) GetConfig() []LevelInfo {
	var lvlInfo []LevelInfo
	switch f.Type {
	case "iron":
		lvlInfo = IronConfig.Info
	case "copper":
		lvlInfo = CopperConfig.Info
	case "gold":
		lvlInfo = GoldConfig.Info
	}

	return lvlInfo
}

func (f *Factory) OreProduced() int {
	lvlInfo := f.GetConfig()
	var result int
	cl := f.GetLevel()
	for i := 0; i < cl; i++ {
		var timeOnLevel time.Duration
		if i+1 == cl {
			timeOnLevel = time.Now().UTC().Sub(f.UpgradeData[i])
		} else {
			timeOnLevel = f.UpgradeData[i+1].Sub(f.UpgradeData[i])
		}

		result += lvlInfo[i].Production * int(timeOnLevel.Seconds()) / 60 // divide by 60 since production rate is recorded in ore/minute
	}

	return result
}

func NewIronFactory() Factory {
	var d []time.Time
	d = append(d, time.Now().UTC())
	return Factory{
		Type:        FactoryType(Iron),
		UpgradeData: d,
	}
}

func NewCopperFactory() Factory {
	var d []time.Time
	d = append(d, time.Now().UTC())
	return Factory{
		Type:        FactoryType(Copper),
		UpgradeData: d,
	}
}

func NewGoldFactory() Factory {
	var d []time.Time
	d = append(d, time.Now().UTC())
	return Factory{
		Type:        FactoryType(Gold),
		UpgradeData: d,
	}
}
