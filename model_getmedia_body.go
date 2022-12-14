/*
 * API whatsgate.ru
 *
 * Интерфейс для взаимодействия с клиентом Whatsapp
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type GetmediaBody struct {
	// Идентификатор Whatsapp ID
	WhatsappID string `json:"WhatsappID,omitempty"`
	// Идентификатор медиафайла, содержится во входящих сообщениях.
	MediaKey string `json:"mediaKey,omitempty"`
}
