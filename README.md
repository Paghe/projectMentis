# 🧠 Mentis — Level Up Your Life

> *“Turn your daily grind into a quest for progress.”*

Mentis is a **gamified productivity app** built to help people with ADHD (and anyone who struggles with consistency) stay motivated, organized, and rewarded.  
It transforms your everyday tasks into **quests**, gives you **XP**, and helps you **build streaks** like a real-life RPG.

---

## 🚀 What’s Inside

Mentis is a **full-stack project** built to learn and experiment with:
- **Go (Gin)** — for the backend REST API  
- **MongoDB** — as the data store for users, tasks, and characters  
- **React Native (Expo)** — for the mobile app interface  
- **ADHD-focused gamification logic** — XP, streaks, and visual feedback loops  

The goal isn’t just to manage to-dos —  
it’s to *feel rewarded* for finishing them.

---

## 🧩 Features

| Feature | Description |
|----------|--------------|
| 🎯 **Task System** | Create, edit, and complete tasks to earn XP and gold |
| 🧙 **Character Progression** | Level up your character as you stay productive |
| 🔥 **Streak Tracking** | Keep your daily momentum — consistency matters |
| 🪙 **Rewards** | Earn gold, XP, and gems for staying on track |
| 💾 **Persistent Storage** | All data stored in MongoDB |
| 📱 **React Native App** | Track progress and manage tasks on the go |
| 🧠 **ADHD-Friendly UX** | Minimal distractions, big visual feedback |

---

## 🧰 Tech Stack

**Backend**
- [Go](https://golang.org/) with [Gin](https://gin-gonic.com/)
- MongoDB via [mongo-go-driver](https://www.mongodb.com/docs/drivers/go/current/)
- REST API structure (`routes.go`, `handlers/`, `models/`)

**Frontend**
- [React Native](https://reactnative.dev/) (via [Expo](https://expo.dev/))
- [Axios](https://axios-http.com/) for API calls
- [React Navigation](https://reactnavigation.org/)
- [React Native Paper](https://callstack.github.io/react-native-paper/) for styling

---

**Database**
- MongoDB Atlas (or local instance)

## 🧠 App Flow

1. **Create a user**  
   → A linked character is automatically generated.

2. **Add tasks (quests)**  
   → Each task belongs to a user and their character.

3. **Complete tasks**  
   → XP and gold are awarded automatically.  
   → The streak and level update.

4. **Check your dashboard**  
   → See progress, stats, and current streak.

It’s a productivity app that *feels like a game*.

---

## ⚙️ How to Run It Locally

### 🐹 Backend Setup (Go + MongoDB)

#### 1. Clone the repo
```bash
https://github.com/Paghe/projectMentis.git
cd mentis/backend
```

#### 2. Install dependencies
```bash
go mod tidy
cd ../cmd/mentis
```

#### 3. Set up MongoDB
Make sure MongoDB is running locally or create a [MongoDB Atlas](https://www.mongodb.com/cloud/atlas) cluster.  
Then set your connection string in `.env`:
```
MONGO_URI=mongodb+srv://<user>:<password>@cluster.mongodb.net
DATABASE_NAME=mentis
```

#### 4. Run the server
```bash
go run main.go
```

✅ The API will be live at:
```
http://localhost:8080
```

#### 5. Test endpoints
Use [Postman](https://www.postman.com/) or curl:
```bash
curl http://localhost:8080/users
```

---

### 📱 Mobile App Setup (React Native)

#### 1. Go to the frontend folder
```bash
cd ../mobile
```

#### 2. Install dependencies
```bash
npm install
```

#### 3. Run Expo
```bash
npm start
```

#### 4. Connect your phone
- Install **Expo Go** on Android or iOS.

---

## 🧪 API Endpoints

| Method | Endpoint | Description |
|---------|-----------|--------------|
| `POST` | `/users` | Create a new user |
| `GET` | `/users` | Get all users |
| `GET` | `/users/:id` | Get user by ID |
| `POST` | `/character` | Create a new character |
| `GET` | `/character/:id` | Get character by ID |
| `POST` | `/task` | Create a task |
| `PUT` | `/task/:id` | Edit a task |
| `PUT` | `/task/:id/complete` | Mark task as complete + add XP |

---

## 🧙 Future Roadmap

- [ ] 🛡️ Authentication (JWT)
- [ ] 🔔 Push Notifications
- [ ] 🧭 XP / Level balancing
- [ ] 🎨 UI polish for ADHD clarity
- [ ] ☁️ Cloud deployment (Render / Azure)
- [ ] 📊 Analytics dashboard

---

## 💡 Why Mentis?

> Because ADHD isn’t about laziness — it’s about managing energy, dopamine, and focus.  

Mentis uses **instant feedback**, **streaks**, and **visual progress** to keep motivation high.  
The app is designed to reward **small wins**, not punish inconsistency.

It’s not just code.  
It’s a tool to *fight procrastination with game mechanics.*

---

## 💬 Contact

👤 **Alessandro Paghera**  
🌍 Based in Germany 🇩🇪  
💻 Software Engineer — R&D @ Schunk  
🧩 Building Mentis to learn Go, MongoDB, and React Native  

---

## 🏁 Closing Words

> _“Discipline is hard. Progress shouldn’t be boring.”_  
> With Mentis, you turn productivity into play.
