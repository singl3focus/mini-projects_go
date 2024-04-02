package pkg

import (
	"os"
	"bufio"
	"path/filepath"
	"strings"
	"log"
	"net/http"
	"encoding/json"
)

func SearchWordFile(w http.ResponseWriter, r *http.Request) {


	folderPath := "C:\\Users\\Дом\\Documents\\GitHub\\mini-projects_go\\web_wordfile_search\\examples"
	searchWord :=  r.URL.Query().Get("word")

	foundWords := []string{}

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Ошибка при обходе файла %q: %v\n", path, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil
		}

		// Проверка, является ли файл обычным файлом, а не каталогом
		if !info.Mode().IsRegular() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			log.Printf("Ошибка при открытии файла %q: %v\n", path, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		// Установка разделителя на пробелы, чтобы разбить строки на слова
		scanner.Split(bufio.ScanWords)

		// Перебор каждого слова в файле
		for scanner.Scan() {
			word := scanner.Text()

			// Поиск искомого слова (без учета регистра)
			if strings.ToLower(word) == strings.ToLower(searchWord) {
				foundWords = append(foundWords, file.Name())
				break
			}
		}

		if err := scanner.Err(); err != nil {
			log.Printf("Ошибка при сканировании файла %q: %v\n", path, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(foundWords); err != nil {
		http.Error(w, "Failed to encode note data", http.StatusInternalServerError)
		return
	}
}
