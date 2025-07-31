package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("new logger: %s", err))
	}

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage:")
		fmt.Println("  ./pathfinderSpellCardGen filter <filter>")
		fmt.Println("  ./pathfinderSpellCardGen print <fields> <filter>")
		fmt.Println("  ./pathfinderSpellCardGen export <filter>")
		fmt.Println("  ./pathfinderSpellCardGen list fields")
		fmt.Println("  ./pathfinderSpellCardGen list options <fieldname>")
		os.Exit(1)
	}

	switch args[0] {
	case "filter":
		if len(args) < 2 {
			fmt.Println("Usage: ./pathfinderSpellCardGen filter <filter>")
			os.Exit(1)
		}

		filter := args[1]
		spells, err := runFilter("", filter, logger.Sugar())
		if err != nil {
			handleError(err)
		}

		fmt.Println("Found these spells:")
		for _, spell := range spells {
			fmt.Println(" -" + spell.Name)
		}
	case "print":
		if len(args) < 3 {
			fmt.Println("Usage: ./pathfinderSpellCardGen print <fields> <filter>")
			os.Exit(1)
		}

		fieldsArg := args[1]
		var fields []string
		if fieldsArg == "ALL" {
			allFields := listFields()
			for field := range allFields {
				fields = append(fields, field)
			}
		} else {
			fields = strings.Split(fieldsArg, ",")
			for i := range fields {
				fields[i] = strings.TrimSpace(fields[i])
			}
		}
		sort.Strings(fields)

		filter := args[2]
		spells, err := runFilter("", filter, logger.Sugar())
		if err != nil {
			handleError(err)
		}

		fmt.Println("Found these spells:")
		for _, spell := range spells {
			val := reflect.ValueOf(*spell)
			fmt.Println(spell.Name)
			for _, field := range fields {
				fmt.Printf("\t%s:%v\n", field, val.FieldByName(field).Interface())
			}
		}
	case "export":
		if len(args) != 2 {
			fmt.Println("Usage: ./pathfinderSpellCardGen export <filter>")
			os.Exit(1)
		}
		filter := args[1]
		err := runExport("", filter, "./output", logger.Sugar())
		if err != nil {
			handleError(err)
		}
	case "list":
		if len(args) < 2 {
			fmt.Println("Usage:")
			fmt.Println("  ./pathfinderSpellCardGen list fields")
			fmt.Println("  ./pathfinderSpellCardGen list options <fieldname>")
			os.Exit(1)
		}
		switch args[1] {
		case "fields":
			if len(args) != 2 {
				fmt.Println("Usage: ./pathfinderSpellCardGen list fields")
				os.Exit(1)
			}
			fields := listFields()
			sortedFields := make([]string, 0, len(fields))
			for f := range fields {
				sortedFields = append(sortedFields, f)
			}
			sort.Strings(sortedFields)
			fmt.Println("Available fields:")
			for _, field := range sortedFields {
				fmt.Println(" -", field)
			}
		case "options":
			if len(args) != 3 {
				fmt.Println("Usage: ./pathfinderSpellCardGen list options <fieldname>")
				os.Exit(1)
			}
			field := args[2]
			options, err := listOptions("", field, logger.Sugar())
			if err != nil {
				handleError(err)
			}

			sortedOptions := make([]string, 0, len(options))
			for f := range options {
				sortedOptions = append(sortedOptions, f)
			}
			sort.Strings(sortedOptions)

			fmt.Printf("Unique values for field '%s':\n", field)
			for _, option := range sortedOptions {
				fmt.Println(" -" + option)
			}

		default:
			fmt.Printf("Unknown list subcommand: %s\n", args[1])
			os.Exit(1)
		}

	default:
		fmt.Printf("Unknown command: %s\n", args[0])
		os.Exit(1)
	}
}

func handleError(err error) {
	_, subErr := os.Stderr.WriteString(err.Error())
	if subErr != nil {
		panic(fmt.Sprintf("writing to stderr: %s", subErr))
	}

	os.Exit(1)
}

func runFilter(csvPathOverride, filterExp string, logger *zap.SugaredLogger) ([]*Spell, error) {
	spells, err := loadSource(csvPathOverride, logger.Named("load_source"))
	if err != nil {
		return nil, err
	}

	spells, err = filter(filterExp, spells)
	if err != nil {
		return nil, err
	}
	return spells, nil
}

func runExport(csvPathOverride, filterExp, outputDir string, logger *zap.SugaredLogger) error {
	spells, err := runFilter(csvPathOverride, filterExp, logger)
	if err != nil {
		return err
	}

	err = convert(spells, outputDir, logger.Named("convert"))
	if err != nil {
		return err
	}

	return nil
}

func listFields() map[string]struct{} {
	typ := reflect.TypeOf(Spell{})

	fields := map[string]struct{}{}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.PkgPath == "" { // PkgPath is empty for exported fields
			fields[field.Name] = struct{}{}
		}
	}

	return fields
}

func listOptions(csvPathOverride, fieldName string, logger *zap.SugaredLogger) (map[string]struct{}, error) {
	fields := listFields()
	_, ok := fields[fieldName]
	if !ok {
		return nil, fmt.Errorf("Field '%s' does not exist\n", fieldName)
	}

	spells, err := loadSource(csvPathOverride, logger.Named("load_source"))
	if err != nil {
		return nil, fmt.Errorf("loading spells: %s", err)
	}

	set := make(map[any]struct{})

	for _, item := range spells {
		val := reflect.ValueOf(*item)
		fieldVal := val.FieldByName(fieldName)

		set[fieldVal.Interface()] = struct{}{}
	}

	opts := map[string]struct{}{}
	for v := range set {
		s := fmt.Sprintf("%v", v)
		opts[s] = struct{}{}
	}

	return opts, nil
}
