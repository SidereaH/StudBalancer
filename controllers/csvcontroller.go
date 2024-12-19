package controllers

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"stud-distributor/database"
	"stud-distributor/distributing"
	"stud-distributor/models"
)

type UserFromCsv struct {
	FirstName string
	LastName  string
	Email     string
	Group     models.Group
}

func RegisterUserByCSV(context *gin.Context) {
	// Извлекаем бинарный файл из тела запроса
	file, _, err := context.Request.FormFile("file")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read file from request: " + err.Error()})
		return
	}
	defer file.Close()

	// Читаем содержимое файла в буфер
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file: " + err.Error()})
		return
	}

	// Парсим CSV из буфера
	records, err := parseCSV(buf)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse CSV: " + err.Error()})
		return
	}
	usersFromCsv := []UserFromCsv{}
	for i, record := range records {
		if i == 0 {
			continue
		}
		if database.ExistsByEmail(record) == false || database.ExistsByPhone(record) == false {
			//регаем если нет типа
			var user models.User
			if err := distributing.CreateUserWithoutDistrib(&user, record); err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to load fileds to user: %s  \n from file: %s ", err.Error(), record)})
				continue
			}
			record := database.Instance.Create(&user)
			if record.Error != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
				context.Abort()
				continue
			}
			var currentGroup models.Group
			err := database.Instance.Where("id = ?", 1).Find(&currentGroup).Error
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
				continue
			}
			resposeUser := UserFromCsv{
				FirstName: user.FirstName,
				LastName:  user.SecondName,
				Email:     user.Email,
				Group:     currentGroup,
			}
			usersFromCsv = append(usersFromCsv, resposeUser)
		}
	}
	// Возвращаем количество обработанных строк
	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Processed %d students successfully!", len(usersFromCsv)),
		"data":    usersFromCsv, // Можно вернуть данные для отладки
	})
}

// parseCSV парсит CSV-данные из io.Reader и возвращает слайс строк
func parseCSV(reader io.Reader) ([][]string, error) {
	csvReader := csv.NewReader(reader)

	// Настройки для CSV, если нужны
	csvReader.Comma = ';'       // Разделитель по умолчанию — запятая
	csvReader.LazyQuotes = true // Учитывать нестрогие кавычки

	// Читаем все записи из CSV
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV: %v", err)
	}

	return records, nil
}
