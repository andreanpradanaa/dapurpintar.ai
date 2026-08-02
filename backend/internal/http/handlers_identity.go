package http

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/identity"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/response"
)

// accountView is the M6 AccountView schema.
type accountView struct {
	ID            string  `json:"id"`
	Email         string  `json:"email"`
	Status        string  `json:"status"`
	Timezone      string  `json:"timezone"`
	EmailVerified bool    `json:"email_verified"`
	CreatedAt     string  `json:"created_at"`
	Profile       *string `json:"profile_id,omitempty"`
}

// userProfileView is the M6 UserProfileView schema.
type userProfileView struct {
	ID          string           `json:"id"`
	AccountID   string           `json:"account_id"`
	DisplayName string           `json:"display_name"`
	Status      string           `json:"status"`
	Timezone    *string          `json:"timezone"`
	Preferences *preferencesView `json:"preferences,omitempty"`
	CreatedAt   string           `json:"created_at"`
}

// preferencesView is the M6 PreferencesView schema.
type preferencesView struct {
	Status      string          `json:"status"`
	Preferences json.RawMessage `json:"preferences"`
	ValidFrom   string          `json:"valid_from"`
}

func toAccountView(a *identity.Account) accountView {
	created := a.CreatedAt.UTC().Format(time.RFC3339)
	return accountView{
		ID:            a.ID,
		Email:         a.Email,
		Status:        string(a.Status),
		Timezone:      a.Timezone,
		EmailVerified: a.EmailVerifiedAt != nil,
		CreatedAt:     created,
	}
}

func toUserProfileView(p *identity.UserProfile) userProfileView {
	created := p.CreatedAt.UTC().Format(time.RFC3339)
	return userProfileView{
		ID:          p.ID,
		AccountID:   p.AccountID,
		DisplayName: p.DisplayName,
		Status:      string(p.Status),
		Timezone:    p.Timezone,
		CreatedAt:   created,
	}
}

func toPreferencesView(p *identity.PreferenceSet) preferencesView {
	return preferencesView{
		Status:      p.Status,
		Preferences: p.Preferences,
		ValidFrom:   p.ValidFrom,
	}
}

// register handles POST /api/v1/accounts.
func (h *Handler) register(c *fiber.Ctx) error {
	var req struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		DisplayName string `json:"display_name"`
		Timezone    string `json:"timezone"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, payloadError(err))
	}

	account, profile, err := h.identity.Register(c.Context(), identity.RegisterInput{
		Email:       req.Email,
		Password:    req.Password,
		DisplayName: req.DisplayName,
		Timezone:    req.Timezone,
	})
	if err != nil {
		return response.Error(c, err)
	}

	session, loginErr := h.identity.Login(c.Context(), identity.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if loginErr == nil {
		h.setSessionCookies(c, session.AccessToken, session.RefreshToken, session.AccessExpiresAt, h.refreshCookieTTL)
	} else {
		h.log.Error("register: auto-login failed", "error", loginErr)
	}

	view := toAccountView(account)
	if profile != nil {
		view.Profile = &profile.ID
	}
	return response.Created(c, view)
}

// login handles POST /api/v1/accounts/login.
func (h *Handler) login(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, payloadError(err))
	}

	session, err := h.identity.Login(c.Context(), identity.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return response.Error(c, err)
	}

	h.setSessionCookies(c, session.AccessToken, session.RefreshToken, session.AccessExpiresAt, h.refreshCookieTTL)
	return response.OK(c, toAccountView(session.Account))
}

// refresh handles POST /api/v1/accounts/refresh.
func (h *Handler) refresh(c *fiber.Ctx) error {
	refreshToken := c.Cookies(h.sessionCookies.RefreshCookieName)
	session, err := h.identity.Refresh(c.Context(), refreshToken)
	if err != nil {
		return response.Error(c, err)
	}

	h.setSessionCookies(c, session.AccessToken, session.RefreshToken, session.AccessExpiresAt, h.refreshCookieTTL)
	return response.OK(c, toAccountView(session.Account))
}

// logout handles POST /api/v1/accounts/logout.
func (h *Handler) logout(c *fiber.Ctx) error {
	refreshToken := c.Cookies(h.sessionCookies.RefreshCookieName)
	if err := h.identity.Logout(c.Context(), refreshToken); err != nil {
		return response.Error(c, err)
	}
	h.sessionCookies.ClearAccess(c)
	h.sessionCookies.ClearRefresh(c)
	return response.NoContent(c)
}

// me handles GET /api/v1/accounts/me.
func (h *Handler) me(c *fiber.Ctx) error {
	account, err := h.identity.GetAccountByID(c.Context(), middlewareSubject(c))
	if err != nil {
		return response.Error(c, err)
	}
	return response.OK(c, toAccountView(account))
}

// getProfile handles GET /api/v1/profile.
func (h *Handler) getProfile(c *fiber.Ctx) error {
	profile, err := h.identity.GetProfileByAccountID(c.Context(), middlewareSubject(c))
	if err != nil {
		return response.Error(c, err)
	}
	view := toUserProfileView(profile)
	view.Preferences = h.activePreferences(c, profile.ID)
	return response.OK(c, view)
}

// updateProfile handles PATCH /api/v1/profile.
func (h *Handler) updateProfile(c *fiber.Ctx) error {
	var req struct {
		DisplayName *string `json:"display_name"`
		Timezone    *string `json:"timezone"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, payloadError(err))
	}

	profile, err := h.identity.UpdateProfile(c.Context(), middlewareSubject(c), identity.UpdateProfileInput{
		DisplayName: req.DisplayName,
		Timezone:    req.Timezone,
	})
	if err != nil {
		return response.Error(c, err)
	}
	view := toUserProfileView(profile)
	view.Preferences = h.activePreferences(c, profile.ID)
	return response.OK(c, view)
}

// updatePreferences handles PATCH /api/v1/profile/preferences.
func (h *Handler) updatePreferences(c *fiber.Ctx) error {
	var req struct {
		Preferences json.RawMessage `json:"preferences"`
		ValidFrom   *string         `json:"valid_from"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, payloadError(err))
	}

	prefs, err := h.identity.UpdatePreferences(c.Context(), middlewareSubject(c), identity.UpdatePreferencesInput{
		Preferences: req.Preferences,
		ValidFrom:   req.ValidFrom,
	})
	if err != nil {
		return response.Error(c, err)
	}
	return response.OK(c, toPreferencesView(prefs))
}

func (h *Handler) activePreferences(c *fiber.Ctx, profileID string) *preferencesView {
	prefs, err := h.identity.ActivePreferences(c.Context(), profileID)
	if err != nil {
		return nil
	}
	view := toPreferencesView(prefs)
	return &view
}

func (h *Handler) setSessionCookies(c *fiber.Ctx, access, refresh string, accessExpires time.Time, refreshTTL time.Duration) {
	h.sessionCookies.SetAccess(c, access, accessExpires)
	h.sessionCookies.SetRefresh(c, refresh, time.Now().UTC().Add(refreshTTL))
}
