# OG2 Coding Challenge

A few notes about my implementation. The challenge originally provided 4 hours for implementing. The version that can be seen in this branch has been refactored further after initial publish. You can find the challenge itself below.

## Idea
The idea behind my implementation is to store information about upgrades of factories as an array of time.Time values which is the time the level has been reached.
From this information the amount of ores that are produced since the user was created can be calculated.
If an upgrade is undergoing can be determined by checking the array values - if there is a value that is in the future it means the upgrade is underway. If all values are in the past it means nothing is being upgraded at the moment.

For the ores I am persisting only the spending of each ore. When the amount produced by the factory is calculated this amount is subtracted which gives the total amount of each ore for the user.

## Implementation
The app is structured in two containers, one for the app itself and one for the DB. 

MongoDB provides persistence, and I've used it so I don't have to handle migrations. Additionally, I've touched upon it before. Also, I had boilerplate code to use and departing from Mongo didn't provide enough value. If working from scratch key-value store would be preferable.

For debugging purposes Mongo Express is also part of the app, deployed as a separate container and accessible on http://localhost:8081.

To attach debugger connect it on port 2345. If you are running on VSCode the debugger configuration is preconfigured, just hit F5 when the app is up.

API endpoints are exposed on http://localhost:5000 and the documentation is exposed via Swagger on http://localhost:5000/swagger/index.html.

To start the app 'cd' to /og2-coding-challenge/src folder and run 'docker compose up'.


# The app requirements

You’ll be building the base for a game that we’ll call OG2. In OG2, the player has three resources (iron, copper, and gold) that they mine through factories. The goal is to increase mining capacity by improving the factories with their extracted resources.

## **Resources and factories**

You have 3 resources: Iron, copper, and gold. You also have three types of factories: Iron, copper, and gold factories (shocking, I know). Factories mine a certain amount of resources per second. They start at level 1 and the user can upgrade them by spending resources. Upgrades increase the resources mined per second. However, upgrading the factories is not an instant process. It takes a certain number of seconds based on the level.

**In this project, the priority is to get factories and resources management working.** Including:

1. Users **should be able to see the level of its factories** 
2. **The resources should accumulate over time**
3. Users **should be able to see** the **production rate per factory and how much the next upgrade would cost**
4. **Users can upgrade their factories** (assuming another update isn’t in progress and they have enough resources). While the upgrade is in progress, it continues producing resources.

## Game interface

This game works exclusively via a **JSON** API. When the program is started for the first time, it should wait until a user is created via the /user endpoint. As soon as a new user is added, the factories and resources should be initialized and accumulated based on the passage of time and the upgrades performed by the user.

You can think of the program you are creating as the backend of the game. Here are the endpoints that we expect:

### POST **/user**

Creates a new user in the game.

There are no requirements in terms of authentication or session management. Feel free to make it as simple as possible by, for example, passing a username in the payload.

### **GET /dashboard**

Displays the number of resources (iron, copper, and gold) a user has. 

Also displays information regarding the factories:

- The level
- The production rate
- If an upgrade is in progress and, if applicable, how long until it’s finished
- If an upgrade isn’t in progress, the cost of the next level

### POST **/upgrade**

Requests an upgrade for a specific factory type (iron, copper, or gold).

## Guidance

We have voluntarily omitted a few elements from this document. For example, the /upgrade POST request doesn’t mention what request payload is expected. We also haven’t given examples of response payloads. This is done to give you flexibility and see how you design APIs.

Regarding user creation and management, we don’t expect any authentication layer or session management. All we are testing is the functionality of the code in a multi-user environment. The only requirement is an endpoint that adds a new user to the game.

It’s important that the application be able to restart without losing any users’ data. For this reason, you should persist users, resources, and factory data. It’s up to you how to proceed, you can use SQL, a KV store, or anything else.

## Factory stats

### Iron factory

| Level | Production | Next upgrade duration | Upgrade cost |
| --- | --- | --- | --- |
| 1 | 10/s | 15s | 300 iron, 100 copper, 1 gold |
| 2 | 20/s | 30s | 800 iron, 250 copper, 2 gold |
| 3 | 40/s | 60s | 1600 iron, 500 copper, 4 gold |
| 4 | 80/s | 90s | 3000 iron, 1000 copper, 8 gold |
| 5 | 150/s | 120s |  |

### Copper factory

| Level | Production | Next upgrade duration | Upgrade cost |
| --- | --- | --- | --- |
| 1 | 3/s | 15s | 200 iron, 70 copper |
| 2 | 7/s | 30s | 400 iron, 150 copper |
| 3 | 14/s | 60s | 800 iron, 300 copper |
| 4 | 30/s | 90s | 1600 iron, 600 copper |
| 5 | 60/s | 120s |  |

### Gold factory

| Level | Production | Next upgrade duration | Upgrade cost |
| --- | --- | --- | --- |
| 1 | 2/m | 15s | 100 copper, 2 gold |
| 2 | 3/m | 30s | 200 copper, 4 gold |
| 3 | 4/m | 60s | 400 copper, 8 gold |
| 4 | 6/m | 90s | 800 copper, 16 gold |
| 5 | 8/m | 120s |  |

---

Good luck 🙌