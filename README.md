# docker-postech-desafio
### Para executar o projeto basta digitar
- docker-compose -d

### Para testar basta executar o seguinte comando curl
- curl --request GET --url http://localhost:8080/languages

### O resultado esperado é o json abaixo
```[
  {
    "id": "1",
    "name": "Java"
  },
  {
    "id": "2",
    "name": "Kotlin"
  },
  {
    "id": "3",
    "name": "Groovy"
  },
  {
    "id": "4",
    "name": "Python"
  },
  {
    "id": "5",
    "name": "Golang"
  }
]
```
