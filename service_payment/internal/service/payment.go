package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"strconv"
	"time"

	v1 "microservices_demo/service_payment/api/v1"
)

func (ps *PaymentService) Charge(ctx context.Context, request *v1.ChargeRequest) (*v1.ChargeResponse, error) {
	amount := request.GetAmount()
	creditCard := request.GetCreditCard()

	cardNumber := creditCard.GetCreditCardNumber()
	cardType, valid := ps.payment.CardValidator(cardNumber)

	if !valid {
		return nil, errors.New("invalid credit card")
	}

	if !(cardType == "visa" || cardType == "mastercard") {
		return nil, errors.Errorf("Sorry, we cannot process %v credit cards. Only VISA or MasterCard is accepted.", cardType)
	}
	currentMonth, _ := strconv.Atoi(time.Now().Format("01"))

	currentYear := time.Now().Year()

	year := creditCard.GetCreditCardExpirationYear()
	month := creditCard.GetCreditCardExpirationMonth()

	if ((currentYear * 12) + currentMonth) > int(year*12+month) {
		return nil, errors.Errorf("Your credit card (ending %v) expired on %d/%d", cardNumber, month, year)
	}

	ps.logger.Sugar().Infof("Transaction processed :%v ending %v Amount: %v%d.%d", cardType, cardNumber, amount.CurrencyCode, amount.Units, amount.Nanos)
	newUUID, _ := uuid.NewUUID()
	return &v1.ChargeResponse{
		TransactionId: newUUID.String(),
	}, nil
}
