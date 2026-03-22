package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// テスト用の環境変数を設定
	os.Setenv("DISCORD_TOKEN", "test_token")
	os.Setenv("DEEPL_AUTH_KEY", "test_key")
	defer func() {
		os.Unsetenv("DISCORD_TOKEN")
		os.Unsetenv("DEEPL_AUTH_KEY")
	}()

	tests := []struct {
		name    string
		wantNil bool
	}{
		{
			name:    "returns config",
			wantNil: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Load(); got == nil && !tt.wantNil {
				t.Error("Load() returned nil")
			}
		})
	}
}

func TestLoad_WithEnvVar(t *testing.T) {
	os.Setenv("DISCORD_TOKEN", "test_token")
	os.Setenv("DEEPL_AUTH_KEY", "test_key")
	defer func() {
		os.Unsetenv("DISCORD_TOKEN")
		os.Unsetenv("DEEPL_AUTH_KEY")
	}()

	cfg := Load()
	if cfg == nil {
		t.Fatal("Load() returned nil")
	}
	if cfg.DiscordToken != "test_token" {
		t.Errorf("DiscordToken = %v, want %v", cfg.DiscordToken, "test_token")
	}
	if cfg.DeepLAuthKey != "test_key" {
		t.Errorf("DeepLAuthKey = %v, want %v", cfg.DeepLAuthKey, "test_key")
	}
}

func TestLoad_MissingDiscordToken(t *testing.T) {
	// DISCORD_TOKENを設定しない
	os.Setenv("DEEPL_AUTH_KEY", "test_key")
	defer os.Unsetenv("DEEPL_AUTH_KEY")

	// log.Fatalはos.Exitを呼ぶため、テスト内では検証できない
	// 代わりに、config.goのロジックを変更してerrorを返すようにするか
	// このテストはスキップ
	t.Skip("log.Fatal causes os.Exit which cannot be tested")
}

func TestLoad_MissingDeepLAuthKey(t *testing.T) {
	// DEEPL_AUTH_KEYを設定しない
	os.Setenv("DISCORD_TOKEN", "test_token")
	defer os.Unsetenv("DISCORD_TOKEN")

	// log.Fatalはos.Exitを呼ぶため、テスト内では検証できない
	t.Skip("log.Fatal causes os.Exit which cannot be tested")
}
