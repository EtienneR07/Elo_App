package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"backend/config"
	"backend/models"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}

func GoogleLogin(c *gin.Context) {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	c.SetCookie("oauth_state", state, 300, "/", "", false, true)

	url := config.GoogleOAuthConfig.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(ctx *gin.Context) {
	stateCookie, err := ctx.Cookie("oauth_state")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "State cookie not found"})
		return
	}

	if ctx.Query("state") != stateCookie {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state parameter"})
		return
	}

	code := ctx.Query("code")
	token, err := config.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token: " + err.Error()})
		return
	}

	client := config.GoogleOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info: " + err.Error()})
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read user info: " + err.Error()})
		return
	}

	var googleUser GoogleUser
	if err := json.Unmarshal(data, &googleUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info: " + err.Error()})
		return
	}

	var user models.User
	err = config.GetDB().Where("email = ? AND o_auth_provider = ?", googleUser.Email, "google").First(&user).Error

	if err != nil {
		user = models.User{
			Email:         googleUser.Email,
			Name:          googleUser.Name,
			OAuthProvider: "google",
			OAuthID:       googleUser.ID,
		}
		if err := config.GetDB().Create(&user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
	}

	jwtToken, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}

	redirectURL := fmt.Sprintf("%s/auth/callback?token=%s", frontendURL, jwtToken)
	ctx.Redirect(http.StatusTemporaryRedirect, redirectURL)
}
