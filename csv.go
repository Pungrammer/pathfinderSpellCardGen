package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/yosssi/gohtml"
	"io/ioutil"
	"os"
)

type spell struct {
	name            string
	school          string
	subSchool       string
	level           string
	castingTime     string
	component       string
	distance        string // would be called range, if it wasn't a keyword
	area            string
	effect          string
	targets         string
	duration        string
	savingThrow     string
	spellResistance string
	description     string
	source          string
}

func fromRecord(record []string) spell {

	return spell{
		name:        record[0],
		school:      record[1],
		subSchool:   record[2],
		level:       record[4],
		castingTime: record[5],
		component:   record[6],
		// 7 is a bool column showing if the spell has costly components
		distance: record[8],
		area:     record[9],
		effect:   record[10],
		targets:  record[11],
		duration: record[12],
		// 13 bool field for dismissible
		// 14 bool field for shapeable
		savingThrow:     record[15],
		spellResistance: record[16],
		// 17 is "description", 18 is "description_formatted", which is HTML formatted, which we want.
		description: record[18],
		source:      record[19],
		// 20 is "full_text", but we don't want that, as the formatting is ugly
		// All other columns are not required right now.
	}
}

func (s *spell) toHTML() string {
	area := handleOptionalField("Area", s.area)
	effect := handleOptionalField("Effect", s.effect)
	targets := handleOptionalField("Targets", s.targets)

	savingThrow := s.savingThrow
	if savingThrow == "" {
		savingThrow = "none"
	}
	spellResistance := s.spellResistance
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
		s.name,
		s.school, s.subSchool, s.level,
		s.castingTime,
		s.component,
		s.distance,
		area,
		effect,
		targets,
		s.duration,
		savingThrow,
		spellResistance,
		s.source,
		s.description,
	)
}

func handleOptionalField(title, value string) string {
	if value != "" {
		return fmt.Sprintf("<b>%s:</b> %s<br>", title, value)
	}
	return ""
}

func fromCSV(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("open file %s:%s", filepath, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("read csv: %s", err)
	}

	wantedSpells := map[string]struct{}{}
	p, err := ioutil.ReadFile("./spells.txt")
	if err != nil {
		return fmt.Errorf("read ./spells.txt: %s", err)
	}
	for _, line := range bytes.Split(p, []byte{'\n'}) {
		if len(line) == 0 {
			// skip empty lines. Might be at end of file to make output of tools like "cat" more pretty.
			continue
		}

		wantedSpells[string(line)] = struct{}{}
	}

	err = os.Mkdir("output", 0755)
	if err != nil {
		if !os.IsExist(err) {
			return fmt.Errorf("create output dir: %s", err)
		}
	}

	header := "<!DOCTYPE html><html lang=\"en\"><head><style>" +
		"@media print {" +
		"  .pagebreak { page-break-before: always; }" +
		"}" +
		"body { font-size: 12px; }" +
		"h1 { font-size: 16px;}" +
		"</style><title></title></head><body>"
	pageBreakDiv := "<div class=\"pagebreak\"> </div>"
	footer := "</body></html>"

	allSpellsContent := ""

	for _, record := range records {
		spell := fromRecord(record)

		if _, ok := wantedSpells[spell.name]; !ok {
			// Skip: Spell is not interesting
			continue
		}

		spellHtml := spell.toHTML()
		allSpellsContent = allSpellsContent + spellHtml + pageBreakDiv

		f, err := os.Create(fmt.Sprintf("./output/%s.html", spell.name))
		if err != nil {
			return fmt.Errorf("create ./output/%s.html: %s", spell.name, err)
		}
		htmlDoc := gohtml.Format(header + spellHtml + footer)
		_, err = f.Write([]byte(htmlDoc))
		if err != nil {
			return fmt.Errorf("write ./output/%s.html: %s", spell.name, err)
		}
	}

	htmlDoc := gohtml.Format(header + allSpellsContent + footer)

	allSpellsFile := "./output/allSpells.html"
	os.Remove(allSpellsFile)
	f, err := os.Create(allSpellsFile)
	if err != nil {
		return fmt.Errorf("create ./output/allSpells.html: %s", err)
	}
	defer func() {
		err = f.Close()
	}()

	_, err = f.Write([]byte(htmlDoc))
	if err != nil {
		return fmt.Errorf("write ./output/allSpells.html: %s", err)
	}
	return nil
}
