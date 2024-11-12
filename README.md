# Groupie Tracker

## Description
Groupie Tracker is a project aimed at receiving and manipulating data from a given API to create a user-friendly website displaying information about bands and artists. The project focuses on visualizing data through various methods such as blocks, cards, tables, lists, pages, and graphics.

## Authors
- [Elhadj Abdoul](https://github.com/abdoulcyf)
- [Reza](https://github.com/Reza-sd)
## Usage: How to Run
Dockerize the  [Groupie Tracker](https://learn.01founders.co/git/ediallo/groupie-tracker) application, by following these steps:
1. Make sure [Docker](https://docs.docker.com/engine/install/) is set on your machine.
1. Clone the repository: ```git@github.com:eediallo/groupie-tracker.git```
2. Navigate to the project directory: `cd groupie-tracker`
3. Build the Docker image: `docker build -t [image name] .`
4. Run the Docker container: `docker run -p 8080:8080 [image name]`
5. Access the application in your web browser at: `http://localhost:8080`

## Implementation Details: Algorithm
Groupie Tracker algorithm works as follows:
- The backend is implemented in Go, adhering to best practices and utilizing only standard Go packages.
- The site and server are designed to handle requests without crashing, ensuring stability.
- All pages are thoroughly tested and functional, with error handling implemented where necessary.
- The project focuses on data manipulation and storage, utilizing JSON files and formats for data representation.
- HTML is used for building the website's frontend.
- Event creation and display are key features, allowing interaction between the client and server.
- The client-server communication follows a request-response model, enabling seamless data exchange.

## Usage Example:
To run the project locally, follow the steps outlined in the Usage section. Once the server is running, you can access the website through your preferred web browser.
