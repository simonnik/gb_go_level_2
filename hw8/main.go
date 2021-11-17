// В качестве завершающего задания нужно выполнить программу поиска дубликатов файлов.
// Дубликаты файлов - это файлы, которые совпадают по имени файла и по его размеру.
// Нужно написать консольную программу, которая проверяет наличие дублирующихся файлов.
// Программа должна работать на локальном компьютере и получать на вход путь до директории.
// Программа должна вывести в стандартный поток вывода список дублирующихся файлов,
// которые находятся как в директории, так и в поддиректориях директории,
// переданной через аргумент командной строки.
// Данная функция должна работать эффективно при помощи распараллеливания программы
// Программа должна принимать дополнительный ключ - возможность удаления обнаруженных дубликатов файлов после поиска.
// Дополнительно нужно придумать, как обезопасить пользователей от случайного удаления файлов.
// В качестве ключей желательно придерживаться общепринятых практик по использованию командных опций.
// Критерии приемки программы:
// 1. Программа компилируется
// 2. Программа выполняет функциональность, описанную выше.
// 3. Программа покрыта тестами
// 4. Программа содержит документацию и примеры использования
// 5. Программа обладает флагом “-h/--help” для краткого объяснения функциональности
// 6. Программа должна уведомлять пользователя об ошибках, возникающих во время выполнения

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()

	if err != nil {
		log.Fatal("logging not initialize")
	}
	sugar := logger.Sugar()
	//defer logger.Sync()

	sugar.Info("Logger initialize")
	var (
		path = flag.String("path", ".", "dir for search")
		file = flag.String("file", "", "file name for search")
		d    = flag.Bool("d", false, "Delete duplicated files?")
	)

	flag.Parse()
	sugar.With(
		zap.String("path", *path),
		zap.String("file", *file),
		zap.Bool("d", *d),
	).Info("Flag parsed")

	sugar.Info("Search for duplicates started")
	duplicateList, err := FindDuplicate(*path, *file, sugar)
	if err != nil {
		sugar.Fatalf("Error: %v", err)
	}
	duplicateCount := len(duplicateList)

	if duplicateCount > 0 {
		fmt.Printf("Found duplicates: %d\n", duplicateCount)
		for i, duplicateName := range duplicateList {
			fmt.Printf("%d. %s\n", i+1, duplicateName)
		}
	} else {
		fmt.Printf("No copy of %q in path: %q\n", *file, *path)
		return
	}

	if *d {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("\nAre you sure you want to delete the duplicates? ('Y' to confirm):\t")
		scanner.Scan()
		userAnswer := scanner.Text()
		if strings.ToLower(userAnswer) == "y" {
			for _, duplicateName := range duplicateList {
				err := os.Remove(duplicateName)
				if err != nil {
					log.Fatalf("Error: %v", err)
				}
				fmt.Println("- deleted:", duplicateName)
			}
		}
	}
}
