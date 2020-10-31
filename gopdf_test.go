package gopdf

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func BenchmarkPdfWithImageHolder(b *testing.B) {

	err := initTesting()
	if err != nil {
		b.Error(err)
		return
	}

	pdf := GoPdf{}
	pdf.Start(Config{PageSize: Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	err = pdf.AddTTFFont("loma", "./test/res/times.ttf")
	if err != nil {
		b.Error(err)
		return
	}

	err = pdf.SetFont("loma", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	bytesOfImg, err := ioutil.ReadFile("./test/res/chilli.jpg")
	if err != nil {
		b.Error(err)
		return
	}

	imgH, err := ImageHolderByBytes(bytesOfImg)
	if err != nil {
		b.Error(err)
		return
	}
	for i := 0; i < b.N; i++ {
		pdf.ImageByHolder(imgH, 20.0, float64(i)*2.0, nil)
	}

	pdf.SetX(250)
	pdf.SetY(200)
	pdf.Cell(nil, "gopher and gopher")

	pdf.WritePdf("./test/out/image_bench.pdf")
}

func initTesting() error {
	err := os.MkdirAll("./test/out", 0777)
	if err != nil {
		return err
	}
	return nil
}

func TestPdfWithImageHolder(t *testing.T) {
	err := initTesting()
	if err != nil {
		t.Error(err)
		return
	}

	pdf := GoPdf{}
	pdf.Start(Config{PageSize: Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	err = pdf.AddTTFFont("loma", "./test/res/times.ttf")
	if err != nil {
		t.Error(err)
		return
	}

	err = pdf.SetFont("loma", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	bytesOfImg, err := ioutil.ReadFile("./test/res/PNG_transparency_demonstration_1.png")
	if err != nil {
		t.Error(err)
		return
	}

	imgH, err := ImageHolderByBytes(bytesOfImg)
	if err != nil {
		t.Error(err)
		return
	}

	err = pdf.ImageByHolder(imgH, 20.0, 20, nil)
	if err != nil {
		t.Error(err)
		return
	}

	err = pdf.ImageByHolder(imgH, 20.0, 200, nil)
	if err != nil {
		t.Error(err)
		return
	}

	pdf.SetX(250)
	pdf.SetY(200)
	pdf.Cell(nil, "gopher and gopher")

	pdf.WritePdf("./test/out/image_test.pdf")
}

/*
func TestBuffer(t *testing.T) {
	b := bytes.NewReader([]byte("ssssssss"))

	b1, err := ioutil.ReadAll(b)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("->%s\n", string(b1))
	b.Seek(0, 0)
	b2, err := ioutil.ReadAll(b)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("+>%s\n", string(b2))
}*/

//https://github.com/signintech/gopdf/issues/143
func TestIssue143(t *testing.T) {
	err := initTesting()
	if err != nil {
		t.Error(err)
		return
	}

	pdf := GoPdf{}

	pdf.Start(Config{PageSize: Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.SetNoCompression()
	pdf.AddPage()

	ttfOpt := defaultTtfFontOption()
	ttfOpt.UseOpenTypeLayout = true

	err = pdf.AddTTFFontWithOption("noto", "./test/res/NotoSansMalayalamUI-Regular.ttf", ttfOpt)
	if err != nil {
		t.Error(err)
		return
	}

	err = pdf.SetFont("noto", "", 14)
	if err != nil {
		t.Fatalf(err.Error())
		return
	}
	pdf.SetX(250)
	pdf.SetY(200)
	/*
		var glyphIndexs = []uint{187, 172}
		//var glyphIndexs = []uint{20, 20}
		var justFakeRunes = []rune{rune(0), rune(0)}

		err = pdf.writeGlyphs(glyphIndexs, justFakeRunes)
		if err != nil {
			t.Fatalf(err.Error())
			return
		}
		_ = glyphIndexs
		_ = justFakeRunes
	*/
	/*
		var glyphIndexs = []uint{72, 72}
		//var glyphIndexs = []uint{20, 20}
		var justFakeRunes = []rune{rune(0), rune(0)}

		err = pdf.writeGlyphs(glyphIndexs, justFakeRunes)
		if err != nil {
			t.Fatalf(err.Error())
			return
		}
		_ = glyphIndexs
		_ = justFakeRunes
	*/
	pdf.Text("ന്മന്മ")
	pdf.Text("ന്മന്മ")

	pdf.WritePdf("./test/out/issue_143.pdf")
}
