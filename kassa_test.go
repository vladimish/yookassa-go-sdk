package yookassa

import (
	"testing"
	yookassa "yookassa/models"
)

// Для тестирования введите ваш идентификатор магазина и приватный ключ.
var shopid, key = "", ""

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

// TestPaymentCreation проверяет создание платежа.
func TestPaymentCreation(t *testing.T) {
	k := NewKassa(shopid, key)
	tAmount := yookassa.Amount{
		Value:    "10.00",
		Currency: "RUB",
	}
	config := yookassa.NewPaymentConfig(
		tAmount,
		yookassa.Redirect{
			Type:      yookassa.TypeRedirect,
			Locale:    "ru_RU",
			Enforce:   true,
			ReturnURL: "https://t.me/ranh_ranepa_bot",
		})

	payment, err := k.SendPaymentConfig(config)
	if err != nil {
		t.Error(err)
	}

	if payment.Amount != tAmount{
		t.Error("Got wrong response.")
	}
}
