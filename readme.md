# 🍲 Recipes API

Welcome to the Recipes API repository! This project is a simple and efficient API designed to manage and share delicious recipes.

# 📋 Table of Contents
- Features
- Installation
- Usage
- API Endpoints
  
# ✨ Features
- Create, read, update, and delete recipes
- Filter recipes by ingredients and categories
- User authentication and authorization
- Docker support for easy setup

# 🚀 Installation
To get started with the Recipes API, follow these steps:
1. Clone the repository:

```bash
git clone https://github.com/ahmedkhaeld/recipes-api.git
cd recipes-api
```
2. Set up environment variables:
Create an .env file and add the necessary environment variables.

```
DB_HOST=your_database_host
DB_USER=your_database_user
DB_PASSWORD=your_database_password
```

3. Run the Docker containers:
```
docker-compose up --build
```


# 📚 Usage
Once the application is up and running, you can access the API at http://localhost:8080. Use tools like Postman or cURL to interact with the API.

#📌 API Endpoints
Here's a list of the main API endpoints:

- GET /recipes: Get all recipes
- POST /recipes: Create a new recipe
- GET /recipes/{id}: Get a specific recipe
- PUT /recipes/{id}: Update a recipe
- DELETE /recipes/{id}: Delete a recipe
