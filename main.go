package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Joziaz/database-GO/models"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	model := models.NewArticleModel()

	for {
		fmt.Print("Write the number of the options: \n" +
			"	[1] Select All\n" +
			"	[2] Select One \n" +
			"	[3] Create \n" +
			"	[4] Update \n" +
			"	[5] Delete \n" +
			"	[0] Close \n",
		)
		scanner.Scan()

		switch scanner.Text() {
		case "1":
			fmt.Println(model.GetAll())

		case "2":
			ID, err := inputID()
			if err != nil {
				log.Printf("This caracther is invalid, error: %s \n", err)
				continue
			}

			art, err := model.GetOne(ID)
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println(art)

		case "3":

			art := models.Article{}
			fillArticle(&art)

			err := model.Create(art)
			if err != nil {
				log.Printf("Error when create: %s", err)
			}

		case "0":
			os.Exit(2)
		default:
			fmt.Println("This caracther is invalid")
		}
	}
}

func fillArticle(article *models.Article) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Insert the title of the article: ")
	scanner.Scan()
	newTitle := scanner.Text()

	fmt.Print("Insert the description of the article: ")
	scanner.Scan()
	newDescription := scanner.Text()

	article.Title = newTitle
	article.Description = newDescription
}

func inputID() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Insert the ID of the article: ")
	scanner.Scan()
	ID, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, err
	}

	return ID, nil
}
