package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/yosssi/gohtml"
	"go.uber.org/zap"
)

const (
	header = "<!DOCTYPE html><html lang=\"en\"><head><style>" +
		"@media print {" +
		"  .pagebreak { page-break-before: always; }" +
		"}" +
		"body { font-size: 12px; }" +
		"h1 { font-size: 16px;}" +
		"</style><title></title></head><body>"
	pageBreakDiv = "<div class=\"pagebreak\"> </div>"
	footer       = "</body></html>"

	allSpellsFile = "allSpells.html"
)

func convert(spells []*Spell, outputDir string, logger *zap.SugaredLogger) error {
	if len(spells) == 0 {
		logger.Errorf("No spells to convert")
		return errors.New("no spells to convert")
	}

	allSpellsContent := ""

	err := os.Mkdir(outputDir, 0755)
	if err != nil {
		if !os.IsExist(err) {
			return fmt.Errorf("create output dir %q: %s", outputDir, err)
		}
	}

	for _, spell := range spells {
		spellHtml := spell.toHTML()
		allSpellsContent = allSpellsContent + spellHtml + pageBreakDiv

		fileName := url.PathEscape(spell.Name)
		f, err := os.Create(fmt.Sprintf("%s/%s.html", outputDir, fileName))
		if err != nil {
			return fmt.Errorf("create %s/%s.html: %s", outputDir, fileName, err)
		}
		htmlDoc := gohtml.Format(header + spellHtml + footer)
		_, err = f.Write([]byte(htmlDoc))
		if err != nil {
			return fmt.Errorf("write %s/%s.html: %s", outputDir, fileName, err)
		}
	}

	htmlDoc := gohtml.Format(header + allSpellsContent + footer)

	allSpellsPath := outputDir + "/" + allSpellsFile
	err = os.Remove(allSpellsPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("remove %q: %s", allSpellsPath, err)
		}
	}
	f, err := os.Create(allSpellsPath)
	if err != nil {
		return fmt.Errorf("create %q: %s", allSpellsPath, err)
	}
	defer func() {
		err = f.Close()
	}()

	_, err = f.Write([]byte(htmlDoc))
	if err != nil {
		return fmt.Errorf("write %q: %s", allSpellsPath, err)
	}
	return nil
}
