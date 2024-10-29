package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Определение флагов командной строки
	afterLines := flag.Int("A", 0, "Печатать +N строк после совпадения")
	beforeLines := flag.Int("B", 0, "Печатать +N строк до совпадения")
	contextLines := flag.Int("C", 0, "Печатать ±N строк вокруг совпадения")
	countOnly := flag.Bool("c", false, "Печатать только количество совпадений")
	ignoreCase := flag.Bool("i", false, "Игнорировать регистр при поиске")
	invertMatch := flag.Bool("v", false, "Исключить строки с совпадениями")
	fixedString := flag.Bool("F", false, "Искать точное совпадение со строкой, не как паттерн")
	lineNumbers := flag.Bool("n", false, "Печатать номера строк")

	flag.Parse()

	// Проверка наличия аргумента - паттерна для поиска
	if flag.NArg() < 1 {
		fmt.Println("Использование: mygrep [опции] паттерн")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Чтение паттерна из аргумента командной строки
	pattern := flag.Arg(0)

	// Открытие файла для чтения или использование stdin
	var input io.Reader
	if flag.NArg() == 2 {
		fileName := flag.Arg(1)
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("Ошибка при открытии файла: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}

	scanner := bufio.NewScanner(input)
	matchingLines := 0
	var lines []string
	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		// Применение флага -F (точное совпадение)
		if *fixedString {
			for _, line := range lines {
				if strings.Contains(line, pattern) != *invertMatch {
					printLine(line, *lineNumbers, lineNumber)
					matchingLines++
				}
			}
		} else {
			// Применение флага -i (игнорирование регистра)
			if *ignoreCase {
				lineLower := strings.ToLower(line)
				patternLower := strings.ToLower(pattern)
				if strings.Contains(lineLower, patternLower) != *invertMatch {
					printLine(line, *lineNumbers, lineNumber)
					matchingLines++
				}
			} else {
				// Поиск совпадения (с учетом флага -v)
				if (strings.Contains(line, pattern) && !*invertMatch) || (!strings.Contains(line, pattern) && *invertMatch) {
					printLine(line, *lineNumbers, lineNumber)
					matchingLines++
				}
			}
		}

		// Печать +N строк после совпадения
		if *afterLines > 0 && matchingLines > 0 && matchingLines <= *afterLines {
			lines = appendLine(lines, line, *lineNumbers, lineNumber)
		}

		// Печать +N строк до совпадения
		if *beforeLines > 0 && matchingLines > 0 && matchingLines <= *beforeLines+1 {
			lines = appendLine(lines, line, *lineNumbers, lineNumber)
		}

		// Печать ±N строк вокруг совпадения
		if *contextLines > 0 && matchingLines > 0 && matchingLines <= *contextLines*2+1 {
			lines = appendLine(lines, line, *lineNumbers, lineNumber)
		}

		// Удаление старых строк
		if len(lines) > *contextLines*2+1 {
			lines = lines[1:]
		}
	}

	// Печать результата
	for _, line := range lines {
		fmt.Println(line)
	}

	// Печать количества совпадений, если флаг -c установлен
	if *countOnly {
		fmt.Println("Количество совпадений:", matchingLines)
	}
}

// Функция для добавления строки в список с учетом номера строки
func appendLine(lines []string, line string, lineNumbers bool, lineNumber int) []string {
	if lineNumbers {
		lines = append(lines, fmt.Sprintf("%d: %s", lineNumber, line))
	} else {
		lines = append(lines, line)
	}
	return lines
}

// Функция для вывода строки с учетом номера строки
func printLine(line string, lineNumbers bool, lineNumber int) {
	if lineNumbers {
		fmt.Printf("%d: %s\n", lineNumber, line)
	} else {
		fmt.Println(line)
	}
}
