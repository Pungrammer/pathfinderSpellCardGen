package main

import (
	"fmt"
	"strconv"
)

type CSVLevel string

func sanitizeSpellLevel(value string) (string, error) {
	if value == "" {
		return "NULL", nil
	}

	if value == "NULL" {
		return value, nil
	}

	_, err := strconv.Atoi(value)
	if err != nil {
		return "", err
	}
	return value, nil
}

func emptyToNull(value string) string {
	if value == "" {
		return "NULL"
	}
	return value
}

type CSVSpell struct {
	Name                 string `csv:"name"`
	School               string `csv:"school"`
	Subschool            string `csv:"subschool"`
	Descriptor           string `csv:"descriptor"`
	SpellLevel           string `csv:"spell_level"`
	CastingTime          string `csv:"casting_time"`
	Components           string `csv:"components"`
	CostlyComponents     string `csv:"costly_components"`
	Range                string `csv:"range"`
	Area                 string `csv:"area"`
	Effect               string `csv:"effect"`
	Targets              string `csv:"targets"`
	Duration             string `csv:"duration"`
	Dismissible          string `csv:"dismissible"`
	Shapeable            string `csv:"shapeable"`
	SavingThrow          string `csv:"saving_throw"`
	SpellResistance      string `csv:"spell_resistance"`
	Description          string `csv:"description"`
	DescriptionFormatted string `csv:"description_formatted"`
	Source               string `csv:"source"`
	FullText             string `csv:"full_text"`
	Verbal               string `csv:"verbal"`
	Somatic              string `csv:"somatic"`
	Material             string `csv:"material"`
	Focus                string `csv:"focus"`
	DivineFocus          string `csv:"divine_focus"`
	SorcererLevel        string `csv:"sor"`
	WizardLevel          string `csv:"wiz"`
	ClericLevel          string `csv:"cleric"`
	DruidLevel           string `csv:"druid"`
	RangerLevel          string `csv:"ranger"`
	BardLevel            string `csv:"bard"`
	PaladinLevel         string `csv:"paladin"`
	AlchemistLevel       string `csv:"alchemist"`
	SummonerLevel        string `csv:"summoner"`
	WitchLevel           string `csv:"witch"`
	InquisitorLevel      string `csv:"inquisitor"`
	OracleLevel          string `csv:"oracle"`
	AntipaladinLevel     string `csv:"antipaladin"`
	MagusLevel           string `csv:"magus"`
	AdeptLevel           string `csv:"adept"`
	DeityLevel           string `csv:"deity"`
	SLALevel             string `csv:"SLA_Level"`
	Domain               string `csv:"domain"`
	ShortDescription     string `csv:"short_description"`

	// Descriptors
	Acid              string `csv:"acid"`
	Air               string `csv:"air"`
	Chaotic           string `csv:"chaotic"`
	Cold              string `csv:"cold"`
	Curse             string `csv:"curse"`
	Darkness          string `csv:"darkness"`
	Death             string `csv:"death"`
	Disease           string `csv:"disease"`
	Earth             string `csv:"earth"`
	Electricity       string `csv:"electricity"`
	Emotion           string `csv:"emotion"`
	Evil              string `csv:"evil"`
	Fear              string `csv:"fear"`
	Fire              string `csv:"fire"`
	Force             string `csv:"force"`
	Good              string `csv:"good"`
	LanguageDependent string `csv:"language_dependent"`
	Lawful            string `csv:"lawful"`
	Light             string `csv:"light"`
	MindAffecting     string `csv:"mind_affecting"`
	Pain              string `csv:"pain"`
	Poison            string `csv:"poison"`
	Shadow            string `csv:"shadow"`
	Sonic             string `csv:"sonic"`
	Water             string `csv:"water"`

	LinkText      string `csv:"linktext"`
	ID            string `csv:"id"`
	MaterialCosts string `csv:"material_costs"`
	Bloodline     string `csv:"bloodline"`
	Patron        string `csv:"patron"`
	MythicText    string `csv:"mythic_text"`
	Augmented     string `csv:"augmented"`
	Mythic        string `csv:"mythic"`
	Bloodrager    string `csv:"bloodrager"`
	Shaman        string `csv:"shaman"`
	Psychic       string `csv:"psychic"`
	Medium        string `csv:"medium"`
	Mesmerist     string `csv:"mesmerist"`
	Occultist     string `csv:"occultist"`
	Spiritualist  string `csv:"spiritualist"`
	Skald         string `csv:"skald"`
	Investigator  string `csv:"investigator"`
	Hunter        string `csv:"hunter"`

	// Oddballs
	HauntStatistics   string `csv:"haunt_statistics"`
	Ruse              string `csv:"ruse"`
	Draconic          string `csv:"draconic"`
	Meditative        string `csv:"meditative"`
	SummonerUnchained string `csv:"summoner_unchained"`
}

