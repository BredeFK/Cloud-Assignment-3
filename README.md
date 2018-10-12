# Assignment 3: Chat Bot, Currency Conversion, Docker

## Note:
[Fixer](http://fixer.io/) now require an API Access key, so this assignment can't return currency rates anymore.

## Authors
| Name                  | Github                                            | Email                                                 |
|-----------------------|---------------------------------------------------|-------------------------------------------------------|
| Brede Fritjof Klausen | [BredeFK](https://github.com/BredeFK)             | [bredefk@stud.ntnu.no](mailto:bredefk@stud.ntnu.no)   |
| Johan Aanesen         | [JohanAanesen](https://github.com/JohanAanesen)   | [johanaan@stud.ntnu.no](mailto:johanaan@stud.ntnu.no) |

## Packages needed:
```
go get github.com/bwmarrin/discordgo
go get gopkg.in/mgo.v2
```

## Root url:
https://cloudtech3.herokuapp.com/

## Instructions
This assignment, Assignment 3, has two aspects, and builds on the services developed in Assignment 2.  

**Aspect 1:** Develop a service that allows users to interact with a bot on a instant messaging system of your choice (e.g. Slack), and provide currency conversions. The bot will form an interactive user-interface that relies on the currency conversion service that you have built in Assignment 2.

Note that the messaging system needs to be able to support both incoming and outgoing webhooks without any additional layer of complexity (e.g. authentication). Check https://dialogflow.com/docs/integrations/ for details.

**Aspect 2:** Assignment 3 functionality must be deployed on mLab (for database functionality) and on Heroku (for computational functionality) and on Dialogflow (for the bot integration). However, you are also required to prepare Dockerfiles and docker-compose configuration (for your code & MongoDB), such that the solution to Assignment 3 could potentially be re-deployed on an alternative cloud provider that supports Docker containers, i.e. not relying on mLabs or on Heroku.

## Service specification
### Bot

The bot should be able to answer simple questions about the current currency conversion rates. Eg.
```
"What is the current exchange rate between Norwegian Kroner and Euro?"
"What is the exchange rate between USD and NOK?"
"What is the exchange rate between Euro and Norwegian kroner?"
```

### Response
The bot should respond in English, with the current exchange rates.

## Resources
- http://fixer.io/  Currency ticker data
- https://mlab.com/  MongoDB cloud hosting (choose the FREE plan)
- http://gopkg.in/mgo.v2 MongoDB Go driver
- Note: most likely we will use Heroku-based deployment, but, we need to test it first.
- https://dialogflow.com Tools and development console for developing the bot. The tutorial about Dialogflow will be posted online on BlackBoard.
- https://dialogflow.com/docs/integrations/ DialogFlow Integrations
