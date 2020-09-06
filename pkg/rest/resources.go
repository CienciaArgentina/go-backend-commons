package rest

const (
	// Email.
	SendEmail = "/email"

	// Enigma.
	EnigmaSignUp                  = "/users"
	EnigmaLogin                   = "/users/login"
	EnigmaConfimPasswordReset     = "/users/confirmpasswordreset"
	EnigmaSendPasswordReset       = "/users/sendpasswordreset"
	EnigmaConfirmEmail            = "/users/confirmemail"
	EnigmaSendConfirmationEmail   = "/users/sendconfirmationemail"
	EnigmaResendConfirmationEmail = "/users/resendconfirmationemail"
	EnigmaForgotUsername          = "/users/forgotusername"
)
