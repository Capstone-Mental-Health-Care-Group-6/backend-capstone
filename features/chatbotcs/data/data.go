package data

import (
	"FinalProject/features/chatbotcs"
	"strings"
	"sync"
)

type ChatbotCsData struct {
	Questions map[string]string
	mu        sync.RWMutex
}

func New(questions map[string]string) chatbotcs.ChatbotCsDataInterface {
	return &ChatbotCsData{
		Questions: map[string]string{
			"bagaimana cara melakukan pendaftaran?":          "Untuk mendaftar, kunjungi halaman pendaftaran dan isi formulir dengan informasi yang diperlukan.",
			"apa langkah-langkah untuk login?":               "Langkah-langkah login melibatkan memasukkan email dan password yang sudah terdaftar.",
			"bagaimana cara mengubah kata sandi?":            "Anda dapat mengubah kata sandi di halaman pengaturan akun setelah berhasil login.",
			"apa yang harus dilakukan jika lupa kata sandi?": "Klik 'Lupa Kata Sandi' pada halaman login dan ikuti petunjuk untuk mereset kata sandi.",
		},
	}
}

func (d *ChatbotCsData) GetAnswer(question string) string {
	d.mu.RLock()
	defer d.mu.RUnlock()

	answer, ok := d.Questions[strings.ToLower(question)]
	if ok {
		return answer
	}

	return "Maaf, saya tidak mengerti pertanyaan Anda."
}
