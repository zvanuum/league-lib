package leaguelib

// SEASONS
const (
	Preseason3    = 0
	Season3       = 1
	Preseason2014 = 2
	Season2014    = 3
	Preseason2015 = 4
	Season2015    = 5
	Preseason2016 = 6
	Season2016    = 7
	Preseason2017 = 8
	Season2017    = 9
	Preseason2018 = 10
	Season2018    = 11
)

// MATCHMAKING QUEUES - queue, queueId, gameQueueConfigId
const (
	// 	Custom games
	CustomGames = 0
	// 	Summoner's Rift	One for All games
	SummonersRiftOneForAll = 70
	// Howling Abyss	1v1 Snowdown Showdown games
	HowlingAbyss1v1Snowdown = 72
	// Howling Abyss	2v2 Snowdown Showdown games
	HowlingAbyss2v2Snowdown = 73
	// Summoner's Rift	6v6 Hexakill games
	SummonersRiftHexakill = 75
	// Summoner's Rift	Ultra Rapid Fire games
	SummonersRiftURF = 76
	// Howling Abyss	One For All: Mirror Mode games
	HowlingAByssOneForAllMirror = 78
	// Summoner's Rift	Co-op vs AI Ultra Rapid Fire games
	CoopVsAIURF = 83
	// Twisted Treeline	6v6 Hexakill games
	TwistedTreelineHexakill = 98
	// Butcher's Bridge	5v5 ARAM games
	ButchersBridgeARAM = 100
	// Summoner's Rift	Nemesis games
	SummonersRiftNemesis = 310
	// Summoner's Rift	Black Market Brawlers games
	SummonersRiftBlackMarket = 313
	// Crystal Scar	Definitely Not Dominion games
	DefinitelyNotDominion = 317
	// Summoner's Rift	All Random games
	SummonersRiftAllRandom = 325
	// Summoner's Rift	5v5 Draft Pick games
	SummonersRiftDraft = 400
	// Summoner's Rift	5v5 Ranked Solo games
	SummonersRiftRankedSolo = 420
	// Summoner's Rift	5v5 Blind Pick games
	SummonersRiftBlind = 430
	//	Summoner's Rift	5v5 Ranked Flex games
	SummonersRiftFlex = 440
	// Howling Abyss	5v5 ARAM games
	HowlingAbyssARAM = 450
	// Twisted Treeline	3v3 Blind Pick games
	TwistedTreelineBlind = 460
	// Twisted Treeline	3v3 Ranked Flex games
	TwistedTreelineFlex = 470
	// Summoner's Rift	Blood Hunt Assassin games
	BloodHuntAssassins = 600
	// Cosmic Ruins	Dark Star: Singularity games
	CosmicRuins = 610
	// Twisted Treeline	Co-op vs. AI Intermediate Bot games
	TwistedTreelineCoopVsAIIntermediate = 800
	// 	Twisted Treeline	Co-op vs. AI Intro Bot games
	TwistedTreelineCoopVsAIInto = 810
	// 	Twisted Treeline	Co-op vs. AI Beginner Bot games
	TwistedTreelineCoopVsAIBeginner = 820
	// Summoner's Rift	Co-op vs. AI Intro Bot games
	SummonersRiftCoopVsAIIntro = 830
	// Summoner's Rift	Co-op vs. AI Beginner Bot games
	SummonersRiftCoopVsAIBeginner = 840
	// Summoner's Rift	Co-op vs. AI Intermediate Bot games
	SummonersRiftCoopVsAIIntermediate = 850
	// Summoner's Rift	ARURF games
	SummonersRiftARURF = 900
	// Crystal Scar	Ascension games
	Ascension = 910
	// Howling Abyss	Legend of the Poro King games
	HowlingAbyssPoroKing = 920
	// Summoner's Rift	Nexus Siege games
	NexusSiege = 940
	// Summoner's Rift	Doom Bots Voting games
	DoomBotsVoting = 950
	// Summoner's Rift	Doom Bots Standard games
	DoomBotsStandard = 960
	// Valoran City Park	Star Guardian Invasion: Normal games
	StarGuardianNormal = 980
	// Valoran City Park	Star Guardian Invasion: Onslaught games
	StarGuardianOnslaught = 990
	// 	Overcharge	PROJECT: Hunters games
	ProjectHuntersOvercharge = 1000
	// 	Summoner's Rift	Snow ARURF games
	SummonersRiftSnowARURF = 1010
)

// MAPS - mapId
const (
	// SummonersRiftOriginal - Summoner's Rift Original Summer Variant
	SummonersRiftOriginal = 1
	// SummonersRiftAutumn - Summoner's Rift Original Autumn Variant
	SummonersRiftAutumn = 2
	// ProvingGrounds - Proving Grounds Tutorial Map
	ProvingGrounds = 3
	// TwistedTreelineOriginal - Twisted Treeline	Original Version
	TwistedTreelineOriginal = 4
	// CrystalScar - The Crystal Scar Dominion Map
	CrystalScar = 8
	// TwistedTreeline - Twisted Treeline Current Version
	TwistedTreeline = 10
	// SummonersRift - Summoner's Rift Current Version
	SummonersRift = 11
	// HowlingAbyss - Howling Abyss ARAM Map
	HowlingAbyss = 12
	// ButchersBridge - Butcher's Bridge ARAM Map
	ButchersBridge = 14
	// DarkStar - Cosmic Ruins Dark Star: Singularity Map
	DarkStar = 16
	// StarGuardian - Valoran City Park Star Guardian Invasion Map
	StarGuardian = 18
	// ProjectHunters - Substructure 43	PROJECT: Hunters Map
	ProjectHunters = 19
)

// GAME MODES - gameMode
const (
	// Classic Summoner's Rift and Twisted Treeline games
	ClassicMode = "CLASSIC"
	// Dominion/Crystal Scar games
	OdinMode = "ODIN"
	// ARAM games
	AramMode = "ARAM"
	// Tutorial games
	TutorialMode = "TUTORIAL"
	// URF games
	UrfMode = "URF"
	// Doom Bot games
	DoombotsMode = "DOOMBOTSTEEMO"
	// One for All games
	OneForAllMode = "ONEFORALL"
	// Ascension games
	AscensionMode = "ASCENSION"
	// Snowdown Showdown games
	FirstBloodMode = "FIRSTBLOOD"
	// Legend of the Poro King games
	KingPoroMode = "KINGPORO"
	// Nexus Siege games
	SiegeMode = "SIEGE"
	// Blood Hunt Assassin games
	AssassinateMode = "ASSASSINATE"
	// All Random Summoner's Rift games
	ArsrMode = "ARSR"
	// Dark Star: Singularity games
	DarkStarMode = "DARKSTAR"
	// Star Guardian Invasion games
	StarGuardianMode = "STARGUARDIAN"
	// PROJECT: Hunters games
	ProjectMode = "PROJECT"
)

// GAME TYPES - gameType
const (
	// Custom games
	Custom = "CUSTOM_GAME"
	// Tutorial games
	Tutorial = "TUTORIAL_GAME"
	// All other games
	Matched = "MATCHED_GAME"
)

// REGIONS
const (
	Brazil        = "br"
	EuropeNE      = "eune"
	EuropeW       = "euw"
	Korea         = "kr"
	LatinAmericaN = "lan"
	LatinAmericaS = "las"
	NorthAmerica  = "na"
	Oceania       = "oce"
	Turkey        = "tr"
	Russia        = "ru"
	Japan         = "jp"
)
