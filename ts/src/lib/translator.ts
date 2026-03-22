import { config } from '../config/index.js';
import { DeepLClient } from 'deepl-node';

const translator = new DeepLClient(config.DEEPL_AUTH_KEY);

/**
 * Check if the text is Japanese
 */
function isJapanese(text: string): boolean {
  const japaneseRegex = /[\u3040-\u309F\u30A0-\u30FF\u4E00-\u9FFF]/;
  return japaneseRegex.test(text);
}

/**
 * Translate text to target language based on source language detection
 * - If source is Japanese, translate to English
 * - Otherwise, translate to Japanese
 */
export async function translateText(text: string): Promise<{ original: string; translated: string; sourceLang: string; targetLang: string }> {
  try {
    const isSourceJapanese = isJapanese(text);
    const sourceLang = isSourceJapanese ? 'ja' : 'en';
    const targetLang = isSourceJapanese ? 'en-US' : 'ja';

    const result = await translator.translateText(text, sourceLang, targetLang);

    return {
      original: text,
      translated: result.text,
      sourceLang,
      targetLang,
    };
  } catch (error) {
    console.error('Translation error:', error);
    throw new Error('Failed to translate text');
  }
}
