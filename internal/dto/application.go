package dto

type GetApplicationRequestDTO struct {
	ApplicationID string
}

type GetApplicationResponseDTO struct {
	Name string
	Code string
}
