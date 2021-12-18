package usercontroller

import (
	"net/http"
	"os"
	"time"

	"github.com/conan080262/gin-basic-api.git/configs"
	"github.com/conan080262/gin-basic-api.git/models"
	"github.com/conan080262/gin-basic-api.git/utility"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/matthewhartstonge/argon2"
)

func GetAll(c *gin.Context) {

	var users []models.User
	configs.DB.Find(&users)
	// configs.DB.Select("id").Find(&users)
	// configs.DB.Raw("select * from users order by id desc").Scan(&users)
	c.JSON(200, gin.H{
		"data": users,
	})
}

func Register(c *gin.Context) {
	var inputJson InputRegister
	if err := c.ShouldBindJSON(&inputJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Fullname: inputJson.Fullname,
		Email:    inputJson.Email,
		Password: inputJson.Password,
	}

	userExist := configs.DB.Where("email = ?", inputJson.Email).First(&user)

	if userExist.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "มีผู้ใช้งานอีเมลนี้ในระบบแล้ว"})
		return
	}

	result := configs.DB.Debug().Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(201, gin.H{
		"data": user,
		"row":  result.RowsAffected,
	})
}

func Login(c *gin.Context) {

	var inputJson InputLogin
	if err := c.ShouldBindJSON(&inputJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:    inputJson.Email,
		Password: inputJson.Password,
	}

	userAccount := configs.DB.Where("email = ?", inputJson.Email).First(&user)

	if userAccount.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบผู้ใช้งานนี้ในระบบ"})
		return
	}

	ok, _ := argon2.VerifyEncoded([]byte(inputJson.Password), []byte(user.Password))

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "รหัสผ่านไม่ถูกต้อง"})
		return
	}

	//JWT Token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 2).Unix(),
	})

	jwtSecret := os.Getenv("JWT_SECRET")
	accessToken, _ := claims.SignedString([]byte(jwtSecret))
	c.JSON(200, gin.H{
		"data":        user,
		"message":     "เข้าสู่ระบบสำเร็จ",
		"accesstoken": accessToken,
	})
}

func GetById(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := configs.DB.First(&user, id) // ใช้สำหรับ P Key เท่านั้น
	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูล"})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}

func SearchByFullname(c *gin.Context) {
	// id := c.DefaultQuery("firstname", "Guest")
	fullname := c.Query("fullname") //?fullname = john

	users := []models.User{}

	result := configs.DB.Where("fullname ILIKE ?", "%"+fullname+"%").Scopes(utility.Paginate(c)).Find(&users)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูล"})
		return
	}

	c.JSON(200, gin.H{
		"data": users,
	})
}

func GetProfile(c *gin.Context) {
	user := c.MustGet("user")
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
