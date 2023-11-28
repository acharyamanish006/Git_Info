# Git_Info

# NOTICE
The project is not complete!


# How To Run

1. Clone this repository to your local machine using the following command:
2. The program will then fetch all repositories from that user and print out their names, descriptions, URLs, etc...
3. First build the project, and run the following shell script:
   ```bash
   ./build.sh
   ```
4. Then the executable file will be created [git_info]
   ![Screenshot from 2023-11-28 20-10-54](https://github.com/acharyamanish006/Git_Info/assets/100832817/921a535e-5840-4c9f-af80-889c851fd6b7)
    ```bash
   ./git_info -u {git_username}

   ```
6. To get user followers use -F flag
   ![Screenshot from 2023-11-28 20-11-57](https://github.com/acharyamanish006/Git_Info/assets/100832817/d7c45555-7b60-4d19-a702-70aa547f983f)
   ```bash
   ./git_info -u {git_username} -F
   ```
   
8. To get the user following use -f flag
   ![Screenshot from 2023-11-28 20-11-35](https://github.com/acharyamanish006/Git_Info/assets/100832817/cd509196-c4d2-490d-abdc-0a03e9f81f53)
    ```bash
   ./git_info -u {git_username} -f
   ```
10. To get help
    ![Screenshot from 2023-11-28 20-11-21](https://github.com/acharyamanish006/Git_Info/assets/100832817/0f592eee-7764-4c03-8613-15e323eaeb92)
    ```bash
     ./git_info -h
      ```
11. To get user Contribution use -c flag
      
      ````bash
       ./git_info -u {git_username} -c
