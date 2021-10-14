package yookassa

import (
	"time"
)

type basePayment struct {
	// Сумма платежа.
	// Иногда партнеры ЮKassa берут с пользователя дополнительную комиссию, которая не входит в эту сумму.
	Amount Amount `json:"amount"`
	// Описание транзакции (не более 128 символов), которое вы увидите в личном кабинете ЮKassa,
	// а пользователь — при оплате. Например: «Оплата заказа № 72 для user@yoomoney.ru».
	Description string `json:"description"`
	// Получатель платежа.
	Recipient struct {
		// Идентификатор магазина в ЮKassa.
		AccountId string `json:"account_id"`
		// Идентификатор субаккаунта.
		// Используется для разделения потоков платежей в рамках одного аккаунта.
		GatewayId string `json:"gateway_id"`
	} `json:"recipient"`
	// Данные о сделке, в составе которой проходит платеж.
	// Присутствует, если вы проводите Безопасную сделку.
	Deal struct {
		// Идентификатор сделки.
		ID string `json:"id"`
		// Данные о распределении денег.
		Settlements []struct {
			// Тип операции. Фиксированное значение: payout — выплата продавцу.
			Type string `json:"type"`
			// Сумма вознаграждения продавца.
			Amount Amount `json:"amount"`
		} `json:"settlements"`
	} `json:"deal"`
	// Идентификатор покупателя в вашей системе, например электронная почта или номер телефона.
	// Не более 200 символов.
	// Присутствует, если вы хотите запомнить банковскую карту и отобразить ее при повторном платеже в виджете ЮKassa.
	MerchantCustomerId string `json:"merchant_customer_id"`
	// Данные о распределении денег — сколько и в какой магазин нужно перевести.
	// Присутствует, если вы используете Сплитование платежей.
	Transfers []struct {
		// Идентификатор магазина, в пользу которого вы принимаете оплату.
		// Выдается ЮKassa, отображается в разделе Продавцы личного кабинета (столбец shopId).
		AccountID string `json:"account_id"`
		// Сумма, которую необходимо перечислить магазину.
		Amount Amount `json:"amount"`
		// Статус распределения денег между магазинами.
		Status Status `json:"status"`
		// Комиссия за проданные товары и услуги, которая удерживается с магазина в вашу пользу.
		PlatformFeeAmount Amount `json:"platform_fee_amount"`
		// Любые дополнительные данные, которые нужны вам для работы (например, номер заказа).
		// Передаются в виде набора пар «ключ-значение» и возвращаются в ответе от ЮKassa.
		// Ограничения: максимум 16 ключей, имя ключа не больше 32 символов,
		// значение ключа не больше 512 символов, тип данных — строка в формате UTF-8.
		Metadata interface{} `json:"metadata"`
	} `json:"transfers"`
}

