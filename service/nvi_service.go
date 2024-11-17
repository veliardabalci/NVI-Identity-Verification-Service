package services

import (
	"bytes"
	"encoding/xml"
	"net/http"
	"nvi-identity-verification/config"
	"nvi-identity-verification/interfaces"
	"nvi-identity-verification/models"
	"nvi-identity-verification/utils"
)

type DefaultNviService struct{}

func NewNviService() interfaces.NviService {
	return &DefaultNviService{}
}

func (s *DefaultNviService) VerifyIdentity(request models.NviRequest) (bool, error) {
	rawXML := utils.GenerateNviSoapRequest(request.ID, request.Name, request.Surname, request.BirthYear)

	resp, err := sendSoapRequest(config.NviServiceURL, rawXML, "http://tckimlik.nvi.gov.tr/WS/TCKimlikNoDogrula")
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var soapResponse models.NviResponse
	if err := xml.NewDecoder(resp.Body).Decode(&soapResponse); err != nil { // resp yerine resp.Body kullanılıyor
		return false, err
	}

	return soapResponse.TCKimlikNoDogrulaResult == "true", nil
}

func (s *DefaultNviService) VerifyForeignIdentity(request models.ForeignNviRequest) (bool, error) {
	rawXML := utils.GenerateForeignNviSoapRequest(request)

	resp, err := sendSoapRequest(config.ForeignNviServiceURL, rawXML, "http://tckimlik.nvi.gov.tr/WS/YabanciKimlikNoDogrula")
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var soapResponse models.ForeignNviResponse
	if err := xml.NewDecoder(resp.Body).Decode(&soapResponse); err != nil { // resp yerine resp.Body kullanılıyor
		return false, err
	}

	return soapResponse.YabanciKimlikNoDogrulaResult == "true", nil
}

func sendSoapRequest(url string, rawXML []byte, soapAction string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(rawXML))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("SOAPAction", soapAction)

	return client.Do(req)
}
