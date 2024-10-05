package handlers

import (
	"encoding/json"
	"fmt"
	"go-backed/app/services"
	"go-backed/app/types/errors"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
)

// Webhook handles Stripe webhook events
func Webhook(us *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Limit the body size to 64KB (Stripe recommends max of 65536 bytes)
		const MaxBodyBytes = int64(65536)
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)

		// Read the body of the request
		payload, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Unable to read request body"})
			return
		}

		// Parse the event from the request body
		var event stripe.Event
		if err := json.Unmarshal(payload, &event); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse webhook body json: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
			return
		}

		// Handle the specific event types
		switch event.Type {
		case "checkout.session.completed":
			var session stripe.CheckoutSession
			if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing checkout session JSON: %v\n", err)
			}

			// Get the user ID from the session
			sessionEmail := session.CustomerEmail
			if sessionEmail == "" {
				fmt.Fprintf(os.Stderr, "No email found in session: %v\n", session)
				c.JSON(http.StatusBadRequest, gin.H{"error": "No email found in session"})
				return
			}
			user, err := us.GetUserByEmail(sessionEmail)
			if err != nil {
				if err == errors.ErrUserNotFound {
					_, err := us.CreateUser(sessionEmail, "", true)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Error creating new user: %v\n", err)
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating new user"})
						return
					}
				}
				return
			}

			// Update the user's subscription status
			user.IsPremium = true
			err = us.UpdateUser(user)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error updating user: %v\n", err)
			}

		case "payment_intent.succeeded":
			var paymentIntent stripe.PaymentIntent
			if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing payment intent JSON: %v\n", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing payment intent"})
				return
			}
			fmt.Println("PaymentIntent was successful")

		case "payment_method.attached":
			var paymentMethod stripe.PaymentMethod
			if err := json.Unmarshal(event.Data.Raw, &paymentMethod); err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing payment method JSON: %v\n", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing payment method"})
				return
			}
			// Call a function to handle the successful attachment of a PaymentMethod
			// handlePaymentMethodAttached(paymentMethod)
			fmt.Println("PaymentMethod was attached")

		// Handle other event types here...

		default:
			fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
		}

		// Return a 200 status code to acknowledge receipt of the event
		c.Status(http.StatusOK)
	}
}
