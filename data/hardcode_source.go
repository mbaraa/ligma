package data

import (
	"image/color"
	"image/png"
	"os"
	"strconv"

	"github.com/mbaraa/asu_forms/errors"
	"github.com/mbaraa/asu_forms/models"
)

func hex2rgba(hex string, opacity uint16) color.RGBA64 {
	r, _ := strconv.ParseInt(hex[1:3], 16, 32)
	g, _ := strconv.ParseInt(hex[3:5], 16, 32)
	b, _ := strconv.ParseInt(hex[5:7], 16, 32)

	return color.RGBA64{
		R: uint16(r), G: uint16(g), B: uint16(b), A: opacity,
	}
}

var (
	img0, _    = os.Open("./res/pics/society_service.png")
	img, _     = png.Decode(img0)
	ssRawImage = models.NewFormImage(img)

	_ = img0.Close()

	blueGoogle = color.RGBA64{R: 66, G: 133, B: 244, A: 255}

	ssFields = map[string]models.Field{
		"StudentName": models.NewTextField(
			models.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 960, Y: 435},
			ssRawImage,
		),
		"StudentId": models.NewTextField(
			models.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 447, Y: 435},
			ssRawImage, true,
		),
		"AcademicAdvisor": models.NewTextField(
			models.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 904, Y: 492},
			ssRawImage,
		),
		"Major": models.NewTextField(
			models.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 971, Y: 547},
			ssRawImage, true,
		),
		"Date": models.NewTextField(
			models.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 940, Y: 600},
			ssRawImage, true,
		),
		"Semester": models.NewTextField(
			models.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 435, Y: 600},
			ssRawImage, true,
		),
		"ActivityGoal": models.NewTextField(
			models.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 940, Y: 655},
			ssRawImage,
		),
		"TargetedPersonnel": models.NewTextField(
			models.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 803, Y: 1034},
			ssRawImage,
		),
		"ActivityTitle": models.NewTextField(
			models.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 946, Y: 1089},
			ssRawImage, true,
		),
		"DeservedPoints": models.NewTextField(
			models.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 972, Y: 1143},
			ssRawImage, true,
		),
	}

	societyServiceForm = models.NewForm(
		"SocietyService",
		ssRawImage,
		ssFields,
	)
	forms = []*models.Form{
		societyServiceForm,
	}
)

type HardCodeSource struct{}

// ExistsByName reports whether the Form exists or not, and an occurring error
func (s *HardCodeSource) ExistsByName(string) (bool, error) {
	panic(errors.ErrNotImplemented)
}

// Get returns a form depending on its name, and an occurring error
func (s *HardCodeSource) Get(name string) (*models.Form, error) {
	for _, form := range forms {
		if form.GetName() == name {
			return form, nil
		}
	}
	return nil, errors.ErrNoFormFound
}

// GetAll returns all available forms, and an occurring error
func (s *HardCodeSource) GetAll() ([]*models.Form, error) {
	return forms, nil
}

// Count returns the number of available forms, and an occurring error
func (s *HardCodeSource) Count() (int64, error) {
	return int64(len(forms)), nil
}
