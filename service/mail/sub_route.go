package mail

import (
	"mail-service/service/mail/controller"
	"time"

	"github.com/go-chi/chi"
)

var MailServiceSubRoute = chi.NewRouter()

// Init package with sub-router for mails service
func init() {

	go func() {
		for {
			controller.SendMail()
			time.Sleep(1 * time.Minute)
		}
	}()

	// MailServiceSubRoute.Group(func(_ chi.Router) {
	// MailServiceSubRoute.Post("/sendMail", controller.SendMail)
	// })
}
