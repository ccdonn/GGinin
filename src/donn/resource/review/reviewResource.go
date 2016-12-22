package review

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	//	"reflect"
	//	"mime/multipart"
	//	"donn/config"
	"donn/entity"
	"donn/repository/article"
)

const (
	IMAGE_LIMITATION         = 1
	MEGABYTE_SIZE            = 1048576
	FILE_HEADER_CHECK_LENGTH = 8
	JPEG_SIGNATURE           = "ffd8"
	PNG_SIGNATURE            = "89504e470d0a1a0a"
	GIF_SIGNATURE_1          = "474946383761"
	GIF_SIGNATURE_2          = "474946383961"
)

type NewReviewForm struct {
	Uid         string    `form:"uid"`
	Context     string    `form:"context"`
	DisplayTime time.Time `form:"displayTime"`
	ImageFile   os.File
	ImageUrl    string `form:"imageUrl"`
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "indexReview_OK",
	})
}

func Create(c *gin.Context) {

	var form NewReviewForm
	if c.Bind(&form) == nil {
		hasher := sha1.New()

		fmt.Println("uid:" + form.Uid)
		file, _, err := c.Request.FormFile("imageFile")
		if file == nil && err != nil {
			fmt.Printf("ERR: %s", err)
		}

		io.Copy(hasher, file)
		filename := hex.EncodeToString(hasher.Sum(nil)[:])
		fmt.Println("sha1: " + filename)
		file.Seek(0, io.SeekStart)

		var buff bytes.Buffer
		// file size
		fileSize, err := buff.ReadFrom(file)
		fmt.Println("filesize=", fileSize) // this will return you a file size.

		if fileSize > IMAGE_LIMITATION*MEGABYTE_SIZE {
			fmt.Println("Over filesize error")
		}

		if fileSize < FILE_HEADER_CHECK_LENGTH {
			fmt.Println("Less filesize error")
		}

		// file format
		hexFileHeader := hex.EncodeToString(buff.Next(FILE_HEADER_CHECK_LENGTH)[:])
		fmt.Println(hexFileHeader)

		if strings.HasPrefix(hexFileHeader, JPEG_SIGNATURE) {
			fmt.Println("JPEG file")
		} else if strings.HasPrefix(hexFileHeader, PNG_SIGNATURE) {
			fmt.Println("PNG file")
		} else if strings.HasPrefix(hexFileHeader, GIF_SIGNATURE_1) || strings.HasPrefix(hexFileHeader, GIF_SIGNATURE_2) {
			fmt.Println("GIF file")
		} else {
			fmt.Println("File format not support")
		}

		out, err := os.Create("/tmp/" + filename + ".jpg")
		if err != nil {
			fmt.Printf("CreateErr:, %s", err)
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Printf("CopyErr:, %s", err)
		}

		//		dbmap := config.GetDbMap()
		inst := &entity.Article{0, 10, "title", filename + ".jpg"}
		fmt.Println("inst=", inst)
		fmt.Println("&inst=", &inst)
		article.Save(inst)
		//		err = dbmap.Insert(inst)

	}

	c.JSON(http.StatusOK, gin.H{
		"status": "createReview_OK",
	})
}

func Show(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "showReview_OK",
	})
}

func Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "updateReview_OK",
	})
}

func Patch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "patchReview_OK",
	})
}

func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "deleteReview_OK",
	})
}
