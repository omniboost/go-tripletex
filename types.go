package tripletex

import (
	"log"

	"github.com/cydev/zero"
	"github.com/omniboost/go-tripletex/omitempty"
)

type VoucherType struct {
	ID      int    `json:"id"`
	Version int    `json:"version"`
	Name    string `json:"name"`
}

type Postings []Posting

type Posting struct {
	ID                  int         `json:"id,omitempty"`
	Version             int         `json:"version,omitempty"`
	URL                 string      `json:"url"`
	Date                string      `json:"date"`
	Description         string      `json:"description"`
	Account             *Account    `json:"account"`
	Customer            *Customer   `json:"customer,omitempty"`
	Supplier            *Supplier   `json:"supplier,omitempty"`
	Employee            *Employee   `json:"employee,omitempty"`
	Project             *Project    `json:"project,omitempty"`
	Product             *Product    `json:"project,omitempty"`
	Department          *Department `json:"department,omitempty"`
	VATType             *VATType    `json:"vatType,omitempty"`
	Amount              float64     `json:"amount,omitempty"`
	AmountCurrency      float64     `json:"amountCurrency,omitempty"`
	AmountGross         float64     `json:"amountGross"`
	AmountGrossCurrency float64     `json:"amountGrossCurrency"`
	Currency            *Currency   `json:"currency,omitempty"`
	CloseGroup          *CloseGroup `json:"closeGroup,omitempty"`
	InvoiceNumber       string      `json:"invoiceNumber,omitempty"`
	TermOfPayment       string      `json:"termOfPayment,omitempty"`
	Row                 int         `json:"row,omitempty"`
}

func (p Posting) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(p)
}

func (p Posting) IsEmpty() bool {
	return zero.IsZero(p)
}

type Accounts []Account

type Account struct {
	ID                             int       `json:"id,omitempty"`
	Version                        int       `json:"version,omitempty"`
	URL                            string    `json:"url"`
	Number                         int       `json:"number"`
	Name                           string    `json:"name"`
	Description                    string    `json:"description"`
	VATType                        VATType   `json:"vatType"`
	VATLocked                      bool      `json:"vatLocked"`
	Currency                       *Currency `json:"currency,omitempty"`
	IsCloseable                    bool      `json:"isCloseable"`
	IsApplicableForSupplierInvoice bool      `json:"isApplicableForSupplierInvoice"`
	RequireReconciliation          bool      `json:"requireReconciliation"`
	IsInactive                     bool      `json:"isInactive"`
	IsBankAccount                  bool      `json:"isBankAccount"`
	IsInvoiceAccount               bool      `json:"isInvoiceAccount"`
	BankAccountNumber              string    `json:"bankAccountNumber"`
	BankAccountCountry             struct {
		ID      int `json:"id"`
		Version int `json:"version"`
	} `json:"bankAccountCountry"`
	BankName         string `json:"bankName"`
	BankAccountIBAN  string `json:"bankAccountIBAN"`
	BankAccountSWIFT string `json:"bankAccountSWIFT"`
}

func (a Account) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(a)
}

type Customers []Customer

type Customer struct {
	ID                    int             `json:"id,omitempty"`
	Version               int             `json:"version,omitempty"`
	URL                   string          `json:"url"`
	Name                  string          `json:"name"`
	OrganizationNumber    string          `json:"organizationNumber,omitempty"`
	SupplierNumber        int             `json:"supplierNumber,omitempty"`
	CustomerNumber        int             `json:"customerNumber,omitempty"`
	IsSupplier            bool            `json:"isSupplier,omitempty"`
	IsCustomer            bool            `json:"isCustomer,omitempty"`
	IsInactive            bool            `json:"isInactive,omitempty"`
	AccountManager        *AccountManager `json:"accountManager,omitempty"`
	Email                 string          `json:"email,omitempty"`
	InvoiceEmail          string          `json:"invoiceEmail,omitempty"`
	OverdueNoticeEmail    string          `json:"overdueNoticeEmail,omitempty"`
	BankAccounts          []string        `json:"bankAccounts,omitempty"`
	PhoneNumber           string          `json:"phoneNumber,omitempty"`
	PhoneNumberMobile     string          `json:"phoneNumberMobile,omitempty"`
	Description           string          `json:"description,omitempty"`
	IsPrivateIndividual   bool            `json:"isPrivateIndividual,omitempty"`
	SingleCustomerInvoice bool            `json:"singleCustomerInvoice,omitempty"`
	InvoiceSendMethod     string          `json:"invoiceSendMethod,omitempty"`
	EmailAttachmentType   string          `json:"emailAttachmentType,omitempty"`
	PostalAddress         *Address        `json:"postalAddress,omitempty"`
	PhysicalAddress       *Address        `json:"physicalAddress,omitempty"`
	DeliveryAddress       *Address        `json:"deliveryAddress,omitempty"`
	Category1             *Category       `json:"category1,omitempty"`
	Category2             *Category       `json:"category2,omitempty"`
	Category3             *Category       `json:"category3,omitempty"`
	InvoicesDueIn         int             `json:"invoicesDueIn,omitempty"`
	InvoicesDueInType     string          `json:"invoicesDueInType,omitempty"`
}

