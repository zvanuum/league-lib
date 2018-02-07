package model

type Champions struct {
	Data map[string]Champion `json:"data"`
}

type Champion struct {
	ID        int             `json:"id"`
	Key       string          `json:"key"`
	Name      string          `json:"name"`
	Title     string          `json:"title"`
	Tags      []string        `json:"tags,omitempty"`
	Blurb     string          `json:"blurb,omitempty"`
	Lore      string          `json:"lore,omitempty"`
	AllyTips  []string        `json:"allytips,omitempty"`
	EnemyTips []string        `json:"enemytips,omitempty"`
	Partype   string          `json:"partype,omitempty"`
	Info      *Info           `json:"info,omitempty"`
	Stats     *Stats          `json:"stats,omitempty"`
	Image     *Image          `json:"image,omitempty"`
	Skins     *[]Skin         `json:"skin,omitempty"`
	Passive   *Passive        `json:"passive,omitempty"`
	Spells    []ChampionSpell `json:"spells,omitempty"`
}

type Info struct {
	Difficulty int `json:"difficulty,omitempty"`
	Attack     int `json:"attack,omitempty"`
	Defense    int `json:"defense,omitempty"`
	Magic      int `json:"magic,omitempty"`
}

type Stats struct {
	ArmorPerLevel        float64 `json:"armorperlevel,omitempty"`
	HpPerLevel           float64 `json:"hpperlevel,omitempty"`
	AttackDamage         float64 `json:"attackdamage,omitempty"`
	MpPerLevel           float64 `json:"mpperlevel,omitempty"`
	AttackSpeedOffset    float64 `json:"attackspeedoffset,omitempty"`
	Armor                float64 `json:"armor,omitempty"`
	Hp                   float64 `json:"hp,omitempty"`
	HpRegenPerLevel      float64 `json:"hpregendperlevel,omitempty"`
	Spellblock           float64 `json:"spellblock,omitempty"`
	AttackRange          float64 `json:"attackrange,omitempty"`
	Movespeed            float64 `json:"movespeed,omitempty"`
	AttackDamagePerLevel float64 `json:"attackdamageperlevel,omitempty"`
	MpRegenPerLevel      float64 `json:"mpregenperlevel,omitempty"`
	Mp                   float64 `json:"mp,omitempty"`
	SpellblockPerLevel   float64 `json:"spellblockperlevel,omitempty"`
	Crit                 float64 `json:"crit,omitempty"`
	MpRegen              float64 `json:"mpregen,omitempty"`
	AttackSpeedPerLevel  float64 `json:"attackspeedperlevel,omitempty"`
	HpRegen              float64 `json:"hpregen,omitempty"`
	CritPerLevel         float64 `json:"critperlevel,omitempty"`
}

type Image struct {
	Full   string `json:"full,omitempty"`
	Group  string `json:"group,omitempty"`
	Sprite string `json:"sprite,omitempty"`
	H      int    `json:"h,omitempty"`
	W      int    `json:"w,omitempty"`
	Y      int    `json:"y,omitempty"`
	X      int    `json:"x,omitempty"`
}

type Skin struct {
	Num  int    `json:"num,omitempty"`
	Name string `json:"name,omitempty"`
	ID   int    `json:"id,omitempty"`
}

type Passive struct {
	Image                Image  `json:"image,omitempty"`
	SanitizedDescription string `json:"sanitizeDescrption,omitempty"`
	Name                 string `json:"name,omitempty"`
	Description          string `json:"description,omitempty"`
}

type Recommended struct {
	Map      string  `json:"map,omitempty"`
	Blocks   []Block `json:"blocks,omitempty"`
	Champion string  `json:"champion,omitempty"`
	Title    string  `json:"title,omitempty"`
	Priority bool    `json:"priority,omitempty"`
	Mode     string  `json:"mode,omitempty"`
	Type     string  `json:"type,omitempty"`
}

type Block struct {
	Items   []BlockItem `json:"items,omitempty"`
	RecMath bool        `json:"recMath,omitempty"`
	Type    string      `json:"type,omitempty"`
}

type BlockItem struct {
	Count int `json:"count,omitempty"`
	ID    int `json:"id,omitempty"`
}

type ChampionSpell struct {
	CooldownBurn         string        `json:"cooldownBurn,omitempty"`
	Resource             string        `json:"resource,omitempty"`
	LevelTip             LevelTip      `json:"leveltip,omitempty"`
	Vars                 []SpellVars   `json:"vars,omitempty"`
	CostType             string        `json:"costType,omitempty"`
	Image                Image         `json:"image,omitempty"`
	SanitizedDescription string        `json:"sanitizedDescription,omitempty"`
	SanitizedTooltip     string        `json:"sanitizedTooltip,omitempty"`
	Effect               [][]float64   `json:"effect,omitempty"`
	Tooltip              string        `json:"tooltip,omitempty"`
	MaxRank              int           `json:"maxRank,omitempty"`
	CostBurn             string        `json:"costBurn,omitempty"`
	RangeBurn            string        `json:"rangeBurn,omitempty"`
	Range                []interface{} `json:"range,omitempty"`
	Cooldown             []float64     `json:"cooldown,omitempty"`
	Cost                 []int         `json:"cost,omitempty"`
	Key                  string        `json:"key,omitempty"`
	Description          string        `json:"description,omitempty"`
	EffectBurn           []string      `json:"effectBurn,omitempty"`
	AltImages            []Image       `json:"altImages,omitempty"`
	Name                 string        `json:"name,omitempty"`
}

type LevelTip struct {
	Effect []string `json:"effect,omitempty"`
	Label  []string `json:"label,omitempty"`
}

type SpellVars struct {
	RanksWith string    `json:"ranksWith,omitempty"`
	Dyn       string    `json:"dyn,omitempty"`
	Link      string    `json:"link,omitempty"`
	Coeff     []float64 `json:"coeff,omitempty"`
	Key       string    `json:"key,omitempty"`
}
