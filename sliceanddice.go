package main

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math"
	"os"
	"os/exec"
	"sort"

	"github.com/ftrvxmtrx/tga"
)

type sheet struct {
	name  string
	enum  string
	areas []area
}

type area struct {
	name   string
	rect   image.Rectangle
	frames []frame
}

type frame struct {
	index  int
	suffix string
}

type sequence struct {
	name string
	img  *image.NRGBA
}

var sheets = [...]sheet{
	{
		"main_menu_sheet",
		"MainMenuSheet",
		[]area{
			{
				"settings",
				image.Rect(0, 0, 192, 192),
				[]frame{
					{0, ""},
					{1, "_logo_hover"},
					{2, "_hover"},
					{4, "_profile_hover"},
				},
			},
			{
				"quit",
				image.Rect(4928, 0, 5120, 192),
				[]frame{
					{0, ""},
					{3, "_hover"},
				},
			},
			{
				"logo",
				image.Rect(384, 0, 896, 256),
				[]frame{
					{0, ""},
					{1, "_hover"},
					{2, "_settings_hover"},
					{4, "_profile_hover"},
				},
			},
			{
				"top_button",
				image.Rect(1216, 0, 1728, 208),
				[]frame{
					{4, "_profile_hover"},
				},
			},
			{
				"top_button",
				image.Rect(1760, 0, 2272, 208),
				[]frame{
					{20, "_right_hover"},
				},
			},
			{
				"top_button",
				image.Rect(2304, 0, 2816, 208),
				[]frame{
					{0, ""},
					{20, "_hover"},
				},
			},
			{
				"top_button",
				image.Rect(2848, 0, 3360, 208),
				[]frame{
					{20, "_left_hover"},
				},
			},
			{
				"top_button",
				image.Rect(3392, 0, 3904, 208),
				[]frame{},
			},
			{
				"create_lobby",
				image.Rect(80, 1000, 1240, 1240),
				[]frame{
					{0, ""},
					{1, "_logo_hover"},
					{4, "_profile_hover"},
					{5, "_hover"},
					{19, "_singleplayer_hover"},
				},
			},
			{
				"singleplayer",
				image.Rect(80, 1280, 1120, 1480),
				[]frame{
					{0, ""},
					{5, "_create_lobby_hover"},
					{18, "_quick_join_hover"},
					{19, "_hover"},
				},
			},
			{
				"quick_join",
				image.Rect(80, 1520, 1120, 2240),
				[]frame{
					{0, ""},
					{9, "_below_hover"},
					{19, "_singleplayer_hover"},
				},
			},
			{
				"quick_join",
				image.Rect(80, 2280, 1120, 3000),
				[]frame{
					{9, "_hover"},
					{18, "_above_hover"},
				},
			},
			{
				"workshop",
				image.Rect(80, 3040, 1120, 3632),
				[]frame{
					{0, ""},
					{9, "_quick_join_hover"},
					{21, "_hover"},
				},
			},
			{
				"hoiaf_top_1",
				image.Rect(3440, 320, 5040, 480),
				[]frame{
					{0, ""},
					{3, "_quit_hover"},
					{11, "_hover"},
					{12, "_below_hover"},
				},
			},
			{
				"hoiaf_top_10",
				image.Rect(3520, 480, 5040, 600),
				[]frame{
					{0, ""},
					{3, "_quit_hover_1"},
					{13, "_below_hover"},
				},
			},
			{
				"hoiaf_top_10",
				image.Rect(3520, 600, 5040, 720),
				[]frame{
					{3, "_quit_hover_2"},
					{13, "_hover"},
				},
			},
			{
				"hoiaf_top_10",
				image.Rect(3520, 720, 5040, 840),
				[]frame{
					{3, "_quit_hover_3"},
					{13, "_above_hover"},
				},
			},
			{
				"hoiaf_top_10",
				image.Rect(3520, 840, 5040, 960),
				[]frame{
					{3, "_quit_hover_4"},
				},
			},
			{
				"hoiaf_top_10",
				image.Rect(3520, 960, 5040, 1080),
				[]frame{
					{3, "_quit_hover_5"},
				},
			},
			{
				"hoiaf_top_10",
				image.Rect(3520, 1080, 5040, 1200),
				[]frame{
					{3, "_quit_hover_6"},
				},
			},
			{
				"hoiaf_top_10",
				image.Rect(3520, 1200, 5040, 1320),
				[]frame{
					{3, "_quit_hover_7"},
				},
			},
			{
				"hoiaf_top_10",
				image.Rect(3520, 1320, 5040, 1440),
				[]frame{
					{3, "_quit_hover_8"},
				},
			},
			{
				"hoiaf_top_10",
				image.Rect(3520, 1440, 5040, 1560),
				[]frame{
					{10, "_hoiaf_timer_hover"},
				},
			},
			{
				"hoiaf_timer",
				image.Rect(3440, 1600, 5040, 1760),
				[]frame{
					{0, ""},
					{8, "_event_timer_hover"},
					{10, "_hover"},
					{15, "_hoiaf_top_10_hover"},
				},
			},
			{
				"event_timer",
				image.Rect(3440, 1832, 5040, 2032),
				[]frame{
					{10, "_hoiaf_timer_hover"},
				},
			},
			{
				"event_timer",
				image.Rect(3440, 2032, 5040, 2232),
				[]frame{
					{6, "_below_hover"},
				},
			},
			{
				"event_timer",
				image.Rect(3440, 2232, 5040, 2432),
				[]frame{
					{0, ""},
					{6, "_hover"},
					{7, "_above_hover"},
					{17, "_news_hover"},
				},
			},
			{
				"news",
				image.Rect(3440, 2472, 5040, 3392),
				[]frame{
					{0, ""},
					{6, "_event_timer_hover"},
					{16, "_update_hover"},
					{17, "_hover"},
				},
			},
			{
				"update",
				image.Rect(3440, 3432, 5040, 3632),
				[]frame{
					{0, ""},
					{16, "_hover"},
					{17, "_news_hover"},
				},
			},
			{
				"ticker_left",
				image.Rect(0, 3680, 1200, 3840),
				[]frame{
					{0, ""},
					{21, "_workshop_hover"},
				},
			},
			{
				"ticker_right",
				image.Rect(3320, 3680, 5120, 3840),
				[]frame{
					{0, ""},
					{16, "_update_hover"},
				},
			},
			{
				"ticker_mid",
				image.Rect(2480, 3680, 2640, 3840),
				[]frame{
					{0, ""},
				},
			},
			{
				"top_bar",
				image.Rect(2048, 0, 3072, 192),
				[]frame{
					{22, ""},
					{25, "_button_glow"},
				},
			},
			{
				"top_bar_left",
				image.Rect(0, 0, 1920, 192),
				[]frame{
					{22, ""},
					{23, "_settings_glow"},
					{24, "_logo_glow"},
					{25, "_profile_glow"},
				},
			},
			{
				"top_bar_right",
				image.Rect(3200, 0, 5120, 192),
				[]frame{
					{22, ""},
					{23, "_quit_glow"},
					{25, "_hoiaf_glow"},
				},
			},
		},
	},
	{
		"commander_profile_sheet",
		"CommanderProfileSheet",
		[]area{
			{
				"profile",
				image.Rect(80, 320, 1240, 896),
				[]frame{
					{0, ""},
					{1, "_logo_hover"},
					{2, "_settings_hover"},
					{4, "_hover"},
					{5, "_create_lobby_hover"},
				},
			},
		},
	},
}