func (c Customer) IsEmpty() bool {
	return zero.IsZero(c)
}

type Supplier struct {
	// 		ID                  int      `json:"id"`
	// 		Version             int      `json:"version"`
	// 		Name                string   `json:"name"`
	// 		OrganizationNumber  string   `json:"organizationNumber"`
	// 		SupplierNumber      int      `json:"supplierNumber"`
	// 		CustomerNumber      int      `json:"customerNumber"`
	// 		IsCustomer          bool     `json:"isCustomer"`
	// 		Email               string   `json:"email"`
	// 		BankAccounts        []string `json:"bankAccounts"`
	// 		InvoiceEmail        string   `json:"invoiceEmail"`
	// 		OverdueNoticeEmail  string   `json:"overdueNoticeEmail"`
	// 		PhoneNumber         string   `json:"phoneNumber"`
	// 		PhoneNumberMobile   string   `json:"phoneNumberMobile"`
	// 		Description         string   `json:"description"`
	// 		IsPrivateIndividual bool     `json:"isPrivateIndividual"`
	// 		ShowProducts        bool     `json:"showProducts"`
	// 		AccountManager      struct {
	// 			ID                       int    `json:"id"`
	// 			Version                  int    `json:"version"`
	// 			FirstName                string `json:"firstName"`
	// 			LastName                 string `json:"lastName"`
	// 			EmployeeNumber           string `json:"employeeNumber"`
	// 			DateOfBirth              string `json:"dateOfBirth"`
	// 			Email                    string `json:"email"`
	// 			PhoneNumberMobileCountry struct {
	// 				ID      int `json:"id"`
	// 				Version int `json:"version"`
	// 			} `json:"phoneNumberMobileCountry"`
	// 			PhoneNumberMobile      string `json:"phoneNumberMobile"`
	// 			PhoneNumberHome        string `json:"phoneNumberHome"`
	// 			PhoneNumberWork        string `json:"phoneNumberWork"`
	// 			NationalIdentityNumber string `json:"nationalIdentityNumber"`
	// 			Dnumber                string `json:"dnumber"`
	// 			InternationalID        struct {
	// 				IntAmeldingType string `json:"intAmeldingType"`
	// 				Country         struct {
	// 					ID      int `json:"id"`
	// 					Version int `json:"version"`
	// 				} `json:"country"`
	// 				Number string `json:"number"`
	// 			} `json:"internationalId"`
	// 			BankAccountNumber     string `json:"bankAccountNumber"`
	// 			Iban                  string `json:"iban"`
	// 			Bic                   string `json:"bic"`
	// 			CreditorBankCountryID int    `json:"creditorBankCountryId"`
	// 			UsesAbroadPayment     bool   `json:"usesAbroadPayment"`
	// 			UserType              string `json:"userType"`
	// 			Comments              string `json:"comments"`
	// 			Address               struct {
	// 				ID           int    `json:"id"`
	// 				Version      int    `json:"version"`
	// 				AddressLine1 string `json:"addressLine1"`
	// 				AddressLine2 string `json:"addressLine2"`
	// 				PostalCode   string `json:"postalCode"`
	// 				City         string `json:"city"`
	// 				Country      struct {
	// 					ID      int `json:"id"`
	// 					Version int `json:"version"`
	// 				} `json:"country"`
	// 			} `json:"address"`
	// 			Department struct {
	// 				ID               int    `json:"id"`
	// 				Version          int    `json:"version"`
	// 				Name             string `json:"name"`
	// 				DepartmentNumber string `json:"departmentNumber"`
	// 			} `json:"department"`
	// 			Employments []struct {
	// 				ID           int    `json:"id"`
	// 				Version      int    `json:"version"`
	// 				EmploymentID string `json:"employmentId"`
	// 				StartDate    string `json:"startDate"`
	// 				EndDate      string `json:"endDate"`
	// 				Division     struct {
	// 					ID                 int    `json:"id"`
	// 					Version            int    `json:"version"`
	// 					Name               string `json:"name"`
	// 					StartDate          string `json:"startDate"`
	// 					EndDate            string `json:"endDate"`
	// 					OrganizationNumber string `json:"organizationNumber"`
	// 					Municipality       struct {
	// 						ID      int `json:"id"`
	// 						Version int `json:"version"`
	// 					} `json:"municipality"`
	// 				} `json:"division"`
	// 				LastSalaryChangeDate     string `json:"lastSalaryChangeDate"`
	// 				NoEmploymentRelationship bool   `json:"noEmploymentRelationship"`
	// 				IsMainEmployer           bool   `json:"isMainEmployer"`
	// 				TaxDeductionCode         string `json:"taxDeductionCode"`
	// 				EmploymentDetails        []struct {
	// 					ID                 int    `json:"id"`
	// 					Version            int    `json:"version"`
	// 					Date               string `json:"date"`
	// 					EmploymentType     string `json:"employmentType"`
	// 					MaritimeEmployment struct {
	// 						ShipRegister string `json:"shipRegister"`
	// 						ShipType     string `json:"shipType"`
	// 						TradeArea    string `json:"tradeArea"`
	// 					} `json:"maritimeEmployment"`
	// 					RemunerationType   string `json:"remunerationType"`
	// 					WorkingHoursScheme string `json:"workingHoursScheme"`
	// 					ShiftDurationHours int    `json:"shiftDurationHours"`
	// 					OccupationCode     struct {
	// 						ID      int    `json:"id"`
	// 						Version int    `json:"version"`
	// 						NameNO  string `json:"nameNO"`
	// 						Code    string `json:"code"`
	// 					} `json:"occupationCode"`
	// 					PercentageOfFullTimeEquivalent int `json:"percentageOfFullTimeEquivalent"`
	// 					AnnualSalary                   int `json:"annualSalary"`
	// 					HourlyWage                     int `json:"hourlyWage"`
	// 					PayrollTaxMunicipalityID       struct {
	// 						ID      int `json:"id"`
	// 						Version int `json:"version"`
	// 					} `json:"payrollTaxMunicipalityId"`
	// 				} `json:"employmentDetails"`
	// 			} `json:"employments"`
	// 			HolidayAllowanceEarned struct {
	// 				Year                   int `json:"year"`
	// 				Amount                 int `json:"amount"`
	// 				Basis                  int `json:"basis"`
	// 				AmountExtraHolidayWeek int `json:"amountExtraHolidayWeek"`
	// 			} `json:"holidayAllowanceEarned"`
	// 		} `json:"accountManager"`
	// 		PostalAddress struct {
	// 			ID           int    `json:"id"`
	// 			Version      int    `json:"version"`
	// 			AddressLine1 string `json:"addressLine1"`
	// 			AddressLine2 string `json:"addressLine2"`
	// 			PostalCode   string `json:"postalCode"`
	// 			City         string `json:"city"`
	// 			Country      struct {
	// 				ID      int `json:"id"`
	// 				Version int `json:"version"`
	// 			} `json:"country"`
	// 		} `json:"postalAddress"`
	// 		PhysicalAddress struct {
	// 			ID           int    `json:"id"`
	// 			Version      int    `json:"version"`
	// 			AddressLine1 string `json:"addressLine1"`
	// 			AddressLine2 string `json:"addressLine2"`
	// 			PostalCode   string `json:"postalCode"`
	// 			City         string `json:"city"`
	// 			Country      struct {
	// 				ID      int `json:"id"`
	// 				Version int `json:"version"`
	// 			} `json:"country"`
	// 		} `json:"physicalAddress"`
	// 		DeliveryAddress struct {
	// 			ID       int `json:"id"`
	// 			Version  int `json:"version"`
	// 			Employee struct {
	// 				ID                       int    `json:"id"`
	// 				Version                  int    `json:"version"`
	// 				FirstName                string `json:"firstName"`
	// 				LastName                 string `json:"lastName"`
	// 				EmployeeNumber           string `json:"employeeNumber"`
	// 				DateOfBirth              string `json:"dateOfBirth"`
	// 				Email                    string `json:"email"`
	// 				PhoneNumberMobileCountry struct {
	// 					ID      int `json:"id"`
	// 					Version int `json:"version"`
	// 				} `json:"phoneNumberMobileCountry"`
	// 				PhoneNumberMobile      string `json:"phoneNumberMobile"`
	// 				PhoneNumberHome        string `json:"phoneNumberHome"`
	// 				PhoneNumberWork        string `json:"phoneNumberWork"`
	// 				NationalIdentityNumber string `json:"nationalIdentityNumber"`
	// 				Dnumber                string `json:"dnumber"`
	// 				InternationalID        struct {
	// 					IntAmeldingType string `json:"intAmeldingType"`
	// 					Country         struct {
	// 						ID      int `json:"id"`
	// 						Version int `json:"version"`
	// 					} `json:"country"`
	// 					Number string `json:"number"`
	// 				} `json:"internationalId"`
	// 				BankAccountNumber     string `json:"bankAccountNumber"`
	// 				Iban                  string `json:"iban"`
	// 				Bic                   string `json:"bic"`
	// 				CreditorBankCountryID int    `json:"creditorBankCountryId"`
	// 				UsesAbroadPayment     bool   `json:"usesAbroadPayment"`
	// 				UserType              string `json:"userType"`
	// 				Comments              string `json:"comments"`
	// 				Address               struct {
	// 					ID           int    `json:"id"`
	// 					Version      int    `json:"version"`
	// 					AddressLine1 string `json:"addressLine1"`
	// 					AddressLine2 string `json:"addressLine2"`
	// 					PostalCode   string `json:"postalCode"`
	// 					City         string `json:"city"`
	// 					Country      struct {
	// 						ID      int `json:"id"`
	// 						Version int `json:"version"`
	// 					} `json:"country"`
	// 				} `json:"address"`
	// 				Department struct {
	// 					ID               int    `json:"id"`
	// 					Version          int    `json:"version"`
	// 					Name             string `json:"name"`
	// 					DepartmentNumber string `json:"departmentNumber"`
	// 				} `json:"department"`
	// 				Employments []struct {
	// 					ID           int    `json:"id"`
	// 					Version      int    `json:"version"`
	// 					EmploymentID string `json:"employmentId"`
	// 					StartDate    string `json:"startDate"`
	// 					EndDate      string `json:"endDate"`
	// 					Division     struct {
	// 						ID                 int    `json:"id"`
	// 						Version            int    `json:"version"`
	// 						Name               string `json:"name"`
	// 						StartDate          string `json:"startDate"`
	// 						EndDate            string `json:"endDate"`
	// 						OrganizationNumber string `json:"organizationNumber"`
	// 						Municipality       struct {
	// 							ID      int `json:"id"`
	// 							Version int `json:"version"`
	// 						} `json:"municipality"`
	// 					} `json:"division"`
	// 					LastSalaryChangeDate     string `json:"lastSalaryChangeDate"`
	// 					NoEmploymentRelationship bool   `json:"noEmploymentRelationship"`
	// 					IsMainEmployer           bool   `json:"isMainEmployer"`
	// 					TaxDeductionCode         string `json:"taxDeductionCode"`
	// 					EmploymentDetails        []struct {
	// 						ID                 int    `json:"id"`
	// 						Version            int    `json:"version"`
	// 						Date               string `json:"date"`
	// 						EmploymentType     string `json:"employmentType"`
	// 						MaritimeEmployment struct {
	// 							ShipRegister string `json:"shipRegister"`
	// 							ShipType     string `json:"shipType"`
	// 							TradeArea    string `json:"tradeArea"`
	// 						} `json:"maritimeEmployment"`
	// 						RemunerationType   string `json:"remunerationType"`
	// 						WorkingHoursScheme string `json:"workingHoursScheme"`
	// 						ShiftDurationHours int    `json:"shiftDurationHours"`
	// 						OccupationCode     struct {
	// 							ID      int    `json:"id"`
	// 							Version int    `json:"version"`
	// 							NameNO  string `json:"nameNO"`
	// 							Code    string `json:"code"`
	// 						} `json:"occupationCode"`
	// 						PercentageOfFullTimeEquivalent int `json:"percentageOfFullTimeEquivalent"`
	// 						AnnualSalary                   int `json:"annualSalary"`
	// 						HourlyWage                     int `json:"hourlyWage"`
	// 						PayrollTaxMunicipalityID       struct {
	// 							ID      int `json:"id"`
	// 							Version int `json:"version"`
	// 						} `json:"payrollTaxMunicipalityId"`
	// 					} `json:"employmentDetails"`
	// 				} `json:"employments"`
	// 				HolidayAllowanceEarned struct {
	// 					Year                   int `json:"year"`
	// 					Amount                 int `json:"amount"`
	// 					Basis                  int `json:"basis"`
	// 					AmountExtraHolidayWeek int `json:"amountExtraHolidayWeek"`
	// 				} `json:"holidayAllowanceEarned"`
	// 			} `json:"employee"`
	// 			AddressLine1 string `json:"addressLine1"`
	// 			AddressLine2 string `json:"addressLine2"`
	// 			PostalCode   string `json:"postalCode"`
	// 			City         string `json:"city"`
	// 			Country      struct {
	// 				ID      int `json:"id"`
	// 				Version int `json:"version"`
	// 			} `json:"country"`
	// 			Name string `json:"name"`
	// 		} `json:"deliveryAddress"`
	// 		Category1 struct {
	// 			ID          int    `json:"id"`
	// 			Version     int    `json:"version"`
	// 			Name        string `json:"name"`
	// 			Number      string `json:"number"`
	// 			Description string `json:"description"`
	// 			Type        int    `json:"type"`
	// 		} `json:"category1"`
	// 		Category2 struct {
	// 			ID          int    `json:"id"`
	// 			Version     int    `json:"version"`
	// 			Name        string `json:"name"`
	// 			Number      string `json:"number"`
	// 			Description string `json:"description"`
	// 			Type        int    `json:"type"`
	// 		} `json:"category2"`
	// 		Category3 struct {
	// 			ID          int    `json:"id"`
	// 			Version     int    `json:"version"`
	// 			Name        string `json:"name"`
	// 			Number      string `json:"number"`
	// 			Description string `json:"description"`
	// 			Type        int    `json:"type"`
	// 		} `json:"category3"`
}

