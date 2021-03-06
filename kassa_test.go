package yookassa

import (
	"testing"
)

// Для тестирования введите ваш идентификатор магазина и приватный ключ.
var shopid, key = "830911", "test_RYPy6ussTGgAD_MNaFDeGTV6G4wyqWW34H8_5ufNx7g"

// TestKassa_Ping проверяет работоспособность функции Ping.
func TestKassa_Ping(t *testing.T) {
	k := NewKassa(shopid, key)
	ok, err := k.Ping()
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("Can't connect to yookassa")
	}
}

// TestKassa_SendPaymentConfig проверяет создание платежа.
func TestKassa_SendPaymentConfig(t *testing.T) {
	k := NewKassa(shopid, key)
	tAmount := Amount{
		Value:    "10.00",
		Currency: "RUB",
	}
	config := NewPaymentConfig(
		tAmount,
		Redirect{
			Type:      TypeRedirect,
			Locale:    "ru_RU",
			Enforce:   true,
			ReturnURL: "https://t.me/ranh_ranepa_bot",
		})

	payment, err := k.SendPaymentConfig(config)
	if err != nil {
		t.Error(err)
	}

	if payment.Amount != tAmount {
		t.Error("Got wrong response.")
	}
}

func TestKassa_GetPaymentInfo(t *testing.T) {
	k := NewKassa(shopid, key)
	tAmount := Amount{
		Value:    "10.00",
		Currency: "RUB",
	}
	config := NewPaymentConfig(
		tAmount,
		Redirect{
			Type:      TypeRedirect,
			Locale:    "ru_RU",
			Enforce:   true,
			ReturnURL: "https://t.me/ranh_ranepa_bot",
		})

	payment, err := k.SendPaymentConfig(config)
	if err != nil {
		t.Error(err)
	}

	id := payment.Id

	payment2, err := k.GetPayment(id)

	t.Log(payment)
	t.Log(payment2)

	if payment.Id != payment2.Id {
		t.Error("Got wrong response.")
	}
}

func TestKassa_GetPaymentErrorsControl(t *testing.T) {
	k := NewKassa(shopid, key)
	config := NewPaymentConfig(
		Amount{
			Value: "-5",
			Currency: "XMR",
		},
		Embedded{})

	_, err := k.SendPaymentConfig(config)
	if err == nil {
		t.Error("Error is not handled.")
	}
}
