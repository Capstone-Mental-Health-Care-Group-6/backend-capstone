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
			"Bagaimana Cara Mengubah Biodata Diri?":    "Buka aplikasi. Masuk ke akun Anda. Cari 'Profil' atau pada  menu. Klik opsi 'Edit Profil'. Isi data baru dan simpan perubahan. Apakah informasi yang saya berikan sudah jelas?.",
			"Apa Fitur yang Tersedia?":                 "Kami memiliki beberapa fitur yang tersedia seperti Artikel, Buat Janji, Chatbot dan Konseling. Namun, kami memiliki fitur unggulan yaitu konseling dengan dokter spesialis yang tersedia.",
			"Panduan Untuk Buat Janji Dengan Konselor": "Klik menu 'Buat Janji'. Pilih konselor yang ingin Anda panggil. Pilih tanggal dan waktu. Klik tombol 'Buat Janji'.",
			"Panduan Untuk Chatbot":                    "Klik menu 'Chatbot'. Pilih pertanyaan yang ingin Anda tanyakan. AI akan membalas pertanyaan Anda.",
			"Panduan Untuk Konseling":                  "Klik menu 'Konseling'. Pastikan Anda Sudah membuat Janji Dengan Dokter, Pilih Dokter yang ingin Anda konsultasi. Anda akan berkomunikasi melalui chat dan video call.",
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