type Employees []Employee

type Employee struct {
	ID                       int    `json:"id"`
	Version                  int    `json:"version"`
	URL                      string `json:"url"`
	FirstName                string `json:"firstName"`
	LastName                 string `json:"lastName"`
	EmployeeNumber           string `json:"employeeNumber"`
	DateOfBirth              string `json:"dateOfBirth"`
	Email                    string `json:"email"`
	PhoneNumberMobileCountry struct {
		ID  int    `json:"id"`
		URL string `json:"url"`
	} `json:"phoneNumberMobileCountry"`
	PhoneNumberMobile      string `json:"phoneNumberMobile"`
	PhoneNumberHome        string `json:"phoneNumberHome"`
	PhoneNumberWork        string `json:"phoneNumberWork"`
	NationalIdentityNumber string `json:"nationalIdentityNumber"`
	Dnumber                string `json:"dnumber"`
	InternationalID        struct {
		IntAmeldingType interface{} `json:"intAmeldingType"`
		Country         interface{} `json:"country"`
		Number          string      `json:"number"`
	} `json:"internationalId"`
	BankAccountNumber            string `json:"bankAccountNumber"`
	Iban                         string `json:"iban"`
	Bic                          string `json:"bic"`
	CreditorBankCountryID        int    `json:"creditorBankCountryId"`
	UsesAbroadPayment            bool   `json:"usesAbroadPayment"`
	AllowInformationRegistration bool   `json:"allowInformationRegistration"`
	IsContact                    bool   `json:"isContact"`
	Comments                     string `json:"comments"`
	Address                      struct {
		ID  int    `json:"id"`
		URL string `json:"url"`
	} `json:"address"`
	Department struct {
		ID  int    `json:"id"`
		URL string `json:"url"`
	} `json:"department"`
	Employments []struct {
		ID  int    `json:"id"`
		URL string `json:"url"`
	} `json:"employments"`
	HolidayAllowanceEarned struct {
		Year                   int     `json:"year"`
		Amount                 float64 `json:"amount"`
		Basis                  float64 `json:"basis"`
		AmountExtraHolidayWeek float64 `json:"amountExtraHolidayWeek"`
	} `json:"holidayAllowanceEarned"`
	EmployeeCategory interface{} `json:"employeeCategory"`
}

