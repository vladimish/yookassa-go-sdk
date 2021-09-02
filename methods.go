package yookassa

type PaymentMethodType string

const ( // Payment type.
	PaymentTypeAlfabank      PaymentMethodType = "alfabank"
	PaymentTypeMobileBalance PaymentMethodType = "mobile_balance"
	PaymentTypeBankCard      PaymentMethodType = "bank_card"
	PaymentTypeInstallments  PaymentMethodType = "installments"
	PaymentTypeCash          PaymentMethodType = "cash"
	PaymentTypeB2BSberbank   PaymentMethodType = "b2b_sberbank"
	PaymentTypeSberbank      PaymentMethodType = "sberbank"
	PaymentTypeTinkoffBank   PaymentMethodType = "tinkoff_bank"
	PaymentTypeYooMoney      PaymentMethodType = "yoo_money"
	PaymentTypeApplePay      PaymentMethodType = "apple_pay"
	PaymentTypeGooglePay     PaymentMethodType = "google_pay"
	PaymentTypeQiwi          PaymentMethodType = "qiwi"
	PaymentTypeWeChat        PaymentMethodType = "wechat"
	PaymentTypeWebmoney      PaymentMethodType = "webmoney"
)

type PaymentMethoder interface {
}

type basePaymentMethod struct {
	// Код способа оплаты.
	Type PaymentMethodType `json:"type"`
	// Идентификатор способа оплаты.
	ID string `json:"id"`
	// С помощью сохраненного способа оплаты можно проводить безакцептные списания.
	Saved bool `json:"saved"`
	// Название способа оплаты.
	Title string `json:"title"`
}

type Alfabank struct {
	basePaymentMethod
	// Логин пользователя в Альфа-Клике (привязанный телефон или дополнительный логин).
	Login string `json:"login"`
}

type MobileBalance struct {
	basePaymentMethod
	// Телефон, с баланса которого осуществляется платеж. Указывается в формате ITU-T E.164, например 79000000000.
	Phone string `json:"phone"`
}

type BankCard struct {
	basePaymentMethod
	// Данные банковской карты.
	Card Card `json:"card"`
}

type Installments struct {
	basePaymentMethod
	// Телефон, с баланса которого осуществляется платеж. Указывается в формате ITU-T E.164, например 79000000000.
	// Поле можно оставить пустым: пользователь сможет заполнить его при оплате на стороне ЮKassa.
	Phone string `json:"phone"`
}

type Cash struct {
	basePaymentMethod
}

type B2BSberbank struct {
	basePaymentMethod
	PayerBankDetails struct {
		// Полное наименование организации.
		FullName string `json:"full_name"`
		// Сокращенное наименование организации.
		ShortName string `json:"short_name"`
		// Адрес организации.
		Address string `json:"address"`
		// Индивидуальный налоговый номер (ИНН) организации.
		INN string `json:"inn"`
		// Код причины постановки на учет (КПП) организации.
		KPP string `json:"kpp"`
		// Наименование банка организации.
		BankName string `json:"bank_name"`
		// Отделение банка организации.
		BankBranch string `json:"bank_branch"`
		// Банковский идентификационный код (БИК) банка организации.
		BankBIK string `json:"bank_bik"`
		// Номер счета организации.
		Account string `json:"account"`
	}
	// Назначение платежа (не больше 210 символов).
	PaymentPurpose string `json:"payment_purpose"`
	// Данные о налоге на добавленную стоимость (НДС). Платеж может облагаться и не облагаться НДС.
	// Товары могут облагаться по одной ставке НДС или по разным.
	VATData struct {
		// Код способа расчета НДС.
		Type string `json:"type"`
		// Сумма НДС.
		Amount Amount `json:"amount"`
		// Налоговая ставка (в процентах).
		Rate string `json:"rate"`
	} `json:"vat_data"`
}

type Sberbank struct {
	basePaymentMethod
	// Телефон пользователя, на который зарегистрирован аккаунт в СберБанк Онлайн/SberPay.
	// Указывается в формате ITU-T E.164, например 79000000000.
	Phone string `json:"phone"`
}

type TinkoffBank struct {
	basePaymentMethod
}

type YooMoney struct {
	basePaymentMethod
	// Номер кошелька ЮMoney, из которого заплатил пользователь.
	AccountNumber string `json:"account_number"`
}

type ApplePay struct {
	basePaymentMethod
	// Содержимое поля paymentData из объекта PKPaymentToken (платежный токен Apple Pay). Приходит в формате Base64.
	PaymentData string `json:"payment_data"`
}

type GooglePay struct {
	basePaymentMethod
	// Требуется криптограмма Payment Token Cryptography для проведения оплаты через Google Pay.
	// Чтобы ее получить, необходимо вызвать метод PaymentData.getPaymentMethodToken().getToken()
	// в мобильном приложении на устройстве пользователя.
	PaymentMethodToken string `json:"payment_method_token"`
}

type Qiwi struct {
	basePaymentMethod
	// Телефон, на который зарегистрирован аккаунт в QIWI. Указывается в формате ITU-T E.164, например 79000000000.
	Phone string `json:"phone"`
}

type WeChat struct {
	basePaymentMethod
}

type WebMoney struct {
	basePaymentMethod
}
