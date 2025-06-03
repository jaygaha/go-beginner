package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/jaygaha/go-beginner/cmd/20_web_frameworks/beego/controllers"
)

func init() {
	beego.Router("/", &controllers.TicketController{}, "get:Get")

	// APIs
	beego.BConfig.CopyRequestBody = true
	beego.Router("/tickets", &controllers.TicketController{}, "get:GetAllTickets;post:CreateTicket")
	beego.Router("/tickets/:id", &controllers.TicketController{}, "get:GetTicket;put:UpdateTicket;delete:DeleteTicket")
}