func (e Employee) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(e)
}

func (e Employee) IsEmpty() bool {
	return zero.IsZero(e)
}

type Project struct {
}

type Product struct {
}

type Departments []Department

type Department struct {
	ID                int         `json:"id,omitempty"`
	Version           int         `json:"version,omitempty"`
	URL               string      `json:"url,omitempty"`
	Name              string      `json:"name,omitempty"`
	DepartmentNumber  string      `json:"departmentNumber,omitempty"`
	DepartmentManager interface{} `json:"departmentManager,omitempty"`
	DisplayName       string      `json:"displayName,omitempty"`
	IsInactive        bool        `json:"isInactive,omitempty"`
}

func (d Department) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(d)
}

func (d Department) IsEmpty() bool {
	return zero.IsZero(d)
}

type VATTypes []VATType

type VATType struct {
	ID         int     `json:"id"`
	Version    int     `json:"version"`
	URL        string  `json:"url"`
	Name       string  `json:"name"`
	Number     string  `json:"number"`
	Percentage float64 `json:"percentage"`
}

type Currency struct {
	// ID          int    `json:"id,omitempty"`
	// Version     int    `json:"version,omitempty"`
	// Code        string `json:"code"`
	// Description string `json:"description,omitempty"`
	// Factor      int    `json:"factor,omitempty"`
}

