import { View, Text } from "react-native";

interface TaskCardProps {
    title: string;
    completed: boolean;
}

const TaskCard = ({ title, completed }: TaskCardProps) => {
    return (
        <Text style={{backgroundColor: completed ? "lightgreen" : "lightgray"}}>
            {title} {completed ? "✅" : "❌"}
        </Text>
    );
}

export default TaskCard;
