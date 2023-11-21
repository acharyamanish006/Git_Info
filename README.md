# Git_Info

# NOTICE
Project is not complete !

# How To Run
1. Clone this repository to your local machine using the following command: 
2. The program will then fetch all repositories from that user and print out their names, description, url, etc...
3. ```bash
   go run main.go -u {git_username}
   ```
   #To get user followers use -F flag
4. ```bash
   go run main.go -u {git_username} -F
   ```
   #To get user following use -f flag
   
5. ```bash
   go run main.go -u {git_username} -f
   ```
   #To get help
6. ```bash
   go run main.go -h
   ```
   #To get user Contribution use -c flag
7. ```bash
   go run main.go -u {git_username} -c
