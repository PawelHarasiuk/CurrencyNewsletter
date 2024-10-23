# Currency Newsletter
You can access website here: <a href="http://frontent-bucket.s3-website.eu-central-1.amazonaws.com">Newsletter</a>

## Overview
This project allows users to submit or remove their email addresses through a frontend web app. Once subscribed, users receive a daily email with the latest currency exchange rates, fetched from an external API. The email is sent automatically every 24 hours.

## Technologies
- Go - used to write newsletter and to write handling request from frontend
- Docker - used docker image to deploy lambda on aws
- AWS
  - Lambda - sending mails and updating s3 file with go on backend
  - ECR - storing docker images
  - EventBridge - scheduling lambda triggering that sends mails
  - API Gateway - triggering request from frontend to Lambda
  - S3 - used to host static website and for storing file with emails
- JavaScript/HTML/CSS - basic frontend 

## Architecture
Here is architecture of the system I have designed. It not perfect and maybe some services I used not how they supposed to be used, but I enjoyed the proccess of designing and learning experience, and developed deeper understanding of used services.

<br>
<img width="479" alt="Zrzut ekranu 2024-10-23 o 19 03 36" src="https://github.com/user-attachments/assets/4aad11dd-e4c0-42de-a308-ba9d3ac072a5">
<br>

## Email Example 
Here is example email received by newsletter 

<br>
<img width="479" alt="Zrzut ekranu 2024-10-23 o 19 02 37" src="https://github.com/user-attachments/assets/fa250932-7d20-426d-8329-943d44c0ea3c">
<br>

## To add / To improve
- better frontend - make it prettier and maybe add authentication
- add feedback to user after submit
- choosing currencies to follow instead of fixed rate of USD and EUR
- database instead of file on s3 (I am aware that using file as db is not smart)
