import { useEffect, useState } from "react";
import { getTasks} from "../api/taskApi";
import { Task } from "../types/task";

export const useTasks = () => {
    const [data, setData] =  useState<Task[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(false);
    useEffect(() => {
        async function fetchData() {
            try{
                const result = await getTasks();
                setData(result.tasks);
                setLoading(false);
            } catch(error) {
                console.log("❌ API error:", error)
                setError(true);
            }
            
        }
        fetchData();
    }, []);
    return {data, loading, error};
}