# Notification service

Service on Gin

## Instalation

Instructions on how to get a copy of the project and running on your local machine.

```bash
https://github.com/ssqwet1324/notification-service.git
```
## Create a database using a file docker-compose.yml

Change the data for the database if necessary.

## Go to the file .env and sign your data for the database, mail and telegram token, SMTP, host, testEmail, chatId

![alt text](https://i.imgur.com/dBHmItk.png)

## Run the code and go to postman

Create a POST request, in the field with the request enter:

```bash
http://localhost:8080/notifications
```

![alt text](https://i.imgur.com/0P6tz4k.png)

If you need a port from 8080, change it to another one.

In the code field, enter the following:

```bash
{
    "id": 1,
    "message": "yourmessage"
}
```
![alt text](https://i.imgur.com/6eouPm7.png)

Done. Now click the send button and check your mail. To receive a message in Telegram, enter /start in the bot and send another request.







