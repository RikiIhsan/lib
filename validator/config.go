package validator

import "github.com/gofiber/fiber/v2"

type Translator struct {
	Tag     string
	Message string
}

type ValidatorError struct {
	FailedField string
	Tag         string
	Message     string
}

type Validators struct {
	Ctx            *fiber.Ctx
	Data           interface{}
	Error          bool
	ValidationsErr []ValidatorError
}