type Payment struct {
	// Идентификатор платежа в ЮKassa.
	Id string `json:"id"`

	// В случае, если сервер пришлёт ошибку, это может быть записано в Type.
	Type string `json:"type"`

	// Сумма платежа.
	// Иногда партнеры ЮKassa берут с пользователя дополнительную комиссию, которая не входит в эту сумму.
	Amount Amount `json:"amount"`
	// Описание транзакции (не более 128 символов), которое вы увидите в личном кабинете ЮKassa,
	// а пользователь — при оплате. Например: «Оплата заказа № 72 для user@yoomoney.ru».
	Description string `json:"description"`
	// Получатель платежа.
	Recipient struct {
		// Идентификатор магазина в ЮKassa.
		AccountId string `json:"account_id"`
		// Идентификатор субаккаунта.
		// Используется для разделения потоков платежей в рамках одного аккаунта.
		GatewayId string `json:"gateway_id"`
	} `json:"recipient"`
	// Данные о сделке, в составе которой проходит платеж.
	// Присутствует, если вы проводите Безопасную сделку.
	Deal struct {
		// Идентификатор сделки.
		ID string `json:"id"`
		// Данные о распределении денег.
		Settlements []struct {
			// Тип операции. Фиксированное значение: payout — выплата продавцу.
			Type string `json:"type"`
			// Сумма вознаграждения продавца.
			Amount Amount `json:"amount"`
		} `json:"settlements"`
	} `json:"deal"`
	// Идентификатор покупателя в вашей системе, например электронная почта или номер телефона.
	// Не более 200 символов.
	// Присутствует, если вы хотите запомнить банковскую карту и отобразить ее при повторном платеже в виджете ЮKassa.
	MerchantCustomerId string `json:"merchant_customer_id"`
	// Данные о распределении денег — сколько и в какой магазин нужно перевести.
	// Присутствует, если вы используете Сплитование платежей.
	Transfers []struct {
		// Идентификатор магазина, в пользу которого вы принимаете оплату.
		// Выдается ЮKassa, отображается в разделе Продавцы личного кабинета (столбец shopId).
		AccountID string `json:"account_id"`
		// Сумма, которую необходимо перечислить магазину.
		Amount Amount `json:"amount"`
		// Статус распределения денег между магазинами.
		Status Status `json:"status"`
		// Комиссия за проданные товары и услуги, которая удерживается с магазина в вашу пользу.
		PlatformFeeAmount Amount `json:"platform_fee_amount"`
		// Любые дополнительные данные, которые нужны вам для работы (например, номер заказа).
		// Передаются в виде набора пар «ключ-значение» и возвращаются в ответе от ЮKassa.
		// Ограничения: максимум 16 ключей, имя ключа не больше 32 символов,
		// значение ключа не больше 512 символов, тип данных — строка в формате UTF-8.
		Metadata interface{} `json:"metadata"`
	} `json:"transfers"`
	//// Поля, содержащиеся как в Payment, так и в PaymentConfig.
	//basePayment

	// Статус платежа.
	// Подробнее про жизненный цикл платежа: https://yookassa.ru/developers/payments/payment-process#lifecycle.
	Status Status `json:"status"`
	// Сумма платежа, которую получит магазин, — значение amount
	// за вычетом комиссии ЮKassa.
	// Если вы партнер и для аутентификации запросов используете OAuth-токен,
	// запросите у магазина право на получение информации о комиссиях при платежах.
	IncomeAmount Amount `json:"income_amount"`
	// Способ оплаты, который был использован для этого платежа.
	PaymentMethod PaymentMethoder `json:"payment_method"`
	// Время подтверждения платежа. Указывается по UTC и передается в формате ISO 8601.
	CapturedAt time.Time `json:"captured_at"`
	// Время создания заказа. Указывается по UTC и передается в формате ISO 8601. Пример: 2017-11-03T11:52:31.827Z
	CreatedAt time.Time `json:"created_at"`
	// Время, до которого вы можете бесплатно отменить или подтвердить платеж.
	// В указанное время платеж в статусе waiting_for_capture
	// будет автоматически отменен.
	ExpiresAt time.Time `json:"expires_at"`
	// Выбранный способ подтверждения платежа.
	// Присутствует, когда платеж ожидает подтверждения от пользователя.
	// Подробнее о сценариях подтверждения: https://yookassa.ru/developers/payments/payment-process#user-confirmation
	Confirmation Confirmer `json:"confirmation"`
	// Признак тестовой операции.
	Test bool `json:"test"`
	// Сумма, которая вернулась пользователю.
	// Присутствует, если у этого платежа есть успешные возвраты.
	RefundedAmount Amount `json:"refunded_amount"`
	// Признак оплаты заказа.
	Paid bool `json:"paid"`
	// Возможность провести возврат по API.
	Refundable bool `json:"refundable"`
	// Статус доставки данных для чека в онлайн-кассу.
	// Присутствует, если вы используете решение ЮKassa для работы по 54-ФЗ.
	ReceiptRegistration Status `json:"receipt_registration"`
	// Любые дополнительные данные, которые нужны вам для работы (например, номер заказа).
	// Передаются в виде набора пар «ключ-значение» и возвращаются в ответе от ЮKassa.
	// Ограничения: максимум 16 ключей, имя ключа не больше 32 символов,
	// значение ключа не больше 512 символов, тип данных — строка в формате UTF-8.
	Metadata interface{} `json:"metadata"`
	// Комментарий к статусу canceled: кто отменил платеж и по какой причине.
	// Подробнее про неуспешные платежи:
	// https://yookassa.ru/developers/payments/declined-payments
	CancellationDetails struct {
		// Участник процесса платежа, который принял решение об отмене транзакции.
		// Подробнее про инициаторов отмены платежа:
		// https://yookassa.ru/developers/payments/declined-payments#cancellation-details-party
		Party string `json:"party"`
		// Причина отмены платежа.
		// Перечень и описание возможных значений:
		// https://yookassa.ru/developers/payments/declined-payments#cancellation-details-reason
		Reason string `json:"reason"`
	} `json:"cancellation_details"`
	// Данные об авторизации платежа.
	AuthorizationDetails struct {
		// Retrieval Reference Number — уникальный идентификатор транзакции в системе эмитента.
		// Используется при оплате банковской картой.
		RRN string `json:"rrn"`
		// Код авторизации банковской карты.
		// Выдается эмитентом и подтверждает проведение авторизации.
		AuthCode string `json:"auth_code"`
	} `json:"authorization_details"`
}

