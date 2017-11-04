package app

import (
	"github.com/gin-gonic/gin"
)

func GetInstructions(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "GET api/v1/instructions"})
}

func GetInstruction(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "GET api/v1/instructions/1"})
}

func PostInstruction(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "POST api/v1/instructions"})

}

func UpdateInstruction(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "PUT api/v1/instructions/1"})

}

func DeleteInstruction(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "DELETE api/v1/instructions/1"})
}
