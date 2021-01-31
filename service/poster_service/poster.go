package poster_service

import (
	"github.com/deoooo/gin_demo/pkg/file"
	"github.com/deoooo/gin_demo/pkg/qrcode"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

type Poster struct {
	PosterName string
	Qr         *qrcode.QrCode
}

func NewPoster(posterName string, qr *qrcode.QrCode) *Poster {
	return &Poster{
		PosterName: posterName,
		Qr:         qr,
	}
}

func GetPosterFlag() string {
	return "poster"
}

func (p *Poster) CheckMergedImage(path string) bool {
	return file.CheckNotExist(path+p.PosterName) != true
}

func (p *Poster) OpenMergedImage(path string) (*os.File, error) {
	f, err := file.MustOpen(p.PosterName, path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

type Rect struct {
	Name string
	X0   int
	Y0   int
	X1   int
	Y1   int
}

type Pt struct {
	X int
	Y int
}

type PosterBg struct {
	Name string
	*Poster
	*Rect
	*Pt
}

func NewPosterBg(name string, p *Poster, rect *Rect, pt *Pt) *PosterBg {
	return &PosterBg{
		Name:   name,
		Poster: p,
		Rect:   rect,
		Pt:     pt,
	}
}

func (p *PosterBg) Generate() (string, string, error) {
	fullPath := qrcode.GetQrCodeFullPath()
	fileName, _, err := p.Qr.Encode(fullPath)
	if err != nil {
		return "", "", err
	}

	if !p.CheckMergedImage(fullPath) {
		mergedF, err := p.OpenMergedImage(fullPath)
		if err != nil {
			return "", "", err
		}
		defer mergedF.Close()

		bgF, err := file.MustOpen(p.Name, fullPath)
		if err != nil {
			return "", "", err
		}
		defer bgF.Close()

		qrF, err := file.MustOpen(fileName, fullPath)
		if err != nil {
			return "", "", err
		}
		defer qrF.Close()

		bgImage, err := jpeg.Decode(bgF)
		if err != nil {
			return "", "", err
		}
		qrImage, err := jpeg.Decode(qrF)
		if err != nil {
			return "", "", err
		}

		jpg := image.NewRGBA(image.Rect(p.Rect.X0, p.Rect.Y0, p.Rect.X1, p.Rect.Y1))

		draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)
		draw.Draw(jpg, jpg.Bounds(), qrImage, qrImage.Bounds().Min.Sub(image.Pt(p.Pt.X, p.Pt.Y)), draw.Over)

		jpeg.Encode(mergedF, jpg, nil)
	}
	return fileName, fullPath, nil
}
