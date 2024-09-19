package v1

import (
	"api-gateway/api/handler/models"
	"api-gateway/pkg/jwt"
	"api-gateway/protos/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags users
// @Accept json
// @Produce json
// @Param register body models.RegisterUserReq true "User Registration Data"
// @Success 200 {object} models.RegisterUserRes "message: Registration successful"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users [post]
func (h *handlerV1) Register(c *gin.Context) {
	req := models.RegisterUserReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data: " + err.Error()})
		return
	}

	res, err := h.service.User().RegisterUser(c, &user.RegisterReq{
		Name:     req.Username,
		Email:    req.Email,
		Password: req.Password,
		Age:      int32(req.Age),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.RegisterUserRes{
		Message: res.Message,
	})
}

// VerifyUser godoc
// @Summary Verify a user's email address
// @Description Verifies a user via a verification code
// @Tags users
// @Accept json
// @Produce json
// @Param verify body models.VerifyUserReq true "Verification"
// @Success 200 {object} models.VerifyUserRes "message: User verified"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Verification User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users/verify [post]
func (h *handlerV1) VerifyUser(c *gin.Context) {
	req := models.VerifyUserReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data: " + err.Error()})
		return
	}
	fmt.Println(req)
	res, err := h.service.User().VerifyUser(c, &user.VerifyReq{
		Email: req.Email,
		Code:  req.Code,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.VerifyUserRes{
		UserId:    int(res.UserId),
		Token:     res.Token,
		Message:   res.Message,
		ExpiresIn: res.ExpireTime,
	})
}

// Login godoc
// @Summary Login user
// @Description Authenticate user with email and password
// @Tags users
// @Accept json
// @Produce json
// @Param login body models.LoginUserReq true "User Login Data"
// @Success 200 {object} models.LoginUserRes "message: Login User"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users/login [post]
func (h *handlerV1) Login(c *gin.Context) {
	req := models.LoginUserReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data: " + err.Error()})
		return
	}

	res, err := h.service.User().LoginUser(c, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.LoginUserRes{
		Userid:    int(res.UserId),
		Token:     res.Token,
		Message:   res.Message,
		ExpiresIn: res.ExpireTime,
	})
}

// GetProfile godoc
// @Summary Get user profile
// @Description Retrieve the current user's profile data by userId
// @Tags users
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {object} models.ProfileUserRes "User profile data"
// @Failure 400 {object} map[string]string "User ID is required"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users/profile/{userId} [get]
// @Security BearerAuth
func (h *handlerV1) GetProfile(c *gin.Context) {
	token := c.GetHeader("Authoriz")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	res, err := h.service.User().GetUser(c, &user.GetUserReq{UserId: int32(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user profile: " + err.Error()})
		return
	}

	b, s, err := jwt.VerifyToken(token)
	if !b || s != res.Response[0].Email || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	profile := models.ProfileUserRes{
		UserId:   int(res.UserId),
		Username: res.Response[0].Name,
		Email:    res.Response[0].Email,
		Age:      int(res.Response[0].Age),
	}

	c.JSON(http.StatusOK, profile)
}
