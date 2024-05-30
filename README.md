# User Signup and Login System

This project implements a basic user signup and login system using a backend server in Go and a frontend interface in React. The backend server listens on port 8080, communicates with a MongoDB database, and stores encrypted user credentials securely.

## Project Structure

The project is divided into two main components:

### Backend Server (Go)

- Listens on port 8080.
- Handles user signup and login requests.
- Encrypts passwords using bcrypt before storing them in MongoDB.

### Frontend (React)

- Provides a minimalistic interface for user signup and login.
- `/login` route displays the login page.
- `/user/signup` route displays the signup page.

## Backend Implementation (Go)

### Setup Instructions

1. **Database Configuration**
   - Set up MongoDB and obtain connection URI.
   - Update `store/store.go` with your MongoDB URI.

2. **Running the Server**
   - Install dependencies: 
     ```bash
     go mod tidy
     ```
   - Start the server: 
     ```bash
     go run main.go
     ```

### API Endpoints

- **POST `/user/signup`**
  - Accepts userId and password.
  - Encrypts the password using bcrypt.
  - Stores the encrypted credentials in MongoDB.

- **POST `/user/login`**
  - Accepts userId and password.
  - Validates the credentials against MongoDB.
  - Issues a session token upon successful login.


## Frontend Implementation (React)

### Routes

- **`/login`**
  - Displays a form for user login.
  - Submits login credentials to the Go backend server.

- **`/user/signup`**
  - Displays a form for user signup.
  - Submits signup details to the Go backend server.

### Setup Instructions

1. **Install Dependencies**
   - Navigate to the `frontend` directory: 
     ```bash
     cd frontend
     ```
   - Install dependencies: 
     ```bash
     npm install
     ```

2. **Running the Frontend**
   - Start the frontend server: 
     ```bash
     npm start
     ```

### User Interaction

- Users navigate to `/login` to access the login form.
- Users navigate to `/user/signup` to access the signup form.
- Upon form submission, data is sent to the backend for processing and storage.
