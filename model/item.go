package model

type Items struct {
	Data map[string]Item `json:"data"`
	*VersionedData
	Tree   []ItemTree `json:"tree,omitempty"`
	Groups []Group    `json:"groups,omitempty"`
}

type Item struct {
	Gold             Gold               `json:"gold,omitempty"`
	PlainText        string             `json:"plaintext,omitempty"`
	HideFromAll      bool               `json:"hideFromAll,omitempty"`
	InStore          bool               `json:"inStore,omitempty"`
	Into             []string           `json:"into,omitempty"`
	ID               int                `json:"id"`
	Stats            InventoryDataStats `json:"stats,omitempty"`
	Colloq           string             `json:"colloq,omitempty"`
	Maps             map[string]bool    `json:"maps,omitempty"`
	SpecialRecipe    int                `json:"specialRecipe,omitempty"`
	Image            Image              `json:"image,omitempty"`
	Description      string             `json:"description,omitempty"`
	Tags             []string           `json:"tags,omitempty"`
	Effect           map[string]string  `json:"effect,omitempty"`
	RequiredChampion string             `json:"requiredChampion,omitempty"`
	From             []string           `json:"from,omitempty"`
	Group            string             `json:"group,omitempty"`
	ConsumeOnFull    bool               `json:"consumeOnFull,omitempty"`
	Name             string             `json:"name"`
}

type Gold struct {
	Sell        int  `json:"sell,omitempty"`
	Total       int  `json:"total,omitempty"`
	Base        int  `json:"base,omitempty"`
	Purchasable bool `json:"purchasable,omitempty"`
}

type InventoryDataStats struct {
	PercentCritDamageMod     float64 `json:"PercentCritDamageMod,omitempty"`
	PercentSpellBlockMod     float64 `json:"PercentSpellBlockMod,omitempty"`
	PercentHPRegenMod        float64 `json:"PercentHPRegenMod,omitempty"`
	PercentMovementSpeedMod  float64 `json:"PercentMovementSpeedMod,omitempty"`
	FlatSpellBlockMod        float64 `json:"FlatSpellBlockMod,omitempty"`
	FlatCritDamageMod        float64 `json:"FlatCritDamageMod,omitempty"`
	FlatEnergyPoolMod        float64 `json:"FlatEnergyPoolMod,omitempty"`
	PercentLifeStealMod      float64 `json:"PercentLifeStealMod,omitempty"`
	FlatMPPoolMod            float64 `json:"FlatMPPoolMod,omitempty"`
	FlatMovementSpeedMod     float64 `json:"FlatMovementSpeedMod,omitempty"`
	PercentAttackSpeedMod    float64 `json:"PercentAttackSpeedMod,omitempty"`
	FlatBlockMod             float64 `json:"FlatBlockMod,omitempty"`
	PercentBlockMod          float64 `json:"PercentBlockMod,omitempty"`
	FlatEnergyRegenMod       float64 `json:"FlatEnergyRegenMod,omitempty"`
	PercentSpellVampMod      float64 `json:"PercentSpellVampMod,omitempty"`
	FlatMPRegenMod           float64 `json:"FlatMPRegenMod,omitempty"`
	PercentDodgeMod          float64 `json:"PercentDodgeMod,omitempty"`
	FlatAttackSpeedMod       float64 `json:"FlatAttackSpeedMod,omitempty"`
	FlatArmorMod             float64 `json:"FlatArmorMod,omitempty"`
	FlatHPRegenMod           float64 `json:"FlatHPRegenMod,omitempty"`
	PercentMagicDamageMod    float64 `json:"PercentMagicDamageMod,omitempty"`
	PercentMPPoolMod         float64 `json:"PercentMPPoolMod,omitempty"`
	FlatMagicDamageMod       float64 `json:"FlatMagicDamageMod,omitempty"`
	PercentMPRegenMod        float64 `json:"PercentMPRegenMod,omitempty"`
	PercentPhysicalDamageMod float64 `json:"PercentPhysicalDamageMod,omitempty"`
	FlatPhysicalDamageMod    float64 `json:"FlatPhysicalDamageMod,omitempty"`
	PercentHPPoolMod         float64 `json:"PercentHPPoolMod,omitempty"`
	PercentArmorMod          float64 `json:"PercentArmorMod,omitempty"`
	PercentCritChanceMod     float64 `json:"PercentCritChanceMod,omitempty"`
	PercentEXPBonus          float64 `json:"PercentEXPBonus,omitempty"`
	FlatHPPoolMod            float64 `json:"FlatHPPoolMod,omitempty"`
	FlatCritChanceMod        float64 `json:"FlatCritChanceMod,omitempty"`
	FlatEXPBonus             float64 `json:"FlatEXPBonus,omitempty"`
}

type ItemTree struct {
	Header string   `json:"header,omitempty"`
	Tags   []string `json:"tags,omitempty"`
}

type Group struct {
	MaxGroupOwnable string `json:"MaxGroupOwnable,omitempty"`
	Key             string `json:"key,omitempty"`
}
