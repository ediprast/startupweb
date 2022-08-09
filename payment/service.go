package payment

import (
	"bwastartup/user"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type service struct {
}

type Service interface {
	GetPaymentUrl(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentUrl(transaction Transaction, user user.User) (string, error) {
	// var snapClient snap.Client

	midtrans.ServerKey = "SB-Mid-server-HsGRTdQ3_19-Y33RUyHxi_Bt"
	midtrans.Environment = midtrans.Sandbox

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.Id),
			GrossAmt: int64(transaction.Amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
		},
	}

	// 3. Request create Snap transaction to Midtrans
	snapResp, err := snap.CreateTransaction(req)

	if err != nil {
		return "", err
	}

	return snapResp.RedirectURL, nil

}
