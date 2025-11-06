import axios from "axios";

const client = axios.create({
    baseURL: "http://192.168.178.21:8080",
    timeout: 5000,
})

export default client;