type Spell struct {
	Name                 string
	School               string
	Subschool            string
	Descriptor           string
	SpellLevel           string
	CastingTime          string
	Components           string
	CostlyComponents     string
	Range                string
	Area                 string
	Effect               string
	Targets              string
	Duration             string
	Dismissible          string
	Shapeable            string
	SavingThrow          string
	SpellResistance      string
	Description          string
	DescriptionFormatted string
	Source               string
	FullText             string
	Verbal               string
	Somatic              string
	Material             string
	Focus                string
	DivineFocus          string
	SLALevel             string
	DeityLevel           string
	Domain               string
	ShortDescription     string

	// Class levels
	SorcererLevel     string
	WizardLevel       string
	ClericLevel       string
	DruidLevel        string
	RangerLevel       string
	BardLevel         string
	PaladinLevel      string
	AlchemistLevel    string
	SummonerLevel     string
	WitchLevel        string
	InquisitorLevel   string
	OracleLevel       string
	AntipaladinLevel  string
	MagusLevel        string
	AdeptLevel        string
	MythicLevel       string
	BloodragerLevel   string
	ShamanLevel       string
	PsychicLevel      string
	MediumLevel       string
	MesmeristLevel    string
	OccultistLevel    string
	SpiritualistLevel string
	SkaldLevel        string
	InvestigatorLevel string
	HunterLevel       string

	// Descriptors
	DescriptorAcid              string
	DescriptorAir               string
	DescriptorChaotic           string
	DescriptorCold              string
	DescriptorCurse             string
	DescriptorDarkness          string
	DescriptorDeath             string
	DescriptorDisease           string
	DescriptorEarth             string
	DescriptorElectricity       string
	DescriptorEmotion           string
	DescriptorEvil              string
	DescriptorFear              string
	DescriptorFire              string
	DescriptorForce             string
	DescriptorGood              string
	DescriptorLanguageDependent string
	DescriptorLawful            string
	DescriptorLight             string
	DescriptorMindAffecting     string
	DescriptorPain              string
	DescriptorPoison            string
	DescriptorShadow            string
	DescriptorSonic             string
	DescriptorWater             string

	LinkText      string
	ID            string
	MaterialCosts string
	Bloodline     string
	Patron        string
	MythicText    string
	Augmented     string

	// Oddballs
	HauntStatistics   string
	Ruse              string
	Draconic          string
	Meditative        string
	SummonerUnchained string
}

