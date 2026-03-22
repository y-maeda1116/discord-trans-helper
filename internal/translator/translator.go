package translator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// DeepL API endpoints
const (
	deepLAPIURL = "https://api-free.deepl.com/v2/translate"
	// DeepL ProのURL: "https://api.deepl.com/v2/translate"
)

// DeepLRequest represents a request to DeepL API
type DeepLRequest struct {
	Text       []string `json:"text"`
	SourceLang string   `json:"source_lang,omitempty"`
	TargetLang string   `json:"target_lang"`
}

// DeepLResponse represents a response from DeepL API
type DeepLResponse struct {
	Translations []struct {
		Text       string `json:"text"`
		DetectedSourceLang string `json:"detected_source_language"`
	} `json:"translations"`
}

// DetectLanguage detects the language of the text
func DetectLanguage(text string) string {
	// 簡易的な言語判定（実際にはDeepL APIを使用）
	textLower := strings.ToLower(text)
	if strings.ContainsAny(textLower, "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをん") {
		return "JA" // 日本語
	}
	return "EN" // 英語（デフォルト）
}

// Translate translates text using DeepL API
func Translate(text, apiKey string) (string, string, error) {
	sourceLang := DetectLanguage(text)
	targetLang := "JA" // デフォルトは日本語
	if sourceLang == "JA" {
		targetLang = "EN"
	}

	reqBody := DeepLRequest{
		Text:       []string{text},
		SourceLang: sourceLang,
		TargetLang: targetLang,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", deepLAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "DeepL-Auth-Key "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("DeepL API error: %s", string(body))
	}

	var result DeepLResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(result.Translations) == 0 {
		return "", "", fmt.Errorf("no translations returned")
	}

	return result.Translations[0].Text, result.Translations[0].DetectedSourceLang, nil
}

// HandleTranslate handles the translate slash command
func HandleTranslate(s *discordgo.Session, i *discordgo.InteractionCreate, apiKey string) {
	// Get the message from interaction data
	message := i.ApplicationCommandData().Resolved.Messages[i.ApplicationCommandData().TargetID]
	content := message.Content

	if content == "" {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Content: "翻訳対象のメッセージが見つかりません",
			},
		})
		return
	}

	// Send thinking response
	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
		},
	}); err != nil {
		return
	}

	// Translate the message
	translatedText, detectedLang, err := Translate(content, apiKey)
	if err != nil {
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &[]string{fmt.Sprintf("翻訳に失敗しました: %v", err)}[0],
		})
		return
	}

	// Send translated message
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &[]string{fmt.Sprintf("🌐 翻訳結果\n\n原文(%s):\n```\n%s\n```\n\n翻訳:\n```\n%s\n```", detectedLang, content, translatedText)}[0],
	})
}
