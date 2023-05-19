package handlers

import (
	database "contacts/databases"
	"contacts/models"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	_ "github.com/nats-io/nats.go"
)

type Message struct {
	Topic string
	Data  []byte
}
type Contact struct {
	IContact database.IContact
}

var ChMessage chan Message

func init() {
	ChMessage = make(chan Message)
	go PublishData()
}

func PublishData() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()
	for msg := range ChMessage {
		nc.Publish(msg.Topic, msg.Data)
	}
}

func (c *Contact) Create(ctx *gin.Context) {

	contact := new(models.Contact)

	err := ctx.Bind(contact)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		ctx.Abort()
	}

	err = contact.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		ctx.Abort()
	}

	result, err := c.IContact.Create(contact)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		ctx.Abort()
	}

	fmt.Println("http://localhost:58081/v1/private/service/email/" + fmt.Sprint(contact.ID))

	resp, err := http.Get("http://localhost:58081/v1/private/service/email/" + fmt.Sprint(contact.ID))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "unable to send email")
		ctx.Abort()
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "unable to send email")
		ctx.Abort()
	}

	mresult := make(map[string]any)
	mresult["data"] = result
	mresult["email"] = string(buf)

	ctx.JSON(http.StatusCreated, mresult)

	// buf, err := io.ReadAll(ctx.Request.Body)
	// if err == nil {
	// 	//ctx.Request.Body.Read(buf)
	// 	json.Unmarshal(buf, contact)
	// }

}

func (c *Contact) CreateUseMessageBroker(ctx *gin.Context) {

	contact := new(models.Contact)

	err := ctx.Bind(contact)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		ctx.Abort()
	}

	err = contact.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		ctx.Abort()
	}

	result, err := c.IContact.Create(contact)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		ctx.Abort()
	}

	// nats code

	//nc, _ := nats.Connect(nats.DefaultURL)

	data, _ := contact.ToJson()

	//nc.Publish("contact.create", []byte(data))

	ChMessage <- Message{Topic: "contact.create", Data: []byte(data)}

	ctx.JSON(http.StatusCreated, result)

	// buf, err := io.ReadAll(ctx.Request.Body)
	// if err == nil {
	// 	//ctx.Request.Body.Read(buf)
	// 	json.Unmarshal(buf, contact)
	// }

}

func (c *Contact) GetByID() func(*gin.Context) {
	return func(ctx *gin.Context) {

		id, ok := ctx.Params.Get("id")
		if !ok {
			ctx.String(http.StatusBadRequest, "invalid id parameter")
			ctx.Abort()
		}

		contact, err := c.IContact.Get(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			ctx.Abort()
		}

		ctx.JSON(http.StatusOK, contact)
		ctx.Abort()
	}
}
