package v0_X_0

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/lavanet/lava/app/upgrades"
)

const UpgradeName = "v0.X.0"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,           // upgrade name defined few lines above
	CreateUpgradeHandler: CreateUpgradeHandler,  // create CreateUpgradeHandler in upgrades.go below
	StoreUpgrades:        store.StoreUpgrades{}, // StoreUpgrades has 3 fields: Added/Renamed/Deleted any module that fits these description should be added in the way below
}
