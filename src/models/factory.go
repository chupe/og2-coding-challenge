package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Factory struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type FactoryType        `json:"type" bson:"type" validate:"required,alpha" example:"iron"`
	// The idea is to use this information to calculate the ores information on each user load by calculating how much time did each factory spend on a level
	// Time to upgrade is calculated as current time - map value with highest level. If the value is negative there is an upgrade pending.
	// Similarly level is calculated as maximum map value that has time in the past.
	// Using a map is favoring readability vs an array that favors effieciency
	UpgradeData map[int]time.Time `json:"upgradeData" bson:"upgradeData" validate:"required" example:"2021-05-25T00:00:00.0Z"`
}

func (f *Factory) GetLevel() int {
	level := 1
	for k, v := range f.UpgradeData {
		if k > level && v.Before(time.Now().UTC()) {
			level = k
		}
	}

	return level
}

func (f *Factory) GetRate() int {
	// needs to be implemented from config file like other extension functions
	return 10
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

func (f *Factory) OreProduced() int {
	var rate []LevelInfo
	switch f.Type {
	case "iron":
		rate = IronConfig.Info
	case "copper":
		rate = CopperConfig.Info
	case "gold":
		rate = GoldConfig.Info
	}

	var result int
	for l, v := range f.UpgradeData {
		if v.After(time.Now().UTC()) {
			continue
		}
		timeOnLevel := v.Sub(f.UpgradeData[l-1])
		result += rate[l-1].Production * int(timeOnLevel.Seconds()) / 60 // divide by 60 since production rate is recorded in ore/minute
	}

	return result
}

func NewIronFactory() Factory {
	d := make(map[int]time.Time)
	d[1] = time.Now().UTC()
	return Factory{
		Type:        FactoryType(Iron),
		UpgradeData: d,
	}
}

func NewCopperFactory() Factory {
	d := make(map[int]time.Time)
	d[1] = time.Now().UTC()
	return Factory{
		Type:        FactoryType(Copper),
		UpgradeData: d,
	}
}

func NewGoldFactory() Factory {
	d := make(map[int]time.Time)
	d[1] = time.Now().UTC()
	return Factory{
		Type:        FactoryType(Gold),
		UpgradeData: d,
	}
}
