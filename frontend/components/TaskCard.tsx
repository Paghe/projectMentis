import { YStack, XStack, Button, Text } from 'tamagui'

interface TaskCardProps {
  title: string
  description?: string
  completed: boolean
  onToggle?: () => void
}

export default function TaskCard({ title, description, completed, onToggle }: TaskCardProps) {
  return (
    <YStack
      p="$4"
      br="$4"
      bg={completed ? '$green3' : '$gray2'}
      mb="$2"
      bw={1}
      bc={completed ? '$green10' : '$gray7'}
    >
      <Text fontSize={16} fontWeight="600" mb="$2">
        {title}
      </Text>
      {description && (
        <Text fontSize={14} color="$gray10" mb="$2">
          {description}
        </Text>
      )}
      <XStack jc="flex-end">
        <Button size="$2" onPress={onToggle}>
          {completed ? 'Undo' : 'Complete'}
        </Button>
      </XStack>
    </YStack>
  )
}
