package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/gocarina/gocsv"
	"go.uber.org/zap"
)

//go:embed spell_full.csv
var defaultDB []byte

func loadSource(csvPathOverride string, logger *zap.SugaredLogger) ([]*Spell, error) {
	fileBytes := defaultDB
	if csvPathOverride != "" {
		logger.Infof("Loading CSV from %q...", csvPathOverride)
		file, err := os.Open(csvPathOverride)
		if err != nil {
			return nil, fmt.Errorf("open file %s: %w", csvPathOverride, err)
		}
		defer func() {
			err := file.Close()
			if err != nil {
				panic(err)
			}
		}()
		fileBytes, err = io.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("read file %s: %w", csvPathOverride, err)
		}
	} else {
		logger.Debugf("Loading default DB...")
	}

	r := bytes.NewReader(fileBytes)
	rawSpells := make([]*CSVSpell, 0)
	err := gocsv.Unmarshal(r, &rawSpells)
	if err != nil {
		return nil, fmt.Errorf("unmarshal rawSpells: %w", err)
	}

	spells := make([]*Spell, len(rawSpells))
	for i, raw := range rawSpells {
		s, err := fromCSV(*raw)
		if err != nil {
			return nil, fmt.Errorf("unmarshal spell %q: %w", raw.Name, err)
		}
		spells[i] = &s
	}

	sort.Slice(spells, func(i, j int) bool {
		return spells[i].Name < spells[j].Name
	})

	logger.Debugf("Loaded %d spells", len(spells))
	return spells, nil
}
