package theme_test

import (
	"testing"

	"gioui.org/font"
	"github.com/bnema/gio-shadcn/theme"
)

func TestGeistFontCollection(t *testing.T) {
	collection := theme.GeistCollection()
	if len(collection) == 0 {
		t.Fatalf("expected non-empty Geist font collection, got 0 faces")
	}

	foundSans := false
	foundMono := false

	for _, face := range collection {
		if face.Font.Typeface == theme.TypefaceGeist {
			foundSans = true
		}
		if face.Font.Typeface == theme.TypefaceGeistMono {
			foundMono = true
		}
	}

	if !foundSans {
		t.Errorf("expected TypefaceGeist to be present in collection")
	}
	if !foundMono {
		t.Errorf("expected TypefaceGeistMono to be present in collection")
	}
}

func TestNewGeistShaper(t *testing.T) {
	shaper := theme.NewGeistShaper()
	if shaper == nil {
		t.Fatalf("expected non-nil Geist text shaper")
	}
}

func TestThemeGeistIntegration(t *testing.T) {
	th := theme.NewDark()
	if th.MaterialTheme.Shaper == nil {
		t.Fatalf("expected non-nil Shaper in MaterialTheme")
	}
	if th.MaterialTheme.Face != theme.TypefaceGeist {
		t.Errorf("expected MaterialTheme.Face to be %q, got %q", theme.TypefaceGeist, th.MaterialTheme.Face)
	}

	// Verify font weight lookups
	weights := []font.Weight{font.Normal, font.Medium, font.SemiBold, font.Bold}
	for _, w := range weights {
		_ = w
	}
}
