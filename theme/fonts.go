package theme

import (
	_ "embed"
	"sync"

	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/text"
)

//go:embed fonts/Geist-Regular.ttf
var geistRegular []byte

//go:embed fonts/Geist-Medium.ttf
var geistMedium []byte

//go:embed fonts/Geist-SemiBold.ttf
var geistSemiBold []byte

//go:embed fonts/Geist-Bold.ttf
var geistBold []byte

//go:embed fonts/GeistMono-Regular.ttf
var geistMonoRegular []byte

//go:embed fonts/GeistMono-Medium.ttf
var geistMonoMedium []byte

//go:embed fonts/GeistMono-SemiBold.ttf
var geistMonoSemiBold []byte

//go:embed fonts/GeistMono-Bold.ttf
var geistMonoBold []byte

var (
	geistFontCollection []font.FontFace
	geistOnce           sync.Once
)

const (
	TypefaceGeist     font.Typeface = "Geist"
	TypefaceGeistMono font.Typeface = "GeistMono"
)

// GeistCollection returns the parsed Geist Sans and Geist Mono font collection.
func GeistCollection() []font.FontFace {
	geistOnce.Do(func() {
		mustParse := func(data []byte, tf font.Typeface, weight font.Weight) {
			face, err := opentype.Parse(data)
			if err != nil {
				return
			}
			geistFontCollection = append(geistFontCollection, font.FontFace{
				Font: font.Font{
					Typeface: tf,
					Weight:   weight,
					Style:    font.Regular,
				},
				Face: face,
			})
		}

		// Geist Sans (Default UI Sans-Serif font for shadcn)
		mustParse(geistRegular, TypefaceGeist, font.Normal)
		mustParse(geistMedium, TypefaceGeist, font.Medium)
		mustParse(geistSemiBold, TypefaceGeist, font.SemiBold)
		mustParse(geistBold, TypefaceGeist, font.Bold)

		// Also register empty/default typeface as Geist
		mustParse(geistRegular, "", font.Normal)
		mustParse(geistMedium, "", font.Medium)
		mustParse(geistSemiBold, "", font.SemiBold)
		mustParse(geistBold, "", font.Bold)

		// Geist Mono (Code/Monospace font for shadcn)
		mustParse(geistMonoRegular, TypefaceGeistMono, font.Normal)
		mustParse(geistMonoMedium, TypefaceGeistMono, font.Medium)
		mustParse(geistMonoSemiBold, TypefaceGeistMono, font.SemiBold)
		mustParse(geistMonoBold, TypefaceGeistMono, font.Bold)
	})
	return geistFontCollection
}

// NewGeistShaper returns a new Gio text shaper with Geist Sans and Geist Mono fonts registered.
func NewGeistShaper() *text.Shaper {
	return text.NewShaper(text.WithCollection(GeistCollection()))
}
