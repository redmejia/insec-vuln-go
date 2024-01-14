# insec-vuln-go

# En

Insec is a web application created for the purpose of learning. This is version number one and is designed for students who are starting their career in cyber security, networking security or fullstack development.
For version number one you can practice vulnerabilities such as sql injection and Cross Site Request Forgery (CSRF).

Note: Insec app has been intentionally developed with these vulnerabilities for learning purposes. With this I want to show that there is no bad programming language. I want to emphasize that errors can occur if code is written incorrectly and this can cause big problems if it is deployed in production. I encourage not only to practice these vulnerabilities but also to propose solutions and try to repair these errors.

# Es

Insec es un web aplicación creada con el motivo de aprendizaje. Esta es la version numero uno y  esta disenada para estudiantes que estan empezando su carrera en seguridata informatica o desarrollo fullstack.
Para la version numero uno puedes practicar vulnerabilidades como sql injection and Cross Site Request Forgery (CSRF). 

Nota: Insec app hacido desarrollada intencionalmente con estas vulnerabilidades con el fin de aprendizaje. Con esto quiero mostrar que no hay lenguaje de programacion malo yo quiero hacer enficis de que errores pueden ocurrir si se escribe codigo incorrectamente y esto puede causar grandes problamas si se desplieaga en producion. Yo insentivo no solo a pracricar estas vulnerabilidades si no a tambien proproner solucines y tratar de reparar estos errores.

Requirement 
- Go(Golang) install instructin -> https://go.dev/doc/install
- Docker Install documentation -> https://docs.docker.com/engine/install/
- Reactjs -> https://react.dev


Build and Run
1. Build and run container 
```bash
    make build_up
```
Restart 
2. Restart continers

```bash
    make restart
```

Database
- User
    postgres
- Password
    password
- Database Name 
    db-vulngo

To persist data on your local machine create a directory on your root application

```bash
    mkdir -p vuln-go/data 
```



