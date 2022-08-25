package controllers

import (
	"encoding/xml"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/leovillar/api-afip-wsfe/models"
	"os"
)

func FEDummy(c *fiber.Ctx) error {

	resp, err := Dummy()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(resp.Body.FEDummyResponse.FEDummyResult)

}

func Dummy() (models.RespFEDummy, error) {

	var newFEDummy models.ReqFEDummy
	var resp models.RespFEDummy

	payload := newFEDummy.NewFEDummy()
	soapAction := "http://ar.gov.afip.dif.FEV1/FEDummy"
	byteData, err := PostSoap(payload, os.Getenv("WSFE_WSDL"), soapAction)
	if err != nil {
		return resp, errors.New("PostSoap --> " + err.Error())
	}

	err = xml.Unmarshal(byteData, &resp)

	return resp, err

}
