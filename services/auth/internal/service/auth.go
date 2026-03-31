package auth

import (
	"booking-system/services/auth/internal/jwt"
	"booking-system/services/auth/internal/models"
	"booking-system/services/auth/internal/storage"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
)

type Auth interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uid int64, err error)
	User(ctx context.Context, email string) (user models.User, err error)
	IsAdmin(ctx context.Context, userID int64) (isAdmin bool, err error)
}

type AuthService struct {
	log      *slog.Logger
	auth     Auth
	tokenTTL time.Duration
	secret   string
}

// New returns a new instance of the Auth service
func New(
	log *slog.Logger,
	auth Auth,
	tokenTTL time.Duration,
	secret string,
) *AuthService {
	return &AuthService{
		log:      log,
		auth:     auth,
		tokenTTL: tokenTTL,
		secret:   secret,
	}
}

// Login checks if user with given credentials exists in the system
//
// If user exists, but password is incorrect, returns error.
// If user doesn't exist, returns error.
func (a *AuthService) Login(
	ctx context.Context,
	email string,
	password string,
) (string, error) {
	const op = "auth.Login"
	log := a.log.With(slog.String("op", op), slog.String("email", email))

	log.Info("logging in")

	user, err := a.auth.User(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			a.log.Warn("user not found", slog.String("err", err.Error()))

			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		a.log.Error("failed to get user", slog.String("err", err.Error()))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		log.Error("invalid credentials", slog.String("err", err.Error()))

		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	log.Info("user logged in successfully")

	token, err := jwt.NewToken(user, a.tokenTTL, a.secret)
	if err != nil {
		a.log.Error("failed to create token", slog.String("err", err.Error()))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}


// RegisterNewUser registers new user in the system and returns user ID.
// If user with given username already exists, returns error.
func (a *AuthService) RegisterNewUser(
	ctx context.Context,
	email string,
	password string,
) (int64, error) {
	const op = "auth.RegisterNewUser"

	log := a.log.With(slog.String("op", op), slog.String("email", email))

	log.Info("registering user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", slog.String("err", err.Error()))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := a.auth.SaveUser(ctx, email, passHash)
	if err != nil {
		if errors.Is(err, storage.ErrUserExists) {
			log.Warn("user already exists", slog.String("err", err.Error()))

			return 0, fmt.Errorf("%s: %w", op, ErrUserExists)
		}
		log.Error("failed to save user", slog.String("err", err.Error()))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user registered")

	return id, err
}

func (a *AuthService) IsAdmin(
	ctx context.Context,
	userID int64,
) (bool, error) {
	const op = "auth.IsAdmin"

	log := a.log.With(slog.String("op", op), slog.Int64("user_id", userID))

	log.Info("checking if user is admin")

	isAdmin, err := a.auth.IsAdmin(ctx, userID)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("checked if user is admin", slog.Bool("is_admin", isAdmin))

	return isAdmin, nil
}
