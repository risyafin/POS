# About This Project
Point of sale  is a simple marketplace RESTful API backend development

What tech stack do I use?

* Go Programming Language for common language
* MySQL for Database Management
* gorm.io/gorm for Object Relational Mapping (ORM)
* gorilla/mux for Route Management
* joho/godotenv for Env Management
* github.com/golang-jwt/jwt for Authorization
* golang.org/x/crypto for Generate and Hash Password

# Installation
This is the instruction how you can test _go-store_ on your local computer:
## Prerequisites
- Make sure that you have installed MySQL on your computer.
  ```sh
  mysql --version
  ``` 
  If not installed, you can get it on [MySQL Documentaion](https://dev.mysql.com/doc/mysql-installation-excerpt/8.0/en/)
- Make sure that you have installed Go on your computer.
  ```sh
  go version
  ```
  If not installed, you can get it on [Go Documentaion](https://go.dev/doc/install)
- Make sure that you have installed Git on your computer.
  ```sh
  git --version
  ```
  If not installed, you can get it on [Git Documentaion](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) <br />

  ## Clone and Prepare Project
- Clone this project.
  ```sh
  git clone https://github.com/risyafin/POS.git
  ```
- Create new env file on project directory. Write the following into the env file.
  Example:
  ```
  DB_USERNAME=root
  DB_PASSWORD=root
  DB_HOST=127.0.0.1:3306
  DB_NAME=pos
  SERVER_PORT=8080
  ```
  Adjust to your environmental values. Then, **save file with name ```.env```** <br />

  ## Run Project
- Make sure your MySQL service is running. You can confirm this with a command:<br />
```Linux:```
  ```sh
  sudo systemctl status mysql
  ```
  ```Windows:```
  ```sh
  mysqladmin -u root -p status
  ```
- Run Project on project directory path.
  ```sh
  go run .
  ```
  After the output write "server starting...." response, you can use this project. Good Luck! <br />

 Project Usage
> [!NOTE]
> Some endpoint is need authorization to be accessed. You can get the token JWT by login, don't forget to register new account first if no user exist on database.
## Admin Route
### POST
- ```/register```, Register new admin. Example body:
```json
{
	"name": "Samsul Arifin",
	"username": "SamsulArifin",
	"password": "samsul1234",
	"branchId": 1
}
```
### POST
- ```/login```, login admin. Example body:
```json
{
	"username": "SamsulArifin",
	"password": "samsul1234"
}
```
