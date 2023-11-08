package service

import (
	"mail-service/pkg/router"
	"mail-service/service/index"
	mail "mail-service/service/mail"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {
	// Set Endpoint for admin
	router.Router.Get(router.RouterBasePath+"/", index.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", index.GetHealth)
	router.Router.Mount(router.RouterBasePath+"/mails", mail.MailServiceSubRoute)
}
