# PT Bejana Investidata Globalindo (BIG IO)

## Made to Fullfill Backend Developer qualifications

I want to say thanks to all BIG IO teams that give me chance to join your Backend Engineering recruitment process. I will not dissapoint you.

## Tech Stacks

1. Go Programming Language -> [golang](https://go.dev/)
2. Echo HTTP Framework -> [echo](https://echo.labstack.com/)
3. GORM Object Relational Mapper -> [gorm](https://gorm.io/index.html)
4. PostgreSQL SQL-based database server-> [postgres](https://www.postgresql.org/)
5. Docker container -> [docker](https://www.docker.com/)
6. Nginx reverse proxy -> [nginx](https://www.nginx.com/)
7. Ubuntu Operating System -> [ubuntu](https://ubuntu.com/)
8. AWS EC2 Instance -> [aws](https://aws.amazon.com/)
9. Let's Encrypt CERBOT -> [certbot](https://certbot.eff.org/)

## What's provided by this app?

as defined in the task, below are the specs needed to be fulfilled (the list is written in Bahasa and unchanged).
1. Terdapat 3 role user, yaitu guru, murid, dan admin
2. Admin bertugas untuk memasukkan data guru dan siswa.
3. Guru bertugas untuk memasukkan nilai siswa yang diampunya.
4. Siswa dapat melihat data nilainya sendiri yang sudah diinputkan oleh guru

This app is already fullfil all the specs defined above. However, i've add some specs or utility that i think that may be needed

## How to run locally?

To run this app locally, first things first, you need to install some dependencies. The dependencies are:

1. PostgreSQL
    - I'm gonna assume you already have knowledge how to install PostgreSQL. So please take a look at `.env.example` file, on PostgreSQL configurations. There is some basic creds to connect to database. Make sure you fill them right. If you want the tutorial how to do so, i have a nice one [here]("https://blog.trusty.my.id/?p=997")
2. Go Programming Language
    - Please install Go Programming Language at version 1.18.3 on your system. Here is the useful link: [golang]("https://go.dev/")
3. Makefile
    - Please ensure that you able to run Makefile. [makefile]("https://opensource.com/article/18/8/what-how-makefile")
4. MODD
    - Please ensure that you already install MODD, [modd]("https://github.com/cortesi/modd")

After all the dependencies have been installed, the steps to run this apps are:

1. Clone this repository
2. Create `.env` file based on template defined on `.env.example` file
3. Install Go Module dependencies
    - Run: `go mod tidy`
4. Initialize database tables
    - Run: `make migrate`
5. Initialize Admin credentials
    - Run: `go run main.go init-admin {name} {password}`
    - Field `name` and `password` must be defined and will be used as admin credentials
6. Run the app
    - Run: `make run`
7. The server should be listening on port 5000 or other port that you define on your .env file

## Can i have the request collection?

Sure. You can use my request collection using Insomnia. The collection are available on file `request-collection-insomnia.json`

## I have a question, can i have your contact?

You can open an issue on this Github repository.

## I dont want to install this app by myself. Do you have a deployed server?

Sure! I do have skills in DevOps too, use this `https://bigio.trusty.my.id` as host value on your Insomnia request collection.
Also, the default Admin creds for the deployed server are:
```
id: 1663323245426115082
password: admin
```