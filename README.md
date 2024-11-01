# API Rate Limiting Server

This project implements an HTTP server with rate limiting. It ensures no more than 5 requests per second can be processed.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Endpoints](#endpoints)
- [Usage](#usage)
- [License](#license)

## Introduction
The API rate limiting server limits the number of requests per second to prevent abuse and manage server load effectively. It provides detailed logging and health check endpoints to monitor server status.

## Features
- Rate limiting to 5 requests per second
- Health check endpoint to monitor server status
- Detailed request logging

## Installation
1. Clone the repository:

   git clone https://github.com/yourusername/api-rate-limiting-server.git
   cd api-rate-limiting-server
   
2. Install dependencies:
   go mod tidy

## Endpoints
1. /: Default route returns a "stupid request" message with status 400.

2. /healthcheck: Endpoint to check the health status of the server. Returns "Health Passed" if the /healthcheck file exists, otherwise "Health Failed".

3. /api/v1/product: Example endpoint for the product API. Returns "This is a Product API" if the /api/v1/product file exists, otherwise "There is no Product API".

## Usage
1. Run the API rate limiting server:
   go run main.go
2. Access the server:
   curl -w %{http_code}http://localhost:8080
   



## License
This project is licensed under the MIT License. See the LICENSE file for details.

This `README.md` should cover the current state of your project, focusing on the API rate limiting aspect. Ready to roll? 🚀
```markdown

This should cover everything your project entails. You can add, modify, or remove sections as needed. 🚀
