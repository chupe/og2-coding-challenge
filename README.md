# Go interview (asynchronous)

<aside>
💡 A minimalistic version of this project can be built within the time you have, but a well-rounded one requires a lot of work. We understand you do not have enough time to build something great, so don’t worry and do your best. Good luck!

If something happens and you can’t get a working version, please still send us your work. We know problems happen, especially during stressful times. We’ll take that into account when reviewing your work.

</aside>

You’ll be building the base for a game that we’ll call OG2. In OG2, the player has three resources (iron, copper, and gold) that they mine through factories. The goal is to increase mining capacity by improving the factories with their extracted resources.

## **Ressources and factories**

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