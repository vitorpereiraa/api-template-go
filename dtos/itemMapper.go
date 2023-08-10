package dtos

import "github.com/vitorpereira/api-template-go/domain"

func ItemToDto(item domain.Item) ItemDTO {
	return ItemDTO{
		Description: item.Description(),
		DueDate:     item.DueDate(),
	}
}

func ItemListToDtoList(items []domain.Item) (itemsDto []ItemDTO) {
	for _, item := range items {
		itemsDto = append(itemsDto, ItemToDto(item))
	}
	return
}
