<h1>Taha Tahviyeh Telegram Bot</h1>

## Overview
Taha Tahviyeh Telegram Bot is a Go-based chatbot designed to help Taha Tahviyeh agents recommend and manage products efficiently. This bot allows agents to search for products, view details, and manage product inventory via a Telegram interface. Built with modern technologies, the bot provides an intuitive experience for both customers and administrators.

## Table Of Content
<!-- TOC -->
  * [Overview](#overview)
  * [Table Of Content](#table-of-content)
  * [Features](#features)
    * [For Normal Users:](#for-normal-users)
    * [For Admins & Super Admins:](#for-admins--super-admins)
  * [Tech Stack](#tech-stack)
    * [Language:](#language)
    * [Key Libraries:](#key-libraries)
    * [Tools:](#tools)
  * [Project Structure](#project-structure)
  * [Installation & Setup](#installation--setup)
    * [Prerequisites](#prerequisites)
    * [Steps](#steps)
  * [Usage](#usage)
  * [Contributing](#contributing)
  * [License](#license)
  * [Contact](#contact)
<!-- TOC -->

## Features
### For Normal Users:
- 🔍 **Search** for products by title or metadata (brand & product type)
- 📄 **View Product Details**
- 📂 **Access Product Files**
- ❓ **Browse FAQs**
- 📖 **Get Help Information**
- 📞 **Contact Support**

### For Admins & Super Admins:
- ➕ **Add, Edit, Remove Products**
- 🔄 **Manage Product Types & Brands**
- 📝 **Modify FAQs, About, and Help Texts**
- 📦 **Update Product Files & Metadata**

## Tech Stack
### Language:
- 🟦 Go

### Key Libraries:
- [GORM](https://gorm.io/) - ORM for PostgreSQL/Supabase
- [Telegram Bot API](https://github.com/go-telegram-bot-api/telegram-bot-api) - Telegram bot framework
- [Fiber](https://gofiber.io/) - Web framework for handling API requests
- [MinIO](https://min.io/) - Object storage for product files

### Tools:
- 🐳 Docker (for containerization)
- 🛠️ Makefile (for automation)

## Project Structure
```
.
├── app/               # Core application logic
├── build/             # Docker-related configurations
├── cmd/               # Entry point (main.go)
├── config/            # Configuration management
├── internal/          # Business logic
├── pkg/               # Utilities, database, and bot-related helpers
├── server/            # Telegram bot server
├── README.md          # Documentation
├── LICENSE            # License file
├── Makefile           # Build automation
└── go.mod             # Go dependencies
```

## Installation & Setup
### Prerequisites
Ensure you have the following installed:
- Go
- Docker & Docker Compose
- PostgreSQL (or Supabase for cloud-based storage)
- MinIO (or an S3-compatible object storage)

### Steps
1. **Clone the repository:**
   ```sh
   git clone https://github.com/your-repo/taha-tahviyeh-bot.git
   cd taha-tahviyeh-bot
   ```
2. **Set up configuration:**
  - Copy `config.sample.json` to `config.json`
  - Update database credentials, Telegram bot token, and MinIO settings
  - Update the `.env` file in the `build/minio` directory with your credentials:
     ```sh
     MINIO_ROOT_USER=<your-root-user>
     MINIO_ROOT_PASSWORD=<your-root-password>
     ```
3. **Run the application:**
   ### Development
     - Start services:
       ```sh
       make up
       ```
     - Start services with logs:
       ```sh
       make uplog
       ```
     - Stop services:
       ```sh
       make down
       ```
     - View logs:
       ```sh
       make logs
       ```
     - Build services:
       ```sh
       make build
       ```
     - Rebuild and restart:
       ```sh
       make rebuild
       ```
   ### Deployment
     - Deploy:
       ```sh
       make deploy
       ```
     - Stop deployed services:
       ```sh
       make deploy-down
       ```
     - View deployment logs:
       ```sh
       make deploy-logs
       ```
     - View logs for a specific container:
       ```sh
       make deploy-log 
       ```
     - Build deployment services:
       ```sh
       make deploy-build
       ```

## Usage
1. Add the bot to your Telegram contacts.
2. Use `/start` to begin interaction.
3. Admins can manage products using inline commands or menus.

## Contributing
Contributions are welcome! Feel free to fork the repository and submit pull requests.

## License
This project is licensed under the MIT License.

## Contact
For any inquiries or support, contact **Taha Tahviyeh** or open an issue in the repository.