type Item struct {
	Description              string `json:"description,omitempty"`
	Quantity                 string `json:"quantity,omitempty"`
	Amount                   Amount `json:"amount,omitempty"`
	VatCode                  int    `json:"vat_code,omitempty"`
	PaymentSubject           string `json:"payment_subject,omitempty"`
	PaymentMode              string `json:"payment_mode,omitempty"`
	ProductCode              string `json:"product_code,omitempty"`
	CountryOfOriginCode      string `json:"country_of_origin_code,omitempty"`
	CustomsDeclarationNumber string `json:"customs_declaration_number,omitempty"`
	Excise                   string `json:"excise,omitempty"`
}

type PaymentConfig struct {
	// Сумма платежа.
	// Иногда партнеры ЮKassa берут с пользователя дополнительную комиссию, которая не входит в эту сумму.
	Amount Amount `json:"amount,omitempty"`
	// Описание транзакции (не более 128 символов), которое вы увидите в личном кабинете ЮKassa,
	// а пользователь — при оплате. Например: «Оплата заказа № 72 для user@yoomoney.ru».
	Description string `json:"description,omitempty"`
	// Получатель платежа.
	Recipient *struct {
		// Идентификатор магазина в ЮKassa.
		AccountId string `json:"account_id,omitempty"`
		// Идентификатор субаккаунта.
		// Используется для разделения потоков платежей в рамках одного аккаунта.
		GatewayId string `json:"gateway_id,omitempty"`
	} `json:"recipient,omitempty"`
	// Данные о сделке, в составе которой проходит платеж.
	// Присутствует, если вы проводите Безопасную сделку.
	Deal *struct {
		// Идентификатор сделки.
		ID string `json:"id,omitempty"`
		// Данные о распределении денег.
		Settlements []struct {
			// Тип операции. Фиксированное значение: payout — выплата продавцу.
			Type string `json:"type,omitempty"`
			// Сумма вознаграждения продавца.
			Amount Amount `json:"amount,omitempty"`
		} `json:"settlements,omitempty"`
	} `json:"deal,omitempty"`
	// Идентификатор покупателя в вашей системе, например электронная почта или номер телефона.
	// Не более 200 символов.
	// Присутствует, если вы хотите запомнить банковскую карту и отобразить ее при повторном платеже в виджете ЮKassa.
	MerchantCustomerId string `json:"merchant_customer_id,omitempty"`
	// Данные о распределении денег — сколько и в какой магазин нужно перевести.
	// Присутствует, если вы используете Сплитование платежей.
	Transfers []struct {
		// Идентификатор магазина, в пользу которого вы принимаете оплату.
		// Выдается ЮKassa, отображается в разделе Продавцы личного кабинета (столбец shopId).
		AccountID string `json:"account_id,omitempty"`
		// Сумма, которую необходимо перечислить магазину.
		Amount Amount `json:"amount,omitempty"`
		// Статус распределения денег между магазинами.
		Status Status `json:"status,omitempty"`
		// Комиссия за проданные товары и услуги, которая удерживается с магазина в вашу пользу.
		PlatformFeeAmount Amount `json:"platform_fee_amount,omitempty"`
		// Любые дополнительные данные, которые нужны вам для работы (например, номер заказа).
		// Передаются в виде набора пар «ключ-значение» и возвращаются в ответе от ЮKassa.
		// Ограничения: максимум 16 ключей, имя ключа не больше 32 символов,
		// значение ключа не больше 512 символов, тип данных — строка в формате UTF-8.
		Metadata interface{} `json:"metadata,omitempty"`
	} `json:"transfers,omitempty"`
	// basePayment

	// Данные для формирования чека в онлайн-кассе (для соблюдения 54-ФЗ ).
	// Необходимо передавать, если вы отправляете данные для формирования чеков по одному из сценариев:
	// Платеж и чек одновременно или Сначала чек, потом платеж.
	Receipt *struct {
		Customer *struct {
			FullName string `json:"full_name,omitempty"`
			INN      string `json:"inn,omitempty"`
			Email    string `json:"email,omitempty"`
			Phone    string `json:"phone,omitempty"`
		} `json:"customer,omitempty"`
		Items         []Item `json:"items"`
		TaxSystemCode int    `json:"tax_system_code,omitempty"`
		Phone         string `json:"phone,omitempty"`
		Email         string `json:"email,omitempty"`
	} `json:"receipt,omitempty"`
	// Одноразовый токен для проведения оплаты, сформированный с помощью веб или мобильного SDK.
	PaymentToken string `json:"payment_token,omitempty"`
	// Идентификатор сохраненного способа оплаты.
	PaymentMethodID string `json:"payment_method_id,omitempty"`
	// Данные для оплаты конкретным способом.
	// Вы можете не передавать этот объект в запросе.
	// В этом случае пользователь будет выбирать способ оплаты на стороне ЮKassa.
	PaymentMethodData PaymentMethoder `json:"payment_method_data,omitempty"`
	// Данные, необходимые для инициирования выбранного сценария подтверждения платежа пользователем.
	Confirmation Confirmer `json:"confirmation,omitempty"`
	// Сохранение платежных данных (с их помощью можно проводить повторные безакцептные списания.
	// Значение true инициирует создание многоразового payment_method.
	SavePaymentMethod bool `json:"save_payment_method,omitempty"`
	// Автоматический прием поступившего платежа.
	Capture bool `json:"capture,omitempty"`
	// IPv4 или IPv6-адрес пользователя. Если не указан, используется IP-адрес TCP-подключения.
	ClientIP string `json:"client_ip,omitempty"`
	// Любые дополнительные данные, которые нужны вам для работы (например, номер заказа).
	// Передаются в виде набора пар «ключ-значение» и возвращаются в ответе от ЮKassa.
	// Ограничения: максимум 16 ключей, имя ключа не больше 32 символов,
	// значение ключа не больше 512 символов, тип данных — строка в формате UTF-8.
	Metadata interface{} `json:"metadata,omitempty"`
	//TODO: add airlines
	Airline interface{} `json:"airline,omitempty"`
}

func NewPaymentConfig(amount Amount, confirmation Confirmer) *PaymentConfig {
	config := PaymentConfig{Amount: amount, Confirmation: confirmation}
	return &config
}
