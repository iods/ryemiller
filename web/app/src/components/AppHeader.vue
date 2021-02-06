<template>

  <header>
    <img alt="Rye Miller" src="../assets/logo.png" height="50" />
    <p>Rye Miller :: Developer</p>

    <nav>
      <ul>
        <li><router-link to="/">Home</router-link></li>
        <li><router-link to="/curriculum-vitae">Curriculum Vitae</router-link></li>
        <li><router-link to="/health">Health Dashboard</router-link></li>
      </ul>
    </nav>
  </header>

</template>


<style type="text/css">

header {
  border-bottom: 1px solid #DDD;
  display: flex;
  padding: .5rem 1rem;
}

header p {
  margin-left: 1rem;
}

nav {
  margin-left: auto;
}

nav ul {
  list-style: none;
}

nav ul li {
  display: inline-flex;
  margin-left: 1rem;
}


</style>