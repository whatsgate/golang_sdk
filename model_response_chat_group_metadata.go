/*
 * API whatsgate.ru
 *
 * Интерфейс для взаимодействия с клиентом Whatsapp
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Метаданные группы, если чат является группой
type ResponseChatGroupMetadata struct {
	// Идентификатор группы в формате WhatsApp
	Id string `json:"id,omitempty"`
	// Дата создания группы в формате Unix Time Stamp
	Creation int32 `json:"creation,omitempty"`
	// Идентификатор создателя группы в формате WhatsApp
	Owner string `json:"owner,omitempty"`
	// Количество участников в группе
	Size int32 `json:"size,omitempty"`
	// Массив участников
	Participants []ResponseChatGroupMetadataParticipants `json:"participants,omitempty"`
}
