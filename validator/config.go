package validator

import "github.com/gofiber/fiber/v2"

type Translator struct {
	Tag     string
	Message string
}

type ValidatorError struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Message     string `json:"message"`
}

type Validators struct {
	Ctx            *fiber.Ctx
	Data           interface{}
	Error          bool
	ValidationsErr []ValidatorError
}
