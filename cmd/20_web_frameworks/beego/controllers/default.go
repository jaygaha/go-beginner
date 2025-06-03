package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"slices"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/jaygaha/go-beginner/cmd/20_web_frameworks/beego/models"
)

// In-memory storage for tickets
var tickets = []models.Ticket{}
var nextId int = 1

type TicketController struct {
	beego.Controller
}

// Get renders the frontend index page
func (c *TicketController) Get() {
	c.TplName = "index.tpl"
}

// GetAllTickets  retrieves all tickets
func (c *TicketController) GetAllTickets() {
	c.Data["json"] = tickets
	c.ServeJSON()
}

// GetTicket retrieves a ticket by ID
func (c *TicketController) GetTicket() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid ID"}
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}

	for _, ticket := range tickets {
		if ticket.Id == id {
			c.Data["json"] = ticket
			c.ServeJSON()
			return
		}
	}

	c.Data["json"] = map[string]string{"error": "Ticket not found"}
	c.Ctx.Output.SetStatus(404)
	c.ServeJSON()
}

// CreateTicket creates a new ticket
func (c *TicketController) CreateTicket() {
	ticket := models.Ticket{}

	// Check if the body is empty
	if len(c.Ctx.Input.RequestBody) == 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Request body is empty"}
		c.ServeJSON()
		return
	}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ticket)
	if err != nil {
		fmt.Println("Error binding JSON:", err)
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}

	ticket.Id = nextId
	nextId++

	now := time.Now()
	ticket.CreatedAt = &now
	ticket.UpdatedAt = &now
	if ticket.CreatedBy == "" {
		ticket.CreatedBy = "anonymous" // Default value if not provided
	}

	tickets = append(tickets, ticket)
	c.Data["json"] = ticket
	c.Ctx.Output.SetStatus(201)
	c.ServeJSON()
}

// UpdateTicket updates an existing ticket
func (c *TicketController) UpdateTicket() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid ID"}
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}

	var updatedTicket models.Ticket
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &updatedTicket); err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	for i, ticket := range tickets {
		if ticket.Id == id {
			updatedTicket.Id = id
			// Preserve original CreatedBy and CreatedAt
			updatedTicket.CreatedBy = ticket.CreatedBy
			updatedTicket.CreatedAt = ticket.CreatedAt
			now := time.Now()
			updatedTicket.UpdatedAt = &now
			tickets[i] = updatedTicket
			c.Data["json"] = updatedTicket
			c.ServeJSON()
			return
		}
	}

	c.Data["json"] = map[string]string{"error": "Ticket not found"}
	c.Ctx.Output.SetStatus(404)
	c.ServeJSON()
}

// DeleteTicket deletes a ticket by ID
func (c *TicketController) DeleteTicket() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid ID"}
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}

	for i, ticket := range tickets {
		if ticket.Id == id {
			tickets = slices.Delete(tickets, i, i+1)
			c.Data["json"] = map[string]string{"message": "Ticket deleted"}
			c.ServeJSON()
			return
		}
	}

	c.Data["json"] = map[string]string{"error": "Ticket not found"}
	c.Ctx.Output.SetStatus(404)
	c.ServeJSON()
}
