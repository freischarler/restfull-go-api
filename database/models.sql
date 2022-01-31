use pokeworld

db.createCollection("pokemons")

db.pokemons.insertMany([
    { name: "Squirtle", type: "Water", health: 44 },
    { name: "Poliwag", type: "Water", health: 40 },
    { name: "Bulbasaur", type: "Grass", health: 45 },
    { name: "Pidgey", type: "Flying", health: 40 }
])