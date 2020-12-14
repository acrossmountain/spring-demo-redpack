package apis

import (
	"github.com/go-spring/spring-boot"
)

func init() {

	v1 := SpringBoot.Route("/v1")

	SpringBoot.RegisterBean(new(Account)).Init(func(a *Account) {
		account := v1.Route("/account")
		account.PostMapping("/create", a.Create)
		account.PostMapping("/transfer", a.Transfer)
		account.GetMapping("/envelope/:userId", a.Envelope)
		account.GetMapping("/:accountNo", a.Account)
	})
}
