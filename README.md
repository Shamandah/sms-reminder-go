# 📩 SMS Reminder System (Go + Twilio)

## 📌 Overview

The SMS Reminder System is a Go-based application that sends automated SMS reminders using the Twilio API.

This project is designed to assist diabetic patients by sending scheduled insulin reminders throughout the day.

---

## 🚀 Features

- Send automated SMS reminders
- Uses Twilio API for reliable message delivery
- Secure environment variable configuration
- Clean and simple Go implementation
- Easily customizable reminder times

---

## 🛠 Tech Stack

- Go (Golang)
- Twilio SMS API
- VS Code
- Git & GitHub
Project Structure
sms-reminder-go/
│
├── main.go
├── go.mod
├── .env.example
├── .gitignore
└── README.md
---
⚙️ Setup Instructions

1️⃣ Clone the Repository

```bash
git clone https://github.com/Shamandah/sms-reminder-go.git
cd sms-reminder-go
2. Install Dependencies
go mod tidy
3.Create .env file
TWILIO_ACCOUNT_SID=your_account_sid
TWILIO_AUTH_TOKEN=your_auth_token
TWILIO_PHONE_NUMBER=your_twilio_number
4.Run the Application
go run main.go
If successful you will see: Message sent successfully! SID: SMxxxxxxxxxxxxxxxx



🔐 Security

Sensitive credentials are stored using environment variables to protect API keys and authentication tokens.

🎯 Future Improvements

Add scheduling with cron jobs

Add a web dashboard

Store patient data in a database

Add error logging system

Deploy to cloud (AWS / Render / Railway)

## 📂 Project Structure