func fromCSV(csv CSVSpell) (Spell, error) {
	spell := Spell{
		Name:                 csv.Name,
		School:               csv.School,
		Subschool:            csv.Subschool,
		Descriptor:           csv.Descriptor,
		SpellLevel:           csv.SpellLevel,
		CastingTime:          csv.CastingTime,
		Components:           csv.Components,
		CostlyComponents:     csv.CostlyComponents,
		Range:                csv.Range,
		Area:                 csv.Area,
		Effect:               csv.Effect,
		Targets:              csv.Targets,
		Duration:             csv.Duration,
		Dismissible:          csv.Dismissible,
		Shapeable:            csv.Shapeable,
		SavingThrow:          csv.SavingThrow,
		SpellResistance:      csv.SpellResistance,
		Description:          csv.Description,
		DescriptionFormatted: csv.DescriptionFormatted,
		Source:               csv.Source,
		FullText:             csv.FullText,
		Verbal:               csv.Verbal,
		Somatic:              csv.Somatic,
		Material:             csv.Material,
		Focus:                csv.Focus,
		DivineFocus:          csv.DivineFocus,
		
		SorcererLevel:     csv.SorcererLevel,
		WizardLevel:       csv.WizardLevel,
		ClericLevel:       csv.ClericLevel,
		DruidLevel:        csv.DruidLevel,
		RangerLevel:       csv.RangerLevel,
		BardLevel:         csv.BardLevel,
		PaladinLevel:      csv.PaladinLevel,
		AlchemistLevel:    csv.AlchemistLevel,
		SummonerLevel:     csv.SummonerLevel,
		WitchLevel:        csv.WitchLevel,
		InquisitorLevel:   csv.InquisitorLevel,
		OracleLevel:       csv.OracleLevel,
		AntipaladinLevel:  csv.AntipaladinLevel,
		MagusLevel:        csv.MagusLevel,
		AdeptLevel:        csv.AdeptLevel,
		MythicLevel:       csv.Mythic,
		BloodragerLevel:   csv.Bloodrager,
		ShamanLevel:       csv.Shaman,
		PsychicLevel:      csv.Psychic,
		MediumLevel:       csv.Medium,
		MesmeristLevel:    csv.Mesmerist,
		OccultistLevel:    csv.Occultist,
		SpiritualistLevel: csv.Spiritualist,
		SkaldLevel:        csv.Skald,
		InvestigatorLevel: csv.Investigator,
		HunterLevel:       csv.Hunter,

		SLALevel:         csv.SLALevel,
		DeityLevel:       csv.DeityLevel,
		Domain:           csv.Domain,
		ShortDescription: csv.ShortDescription,

		DescriptorAcid:              csv.Acid,
		DescriptorAir:               csv.Air,
		DescriptorChaotic:           csv.Chaotic,
		DescriptorCold:              csv.Cold,
		DescriptorCurse:             csv.Curse,
		DescriptorDarkness:          csv.Darkness,
		DescriptorDeath:             csv.Death,
		DescriptorDisease:           csv.Disease,
		DescriptorEarth:             csv.Earth,
		DescriptorElectricity:       csv.Electricity,
		DescriptorEmotion:           csv.Emotion,
		DescriptorEvil:              csv.Evil,
		DescriptorFear:              csv.Fear,
		DescriptorFire:              csv.Fire,
		DescriptorForce:             csv.Force,
		DescriptorGood:              csv.Good,
		DescriptorLanguageDependent: csv.LanguageDependent,
		DescriptorLawful:            csv.Lawful,
		DescriptorLight:             csv.Light,
		DescriptorMindAffecting:     csv.MindAffecting,
		DescriptorPain:              csv.Pain,
		DescriptorPoison:            csv.Poison,
		DescriptorShadow:            csv.Shadow,
		DescriptorSonic:             csv.Sonic,
		DescriptorWater:             csv.Water,

		LinkText:          csv.LinkText,
		ID:                csv.ID,
		MaterialCosts:     csv.MaterialCosts,
		Bloodline:         csv.Bloodline,
		Patron:            csv.Patron,
		MythicText:        csv.MythicText,
		Augmented:         csv.Augmented,
		HauntStatistics:   csv.HauntStatistics,
		Ruse:              csv.Ruse,
		Draconic:          csv.Draconic,
		Meditative:        csv.Meditative,
		SummonerUnchained: csv.SummonerUnchained,
	}

	var err error
	spell.SorcererLevel, err = sanitizeSpellLevel(spell.SorcererLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("SorcererLevel: %w", err)
	}
	spell.WizardLevel, err = sanitizeSpellLevel(spell.WizardLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("WizardLevel: %w", err)
	}
	spell.ClericLevel, err = sanitizeSpellLevel(spell.ClericLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("ClericLevel: %w", err)
	}
	spell.DruidLevel, err = sanitizeSpellLevel(spell.DruidLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("DruidLevel: %w", err)
	}
	spell.RangerLevel, err = sanitizeSpellLevel(spell.RangerLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("RangerLevel: %w", err)
	}
	spell.BardLevel, err = sanitizeSpellLevel(spell.BardLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("BardLevel: %w", err)
	}
	spell.PaladinLevel, err = sanitizeSpellLevel(spell.PaladinLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("PaladinLevel: %w", err)
	}
	spell.AlchemistLevel, err = sanitizeSpellLevel(spell.AlchemistLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("AlchemistLevel: %w", err)
	}
	spell.SummonerLevel, err = sanitizeSpellLevel(spell.SummonerLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("SummonerLevel: %w", err)
	}
	spell.WitchLevel, err = sanitizeSpellLevel(spell.WitchLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("WitchLevel: %w", err)
	}
	spell.InquisitorLevel, err = sanitizeSpellLevel(spell.InquisitorLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("InquisitorLevel: %w", err)
	}
	spell.OracleLevel, err = sanitizeSpellLevel(spell.OracleLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("OracleLevel: %w", err)
	}
	spell.AntipaladinLevel, err = sanitizeSpellLevel(spell.AntipaladinLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("AntipaladinLevel: %w", err)
	}
	spell.MagusLevel, err = sanitizeSpellLevel(spell.MagusLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("MagusLevel: %w", err)
	}
	spell.AdeptLevel, err = sanitizeSpellLevel(spell.AdeptLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("AdeptLevel: %w", err)
	}
	spell.MythicLevel, err = sanitizeSpellLevel(spell.MythicLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("MythicLevel: %w", err)
	}
	spell.BloodragerLevel, err = sanitizeSpellLevel(spell.BloodragerLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("BloodragerLevel: %w", err)
	}
	spell.ShamanLevel, err = sanitizeSpellLevel(spell.ShamanLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("ShamanLevel: %w", err)
	}
	spell.PsychicLevel, err = sanitizeSpellLevel(spell.PsychicLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("PsychicLevel: %w", err)
	}
	spell.MediumLevel, err = sanitizeSpellLevel(spell.MediumLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("MediumLevel: %w", err)
	}
	spell.MesmeristLevel, err = sanitizeSpellLevel(spell.MesmeristLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("MesmeristLevel: %w", err)
	}
	spell.OccultistLevel, err = sanitizeSpellLevel(spell.OccultistLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("OccultistLevel: %w", err)
	}
	spell.SpiritualistLevel, err = sanitizeSpellLevel(spell.SpiritualistLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("SpiritualistLevel: %w", err)
	}
	spell.SkaldLevel, err = sanitizeSpellLevel(spell.SkaldLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("SkaldLevel: %w", err)
	}
	spell.InvestigatorLevel, err = sanitizeSpellLevel(spell.InvestigatorLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("InvestigatorLevel: %w", err)
	}
	spell.HunterLevel, err = sanitizeSpellLevel(spell.HunterLevel)
	if err != nil {
		return Spell{}, fmt.Errorf("HunterLevel: %w", err)
	}

	spell.SLALevel, err = sanitizeSpellLevel(spell.SLALevel)
	if err != nil {
		return Spell{}, fmt.Errorf("SLALevel: %w", err)
	}

	spell.DeityLevel = emptyToNull(spell.DeityLevel)

	return spell, nil
}

