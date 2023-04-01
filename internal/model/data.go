package model

import "encoding/xml"

type Contract struct {
	Text          string `xml:",chardata" json:"-"`
	SchemeVersion string `xml:"schemeVersion,attr" json:"schemeVersion"`
	ID            string `xml:"id" json:"id"`
	RegNum        string `xml:"regNum" json:"regNum"`
	Number        string `xml:"number" json:"number"`
	PublishDate   string `xml:"publishDate" json:"publishDate"`
	SignDate      string `xml:"signDate" json:"signDate"`
	VersionNumber string `xml:"versionNumber" json:"versionNumber"`
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
		ID   string `xml:"id" json:"id"`
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
			Quantity float64 `xml:"quantity" json:"quantity"`
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
}

type ContractProcedure struct {
	Text          string `xml:",chardata" json:"-"`
	SchemeVersion string `xml:"schemeVersion,attr" json:"-"`
	ID            string `xml:"id" json:"id"`
	RegNum        string `xml:"regNum" json:"regNum"`
	PublishDate   string `xml:"publishDate" json:"publishDate"`
	VersionNumber string `xml:"versionNumber" json:"versionNumber"`
	Executions    struct {
		Text  string `xml:",chardata" json:"-"`
		Stage struct {
			Text    string `xml:",chardata" json:"-"`
			EndDate string `xml:"endDate" json:"endDate"`
		} `xml:"stage" json:"stage"`
		OrdinalNumber string `xml:"ordinalNumber" json:"ordinalNumber"`
		Execution     []struct {
			Text   string `xml:",chardata" json:"-"`
			PayDoc struct {
				Text           string `xml:",chardata" json:"-"`
				Sid            string `xml:"sid" json:"sid"`
				DocumentName   string `xml:"documentName" json:"documentName"`
				DocumentDate   string `xml:"documentDate" json:"documentDate"`
				DocumentNum    string `xml:"documentNum" json:"documentNum"`
				PayDocTypeInfo struct {
					Text                string `xml:",chardata" json:"-"`
					DocAcceptancePayDoc struct {
						Text                             string `xml:",chardata" json:"-"`
						IsDocAcceptancePayDoc            string `xml:"isDocAcceptancePayDoc" json:"isDocAcceptancePayDoc"`
						PayDocToDocAcceptanceCompliances struct {
							Text          string `xml:",chardata" json:"-"`
							DocAcceptance struct {
								Text         string `xml:",chardata" json:"-"`
								Sid          string `xml:"sid" json:"sid"`
								Name         string `xml:"name" json:"name"`
								DocumentDate string `xml:"documentDate" json:"documentDate"`
								DocumentNum  string `xml:"documentNum" json:"documentNum"`
							} `xml:"docAcceptance" json:"docAcceptance"`
						} `xml:"payDocToDocAcceptanceCompliances" json:"payDocToDocAcceptanceCompliances"`
					} `xml:"docAcceptancePayDoc" json:"docAcceptancePayDoc"`
				} `xml:"payDocTypeInfo" json:"payDocTypeInfo"`
			} `xml:"payDoc" json:"payDoc"`
			Currency struct {
				Text string `xml:",chardata" json:"-"`
				Code string `xml:"code" json:"code"`
				Name string `xml:"name" json:"name"`
			} `xml:"currency" json:"currency"`
			Paid                  string `xml:"paid" json:"paid"`
			PaidRUR               string `xml:"paidRUR" json:"paidRUR"`
			PaidVAT               string `xml:"paidVAT" json:"paidVAT"`
			PaidVATRUR            string `xml:"paidVATRUR" json:"paidVATRUR"`
			ImproperExecutionText string `xml:"improperExecutionText" json:"improperExecutionText"`
		} `xml:"execution" json:"execution"`
	} `xml:"executions" json:"executions"`
	Termination struct {
		Text            string `xml:",chardata" json:"-"`
		Paid            string `xml:"paid" json:"paid"`
		TerminationDate string `xml:"terminationDate" json:"terminationDate"`
		ReasonInfo      string `xml:"reasonInfo" json:"reasonInfo"`
		Reason          struct {
			Text string `xml:",chardata" json:"-"`
			Code string `xml:"code" json:"code"`
			Name string `xml:"name" json:"name"`
		} `xml:"reason" json:"reason"`
		DocTermination struct {
			Text         string `xml:",chardata" json:"-"`
			Code         string `xml:"code" json:"code"`
			Name         string `xml:"name" json:"name"`
			DocumentDate string `xml:"documentDate" json:"documentDate"`
		} `xml:"docTermination" json:"docTermination"`
	} `xml:"termination" json:"termination"`
	PrintForm struct {
		Text         string `xml:",chardata" json:"-"`
		URL          string `xml:"url" json:"url"`
		DocRegNumber string `xml:"docRegNumber" json:"docRegNumber"`
	} `xml:"printForm" json:"printForm"`
	TerminationDocuments struct {
		Text       string `xml:",chardata" json:"-"`
		Attachment struct {
			Text               string `xml:",chardata" json:"-"`
			PublishedContentId string `xml:"publishedContentId" json:"publishedContentId"`
			FileName           string `xml:"fileName" json:"fileName"`
			DocDescription     string `xml:"docDescription" json:"docDescription"`
			DocRegNumber       string `xml:"docRegNumber" json:"docRegNumber"`
			URL                string `xml:"url" json:"url"`
			CryptoSigns        struct {
				Text      string   `xml:",chardata" json:"0"`
				Signature []string `xml:"signature" json:"signature"`
			} `xml:"cryptoSigns" json:"cryptoSigns"`
		} `xml:"attachment" json:"attachment"`
	} `xml:"terminationDocuments" json:"terminationDocuments"`
	PaymentDocuments struct {
		Text       string `xml:",chardata" json:"-"`
		Attachment []struct {
			Text               string `xml:",chardata" json:"-"`
			PublishedContentId string `xml:"publishedContentId" json:"publishedContentId"`
			FileName           string `xml:"fileName" json:"fileName"`
			DocDescription     string `xml:"docDescription" json:"docDescription"`
			DocRegNumber       string `xml:"docRegNumber" json:"docRegNumber"`
			URL                string `xml:"url" json:"url"`
			CryptoSigns        struct {
				Text      string   `xml:",chardata" json:"-"`
				Signature []string `xml:"signature" json:"signature"`
			} `xml:"cryptoSigns" json:"cryptoSigns"`
		} `xml:"attachment" json:"attachment"`
	} `xml:"paymentDocuments" json:"paymentDocuments"`
	ReceiptDocuments struct {
		Text       string `xml:",chardata" json:"-"`
		Attachment struct {
			Text               string `xml:",chardata" json:"-"`
			PublishedContentId string `xml:"publishedContentId" json:"publishedContentId"`
			FileName           string `xml:"fileName" json:"fileName"`
			DocDescription     string `xml:"docDescription" json:"docDescription"`
			DocRegNumber       string `xml:"docRegNumber" json:"docRegNumber"`
			URL                string `xml:"url" json:"url"`
			CryptoSigns        struct {
				Text      string   `xml:",chardata" json:"-"`
				Signature []string `xml:"signature" json:"signature"`
			} `xml:"cryptoSigns" json:"cryptoSigns"`
		} `xml:"attachment" json:"attachment"`
	} `xml:"receiptDocuments" json:"receiptDocuments"`
	ModificationReason      string `xml:"modificationReason" json:"modificationReason"`
	CurrentContractStage    string `xml:"currentContractStage" json:"currentContractStage"`
	Okpd2okved2             string `xml:"okpd2okved2" json:"okpd2okved2"`
	IsEDIBased              string `xml:"isEDIBased" json:"isEDIBased"`
	IsPURorASFKBased        string `xml:"isPURorASFKBased" json:"isPURorASFKBased"`
	IsUnilateralRefusalAuto string `xml:"isUnilateralRefusalAuto" json:"isUnilateralRefusalAuto"`
}

type Data struct {
	XMLName           xml.Name           `xml:"export"`
	Text              string             `xml:",chardata"`
	Xsi               string             `xml:"xsi,attr"`
	Xmlns             string             `xml:"xmlns,attr"`
	Oos               string             `xml:"oos,attr"`
	Contract          *Contract          `xml:"contract" json:"contract,omitempty"`
	ContractProcedure *ContractProcedure `xml:"contractProcedure" json:"contractProcedure,omitempty"`
}
