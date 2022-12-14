/*
 * API whatsgate.ru
 *
 * Интерфейс для взаимодействия с клиентом Whatsapp
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseChatGroupMetadataParticipants struct {
	// Идентификатор контакта в формате WhatsApp
	Id string `json:"id,omitempty"`
	// Является ли участник администратором
	IsAdmin bool `json:"isAdmin,omitempty"`
	// Является ли участник суперадминистратором
	IsSuperAdmin bool `json:"isSuperAdmin,omitempty"`
}
