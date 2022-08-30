package services

import (
	"errors"
	"time"

	"github.com/chupe/og2-coding-challenge/models"
)

type FactoryService struct{}

func NewFactoryService() *FactoryService {
	return &FactoryService{}
}

func (FactoryService) UpgradeFactory(user *models.User, factory string) (*models.User, error) {
	err := deduceOres(user, factory)
	if err != nil {
		return nil, err
	}
	upgradeFactory(user, factory)

	return user, nil
}

func deduceOres(user *models.User, factory string) error {
	cost := models.Ores{}
	switch factory {
	case "iron":
		facLvl := user.IronFactory.GetLevel()
		cost = models.IronConfig.Info[facLvl-1].Cost
	case "copper":
		facLvl := user.CopperFactory.GetLevel()
		cost = models.CopperConfig.Info[facLvl-1].Cost
	case "gold":
		facLvl := user.CopperFactory.GetLevel()
		cost = models.CopperConfig.Info[facLvl-1].Cost
	}

	user.IronSpending += cost.Iron
	user.CopperSpending += cost.Copper
	user.GoldSpending += cost.Gold

	if user.GetIronOre() < 0 || user.GetCopperOre() < 0 || user.GetGoldOre() < 0 {
		return errors.New("not enough resources")
	}

	return nil
}

func upgradeFactory(user *models.User, factory string) {
	switch factory {
	case "iron":
		fac := &user.IronFactory
		lvl := fac.GetLevel()
		fac.UpgradeData = append(fac.UpgradeData, time.Now().UTC().Add(time.Second*time.Duration(models.IronConfig.Info[lvl-1].UpgradeDuration)))
	case "copper":
		fac := &user.IronFactory
		lvl := fac.GetLevel()
		fac.UpgradeData = append(fac.UpgradeData, time.Now().UTC().Add(time.Second*time.Duration(models.IronConfig.Info[lvl-1].UpgradeDuration)))
	case "gold":
		fac := &user.IronFactory
		lvl := fac.GetLevel()
		fac.UpgradeData = append(fac.UpgradeData, time.Now().UTC().Add(time.Second*time.Duration(models.IronConfig.Info[lvl-1].UpgradeDuration)))
	}
}
