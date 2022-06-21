package customers

import (
	"net/http"

	"github.com/Singh555/mycms/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetCustomer(c *gin.Context) {
	id := c.Param("id")

	var customer models.Customer

	if result := h.DB.Order("id DESC").Where("id = ?", id).Select("id", "first_name", "last_name", "email", "mobile", "address", "status", "created_at", "updated_at").Find(&customer); result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while getting customer data", "error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer data found", "data": &customer})
}