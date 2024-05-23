package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func createFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Ошибка чтения файла –", err)
	}
	defer file.Close()
	fmt.Printf("Файл %s успешно создан. \n", filename)
}

func deleteFile(filename string) {
	_, err := os.Stat(filename)
	if err != nil {
		log.Fatalf("Файл '%s' не был найден.", filename)
	}

	var s string
	fmt.Printf("Чтобы удалить файл, введите его название (%s)\n", filename)
	_, err = fmt.Scan(&s)
	if err != nil {
		log.Fatal("Ошибка при считывании названия файла")
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

func readFile(filename string) {
	bytesArr, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Ошибка чтения файла –", err)
	}
	fmt.Println(string(bytesArr))
}

func editFile(filename string) {
	_, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("Файл '%s' не был найден и будет создан\n", filename)
	}

	readFile(filename)

	fmt.Println("Введите данные для добавления в файл:")
	var data string
	fmt.Scan(&data)
	err = os.WriteFile(filename, []byte(data), 0666)
	if err != nil {
		log.Fatal("Ошибка записи в файл")
	}
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
		readFile(*readPtr)
	case *editPtr != "":
		editFile(*editPtr)
	case *deletePtr != "":
		deleteFile(*deletePtr)
	default:
		fmt.Println("Укажите тип операции и/или имя файла, для информации --help")
	}

}
