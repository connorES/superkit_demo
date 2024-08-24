package app

import (
	"superkit_demo/app/events"
	"superkit_demo/plugins/auth"

	"github.com/anthdm/superkit/event"
)

// Events are functions that are handled in separate goroutines.
// They are the perfect fit for offloading work in your handlers
// that otherwise would take up response time.
// - sending email
// - sending notifications (Slack, Telegram, Discord)
// - analytics..

// Register your events here.
func RegisterEvents() {
	event.Subscribe(auth.UserSignupEvent, events.OnUserSignup)
	event.Subscribe(auth.ResendVerificationEvent, events.OnResendVerificationToken)
}
