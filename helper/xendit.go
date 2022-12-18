package helper

import (
	"log"

	"github.com/kelompok4-loyaltypointagent/backend/config"
	"github.com/kelompok4-loyaltypointagent/backend/models"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/balance"
	"github.com/xendit/xendit-go/disbursement"
	"github.com/xendit/xendit-go/invoice"
	"github.com/xendit/xendit-go/payout"
)

//Fitur mengirim uang
func CreatePayoutXendit(transaction models.Transaction) (xendit.Payout, error) {
	xenditConfig := config.LoadXenditConfig()
	xendit.Opt.SecretKey = xenditConfig.Secret

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
	xenditConfig := config.LoadXenditConfig()
	xendit.Opt.SecretKey = xenditConfig.Secret

	customer := xendit.InvoiceCustomer{
		GivenNames:   user.Name,
		Email:        user.Email,
		MobileNumber: transactionDetail.Number,
	}

	data := invoice.CreateParams{
		ExternalID:      transaction.ID.String(),
		PayerEmail:      transactionDetail.Email,
		Amount:          transaction.Amount,
		Customer:        customer,
		Description:     "Invoice from Halo Pulsa",
		InvoiceDuration: 3600,
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		return xendit.Invoice{}, err
	}

	return *resp, nil
}

func CreateDisbursementXendit(transaction models.Transaction, transactionDetail models.TransactionDetail, user models.User) (xendit.Disbursement, error) {
	xenditConfig := config.LoadXenditConfig()
	xendit.Opt.SecretKey = xenditConfig.Secret

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

func GetBalance() (*float64, error) {
	xenditConfig := config.LoadXenditConfig()
	xendit.Opt.SecretKey = xenditConfig.Secret

	getData := balance.GetParams{
		AccountType: "CASH",
	}

	resp, err := balance.Get(&getData)
	if err != nil {
		log.Fatal(err)
	}
	return &resp.Balance, nil

}