func (c *Currency) IsEmpty() bool {
	return c == nil || zero.IsZero(c)
}

type CloseGroup struct {
	// 		ID      int    `json:"id"`
	// 		Version int    `json:"version"`
	// 		Date    string `json:"date"`
}

type Document struct {
	// 	ID       int    `json:"id"`
	// 	Version  int    `json:"version"`
	// 	FileName string `json:"fileName"`
}

type Attachment struct {
	// 	ID       int    `json:"id"`
	// 	Version  int    `json:"version"`
	// 	FileName string `json:"fileName"`
}

type EDIDocument struct {
	// 	ID       int    `json:"id"`
	// 	Version  int    `json:"version"`
	// 	FileName string `json:"fileName"`
}

type Invoice struct {
	ID             int      `json:"id,omitempty"`
	Version        int      `json:"version,omitempty"`
	URL            string   `json:"url,omitempty"`
	InvoiceNumber  int      `json:"invoiceNumber,omitempty"`
	InvoiceDate    string   `json:"invoiceDate,omitempty"`
	Customer       Customer `json:"customer,omitempty"`
	InvoiceDueDate string   `json:"invoiceDueDate,omitempty"`
	KID            string   `json:"kid,omitempty"`
	Comment        string   `json:"comment,omitempty"`
	Orders         Orders   `json:"orders,omitempty"`
	Voucher        Voucher  `json:"voucher,omitempty"`
	// Currency       Currency `json:"currency"`
	InvoiceRemarks string  `json:"invoiceRemarks,omitempty"`
	PaymentTypeID  int     `json:"paymentTypeId,omitempty"`
	PaidAmount     float64 `json:"paidAmount,omitempty"`
	EhfSendStatus  string  `json:"ehfSendStatus,omitempty"`
}

