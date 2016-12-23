package article

import (
	articleDao "donn/repository/article"
	articleResp "donn/response/article"
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func Index(c *gin.Context) {
	articles := articleDao.List()
	fmt.Println(articles)
	c.JSON(http.StatusOK, gin.H{
		"status": "indexArticle_OK",
		"result": articleResp.MakeIndexArticleResp(articles),
	})
}

func Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "createArticle_OK",
	})
}

func Show(c *gin.Context) {

	// check input param

	inst := articleDao.Get(3)
	fmt.Print(inst)

	if inst == nil {
		// exception
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "showArticle_OK",
			//			"result": articleResp.GetArticleResp{inst.Id, inst.Title, inst.Content},
			"res": articleResp.MakeShowArticleResp(inst),
		})
	}
}

func Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "updateArticle_OK",
	})
}

func Patch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "patchArticle_OK",
	})
}

func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "deleteArticle_OK",
	})
}
