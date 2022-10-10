package error_identification

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"strconv"
)

type (
	ErrorIdentification struct {
	}

	ErrorIdentificationInterface interface {
		createId() string
		addToContext(ctx *fiber.Ctx)
		ResignInMiddleware(c *fiber.Ctx) error // resign with middleware in frame
	}
)

const NameFields = "error_identification_debug"
const DefaultError = "-1"

func (e *ErrorIdentification) createId() string {
	return strconv.FormatUint(rand.Uint64(), 10)
}

func (e *ErrorIdentification) addToContext(ctx *fiber.Ctx) {
	ctx.Append(NameFields, e.createId())
}

func (e *ErrorIdentification) ResignInMiddleware(ctx *fiber.Ctx) error {
	e.addToContext(ctx)
	return nil
}

func GetErrIdFrCtx(ctx *fiber.Ctx) (string, error) {
	value := ctx.Get(NameFields, DefaultError)
	if value == DefaultError {
		return "", errors.New("missing ErrorIdentification")
	}

	return value, nil
}

func NewErrorIdentification() ErrorIdentificationInterface {
	return &ErrorIdentification{}
}
