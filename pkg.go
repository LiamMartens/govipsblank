package govipsblank

import "github.com/davidbyttow/govips/v2/vips"

/*
 * Creates a blank image with a single background color in the sRGBA space
 */
func MakeBlankImage(width, height int, color vips.ColorRGBA) (*vips.ImageRef, error) {
	image_ref, err := vips.Black(width, height)
	if err != nil {
		return nil, err
	}

	err = image_ref.ToColorSpace(vips.InterpretationSRGB)
	if err != nil {
		return nil, err
	}

	err = image_ref.AddAlpha()
	if err != nil {
		return nil, err
	}

	err = image_ref.Cast(vips.BandFormatUchar)
	if err != nil {
		return nil, err
	}

	err = image_ref.DrawRect(color, 0, 0, width, height, true)
	if err != nil {
		return nil, err
	}

	return image_ref, nil
}
