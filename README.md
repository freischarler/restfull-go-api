Autor: Martin Omar Paz - 2022
# Introduction
This proyect is for no-country organization.
Reference: https://dev.to/orlmonteverde/api-rest-con-go-golang-y-postgresql-m0o

## Table of contents
* [Requeriments](#requeriments)
* [Run](#run)
* [Contributing](#contributing)
* [Help](#help)

## Requeriments
```sh
postgresql
go
```

## Run
Para correr la aplicación es necesario ejecutar el siguiente comando:
```sh
go run main.go
```
## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Help

**DB CONFIG EXAMPLE**

DB_HOST="127.0.0.1"
DB_PORT=5432
DB_USER="postgres"
DB_PASSWORD="root"
DB_NAME="no-country"

**LOGIN**

`POST`
```
http://localhost:9000/api/v1/users/login/
```

{
    "username": "freis",
    "password": "123",
}

It return:
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.Qi0Tc-jTChzascHaZhl0e6rRaCvAS6OJ8RLsI8Y-R78"
}
```

Use it as
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.Qi0Tc-jTChzascHaZhl0e6rRaCvAS6OJ8RLsI8Y-R22'
```

**GET ALL USERS**

`GET`
```
http://localhost:9000/api/v1/users/ 
```

**ADD USER**

`POST`
```
http://localhost:9000/api/v1/users/ 
```

{
    "first_name": "Martin",
    "last_name": "Paz",
    "username": "freis",
    "email": "martin.paz@live.com",
    "password": "123456"
}

**UPDATE USER**

`PUT`
```
http://localhost:9000/api/v1/users/{id} 
```
example: http://localhost:9000/api/v1/users/1

{
    "first_name": "Nicole",
    "last_name": "Mockert",
    "email": "nic@gmai.com",
    "picture": "DASdsadasdsadsa"
}

**DELETE USER**

`DELETE`
```
http://localhost:9000/api/v1/users/{id}
```

**ADD RECIPE**

`POST`
```
http://localhost:8000/api/v1/posts/ *POST*
```


{
    "recipe_name": "g",
    "recipe_type": "g",
    "thumbnail": "g",
    "ingredients": ["g", "a"],
    "description": ["g.","a."],
    "user_id": 2
}


**LIKE RECIPE**

`PUT`
```
http://localhost:8000/api/v1/posts/like/{id}
```

**TOP 10 RECIPE's**

`POST`
```
http://localhost:8000/api/v1/posts/rank *POST*
```

**COMMENT**

`POST`
```
http://localhost:8000/api/v1/comment/{id}
```

{
    "recipe_id": 1,
    "name": "troll",
    "comment": "horrible"
}

**GET ALL RECIPE's COMMENTS**

`GET`
```
http://localhost:8000/api/v1/comment/{id} *GET*
```

**DELETE COMMENT**

`DELETE`
```
http://localhost:8000/api/v1/comment/{id} *DELETE*
```

**RECIPE EXAMPLES**

{
    "recipe_name": "Aros de Cebolla y Calabazas",
    "recipe_type": "Vegana",
    "thumbnail": "https://dam.cocinafacil.com.mx/wp-content/uploads/2019/01/calabazas-con-aros-de-cebolla.jpg",
    "ingredients": ["Calabacita", "Cebolla", "Harina","Queso Parmesano","Aceite"],
    "description": ["Asa 2 calabacitas a la plancha con 2 cucharadas de aceite de oliva. Reserva.","De una de harina, guarda dos cucharadas de harina y enharina 1 manojo de aros de cebolla en ellas. Mezcla el resto con 3/4 de taza de cerveza y salpimienta. Pasa los aros por la mezcla y fríelos en aceite caliente hasta que doren y dejen de salir burbujas en el aceite.","Sírvelos sobre rebanadas de calabaza y espolvorea 1/2 taza de queso parmesano sobre ellas."],
    "user_id": 1
}

{
    "recipe_name": "Trufas de Chocolate Vegana",
    "recipe_type": "Vegana",
    "thumbnail": "https://dam.cocinafacil.com.mx/wp-content/uploads/2019/01/bolitas-de-chocolate-veganas.jpg",
    "ingredients": ["Datil", "Cacao", "Pepitas de calabazas tostadas y peladas","Queso Miel"],
    "description": ["Coloca 1 taza de datiles sin hueso, 4 cucharadas de cacao, 1 taza de pepitas y 1 taza de miel en el procesador de alimentos. Tritura hasta formar una masa uniforme.","Forma bolitas con la mezcla y mantenlas en refrigeración hasta servir."],
    "user_id": 1
}

{
    "recipe_name": "Palomitas de maiz con chipotle",
    "recipe_type": "Vegana",
    "thumbnail": "https://dam.cocinafacil.com.mx/wp-content/uploads/2019/01/palomitas-con-chipotle-veganas.jpg",
    "ingredients": ["Mantequilla vegetal", "Sal", "Maiz palomero","Chipotle en polvo"],
    "description": ["Coloca en una olla mediana 2 cucharadas de mantequilla con la sal y 1 taza de granos de maíz, tapa y caliente a fuego alto hasta que se abran todos los granos.","Retira a un tazón, agrega 1 cucharadita de chipotle en polvo y enfría algunos minutos. Sirve caliente."],
    "user_id": 1
}




