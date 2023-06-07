# Обёртка для API ЮКассы на go.
Задача этого проекта - просто предоставить оболочку API без каких-либо дополнительных функций.

SDK находится в стадии разработки.

## Пример
Простая программа для создания платежа и проверки его статуса.
```go
package main

import (
"github.com/vladimish/yookassa-go-sdk"
"log"
"time"
)

func main() {
// Создание нового экземпляра Kassa, через который будут осуществляться запросы.
kassa := yookassa.NewKassa("ИДЕНТИФИКАТОР_МАГАЗИНА", "ПРИВАТНЫЙ_КЛЮЧ")

	// Создание и настройка платежа.
	config := yookassa.NewPaymentConfig(
		// Размер платежа.
		yookassa.Amount{
			Value:    "10.00",
			Currency: "RUB",
		},
		// Способ подтверждения платежа пользователем.
		yookassa.Redirect{
			Type:            yookassa.TypeRedirect,
			Locale:          "ru_RU",
			ReturnURL:       "https://example.com",
		})
	config.Description = "Тестовый платёж"

	// Отправка конфига и получение экземпляра платежа в ответ.
	p, err := kassa.SendPaymentConfig(config)
	if err != nil{
		log.Fatal(err)
	}

	// Вывод ссылки для оплаты.
	log.Println(p.Confirmation.(map[string]interface{})["confirmation_url"])

	// Ожидание оплаты и опрос сервера на предмет её совершения.
	// Опрос заканчивается, если истекло время оплаты.
	for i := 0; i < 100; i++ {
		newp, err := kassa.GetPayment(p.Id)
		if err != nil{
			log.Fatal(err)
		}

		if newp.Paid == true{
			log.Println("Оплата прошла успешно.")
			return
		}else{
			log.Println("Пользователь производит оплату...")
			time.Sleep(time.Second * 10)
		}
	}
	log.Println("Оплата отменена.")
}
```
