package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"microservices_demo/service_payment/internal/api/v1"
	"strconv"
	"time"
)

func (ps *PaymentService) Charge(ctx context.Context, request *v1.ChargeRequest) (*v1.ChargeResponse, error) {
	resp := &v1.ChargeResponse{}
	amount := request.GetAmount()
	creditCard := request.GetCreditCard()

	cardNumber := creditCard.GetCreditCardNumber()
	cardType, valid := ps.payment.CardValidator(cardNumber)

	if !valid {
		return resp, status.Errorf(codes.InvalidArgument, "invalid credit card")

	}

	if !(cardType == "visa" || cardType == "mastercard") {
		return resp, errors.Errorf("Sorry, we cannot process %v credit cards. Only VISA or MasterCard is accepted.", cardType)
	}
	currentMonth, _ := strconv.Atoi(time.Now().Format("01"))

	currentYear := time.Now().Year()

	year := creditCard.GetCreditCardExpirationYear()
	month := creditCard.GetCreditCardExpirationMonth()

	if ((currentYear * 12) + currentMonth) > int(year*12+month) {
		return resp, errors.Errorf("Your credit card (ending %v) expired on %d/%d", cardNumber, month, year)
	}

	ps.logger.Sugar().Infof("Transaction processed :%v ending %v Amount: %v%d.%d", cardType, cardNumber, amount.CurrencyCode, amount.Units, amount.Nanos)
	newUUID, _ := uuid.NewUUID()
	return &v1.ChargeResponse{
		TransactionId: newUUID.String(),
	}, nil
}
