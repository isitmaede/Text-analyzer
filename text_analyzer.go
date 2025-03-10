package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// دالة لحساب عدد الكلمات
func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

// دالة لحساب عدد الحروف
func countCharacters(text string) int {
	count := 0
	for _, char := range text {
		if !unicode.IsSpace(char) {
			count++
		}
	}
	return count
}

// دالة لحساب عدد الجمل
func countSentences(text string) int {
	sentences := strings.Count(text, ".") + strings.Count(text, "!") + strings.Count(text, "?")
	return sentences
}

// دالة لحساب تكرار الكلمات
func wordFrequency(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))
	frequency := make(map[string]int)

	for _, word := range words {
		word = strings.Trim(word, ".,!?()[]{}\":;") // إزالة الرموز من الكلمات
		frequency[word]++
	}

	return frequency
}

func main() {
	fmt.Println("📖 أدخل النص المراد تحليله:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	fmt.Println("\n🔍 نتائج التحليل:")
	fmt.Printf("عدد الكلمات: %d\n", countWords(text))
	fmt.Printf("عدد الحروف: %d\n", countCharacters(text))
	fmt.Printf("عدد الجمل: %d\n", countSentences(text))

	fmt.Println("\n📊 الكلمات الأكثر تكرارًا:")
	for word, freq := range wordFrequency(text) {
		if freq > 1 {
			fmt.Printf("%s: %d مرات\n", word, freq)
		}
	}
}
