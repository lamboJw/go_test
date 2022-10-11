package types

type Genre uint8

const (
	Blues Genre = iota
	ClassicRock
	Country
	Dance
	Disco
	Funk
	Grunge
	HipHop
	Jazz
	Metal
	NewAge
	Oldies
	Other
	Pop
	RAndB
	Rap
	Reggae
	Rock
	Techno
	Industrial
	Alternative
	Ska
	DeathMetal
	Pranks
	Soundtrack
	EuroTechno
	Ambient
	TripHop
	Vocal
	JazzFunk
	Fusion
	Trance
	Classical
	Instrumental
	Acid
	House
	Game
	SoundClip
	Gospel
	Noise
	AlternRock
	Bass
	Soul
	Punk
	Space
	Meditative
	InstrumentalPop
	InstrumentalRock
	Ethnic
	Gothic
	Darkwave
	TechnoIndustrial
	Electronic
	PopFolk
	Eurodance
	Dream
	SouthernRock
	Comedy
	Cult
	Gangsta
	Top40
	ChristianRap
	PopFunk
	Jungle
	NativeAmerican
	Cabaret
	NewWave
	Psychadelic
	Rave
	Showtunes
	Trailer
	LoFi
	Tribal
	AcidPunk
	AcidJazz
	Polka
	Retro
	Musical
	RockAndRoll
	HardRock
	Folk
	FolkRock
	NationalFolk
	Swing
	FastFusion
	Bebob
	Latin
	Revival
	Celtic
	Bluegrass
	Avantgarde
	GothicRock
	ProgessiveRock
	PsychedelicRock
	SymphonicRock
	SlowRock
	BigBand
	Chorus
	EasyListening
	Acoustic
	Humour
	Speech
	Chanson
	Opera
	ChamberMusic
	Sonata
	Symphony
	BootyBass
	Primus
	PornGroove
	Satire
	SlowJam
	Club
	Tango
	Samba
	Folklore
	Ballad
	PowerBallad
	RhythmicSoul
	Freestyle
	Duet
	PunkRock
	DrumSolo
	Acapella
	EuroHouse
	DanceHall
	Goa
	DrumAndBass
	ClubHouse
	Hardcore
	Terror
	Indie
	BritPop
	Negerpunk
	PolskPunk
	Beat
	ChristianGangstaRap
	HeavyMetal
	BlackMetal
	Crossover
	ContemporaryChristian
	ChristianRock
	Merengue
	Salsa
	TrashMetal
	Anime
	JPop
	Synthpop
)

