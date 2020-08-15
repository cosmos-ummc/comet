# Submissions

[Link to the demonstration video](https://www.youtube.com/watch?v=gR71QjbmJvo)

[Link to the presentation slides](https://docs.google.com/presentation/d/1YxOG8uyPY6fy_Q7czGD5bJGpYhFRngDTRjdESdkzCLI/edit?usp=sharing)

[Link to the user manual](https://docs.google.com/document/d/1yM8Fy_c3kIlEPvMYso4nH1ZZVMo2aN-5uhWAhb0Vegk/edit?usp=sharing)

[Link to the demonstration video](https://quaranteams-admin.herokuapp.com/#/login)

[Link to the Psychologist / Staff / Admin Login Page](https://quaranteams-admin.herokuapp.com/#/login)

[Link to The Users(Client) Login Page](https://quaranteams-main.herokuapp.com/login) 

[Link to our humanized chat bot](https://t.me/Quaranteams_bot)

[Link to the DASS-21 Original Question Sets](https://journals.plos.org/plosone/article/file?type=supplementary&id=info:doi/10.1371/journal.pone.0219193.s004#:~:text=The%20Depression%2C%20Anxiety%20and%20Stress,into%20subscales%20with%20similar%20content.)

[Link to the IES-R Original Question Sets](https://www.aerztenetz-grafschaft.de/download/IES-R-englisch-5-stufig.pdf)

[Link to the Backend Main Service](https://github.com/cosmos-ummc/comet)

[Link to the Client Website](https://github.com/cosmos-ummc/mayall)

[Link to the Admin Dashboard](https://github.com/cosmos-ummc/butterfly)

[Link to the Backend Chat Service](https://github.com/cosmos-ummc/needle)

[Link to the Telegram Chat Bot](https://github.com/cosmos-ummc/Willman)

# Instructions

This repository is for the main backend services. It covers most of the logic and exposes the APIs for the client websites as well as chat bot to operate. 
The service is written in Golang, communicated via REST API that is routed to the gRPC handlers.
There are several modules included:

- Admin module

This module is used to manage the admins.

- User module

This module is used to manage the users. It includes the APIs for report submission from the chat bot, verification in the main client website, and all other
user associated features. It also provides several variants of queries to allow visualization of the graphs in the front end sites.

- Report module

This module is used to manage the user reports. The reports include the DASS, IES-R and daily reports. All of the reports received will then be processed
to calculate the scoring here based on the standard validated mental health scoring scheme.

- Meeting module

This module is to manage the meetings and their status. When a user created a new meeting, a random consultant will be selected based on availability automatically.

- Chat module
This module includes the management of anonymous user chat rooms and messages. It covers from the identification of similar users to the creation of chat room, messages event emitters,
and to the blocking features (Block users whom I don't like)

- Customized contents module
This module offers a dynamic selection of health feeds based on the users' mental status by analysing the reports submitted.

- Useful health information module
This module provides a set of useful health information for the users such as the health tips, health feeds, games and meditations. The bank of contents is managed
in the admin dashboard by the superuser.

- Scheduler module
This module toggles a scheduler to emit daily messages and reminders to the users.

# Project Setup

- To run the server, first, setup the [Golang](https://golang.org/) version 1.13.

- Install [MongoDB](https://www.mongodb.com/) database server. 

- Next, clone the repository.

- cd to the project directory and run `go build .` to install the dependencies.

- add the environment variables as follows:

```
MONGODB_URL=<MongoDB URL>
ACCESS_SECRET=<Your JWT access secret>
REFRESH_SECRET=<Your JWT refresh secret>>
BOT_TOKEN=<Your Telegram bot token>
```

- run `./main` to start the service locally.
