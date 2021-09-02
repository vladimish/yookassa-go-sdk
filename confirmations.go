package yookassa

type ConfirmationType string

const (
	TypeEmbedded          ConfirmationType = "embedded"
	TypeExternal          ConfirmationType = "external"
	TypeMobileApplication ConfirmationType = "mobile_application"
	TypeQR                ConfirmationType = "qr"
	TypeRedirect          ConfirmationType = "redirect"
)

type Confirmer interface {
}

type Embedded struct {
	// Код сценария подтверждения.
	Type ConfirmationType `json:"type"`
	// Язык интерфейса, писем и смс, которые будет видеть или получать пользователь. Формат соответствует ISO/IEC 15897.
	// Возможные значения: ru_RU, en_US. Регистр важен.
	Locale string `json:"locale"`

	// Токен для инициализации платежного виджета ЮKassa.
	ConfirmationToken string `json:"confirmation_token"`
}

type External struct {
	// Код сценария подтверждения.
	Type ConfirmationType `json:"type"`
	// Язык интерфейса, писем и смс, которые будет видеть или получать пользователь. Формат соответствует ISO/IEC 15897.
	// Возможные значения: ru_RU, en_US. Регистр важен.
	Locale string `json:"locale"`
}

type MobileApplication struct {
	// Код сценария подтверждения.
	Type ConfirmationType `json:"type"`
	// Язык интерфейса, писем и смс, которые будет видеть или получать пользователь. Формат соответствует ISO/IEC 15897.
	// Возможные значения: ru_RU, en_US. Регистр важен.
	Locale string `json:"locale"`

	// Диплинк на мобильное приложение, в котором пользователь подтверждает платеж.
	ConfirmationURL string `json:"confirmation_url"`
}

type QR struct {
	// Код сценария подтверждения.
	Type ConfirmationType `json:"type"`
	// Язык интерфейса, писем и смс, которые будет видеть или получать пользователь. Формат соответствует ISO/IEC 15897.
	// Возможные значения: ru_RU, en_US. Регистр важен.
	Locale string `json:"locale"`

	// Данные для генерации QR-кода.
	ConfirmationData string `json:"confirmation_data"`
}

type Redirect struct {
	// Код сценария подтверждения.
	Type ConfirmationType `json:"type"`
	// Язык интерфейса, писем и смс, которые будет видеть или получать пользователь. Формат соответствует ISO/IEC 15897.
	// Возможные значения: ru_RU, en_US. Регистр важен.
	Locale string `json:"locale"`
	// URL, на который необходимо перенаправить пользователя для подтверждения оплаты.
	ConfirmationURL string `json:"confirmation_url"`
	// Запрос на проведение платежа с аутентификацией по 3-D Secure.
	// Будет работать, если оплату банковской картой вы по умолчанию принимаете без подтверждения платежа пользователем.
	// В остальных случаях аутентификацией по 3-D Secure будет управлять ЮKassa.
	// Если хотите принимать платежи без дополнительного подтверждения пользователем, напишите вашему менеджеру ЮKassa.
	Enforce bool `json:"enforce"`
	// URL, на который вернется пользователь после подтверждения или отмены платежа на веб-странице.
	ReturnURL string `json:"return_url"`
}