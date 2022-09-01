package services

import (
	"errors"
	"time"

	"github.com/chupe/og2-coding-challenge/config"
	"github.com/chupe/og2-coding-challenge/models"
)

type FactoryService struct {
	fc *config.Factories
}

func NewFactoryService(fc *config.Factories) *FactoryService {
	return &FactoryService{
		fc: fc,
	}
}

func (fs *FactoryService) GetConfig(f *models.Factory) ([]config.Level, error) {
	var result []config.Level
	switch string(f.Type) {
	case "iron":
		result = fs.fc.Iron
	case "copper":
		result = fs.fc.Copper
	case "gold":
		result = fs.fc.Gold
	}

	return result, nil
}

func (fs *FactoryService) GetRate(f *models.Factory) (int, error) {
	v, err := fs.GetConfig(f)
	if err != nil {
		return -1, err
	}
	return v[f.GetLevel()-1].Production, nil
}

func (fs *FactoryService) OreProduced(f *models.Factory) (int, error) {
	lvlInfo, err := fs.GetConfig(f)
	if err != nil {
		return -1, err
	}

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

	return result, nil
}

func (fs *FactoryService) UpgradeFactory(user *models.User, factory string) (*models.User, error) {
	fac, err := user.GetFactory(factory)
	if err != nil {
		return nil, err
	}

	if fac.UnderConstruction() {
		return nil, errors.New("factory under construction")
	}

	fc, err := fs.GetConfig(fac)
	if err != nil {
		return nil, err
	}

	lvl := fac.GetLevel()
	if !(len(fc) > lvl) {
		return nil, errors.New("level information not available")
	}

	err = fs.deduceOres(user, models.Cost(fc[fac.GetLevel()-1].Cost))
	if err != nil {
		return nil, err
	}

	fac.UpgradeData = append(
		fac.UpgradeData,
		time.Now().UTC().Add(time.Second*time.Duration(fc[lvl-1].UpgradeDuration)),
	)

	return user, nil
}

func (fs *FactoryService) deduceOres(user *models.User, cost models.Cost) error {
	user.IronSpending += cost.Iron
	user.CopperSpending += cost.Copper
	user.GoldSpending += cost.Gold

	iron, err := fs.OreProduced(&user.IronFactory)
	copper, err := fs.OreProduced(&user.CopperFactory)
	gold, err := fs.OreProduced(&user.GoldFactory)
	if err != nil {
		return errors.New("failed to get amount of ores for user")
	}

	if iron < 0 || copper < 0 || gold < 0 {
		return errors.New("not enough resources")
	}

	return nil
}