func (i Invoice) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(i)
}

func (i Invoice) IsEmpty() bool {
	return zero.IsZero(i)
}

type Vouchers []Voucher

type Voucher struct {
	ID          int          `json:"id,omitempty"`
	Version     int          `json:"version,omitempty"`
	Date        string       `json:"date"`
	Description string       `json:"description"`
	VoucherType *VoucherType `json:"voucherType,omitempty"`
	Postings    Postings     `json:"postings"`
	Document    *Document    `json:"document,omitempty"`
	Attachment  *Attachment  `json:"attachment,omitempty"`
	EDIDocument *EDIDocument `json:"ediDocument,omitempty"`
}

func (v Voucher) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(v)
}

func (v Voucher) IsEmpty() bool {
	log.Fatal(zero.IsZero(v))
	return zero.IsZero(v)
}

type Orders []Order

type Order struct {
	// ID                 int      `json:"id"`
	// Version            int      `json:"version"`
	// URL                string   `json:"url"`
	Customer Customer `json:"customer"`
	// Contact            Contact  `json:"contact"`
	// Attn               Attn     `json:"attn"`
	ReceiverEmail      string  `json:"receiverEmail"`
	OverdueNoticeEmail string  `json:"overdueNoticeEmail"`
	Number             string  `json:"number"`
	Reference          string  `json:"reference"`
	OurContact         Contact `json:"ourContact,omitempty"`
	OurContactEmployee Contact `json:"ourContactEmployee,omitempty"`
	// Department         struct {
	// 	ID               int    `json:"id"`
	// 	Version          int    `json:"version"`
	// 	Name             string `json:"name"`
	// 	DepartmentNumber string `json:"departmentNumber"`
	// } `json:"department"`
	OrderDate string `json:"orderDate"`
	// Project                                     Project    `json:"project"`
	InvoiceComment string `json:"invoiceComment"`
	// Currency       Currency `json:"currency"`
	// InvoicesDueIn                               int        `json:"invoicesDueIn"`
	// InvoicesDueInType                           string     `json:"invoicesDueInType"`
	IsShowOpenPostsOnInvoices bool `json:"isShowOpenPostsOnInvoices"`
	// IsClosed                                    bool       `json:"isClosed"`
	DeliveryDate string `json:"deliveryDate"`
	// DeliveryAddress                             Address    `json:"deliveryAddress"`
	// DeliveryComment                             string     `json:"deliveryComment"`
	IsPrioritizeAmountsIncludingVat bool `json:"isPrioritizeAmountsIncludingVat"`
	// OrderLineSorting                            string     `json:"orderLineSorting"`
	OrderLines OrderLines `json:"orderLines"`
	// IsSubscription                              bool       `json:"isSubscription"`
	// SubscriptionDuration                        int        `json:"subscriptionDuration"`
	// SubscriptionDurationType                    string     `json:"subscriptionDurationType"`
	// SubscriptionPeriodsOnInvoice                int        `json:"subscriptionPeriodsOnInvoice"`
	// SubscriptionInvoicingTimeInAdvanceOrArrears string     `json:"subscriptionInvoicingTimeInAdvanceOrArrears"`
	// SubscriptionInvoicingTime                   int        `json:"subscriptionInvoicingTime"`
	// SubscriptionInvoicingTimeType               string     `json:"subscriptionInvoicingTimeType"`
	// IsSubscriptionAutoInvoicing                 bool       `json:"isSubscriptionAutoInvoicing"`
}

