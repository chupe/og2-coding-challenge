package services

import (
	"errors"
	"time"

	"github.com/chupe/og2-coding-challenge/config"
	"github.com/chupe/og2-coding-challenge/models"
)

type FactoryService struct {
	fc *config.FactoryConfig
}

func NewFactoryService(fc *config.FactoryConfig) *FactoryService {
	return &FactoryService{
		fc: fc,
	}
}

func (fs *FactoryService) GetConfig(f *models.Factory) ([]config.LevelInfo, error) {
	config := *fs.fc
	lvlInfo, ok := config[string(f.Type)]

	if !ok {
		return nil, errors.New("failed to load factory config")
	}
	return lvlInfo, nil
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

	err = fs.deduceOres(user, models.Cost(fc[fac.GetLevel()].Cost))
	if err != nil {
		return nil, err
	}

	err = fs.upgradeFactory(user, fac)
	if err != nil {
		return nil, err
	}

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

func (fs *FactoryService) upgradeFactory(user *models.User, f *models.Factory) error {
	fc, err := fs.GetConfig(f)
	if err != nil {
		return err
	}

	lvl := f.GetLevel()
	if !(len(fc) > lvl) {
		return errors.New("level information not available")
	}
	f.UpgradeData = append(
		f.UpgradeData,
		time.Now().UTC().Add(time.Second*time.Duration(fc[lvl-1].UpgradeDuration)),
	)

	return nil
}