func (g Genre) String() string {
	maps := map[Genre]string{
		Blues:                 "Blues",
		ClassicRock:           "ClassicRock",
		Country:               "Country",
		Dance:                 "Dance",
		Disco:                 "Disco",
		Funk:                  "Funk",
		Grunge:                "Grunge",
		HipHop:                "Hip-Hop",
		Jazz:                  "Jazz",
		Metal:                 "Metal",
		NewAge:                "NewAge",
		Oldies:                "Oldies",
		Other:                 "Other",
		Pop:                   "Pop",
		RAndB:                 "R&B",
		Rap:                   "Rap",
		Reggae:                "Reggae",
		Rock:                  "Rock",
		Techno:                "Techno",
		Industrial:            "Industrial",
		Alternative:           "Alternative",
		Ska:                   "Ska",
		DeathMetal:            "DeathMetal",
		Pranks:                "Pranks",
		Soundtrack:            "Soundtrack",
		EuroTechno:            "Euro-Techno",
		Ambient:               "Ambient",
		TripHop:               "Trip-Hop",
		Vocal:                 "Vocal",
		JazzFunk:              "Jazz+Funk",
		Fusion:                "Fusion",
		Trance:                "Trance",
		Classical:             "Classical",
		Instrumental:          "Instrumental",
		Acid:                  "Acid",
		House:                 "House",
		Game:                  "Game",
		SoundClip:             "SoundClip",
		Gospel:                "Gospel",
		Noise:                 "Noise",
		AlternRock:            "AlternRock",
		Bass:                  "Bass",
		Soul:                  "Soul",
		Punk:                  "Punk",
		Space:                 "Space",
		Meditative:            "Meditative",
		InstrumentalPop:       "InstrumentalPop",
		InstrumentalRock:      "InstrumentalRock",
		Ethnic:                "Ethnic",
		Gothic:                "Gothic",
		Darkwave:              "Darkwave",
		TechnoIndustrial:      "Techno-Industrial",
		Electronic:            "Electronic",
		PopFolk:               "Pop-Folk",
		Eurodance:             "Eurodance",
		Dream:                 "Dream",
		SouthernRock:          "SouthernRock",
		Comedy:                "Comedy",
		Cult:                  "Cult",
		Gangsta:               "Gangsta",
		Top40:                 "Top40",
		ChristianRap:          "ChristianRap",
		PopFunk:               "Pop/Funk",
		Jungle:                "Jungle",
		NativeAmerican:        "NativeAmerican",
		Cabaret:               "Cabaret",
		NewWave:               "NewWave",
		Psychadelic:           "Psychadelic",
		Rave:                  "Rave",
		Showtunes:             "Showtunes",
		Trailer:               "Trailer",
		LoFi:                  "Lo-Fi",
		Tribal:                "Tribal",
		AcidPunk:              "AcidPunk",
		AcidJazz:              "AcidJazz",
		Polka:                 "Polka",
		Retro:                 "Retro",
		Musical:               "Musical",
		RockAndRoll:           "Rock&Roll",
		HardRock:              "HardRock",
		Folk:                  "Folk",
		FolkRock:              "Folk-Rock",
		NationalFolk:          "NationalFolk",
		Swing:                 "Swing",
		FastFusion:            "FastFusion",
		Bebob:                 "Bebob",
		Latin:                 "Latin",
		Revival:               "Revival",
		Celtic:                "Celtic",
		Bluegrass:             "Bluegrass",
		Avantgarde:            "Avantgarde",
		GothicRock:            "GothicRock",
		ProgessiveRock:        "ProgessiveRock",
		PsychedelicRock:       "PsychedelicRock",
		SymphonicRock:         "SymphonicRock",
		SlowRock:              "SlowRock",
		BigBand:               "BigBand",
		Chorus:                "Chorus",
		EasyListening:         "EasyListening",
		Acoustic:              "Acoustic",
		Humour:                "Humour",
		Speech:                "Speech",
		Chanson:               "Chanson",
		Opera:                 "Opera",
		ChamberMusic:          "ChamberMusic",
		Sonata:                "Sonata",
		Symphony:              "Symphony",
		BootyBass:             "BootyBass",
		Primus:                "Primus",
		PornGroove:            "PornGroove",
		Satire:                "Satire",
		SlowJam:               "SlowJam",
		Club:                  "Club",
		Tango:                 "Tango",
		Samba:                 "Samba",
		Folklore:              "Folklore",
		Ballad:                "Ballad",
		PowerBallad:           "PowerBallad",
		RhythmicSoul:          "RhythmicSoul",
		Freestyle:             "Freestyle",
		Duet:                  "Duet",
		PunkRock:              "PunkRock",
		DrumSolo:              "DrumSolo",
		Acapella:              "Acapella",
		EuroHouse:             "Euro-House",
		DanceHall:             "DanceHall",
		Goa:                   "Goa",
		DrumAndBass:           "Drum&Bass",
		ClubHouse:             "Club-House",
		Hardcore:              "Hardcore",
		Terror:                "Terror",
		Indie:                 "Indie",
		BritPop:               "BritPop",
		Negerpunk:             "Negerpunk",
		PolskPunk:             "PolskPunk",
		Beat:                  "Beat",
		ChristianGangstaRap:   "ChristianGangstaRap",
		HeavyMetal:            "HeavyMetal",
		BlackMetal:            "BlackMetal",
		Crossover:             "Crossover",
		ContemporaryChristian: "ContemporaryChristian",
		ChristianRock:         "ChristianRock",
		Merengue:              "Merengue",
		Salsa:                 "Salsa",
		TrashMetal:            "TrashMetal",
		Anime:                 "Anime",
		JPop:                  "JPop",
		Synthpop:              "Synthpop",
	}
	s, ok := maps[g]
	if !ok {
		return "Unknown"
	}
	return s
}
