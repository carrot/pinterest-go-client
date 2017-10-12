package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/models"
)

// MeLikesController is the controller that is responsible for all
// /v1/me/likes/ endpoints in the Pinterest API.
type MeLikesController struct {
	wreckerClient *wrecker.Wrecker
}

// newMeController instantiates a new MeLikesController
func newMeLikesController(wc *wrecker.Wrecker) *MeLikesController {
	return &MeLikesController{
		wreckerClient: wc,
	}
}

// MeLikesFetchOptionals is a struct that represents the optional parameters
// that can be passed to the Fetch endpoint
type MeLikesFetchOptionals struct {
	Cursor string
}

// Fetch loads all of the pins that the authorized user has liked
// Endpoint: [GET] /v1/me/likes/
func (mlc *MeLikesController) Fetch(optionals *MeLikesFetchOptionals) (*[]models.Pin, *models.Page, error) {
	// Build + execute request
	resp := new(models.Response)
	resp.Data = &[]models.Pin{}
	request := mlc.wreckerClient.Get("/me/likes/").
		URLParam("fields", models.PIN_FIELDS).
		Into(resp)
	if optionals.Cursor != "" {
		request.URLParam("cursor", optionals.Cursor)
	}
	httpResp, err := request.Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return nil, nil, err
	}

	// OK
	return resp.Data.(*[]models.Pin), &resp.Page, nil
}