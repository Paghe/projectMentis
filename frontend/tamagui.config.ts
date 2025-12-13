// tamagui.config.ts
import { createTamagui } from 'tamagui';
import { defaultConfig } from '@tamagui/config/v4';

export const config = createTamagui(defaultConfig);

type Conf = typeof config;

declare module 'tamagui' {
  interface TamaguiCustomConfig extends Conf {}
}
