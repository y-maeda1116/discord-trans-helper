import { Client, GatewayIntentBits, Interaction, MessageContextMenuCommandInteraction } from 'discord.js';
import { config } from './config/index.js';
import { translateText } from './lib/translator.js';
import { registerCommands } from './lib/commands.js';

const client = new Client({
  intents: [
    GatewayIntentBits.Guilds,
    GatewayIntentBits.GuildMessages,
    GatewayIntentBits.MessageContent,
  ],
});

// Ready event - bot is online
client.once('ready', async () => {
  console.log(`Logged in as ${client.user?.tag}!`);

  // Register commands when bot is ready
  try {
    await registerCommands();
    console.log('Commands registered successfully');
  } catch (error) {
    console.error('Failed to register commands:', error);
  }
});

// Handle interactions (button clicks, context menu commands)
client.on('interactionCreate', async (interaction: Interaction) => {
  if (!interaction.isMessageContextMenuCommand()) return;

  const ctxInteraction = interaction as MessageContextMenuCommandInteraction;

  if (ctxInteraction.commandName === 'Translate') {
    await handleTranslateInteraction(ctxInteraction);
  }
});

/**
 * Handle Translate interaction
 */
async function handleTranslateInteraction(interaction: MessageContextMenuCommandInteraction) {
  const targetMessage = interaction.targetMessage;
  const textToTranslate = targetMessage.content;

  if (!textToTranslate || textToTranslate.trim() === '') {
    await interaction.reply({
      content: 'このメッセージには翻訳可能なテキストがありません。',
      ephemeral: true,
    });
    return;
  }

  try {
    // Defer reply because translation might take a moment
    await interaction.deferReply({ ephemeral: true });

    // Translate the text
    const result = await translateText(textToTranslate);

    // Send ephemeral response (only visible to the user who clicked)
    await interaction.editReply({
      content: `**翻訳結果:**\n\n${result.translated}`,
    });
  } catch (error) {
    console.error('Error handling translate interaction:', error);
    await interaction.editReply({
      content: '翻訳中にエラーが発生しました。',
    });
  }
}

// Shutdown handlers
const shutdown = (signal: string) => {
  console.log(`Received ${signal}, shutting down gracefully...`);
  client.destroy();
  process.exit(0);
};

process.on('SIGINT', () => shutdown('SIGINT'));
process.on('SIGTERM', () => shutdown('SIGTERM'));

// Login with token
client.login(config.DISCORD_TOKEN);
