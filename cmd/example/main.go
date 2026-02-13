package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// MessageRequest represents Bale API message structure
type MessageRequest struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"` // "Markdown" یا "HTML"
}

func main() {

	botToken := os.Getenv("BALE_TOKEN")
	chatID := os.Getenv("BALE_CHAT_ID")

	// ساخت درخواست
	msg := MessageRequest{
		ChatID:    chatID,
		Text:      "*سلام* و خوش آمدید!\nاین یک پیام تست از گولنگ است.",
		ParseMode: "Markdown", // فعال‌سازی فرمت‌بندی متن
	}

	// تبدیل به JSON
	jsonData, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("خطا در ساخت JSON: %v", err)
	}

	// URL کامل ربات
	url := "https://tapi.bale.ai/bot" + botToken + "/sendMessage"

	// ارسال درخواست
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("خطا در ارسال درخواست: %v", err)
	}
	defer resp.Body.Close()

	// خواندن پاسخ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("خطا در خواندن پاسخ: %v", err)
	}

	// نمایش نتیجه
	fmt.Printf("کد وضعیت: %d\n", resp.StatusCode)
	fmt.Printf("پاسخ سرور: %s\n", string(body))

	if resp.StatusCode == 200 {
		fmt.Println("✅ پیام با موفقیت ارسال شد!")
	} else {
		fmt.Println("❌ خطا در ارسال پیام")
	}
}
