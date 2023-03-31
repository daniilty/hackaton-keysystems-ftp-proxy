package model

import "encoding/xml"

type Data struct {
	XMLName  xml.Name `xml:"export"`
	Text     string   `xml:",chardata"`
	Xsi      string   `xml:"xsi,attr"`
	Xmlns    string   `xml:"xmlns,attr"`
	Oos      string   `xml:"oos,attr"`
	Contract struct {
		Text          string `xml:",chardata" json:"-"`
		SchemeVersion string `xml:"schemeVersion,attr" json:"schemeVersion"`
		ID            string `xml:"id" json:"id"`
		RegNum        string `xml:"regNum" json:"regNum"`
		Number        int    `xml:"number" json:"number"`
		PublishDate   string `xml:"publishDate" json:"publishDate"`
		SignDate      string `xml:"signDate" json:"signDate"`
		VersionNumber int    `xml:"versionNumber" json:"versionNumber"`
		Foundation    struct {
			Text     string `xml:",chardata" json:"-"`
			OosOrder struct {
				Text               string `xml:",chardata" json:"chardata"`
				NotificationNumber string `xml:"notificationNumber" json:"notificationNumber"`
				LotNumber          string `xml:"lotNumber" json:"lotNumbert"`
				Placing            string `xml:"placing" json:"placing"`
			} `xml:"oosOrder" json:"oosOrder"`
		} `xml:"foundation" json:"foundation"`
		Customer struct {
			Text     string `xml:",chardata" json:"-"`
			RegNum   string `xml:"regNum" json:"regNum"`
			FullName string `xml:"fullName" json:"fullName"`
			Inn      string `xml:"inn" json:"inn"`
			Kpp      string `xml:"kpp" json:"kpp"`
		} `xml:"customer" json:"customer"`
		ProtocolDate string  `xml:"protocolDate" json:"protocolDate"`
		DocumentBase string  `xml:"documentBase" json:"documentBase"`
		Price        float64 `xml:"price" json:"price"`
		Currency     struct {
			Text string `xml:",chardata" json:"-"`
			Code string `xml:"code" json:"code"`
			Name string `xml:"name" json:"name"`
		} `xml:"currency" json:"currency"`
		SingleCustomerReason struct {
			Text string `xml:",chardata" json:"-"`
			ID   int    `xml:"id" json:"id"`
			Name string `xml:"name" json:"name"`
		} `xml:"singleCustomerReason" json:"singleCustomerReason"`
		ExecutionDate struct {
			Text  string `xml:",chardata" json:"-"`
			Month int    `xml:"month" json:"month"`
			Year  int    `xml:"year" json:"year"`
		} `xml:"executionDate" json:"executionDate"`
		Finances struct {
			Text          string `xml:",chardata" json:"-"`
			FinanceSource string `xml:"financeSource" json:"financeSource"`
			Extrabudget   struct {
				Text string `xml:",chardata" json:"-"`
				Code string `xml:"code" json:"code"`
				Name string `xml:"name" json:"name"`
			} `xml:"extrabudget" json:"extrabudget"`
			Budget struct {
				Text string `xml:",chardata" json:"-"`
				Code string `xml:"code" json:"code"`
				Name string `xml:"name" json:"name"`
			} `xml:"budget" json:"budget"`
			BudgetLevel string `xml:"budgetLevel" json:"budgetLevel"`
			Budgetary   struct {
				Text          string  `xml:",chardata" json:"-"`
				Month         int     `xml:"month" json:"month"`
				Year          int     `xml:"year" json:"year"`
				SubstageMonth int     `xml:"substageMonth" json:"substageMonth"`
				SubstageYear  int     `xml:"substageYear" json:"substageYear"`
				KBK           string  `xml:"KBK" json:"KBK"`
				Price         float64 `xml:"price" json:"price"`
				Comment       string  `xml:"comment" json:"comment"`
			} `xml:"budgetary" json:"budgetary"`
			Extrabudgetary struct {
				Text          string  `xml:",chardata" json:"-"`
				Month         int     `xml:"month" json:"month"`
				Year          int     `xml:"year" json:"year"`
				SubstageMonth int     `xml:"substageMonth" json:"substageMonth"`
				SubstageYear  int     `xml:"substageYear" json:"substageYear"`
				KOSGU         string  `xml:"KOSGU" json:"KOSGU"`
				Price         float64 `xml:"price" json:"price"`
			} `xml:"extrabudgetary" json:"extrabudgetary"`
		} `xml:"finances" json:"finances"`
		Products struct {
			Text    string `xml:",chardata" json:"-"`
			Product []struct {
				Text string `xml:",chardata" json:"-"`
				Sid  string `xml:"sid" json:"sid"`
				OKPD struct {
					Text string `xml:",chardata" json:"-"`
					Code string `xml:"code" json:"code"`
					Name string `xml:"name" json:"name"`
				} `xml:"OKPD" json:"OKPD"`
				Name string `xml:"name" json:"name"`
				OKEI struct {
					Text         string `xml:",chardata" json:"-"`
					Code         string `xml:"code" json:"code"`
					NationalCode string `xml:"nationalCode" json:"nationalCode"`
				} `xml:"OKEI" json:"OKEI"`
				Price    float64 `xml:"price" json:"price"`
				Quantity int     `xml:"quantity" json:"quantity"`
				Sum      float64 `xml:"sum" json:"sum"`
			} `xml:"product" json:"product"`
		} `xml:"products" json:"products"`
		Suppliers struct {
			Text     string `xml:",chardata" json:"-"`
			Supplier []struct {
				Text             string `xml:",chardata" json:"-"`
				ParticipantType  string `xml:"participantType" json:"participantType"`
				Inn              string `xml:"inn" json:"inn"`
				Kpp              string `xml:"kpp" json:"kpp"`
				OrganizationName string `xml:"organizationName" json:"organizationName"`
				Country          struct {
					Text            string `xml:",chardata" json:"-"`
					CountryCode     string `xml:"countryCode" json:"countryCode"`
					CountryFullName string `xml:"countryFullName" json:"countryFullName"`
				} `xml:"country" json:"country"`
				FactualAddress string `xml:"factualAddress" json:"factualAddress"`
				PostAddress    string `xml:"postAddress" json:"postAddress"`
				ContactInfo    struct {
					Text       string `xml:",chardata" json:"-"`
					LastName   string `xml:"lastName" json:"lastName"`
					FirstName  string `xml:"firstName" json:"firstName"`
					MiddleName string `xml:"middleName" json:"middleName"`
				} `xml:"contactInfo" json:"contactInfo"`
				ContactEMail string `xml:"contactEMail" json:"contactEMail"`
				ContactPhone string `xml:"contactPhone" json:"contactPhone"`
				ContactFax   string `xml:"contactFax" json:"contactFax"`
			} `xml:"supplier" json:"supplier"`
		} `xml:"suppliers" json:"suppliers"`
		Href      string `xml:"href" json:"href"`
		PrintForm struct {
			Text      string `xml:",chardata" json:"-"`
			URL       string `xml:"url" json:"url"`
			Signature struct {
				Text string `xml:",chardata" json:"-"`
				Type string `xml:"type,attr" json:"type"`
			} `xml:"signature" json:"signature"`
		} `xml:"printForm" json:"printForm"`
		Modification struct {
			Text        string `xml:",chardata" json:"modification"`
			Type        string `xml:"type" json:"type"`
			Description string `xml:"description" json:"description"`
			Base        string `xml:"base" json:"base"`
		} `xml:"modification"`
		CurrentContractStage string `xml:"currentContractStage" json:"currentContractStage"`
	} `xml:"contract" json:"contract"`
}
