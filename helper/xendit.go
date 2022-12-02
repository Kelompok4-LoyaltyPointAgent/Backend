package helper

import (
	"github.com/kelompok4-loyaltypointagent/backend/models"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
	"github.com/xendit/xendit-go/payout"
)

//Fitur mengirim uang
func CreatePayoutXendit(transaction models.Transaction) (xendit.Payout, error) {
	xendit.Opt.SecretKey = "xnd_development_375ruRhRvXqSLVTgujv35QbKjxZ14H09PNVFwhwt1bPNdTtwfbaseyL68JAjyzk"

	payoutParams := payout.CreateParams{
		ExternalID: transaction.ID.String(),
		Amount:     transaction.Amount,
		Email:      transaction.Email,
	}

	resp, err := payout.Create(&payoutParams)
	if err != nil {
		return xendit.Payout{}, err
	}

	return *resp, nil

}

func CreateInvoiceXendit(transaction models.Transaction, user models.User) (xendit.Invoice, error) {
	xendit.Opt.SecretKey = "xnd_development_375ruRhRvXqSLVTgujv35QbKjxZ14H09PNVFwhwt1bPNdTtwfbaseyL68JAjyzk"

	customer := xendit.InvoiceCustomer{
		GivenNames:   user.Name,
		Email:        user.Email,
		MobileNumber: transaction.PhoneNumber,
	}

	data := invoice.CreateParams{
		ExternalID: transaction.ID.String(),
		PayerEmail: transaction.Email,
		Amount:     transaction.Amount,
		Customer:   customer,
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		return xendit.Invoice{}, err
	}

	return *resp, nil
}
