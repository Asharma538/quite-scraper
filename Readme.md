# Quite-Scraper

Quite-Scraper is a web application designed to monitor user activity on Instagram using Firebase for authentication and Firestore for data storage.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Installation

### Backend

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/quite-scraper.git
    cd quite-scraper/backend
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up Firebase:
    - Download your Firebase Admin SDK key and save it as `config/secret_key.json`.

4. Run the backend server:
    ```sh
    go run app.go
    ```

### Frontend

1. Navigate to the frontend directory:
    ```sh
    cd ../frontend
    ```

2. Install dependencies:
    ```sh
    npm install
    ```

3. Run the development server:
    ```sh
    npm run dev
    ```

## Usage

1. Start the backend server as described in the installation section.
2. Start the frontend development server.
3. Open your browser and navigate to `http://localhost:3000`.

## Project Structure

```
quite-scraper/
├── backend/
│   ├── config/
│   │   └── secret_key.json
│   ├── controllers/
│   │   └── firebase_controller.go
│   ├── models/
│   ├── app.go
│   └── go.mod
├── frontend/
│   ├── pages/
│   │   └── index.js
│   ├── styles/
│   │   └── Home.module.css
│   ├── .nvmrc
│   ├── package.json
│   └── .gitignore
└── README.md
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

---
This readme is generated completely by Github Copilot! except this line :)
