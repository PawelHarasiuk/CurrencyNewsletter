# Currency Newsletter
You can access website here: <a href="http://frontent-bucket.s3-website.eu-central-1.amazonaws.com">Newsletter</a>
## Overview

This project allows users to submit or remove their email addresses through a frontend web app. Once subscribed, users receive a daily email with the latest currency exchange rates, fetched from an external API. The email is sent automatically every 24 hours.

## Architecture


## Technologies
- Go
- Docker
- AWS: Lambda, ECR, EventBridge, API Gateway, S3
- JavaScript/HTML/CSS

## To add / To improve
- better frontend - make it prettier and maybe add logging
- add feedback to user after submit
- choosing currencies to follow instead of fixed rate of USD and EUR
- database instead of file on s3
