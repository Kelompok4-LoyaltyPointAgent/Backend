package helper

import (
	"github.com/kelompok4-loyaltypointagent/backend/models"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/disbursement"
	"github.com/xendit/xendit-go/invoice"
	"github.com/xendit/xendit-go/payout"
)

//Fitur mengirim uang
func CreatePayoutXendit(transaction models.Transaction) (xendit.Payout, error) {
	xendit.Opt.SecretKey = "xnd_development_375ruRhRvXqSLVTgujv35QbKjxZ14H09PNVFwhwt1bPNdTtwfbaseyL68JAjyzk"

	payoutParams := payout.CreateParams{
		ExternalID: transaction.ID.String(),
		Amount:     transaction.Amount,
	}

	resp, err := payout.Create(&payoutParams)
	if err != nil {
		return xendit.Payout{}, err
	}

	return *resp, nil

}

func CreateInvoiceXendit(transaction models.Transaction, transactionDetail models.TransactionDetail, user models.User) (xendit.Invoice, error) {
	xendit.Opt.SecretKey = "xnd_development_375ruRhRvXqSLVTgujv35QbKjxZ14H09PNVFwhwt1bPNdTtwfbaseyL68JAjyzk"

	customer := xendit.InvoiceCustomer{
		GivenNames:   user.Name,
		Email:        user.Email,
		MobileNumber: transactionDetail.Number,
	}

	data := invoice.CreateParams{
		ExternalID:  transaction.ID.String(),
		PayerEmail:  transactionDetail.Email,
		Amount:      transaction.Amount,
		Customer:    customer,
		Description: "Invoice from Halo Pulsa",
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		return xendit.Invoice{}, err
	}

	return *resp, nil
}

func CreateDisbursementXendit(transaction models.Transaction, transactionDetail models.TransactionDetail, user models.User) (xendit.Disbursement, error) {
	xendit.Opt.SecretKey = "xnd_development_375ruRhRvXqSLVTgujv35QbKjxZ14H09PNVFwhwt1bPNdTtwfbaseyL68JAjyzk"

	createDisbursementParams := disbursement.CreateParams{
		ExternalID:        transaction.ID.String(),
		Amount:            transaction.Amount,
		BankCode:          transaction.Method,
		AccountHolderName: user.Name,
		AccountNumber:     transactionDetail.Number,
		EmailTo:           []string{transactionDetail.Email},
		Description:       "Cashout from Halo Pulsa",
	}

	resp, err := disbursement.Create(&createDisbursementParams)
	if err != nil {
		return xendit.Disbursement{}, err
	}

	return *resp, nil
}