func main() {
	sheetSequences := make([][]sequence, len(sheets))
	for i := 0; ; i++ {
		src, err := readFrame(i)
		if err != nil {
			if os.IsNotExist(err) {
				break
			}

			panic(err)
		}

		for j, s := range sheets {
			for _, a := range s.areas {
				for _, f := range a.frames {
					if f.index != i {
						continue
					}

					name := a.name + f.suffix
					fmt.Printf("cropping %q\n", name)

					sub := src.SubImage(a.rect)
					dst := image.NewNRGBA(sub.Bounds().Sub(sub.Bounds().Min))
					draw.Draw(dst, dst.Rect, sub, sub.Bounds().Min, draw.Src)

					sheetSequences[j] = append(sheetSequences[j], sequence{name, dst})
				}
			}
		}
	}

	for sheetIndex, sequences := range sheetSequences {
		sort.Slice(sequences, func(i, j int) bool {
			return sequences[i].name < sequences[j].name
		})

		sequenceOrder := make([]int, len(sequences))
		sortMethods := []func(i, j int) bool{
			func(i, j int) bool {
				a := &sequences[sequenceOrder[i]]
				b := &sequences[sequenceOrder[j]]
				return a.img.Rect.Dx() < b.img.Rect.Dx()
			},
			func(i, j int) bool {
				a := &sequences[sequenceOrder[i]]
				b := &sequences[sequenceOrder[j]]
				return a.img.Rect.Dy() < b.img.Rect.Dy()
			},
			func(i, j int) bool {
				a := &sequences[sequenceOrder[i]]
				b := &sequences[sequenceOrder[j]]
				ax := a.img.Rect.Dx()
				bx := b.img.Rect.Dx()
				if ay := a.img.Rect.Dy(); ax < ay {
					ax = ay
				}
				if by := b.img.Rect.Dy(); bx < by {
					bx = by
				}
				return ax < bx
			},
			func(i, j int) bool {
				a := &sequences[sequenceOrder[i]]
				b := &sequences[sequenceOrder[j]]
				aa := a.img.Rect.Dx() * a.img.Rect.Dy()
				ba := b.img.Rect.Dx() * b.img.Rect.Dy()
				return aa < ba
			},
		}

		var bestTexture *image.NRGBA
		var bestSheetData []byte
		bestSquareness, bestSize, bestTexSize := 1<<30, 1<<30, 1<<30

		// this is the same (naive) algorithm that mksheet.exe uses, except:
		// - the image height and width are limited to 2^22, not 2^11
		// - we first sort the frames by four different methods (width, height, longest side, and total area) to try to get a better pack
		for tryWidth := 1 << 22; tryWidth >= 4; tryWidth >>= 1 {
			for _, sortMethod := range sortMethods {
				for i := range sequenceOrder {
					sequenceOrder[i] = i
				}
				sort.SliceStable(sequenceOrder, sortMethod)
				tex, sheetData, width, height := packSheet(sequences, sequenceOrder, tryWidth, false)
				if tex == nil {
					continue
				}

				size := width * height
				texSize := tex.Rect.Dx() * tex.Rect.Dy()
				squareness := 1
				if width != height {
					squareness = height/width + width/height
				}

				message := "discarding"

				if texSize < bestTexSize || (texSize == bestTexSize && size < bestSize) || (texSize == bestTexSize && size == bestSize && squareness < bestSquareness) {
					bestTexture, _, _, _ = packSheet(sequences, sequenceOrder, tryWidth, true)
					bestSheetData = sheetData
					bestSize = size
					bestTexSize = texSize
					bestSquareness = squareness
					message = "new best"
				}

				fmt.Printf("Packing option: %dx%d (%d pixels) (%s)\n", width, height, size, message)
			}
		}

		// free up memory, maybe
		for i := range sequences {
			sequences[i].img = nil
		}

		if bestTexture == nil {
			panic("failed to pack sheet")
		}

		fmt.Println("writing files...")

		out, err := os.Create(sheets[sheetIndex].name + ".tga")
		if err != nil {
			panic(err)
		}

		err = tga.Encode(out, bestTexture)
		if err != nil {
			panic(err)
		}

		err = out.Close()
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(sheets[sheetIndex].name+".sht", bestSheetData, 0644)
		if err != nil {
			panic(err)
		}

		enum, err := os.Create(sheets[sheetIndex].name + "_enum.txt")
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(enum, "\tDECLARE_HUD_SHEET( %s )\n", sheets[sheetIndex].enum)
		for _, s := range sequences {
			fmt.Fprintf(enum, "\t\tDECLARE_HUD_SHEET_UV( %s ),\n", s.name)
		}
		fmt.Fprintf(enum, "\tEND_HUD_SHEET( %s );\n", sheets[sheetIndex].enum)

		err = enum.Close()
		if err != nil {
			panic(err)
		}

		fmt.Println("compiling vtf...")

		// sorry about the hard-coded path; I'm lazy
		cmd := exec.Command(`D:\Program Files\Steam\steamapps\common\Alien Swarm Reactive Drop\bin\vtex.exe`, `-nopause`, `-nop4`, `-game`, `D:\Program Files\Steam\steamapps\common\Alien Swarm Reactive Drop\reactivedrop`, `-outdir`, `.`, sheets[sheetIndex].name)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			panic(err)
		}

		fmt.Print("\n\n")
	}

	fmt.Println("done!")
}