func (s *Spell) toHTML() string {
	area := handleOptionalField("Area", s.Area)
	effect := handleOptionalField("Effect", s.Effect)
	targets := handleOptionalField("Targets", s.Targets)

	savingThrow := s.SavingThrow
	if savingThrow == "" {
		savingThrow = "none"
	}
	spellResistance := s.SpellResistance
	if spellResistance == "" {
		spellResistance = "no"
	}

	return fmt.Sprintf(""+
		"<h1>%s</h1>"+
		"<b>School:</b> %s (%s); <b>Level:</b> %s<br>"+
		"<br>"+
		"<b>Casting Time:</b> %s<br>"+
		"<b>Components:</b> %s<br>"+
		"<br>"+
		"<b>Range:</b> %s<br>"+
		"%s"+
		"%s"+
		"%s"+
		"<b>Duration:</b> %s<br>"+
		"<b>Saving Throw:</b> %s<br>"+
		"<b>Spell Resistance:</b> %s<br>"+
		"<b>Source:</b> %s<br>"+
		"<br>"+
		"%s",
		s.Name,
		s.School, s.Subschool, s.SpellLevel,
		s.CastingTime,
		s.Components,
		s.Range,
		area,
		effect,
		targets,
		s.Duration,
		savingThrow,
		spellResistance,
		s.Source,
		s.Description,
	)
}

func handleOptionalField(title, value string) string {
	if value != "" {
		return fmt.Sprintf("<b>%s:</b> %s<br>", title, value)
	}
	return ""
}
