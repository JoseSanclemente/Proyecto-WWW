<template>
    <div class="page">
        <div class="container">
            <div class="category-container">
                <h1>Categorías</h1>
                <ul>
                    <li>
                        <router-link to="/admin/add-category" class="back">Agregar</router-link>
                    </li>
                    <li>
                        <router-link to="/admin/search-category" class="back">Consultar</router-link>
                    </li>
                </ul>
            </div>

            <div class="category-container">
                <h1>Subcategorías</h1>
                <ul>
                    <li>
                        <router-link to="/admin/add-subcategory" class="back">Agregar</router-link>
                    </li>
                    <li>
                        <router-link to="/admin/search-subcategory" class="back">Consultar</router-link>
                    </li>
                </ul>
            </div>

            <div class="category-container">
                <h1>Productos</h1>
                <ul>
                    <li>
                        <router-link to="/admin/add-products" class="back">Agregar</router-link>
                    </li>
                    <li>
                        <router-link to="/admin/search-products" class="back">Consultar</router-link>
                    </li>
                </ul>
            </div>

            <div class="category-container">
                <h1>Descuentos</h1>
                <ul>
                    <li>
                        <router-link to="/admin/add-discount" class="back">Agregar</router-link>
                    </li>
                    <li>
                        <router-link to="/admin/search-discount" class="back">Consultar</router-link>
                    </li>
                </ul>
            </div>

            <div class="category-container">
                <h1>Usuarios</h1>
                <ul>
                    <li>
                        <router-link to="/admin/create-admin" class="back">Agregar</router-link>
                    </li>
                    <li>
                        <router-link to="/admin/search-user" class="back">Consultar</router-link>
                    </li>
                </ul>
            </div>

            <div class="category-container">
                <h1>Reportes</h1>
                <ul>
                    <li>
                        <router-link to="/admin/report" class="back">Consultar</router-link>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>

<style scoped>
.page {
    height: 100vh;
    border: 0.5px solid black;
    background-image: url("../../assets/images/background_palm_tree.jpeg");
}

h1 {
    margin: 10px;
}

ul{
    margin: 10px;
}

.container{
    display: flex;
    flex-direction: row;
    flex-direction: row;
    flex-wrap: wrap;
    width: 90vw;
    height: 33vh;
    margin: 50px auto;
    border-radius: 10px;
}

.category-container {
    color: white;
  background: rgba(0, 0, 0, 0.7);
  border: 2px solid rgb(247, 243, 230);
    display: flex;
    flex-direction: column;
    align-content: center;
    align-items: center;
    margin: 10px;
    border-radius: 10px;
    width: 20vw;
    height: 30vh;
}
</style>
