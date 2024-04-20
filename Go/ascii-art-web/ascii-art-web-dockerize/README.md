# ascii-art-web-dockerize

**[ascii-art-web-dockerize official homepage](https://01.kood.tech/git/root/public/src/branch/master/subjects/ascii-art-web/dockerize)**  
**[ascii-art-web-dockerize official audit page](https://01.kood.tech/git/root/public/src/branch/master/subjects/ascii-art-web/dockerize/audit.md)**  
**[Ascii-art-web official homepage](https://01.kood.tech/git/root/public/src/branch/master/subjects/ascii-art-web)**  
**[Get Docker](https://docs.docker.com/get-docker/)**  

## Description

Our objective was to containerize the ASCII art web application using Docker technology. To achieve this, we crafted a Dockerfile that comprises all the necessary instructions for building an image. The resulting image can then be utilized to instantiate a container.  

## Usage

Prior to starting, it is essential to have Docker installed. To run the program, simply type "sh start.sh" into the terminal and press Enter. After that, open your browser and access the application at <http://localhost:8080>. It's that simple! When you're finished, press in terminal any key to stop the container, and the container and its related image files will be automatically removed.  

## Implementation details: algorithm

You may be curious about what we accomplished. We established a server using GO's ListenAndServe function. Then, we created the website, including an input form where you can enter text and an output page where you can view the ASCII art (HTML). We ensured that the server and website communicate effectively through POST and GET methods. Lastly, we included a feature that allows you to download your ASCII art. Throughout the process, we maintained a user-friendly and easy-to-navigate design through the use of CSS. Last but not least, we packaged our application into a Docker container (Dockerfile).  

## Authors

Created by: Jaan.Ginzul, rols55, \_Parker\_  
