import { z } from 'zod';
import dotenv from 'dotenv';

dotenv.config();

const envSchema = z.object({
  DISCORD_TOKEN: z.string().min(1),
  DEEPL_AUTH_KEY: z.string().min(1),
  CLIENT_ID: z.string().min(1),
});

export const config = envSchema.parse(process.env);
