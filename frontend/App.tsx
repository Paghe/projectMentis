import { TamaguiProvider, createTamagui } from '@tamagui/core'
import { defaultConfig } from '@tamagui/config/v4'
import { Button, YStack } from 'tamagui'
import { config } from './tamagui.config'
import TaskCard from './components/TaskCard'


export default function App() {
  return (
    <TamaguiProvider config={config} defaultTheme="light">
      <YStack padding={16} space={8}>
        <TaskCard title="Buy groceries" completed={false} />
        <TaskCard title="Read a book" description="Finish chapter 5" completed={true} />
        <TaskCard title="Write blog post" completed={false} />
      </YStack>
    </TamaguiProvider>
  )
}