package core_user_domain

import (
	"fmt"
	"net/mail"
	"regexp"

	core_error "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/error"
)

type User struct {
	ID           int
	Version      int
	FullName     string
	PasswordHash string
	Email        string
	PhoneNumber  *string
	Telegram     *string
}

func NewUser(
	id int,
	version int,
	fullName string,
	passwordHash string,
	email string,
	phoneNumber *string,
	telegram *string,
) (User, error) {
	if err := validate(fullName, email, phoneNumber, telegram); err != nil {
		return User{}, fmt.Errorf(
			"validate user domain: %w",
			err,
		)
	}
	return User{
		ID:           id,
		Version:      version,
		FullName:     fullName,
		PasswordHash: passwordHash,
		Email:        email,
		PhoneNumber:  phoneNumber,
		Telegram:     telegram,
	}, nil
}

func NewUninitializedUser(
	fullName string,
	passwordHash string,
	email string,
	phoneNumber *string,
	telegram *string,
) (User, error) {
	if err := validate(fullName, email, phoneNumber, telegram); err != nil {
		return User{}, fmt.Errorf(
			"validate user domain: %w",
			err,
		)
	}
	return User{
		ID:           UninitializedID,
		Version:      UninitializedVersion,
		FullName:     fullName,
		PasswordHash: passwordHash,
		Email:        email,
		Telegram:     telegram,
	}, nil
}

func validate(
	fullName string,
	email string,
	phoneNumber *string,
	telegram *string,
) error {
	fullNameLen := len([]rune(fullName))
	if fullNameLen < 3 || fullNameLen > 100 {
		return fmt.Errorf(
			"invalid `FullName` len: %d: %w",
			fullNameLen,
			core_error.ErrInvalidArgument,
		)
	}

	emailLen := len([]rune(email))
	if emailLen < 5 || emailLen > 100 {
		return fmt.Errorf(
			"invalid `Email` len: %d: %w",
			emailLen,
			core_error.ErrInvalidArgument,
		)
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf(
			"invalid `Email` format: %v: %w",
			err,
			core_error.ErrInvalidArgument,
		)
	}

	if phoneNumber != nil {
		phoneNumberLen := len([]rune(*phoneNumber))
		if phoneNumberLen < 10 || phoneNumberLen > 15 {
			return fmt.Errorf(
				"invalid `PhoneNumber` len: %d: %w",
				phoneNumberLen,
				core_error.ErrInvalidArgument,
			)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*phoneNumber) {
			return fmt.Errorf(
				"invalid `PhoneNumber` format: %w",
				core_error.ErrInvalidArgument,
			)
		}
	}

	if telegram != nil {
		telegramLen := len([]rune(*telegram))
		if telegramLen < 6 || telegramLen > 33 {
			return fmt.Errorf(
				"invalid `Telegram` len: %d: %w",
				telegramLen,
				core_error.ErrInvalidArgument,
			)
		}

		re := regexp.MustCompile(`^\@[A-Za-z0-9_]+$`)

		if !re.MatchString(*telegram) {
			return fmt.Errorf(
				"invalid `Telegram` format: %w",
				core_error.ErrInvalidArgument,
			)
		}
	}

	return nil
}
