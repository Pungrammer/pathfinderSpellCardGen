package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func Test_export(t *testing.T) {

	logger, err := zap.NewProduction()
	require.NoError(t, err)

	type testCase struct {
		expression string
		validate   func(t *testing.T, testDir string)
	}

	testCases := map[string]testCase{
		"by name": {
			expression: `Name == "Detect Magic"`,
			validate: func(t *testing.T, testDir string) {
				dir, err := os.ReadDir(testDir)
				require.NoError(t, err)

				assert.Len(t, dir, 2) // All spells and the spell-specific HTML

				for _, entry := range dir {
					file, err := os.OpenFile(filepath.Join(testDir, entry.Name()), os.O_RDONLY, 0666)
					require.NoError(t, err)

					all, err := io.ReadAll(file)
					require.NoError(t, err)

					s := string(all)

					if entry.Name() == "Detect%20Magic.html" || entry.Name() == "allSpells.html" {
						assert.Contains(t, s, "Detect Magic") // Not perfect, but good enough, I guess
					} else {
						assert.Fail(t, fmt.Sprintf("Unexpected file: %s", entry.Name()))
					}

				}
			},
		},
		"by class": {
			expression: `PaladinLevel != "NULL"`,
			validate: func(t *testing.T, testDir string) {
				dir, err := os.ReadDir(testDir)
				require.NoError(t, err)

				for _, entry := range dir {
					file, err := os.OpenFile(filepath.Join(testDir, entry.Name()), os.O_RDONLY, 0666)
					require.NoError(t, err)

					all, err := io.ReadAll(file)
					require.NoError(t, err)

					s := string(all)
					assert.Contains(t, s, "paladin") // Not perfect, but good enough, I guess
				}
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			dir := os.TempDir()
			testDir := path.Join(dir, "pathfinderSpellCardGen-test", uuid.NewString())
			err := os.MkdirAll(testDir, 0755)
			require.NoError(t, err)
			defer require.NoError(t, os.RemoveAll(testDir))

			err = runExport("", tc.expression, testDir, logger.Sugar())
			require.NoError(t, err)

			readDir, err := os.ReadDir(testDir)
			require.NoError(t, err)

			found := false
			for _, entry := range readDir {
				if entry.Name() == "allSpells.html" {
					found = true
					break
				}
			}
			assert.True(t, found, "allSpells.html not found")

			tc.validate(t, testDir)
		})
	}
}

func Test_listFields(t *testing.T) {
	fields := listFields()
	require.Len(t, fields, 93)
}

func Test_listOptions(t *testing.T) {
	fields := listFields()
	nop := zap.NewNop()

	// Check that the call works for each field
	for s := range fields {
		_, err := listOptions("", s, nop.Sugar())
		require.NoError(t, err)
	}

	// Check a few chosen fields for correctness
	type testData struct {
		field    string
		expected map[string]struct{}
	}

	td := []testData{
		{
			field: "PaladinLevel",
			expected: map[string]struct{}{
				"NULL": {},
				"1":    {},
				"2":    {},
				"3":    {},
				"4":    {},
			},
		},
	}

	for _, tc := range td {
		t.Run(tc.field, func(t *testing.T) {
			opt, err := listOptions("", tc.field, nop.Sugar())
			require.NoError(t, err)

			assert.Equal(t, tc.expected, opt)
		})
	}
}
