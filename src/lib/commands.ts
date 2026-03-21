import { REST, Routes } from 'discord.js';
import { config } from '../config/index.js';

/**
 * Register message context menu command (Translate button)
 */
export async function registerCommands() {
  const rest = new REST({ version: '10' }).setToken(config.DISCORD_TOKEN);

  const commands = [
    {
      name: 'Translate',
      type: 3, // Message context menu
    },
  ];

  try {
    console.log(`Started refreshing application (/) commands.`);

    const data = await rest.put(
      Routes.applicationCommands(config.CLIENT_ID),
      { body: commands }
    );

    console.log(`Successfully reloaded ${data} application (/) commands.`);
  } catch (error) {
    console.error('Error registering commands:', error);
    throw error;
  }
}
