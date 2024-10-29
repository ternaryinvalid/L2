package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Определение структуры Line для хранения строк и их ключей.
type Line struct {
	text string
	key  string
}

func main() {
	inputFileName := flag.String("i", "", "Имя входного файла")
	outputFileName := flag.String("o", "", "Имя выходного файла")
	keyColumn := flag.Int("k", 0, "Колонка для сортировки (по умолчанию: 0)")
	numeric := flag.Bool("n", false, "Сортировать числа")
	reverse := flag.Bool("r", false, "Сортировать в обратном порядке")
	unique := flag.Bool("u", false, "Игнорировать дубликаты")
	monthSort := flag.Bool("M", false, "Сортировать по месяцам")
	ignoreTrailingSpace := flag.Bool("b", false, "Игнорировать конечные пробелы")
	checkSorted := flag.Bool("c", false, "Проверить, отсортированы ли данные")
	numericSuffix := flag.Bool("h", false, "Сортировать числа с суффиксами")

	flag.Parse()

	if *inputFileName == "" {
		fmt.Println("Необходимо указать имя входного файла")
		return
	}

	file, err := os.Open(*inputFileName)
	if err != nil {
		fmt.Printf("Ошибка при открытии входного файла: %v\n", err)
		return
	}
	defer file.Close()

	lines := make([]Line, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, Line{text: text, key: extractKey(text, *keyColumn, *monthSort, *numericSuffix)})
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка при чтении входного файла: %v\n", err)
		return
	}

	if *checkSorted && !isSorted(lines, *reverse) {
		fmt.Println("Данные не отсортированы.")
		return
	}

	if *numericSuffix {
		sort.SliceStable(lines, func(i, j int) bool {
			return compareHumanNumbers(lines[i].key, lines[j].key, *reverse)
		})
	} else {
		sort.SliceStable(lines, func(i, j int) bool {
			return compareStrings(lines[i].key, lines[j].key, *numeric, *reverse, *ignoreTrailingSpace)
		})
	}

	if *unique {
		lines = removeDuplicates(lines)
	}

	outputFile := os.Stdout
	if *outputFileName != "" {
		outputFile, err = os.Create(*outputFileName)
		if err != nil {
			fmt.Printf("Ошибка при создании выходного файла: %v\n", err)
			return
		}
		defer outputFile.Close()
	}

	for _, line := range lines {
		fmt.Fprintln(outputFile, line.text)
	}
}

// Функция compareStrings сравнивает две строки с учетом различных флагов.
func compareStrings(a, b string, numeric, reverse, ignoreTrailingSpace bool) bool {
	if ignoreTrailingSpace {
		a = strings.TrimSpace(a)
		b = strings.TrimSpace(b)
	}

	if numeric {
		numA, errA := strconv.Atoi(a)
		numB, errB := strconv.Atoi(b)

		if errA == nil && errB == nil {
			if reverse {
				return numA > numB
			}
			return numA < numB
		}
	}

	if reverse {
		return a > b
	}
	return a < b
}

// Функция extractKey извлекает ключ из строки на основе заданных параметров.
func extractKey(line string, keyColumn int, monthSort, numericSuffix bool) string {
	fields := strings.Fields(line)
	if keyColumn >= len(fields) {
		return ""
	}
	key := fields[keyColumn]

	if monthSort {
		t, err := time.Parse("January", key)
		if err == nil {
			return t.Format("01")
		}
	}

	if numericSuffix {
		parts := strings.SplitN(key, ".", 2)
		if len(parts) == 2 {
			numericPart, err := strconv.Atoi(parts[0])
			if err == nil {
				return fmt.Sprintf("%06d.%s", numericPart, parts[1])
			}
		}
	}

	return key
}

// Функция compareHumanNumbers сравнивает строки с числами и суффиксами.
func compareHumanNumbers(a, b string, reverse bool) bool {
	numA, suffixA := extractNumericSuffix(a)
	numB, suffixB := extractNumericSuffix(b)

	if numA == numB {
		if reverse {
			return suffixA > suffixB
		}
		return suffixA < suffixB
	}

	if reverse {
		return numA > numB
	}
	return numA < numB
}

// Функция extractNumericSuffix извлекает числовую часть и суффикс из строки.
func extractNumericSuffix(s string) (int, string) {
	parts := strings.SplitN(s, ".", 2)
	num := 0
	if len(parts) > 0 {
		if n, err := strconv.Atoi(parts[0]); err == nil {
			num = n
		}
	}
	suffix := ""
	if len(parts) > 1 {
		suffix = parts[1]
	}
	return num, suffix
}

// Функция isSorted проверяет, отсортированы ли строки.
func isSorted(lines []Line, reverse bool) bool {
	for i := 1; i < len(lines); i++ {
		if (reverse && lines[i].key > lines[i-1].key) || (!reverse && lines[i].key < lines[i-1].key) {
			return false
		}
	}
	return true
}

// Функция removeDuplicates удаляет дубликаты из списка строк.
func removeDuplicates(lines []Line) []Line {
	uniqueLines := make(map[string]struct{})
	result := make([]Line, 0)

	for _, line := range lines {
		if _, found := uniqueLines[line.text]; !found {
			uniqueLines[line.text] = struct{}{}
			result = append(result, line)
		}
	}

	return result
}
