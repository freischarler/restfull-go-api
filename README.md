**OBTENER TODOS LOS USUARIOS**

curl http://localhost:9000/api/v1/users/

**CREAR USUARIO**

http://localhost:9000/api/v1/users/ *POST*

{
    "first_name": "Martin",
    "last_name": "Paz",
    "username": "freischarler",
    "email": "martin.paz@live.com.ar",
    "password": "123456"
}


**ACTUALIZAR USUARIO**

http://localhost:9000/api/v1/users/2 *PUT*

{
    "first_name": "Nicole",
    "last_name": "Mockert",
    "email": "NICOMO",
    "picture": "DASdsadasdsadsa"
}

  

**ELIMINAR USUARIO**

http://localhost:9000/api/v1/users/2 *DELETE*

**CREAR RECETA**

http://localhost:8000/api/v1/posts/ *POST*

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

{
    "recipe_name": "a",
    "recipe_type": a",
    "thumbnail": "https://dam.cocinafacil.com.mx/wp-content/uploads/2019/01/palomitas-con-chipotle-veganas.jpg",
    "ingredients": ["ab", "cd", "d""],
    "description": ["aaa", "bbb"],
    "user_id": 1
}

**LIKE RECETA**

http://localhost:8000/api/v1/posts/like/1 *PUT*

