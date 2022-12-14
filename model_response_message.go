/*
 * API whatsgate.ru
 *
 * Интерфейс для взаимодействия с клиентом Whatsapp
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseMessage struct {
	// Идентификатор сообщения
	Id string `json:"id,omitempty"`
	// Флаг просмотра сообщения
	Ack int32 `json:"ack,omitempty"`
	// Флаг, имеет ли сообщение медиафайл
	HasMedia bool `json:"hasMedia,omitempty"`
	// Ключ медиафайла (при наличии)
	MediaKey string `json:"mediaKey,omitempty"`
	// Текст сообщения
	Body string `json:"body,omitempty"`
	// Тип сообщения
	Type_ string `json:"type,omitempty"`
	// Время сообщения в формате Unix Timestamp
	Timestamp int32 `json:"timestamp,omitempty"`
	// Идентификатор отправителя в формате WhatsApp
	From string `json:"from,omitempty"`
	// Идентификатор получателя в формате WhatsApp
	To string `json:"to,omitempty"`
	// Флаг, было ли сообщение перенаправлено
	IsForwarded bool `json:"isForwarded,omitempty"`
}
