package dto

// swagger:model GetApplicationRequestDTO
type GetApplicationRequestDTO struct {
	// ApplicationID of the application's id
	// in: string
	ApplicationID string
}

// swagger:model GetApplicationResponseDTO
type GetApplicationResponseDTO struct {
	// Name of the application
	// in: string
	Name string
	// Code of the application
	// in: string
	Code string
}