func (o Order) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(o)
}

func (o Order) IsEmpty() bool {
	return zero.IsZero(o)
}

type Contacts []Contact

type Contact struct {
	ID                       int    `json:"id"`
	Version                  int    `json:"version"`
	URL                      string `json:"url"`
	FirstName                string `json:"firstName"`
	LastName                 string `json:"lastName"`
	Email                    string `json:"email"`
	PhoneNumberMobileCountry struct {
		ID      int    `json:"id"`
		Version int    `json:"version"`
		URL     string `json:"url"`
	} `json:"phoneNumberMobileCountry"`
	PhoneNumberMobile string   `json:"phoneNumberMobile"`
	PhoneNumberWork   string   `json:"phoneNumberWork"`
	Customer          Customer `json:"customer,omitempty"`
}

func (c Contact) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(c)
}

func (c Contact) IsEmpty() bool {
	return zero.IsZero(c)
}

type Attn struct {
	ID                       int    `json:"id"`
	Version                  int    `json:"version"`
	URL                      string `json:"url"`
	FirstName                string `json:"firstName"`
	LastName                 string `json:"lastName"`
	Email                    string `json:"email"`
	PhoneNumberMobileCountry struct {
		ID      int `json:"id"`
		Version int `json:"version"`
	} `json:"phoneNumberMobileCountry"`
	PhoneNumberMobile string   `json:"phoneNumberMobile"`
	PhoneNumberWork   string   `json:"phoneNumberWork"`
	Customer          Customer `json:"customer"`
}

type Address struct {
	ID           int      `json:"id"`
	Version      int      `json:"version"`
	URL          string   `json:"url"`
	Employee     Employee `json:"employee,omitempty"`
	AddressLine1 string   `json:"addressLine1"`
	AddressLine2 string   `json:"addressLine2"`
	PostalCode   string   `json:"postalCode"`
	City         string   `json:"city"`
	Country      struct {
		ID            int    `json:"id"`
		Version       int    `json:"version"`
		ISOAlpha2Code string `json:"isoAlpha2Code"`
		ISOAlpha3Code string `json:"isoAlpha3Code"`
	} `json:"country"`
	Name string `json:"name"`
}

func (a Address) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(a)
}

func (a Address) IsEmpty() bool {
	return zero.IsZero(a)
}

type OrderLines []OrderLine

type OrderLine struct {
	ID int `json:"id"`
	// Version int     `json:"version"`
	// URL     string  `json:"url"`
	Product Product `json:"product"`
	// Inventory struct {
	// 	ID              int    `json:"id"`
	// 	Version         int    `json:"version"`
	// 	Name            string `json:"name"`
	// 	Number          string `json:"number"`
	// 	IsMainInventory bool   `json:"isMainInventory"`
	// 	IsInactive      bool   `json:"isInactive"`
	// } `json:"inventory"`
	Description                   string  `json:"description"`
	Count                         int     `json:"count"`
	UnitCostCurrency              int     `json:"unitCostCurrency,omitempty"`
	UnitPriceExcludingVATCurrency float64 `json:"unitPriceExcludingVatCurrency,omitempty"`
	// Currency                      Currency `json:"currency"`
	// Markup                        int     `json:"markup"`
	// Discount                      int     `json:"discount"`
	VATType                       VATType `json:"vatType"`
	UnitPriceIncludingVATCurrency float64 `json:"unitPriceIncludingVatCurrency,omitempty"`
	AmountExcludingVATCurrency    float64 `json:"amountExcludingVatCurrency,omitempty"`
	AmountIncludingVATCurrency    float64 `json:"amountIncludingVatCurrency"`
	// IsSubscription                bool    `json:"isSubscription"`
	// SubscriptionPeriodStart       string  `json:"subscriptionPeriodStart"`
	// SubscriptionPeriodEnd         string  `json:"subscriptionPeriodEnd"`
	// OrderGroup                    struct {
	// 	ID        int    `json:"id"`
	// 	Version   int    `json:"version"`
	// 	Title     string `json:"title"`
	// 	Comment   string `json:"comment"`
	// 	SortIndex int    `json:"sortIndex"`
	// } `json:"orderGroup"`
}

