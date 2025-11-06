import client from "./client";


export const getTasks = async () => {
    try {
        const response = await client.get("/tasks");
        return response.data;
    } catch(error) {
        console.error("❌ API error:", error);
        return null;
    }
  };