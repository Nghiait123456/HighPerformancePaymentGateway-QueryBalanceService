package error_handle

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/high-performance-payment-gateway/balance-service/balance/pkg/external/sentry_fiber"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

type (
	PanicHandle struct {
		alert *AlertAc
		app   *fiber.App
	}

	AlertAc = Sentry
	Sentry  struct {
		Dsn              string
		Debug            bool
		AttachStacktrace bool
		Repanic          bool
		WaitForDelivery  bool
		Timeout          time.Duration
	}

	PanicHandleInterface interface {
		Init()
	}
)

func (p *PanicHandle) Init() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              p.alert.Dsn,
		Debug:            p.alert.Debug,
		AttachStacktrace: p.alert.AttachStacktrace,
	})
	if err != nil {
		errMessage := fmt.Sprintf("init Sentry error")
		log.WithFields(log.Fields{
			"errMessage": errMessage,
		}).Error(errMessage)

		panic(errMessage)
		os.Exit(0)
	}

	//resign middleware capture panic
	p.app.Use(sentry_fiber.New(sentry_fiber.Options{
		Repanic:         p.alert.Repanic,
		WaitForDelivery: p.alert.WaitForDelivery,
		Timeout:         p.alert.Timeout,
	}))
}

func NewPanicHandle(alert *AlertAc, app *fiber.App) PanicHandleInterface {
	return &PanicHandle{
		alert: alert,
		app:   app,
	}
}
