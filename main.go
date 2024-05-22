package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func createFile(filename string) {
	_, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Файл %s успешно создан. \n", filename)
}

func deleteFile(filename string) {
	var s string
	fmt.Printf("Чтобы удалить файл, введите его название (%s)\n", filename)
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatal("Ошибка при считаывании названия файла")
	}
	if s == filename {
		err = os.Remove(filename)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Файл %s успешно удален.\n", filename)
		return
	}
	log.Fatalf("Введенное имя '%s' не совпадает с указанным именем файла '%s'", s, filename)

}

func main() {
	createPtr := flag.String("create", "", "Create new file.")
	readPtr := flag.String("read", "", "Read file data.")
	editPtr := flag.String("edit", "", "Edit file data.")
	deletePtr := flag.String("delete", "", "Delete file.")
	flag.Parse()


	switch {
	case *createPtr != "":
		createFile(*createPtr)

	case *readPtr != "":
		fmt.Println("read")
	case *editPtr != "":
		fmt.Println("Edit")
	case *deletePtr != "":
		deleteFile(*deletePtr)
	default:
		fmt.Println("Не задано имя файла.")
	}

	
}
