package utils

import (
	"fmt"
	"nvi-identity-verification/models"
)

func GenerateNviSoapRequest(id, name, surname, birthYear string) []byte {
	return []byte(fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
		<soap:Body>
			<TCKimlikNoDogrula xmlns="http://tckimlik.nvi.gov.tr/WS">
				<TCKimlikNo>%s</TCKimlikNo>
				<Ad>%s</Ad>
				<Soyad>%s</Soyad>
				<DogumYili>%s</DogumYili>
			</TCKimlikNoDogrula>
		</soap:Body>
	</soap:Envelope>`, id, name, surname, birthYear))
}

func GenerateForeignNviSoapRequest(request models.ForeignNviRequest) []byte {
	return []byte(fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
		<soap:Body>
			<YabanciKimlikNoDogrula xmlns="http://tckimlik.nvi.gov.tr/WS">
				<KimlikNo>%s</KimlikNo>
				<Ad>%s</Ad>
				<Soyad>%s</Soyad>
				<DogumGun>%d</DogumGun>
				<DogumAy>%d</DogumAy>
				<DogumYil>%d</DogumYil>
			</YabanciKimlikNoDogrula>
		</soap:Body>
	</soap:Envelope>`, request.ID, request.Name, request.Surname, request.BirthDay, request.BirthMonth, request.BirthYear))
}
