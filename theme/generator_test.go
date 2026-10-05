package theme

import (
	"encoding/json"
	"go/parser"
	"go/token"
	"image/color"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func validConfig() *Config {
	c := new(Config)
	c.Colors.Light = make(map[string]string)
	c.Colors.Dark = make(map[string]string)
	for _, name := range []string{"background", "foreground", "card", "card-foreground", "popover", "popover-foreground", "primary", "primary-foreground", "secondary", "secondary-foreground", "muted", "muted-foreground", "accent", "accent-foreground", "destructive", "destructive-foreground", "border", "input", "ring"} {
		c.Colors.Light[name] = "#123456"
		c.Colors.Dark[name] = "#654321"
	}
	return c
}

func writeConfig(t *testing.T, c *Config) string {
	t.Helper()
	data, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "theme.json")
	if err := os.WriteFile(path, data, 0600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestJSONThemeInitializesRenderingDefaults(t *testing.T) {
	th, err := NewThemeFromJSON(writeConfig(t, validConfig()))
	if err != nil {
		t.Fatal(err)
	}
	if th.MaterialTheme == nil || th.MaterialTheme.Shaper == nil || th.Radius.RadiusMD <= 0 {
		t.Fatal("loaded theme lacks rendering defaults")
	}
	if err := ValidateTheme(th); err != nil {
		t.Fatal(err)
	}
	before := th.Colors.Background
	th.ToggleDark()
	if th.Colors.Background == before || th.MaterialTheme.Palette.Bg != th.Colors.Background {
		t.Fatal("dark mode did not update loaded theme and material palette")
	}
}

func TestJSONThemeReportsInvalidConfiguration(t *testing.T) {
	for _, mode := range []bool{false, true} {
		for _, name := range []string{"background", "foreground", "card", "card-foreground", "popover", "popover-foreground", "primary", "primary-foreground", "secondary", "secondary-foreground", "muted", "muted-foreground", "accent", "accent-foreground", "destructive", "destructive-foreground", "border", "input", "ring"} {
			c := validConfig()
			if mode {
				delete(c.Colors.Dark, name)
			} else {
				delete(c.Colors.Light, name)
			}
			if _, err := c.ToColorScheme(mode); err == nil {
				t.Fatalf("missing %s accepted", name)
			}
		}
	}
	for _, hex := range []string{"red", "1234567", "#gg3456", "#12gg56", "#1234gg"} {
		if _, err := hexToNRGBA(hex); err == nil {
			t.Fatalf("invalid color %q accepted", hex)
		}
	}
	if _, err := NewThemeFromJSON(filepath.Join(t.TempDir(), "missing.json")); err == nil {
		t.Fatal("missing file accepted")
	}
	path := filepath.Join(t.TempDir(), "invalid.json")
	if err := os.WriteFile(path, []byte("invalid"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadThemeFromJSON(path); err == nil {
		t.Fatal("invalid JSON accepted")
	}
	if _, err := LoadThemeFromJSON(filepath.Dir(path)); err == nil {
		t.Fatal("directory accepted")
	}
	for _, mode := range []bool{false, true} {
		c := validConfig()
		if mode {
			c.Colors.Dark["background"] = "bad"
		} else {
			c.Colors.Light["background"] = "bad"
		}
		if _, err := NewThemeFromJSON(writeConfig(t, c)); err == nil {
			t.Fatal("bad colors accepted")
		}
	}
}

func TestGeneratedThemeSourceParses(t *testing.T) {
	source := GenerateThemeConstants(validConfig())
	if _, err := parser.ParseFile(token.NewFileSet(), "generated.go", source, parser.AllErrors); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(source, "BACKGROUND_LIGHT") || !strings.Contains(source, "BACKGROUND_DARK") {
		t.Fatal("generated source omitted modes")
	}
	got, err := hexToNRGBA("#123456")
	if err != nil || got != (color.NRGBA{R: 18, G: 52, B: 86, A: 255}) {
		t.Fatalf("converted color=%v, err=%v", got, err)
	}
}