type AccountManager struct {
	ID                       int    `json:"id"`
	Version                  int    `json:"version"`
	URL                      string `json:"url"`
	FirstName                string `json:"firstName"`
	LastName                 string `json:"lastName"`
	EmployeeNumber           string `json:"employeeNumber"`
	DateOfBirth              string `json:"dateOfBirth"`
	Email                    string `json:"email"`
	PhoneNumberMobileCountry struct {
		ID      int `json:"id"`
		Version int `json:"version"`
	} `json:"phoneNumberMobileCountry"`
	PhoneNumberMobile      string `json:"phoneNumberMobile"`
	PhoneNumberHome        string `json:"phoneNumberHome"`
	PhoneNumberWork        string `json:"phoneNumberWork"`
	NationalIdentityNumber string `json:"nationalIdentityNumber"`
	Dnumber                string `json:"dnumber"`
	InternationalID        struct {
		IntAmeldingType string `json:"intAmeldingType"`
		Country         struct {
			ID      int `json:"id"`
			Version int `json:"version"`
		} `json:"country"`
		Number string `json:"number"`
	} `json:"internationalId"`
	BankAccountNumber     string     `json:"bankAccountNumber"`
	Iban                  string     `json:"iban"`
	Bic                   string     `json:"bic"`
	CreditorBankCountryID int        `json:"creditorBankCountryId"`
	UsesAbroadPayment     bool       `json:"usesAbroadPayment"`
	UserType              string     `json:"userType"`
	Comments              string     `json:"comments"`
	Address               Address    `json:"address"`
	Department            Department `json:"department"`
	Employments           []struct {
		ID           int    `json:"id"`
		Version      int    `json:"version"`
		EmploymentID string `json:"employmentId"`
		StartDate    string `json:"startDate"`
		EndDate      string `json:"endDate"`
		Division     struct {
			ID                 int    `json:"id"`
			Version            int    `json:"version"`
			Name               string `json:"name"`
			StartDate          string `json:"startDate"`
			EndDate            string `json:"endDate"`
			OrganizationNumber string `json:"organizationNumber"`
			Municipality       struct {
				ID      int `json:"id"`
				Version int `json:"version"`
			} `json:"municipality"`
		} `json:"division"`
		LastSalaryChangeDate     string `json:"lastSalaryChangeDate"`
		NoEmploymentRelationship bool   `json:"noEmploymentRelationship"`
		IsMainEmployer           bool   `json:"isMainEmployer"`
		TaxDeductionCode         string `json:"taxDeductionCode"`
		EmploymentDetails        []struct {
			ID                 int    `json:"id"`
			Version            int    `json:"version"`
			Date               string `json:"date"`
			EmploymentType     string `json:"employmentType"`
			MaritimeEmployment struct {
				ShipRegister string `json:"shipRegister"`
				ShipType     string `json:"shipType"`
				TradeArea    string `json:"tradeArea"`
			} `json:"maritimeEmployment"`
			RemunerationType   string `json:"remunerationType"`
			WorkingHoursScheme string `json:"workingHoursScheme"`
			ShiftDurationHours int    `json:"shiftDurationHours"`
			OccupationCode     struct {
				ID      int    `json:"id"`
				Version int    `json:"version"`
				NameNO  string `json:"nameNO"`
				Code    string `json:"code"`
			} `json:"occupationCode"`
			PercentageOfFullTimeEquivalent int `json:"percentageOfFullTimeEquivalent"`
			AnnualSalary                   int `json:"annualSalary"`
			HourlyWage                     int `json:"hourlyWage"`
			PayrollTaxMunicipalityID       struct {
				ID      int `json:"id"`
				Version int `json:"version"`
			} `json:"payrollTaxMunicipalityId"`
		} `json:"employmentDetails"`
	} `json:"employments"`
	HolidayAllowanceEarned struct {
		Year                   int     `json:"year"`
		Amount                 float64 `json:"amount"`
		Basis                  float64 `json:"basis"`
		AmountExtraHolidayWeek float64 `json:"amountExtraHolidayWeek"`
	} `json:"holidayAllowanceEarned"`
}

type Category struct {
	ID          int    `json:"id"`
	Version     int    `json:"version"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Number      string `json:"number"`
	Description string `json:"description"`
	Type        int    `json:"type"`
}