func readFrame(i int) (*image.NRGBA, error) {
	name := fmt.Sprintf("mainmenu_%04d.png", i)
	fmt.Printf("reading %q\n", name)

	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return nil, err
	}

	return img.(*image.NRGBA), nil
}

func packSheet(sequences []sequence, sequenceOrder []int, width int, copyPixels bool) (*image.NRGBA, []byte, int, int) {
	const padding = 8
	offsets := make([]image.Point, len(sequences))
	row, col, nextRow, maxCol := 0, 0, 0, 0

	for _, i := range sequenceOrder {
		seq := sequences[i]
		if col+seq.img.Rect.Dx() > width {
			col = 0
			row = nextRow
		}

		if col+seq.img.Rect.Dx() > width {
			return nil, nil, 0, 0
		}

		offsets[i].X = col
		offsets[i].Y = row

		if row+seq.img.Rect.Dy()+padding > nextRow {
			nextRow = row + seq.img.Rect.Dy() + padding
		}

		col += seq.img.Rect.Dx() + padding
		if col > maxCol {
			maxCol = col
		}
	}

	// we don't need outer padding because we clamp tex coords
	maxCol -= padding
	nextRow -= padding

	w, h := 1, 1
	for w < maxCol {
		w <<= 1
	}
	for h < nextRow {
		h <<= 1
	}

	sheetData := appendInt(nil, 1) // format version number
	sheetData = appendInt(sheetData, uint32(len(sequences)))

	dst := image.NewNRGBA(image.Rect(0, 0, w, h))
	for i, seq := range sequences {
		rect := seq.img.Rect.Add(offsets[i])
		if copyPixels {
			for offset := padding / 2; offset > 0; offset-- {
				draw.Draw(dst, rect.Add(image.Pt(offset, offset)), seq.img, image.Point{}, draw.Src)
				draw.Draw(dst, rect.Add(image.Pt(-offset, -offset)), seq.img, image.Point{}, draw.Src)
				draw.Draw(dst, rect.Add(image.Pt(offset, -offset)), seq.img, image.Point{}, draw.Src)
				draw.Draw(dst, rect.Add(image.Pt(-offset, offset)), seq.img, image.Point{}, draw.Src)
			}

			for offset := padding / 2; offset > 0; offset-- {
				draw.Draw(dst, rect.Add(image.Pt(-offset, 0)), seq.img, image.Point{}, draw.Src)
				draw.Draw(dst, rect.Add(image.Pt(offset, 0)), seq.img, image.Point{}, draw.Src)
				draw.Draw(dst, rect.Add(image.Pt(0, -offset)), seq.img, image.Point{}, draw.Src)
				draw.Draw(dst, rect.Add(image.Pt(0, offset)), seq.img, image.Point{}, draw.Src)
			}

			draw.Draw(dst, rect, seq.img, image.Point{}, draw.Src)
		}

		sheetData = appendInt(sheetData, uint32(i))
		sheetData = appendInt(sheetData, 1)   // does not loop
		sheetData = appendInt(sheetData, 1)   // number of frames
		sheetData = appendFloat(sheetData, 1) // total sequence time
		sheetData = appendFloat(sheetData, 1) // first (and only) frame time

		// each color channel has a separate UV rectangle, but we are using RGBA so they're all the same
		for j := 0; j < 4; j++ {
			sheetData = appendFloat(sheetData, (float32(rect.Min.X)+0.5)/float32(w))
			sheetData = appendFloat(sheetData, (float32(rect.Min.Y)+0.5)/float32(h))
			sheetData = appendFloat(sheetData, (float32(rect.Max.X)-0.5)/float32(w))
			sheetData = appendFloat(sheetData, (float32(rect.Max.Y)-0.5)/float32(h))
		}
	}

	return dst, sheetData, maxCol, nextRow
}

func appendInt(b []byte, i uint32) []byte {
	var buf [4]byte

	binary.LittleEndian.PutUint32(buf[:], i)

	return append(b, buf[:]...)
}

func appendFloat(b []byte, f float32) []byte {
	return appendInt(b, math.Float32bits(f))
}
