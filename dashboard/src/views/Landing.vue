<template>
  <div class="landing">
    <header>
      <p href class="header_logo">Frens Co.</p>
      <div class="header_cta">
        <a href class="header_button">Services</a>
        <a href class="header_button">About</a>
        <router-link to="/login">
          <a href class="header_login">Log in</a>
        </router-link>
      </div>
    </header>
    <main>
      <section class="banner_section">
        <div class="section_container">
          <h1>Make your future brigther.</h1>
          <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
        </div>
      </section>
    </main>
    <footer></footer>
  </div>
</template>
<script>
export default {};
</script>
<style scoped>
.banner_section {
  width: 100%;
  height: 93vh;
}

.landing {
  background-size: cover;
  background-image: url("../assets/landing_cover_city_3.jpg");
}

.section_container {
  display: flex;
  flex-direction: column;
}

.section_container p {
  color: white;
  font-size: 2em;
  width: 50%;
  border-radius: 3px;
  margin: 1em auto 0 auto;
  font-weight: lighter;
  text-align: center;
  line-height: 40px;
}

.section_container h1 {
  margin: 3em auto 0 auto;
  background-color: rgba(0, 0, 0, 0.5);
  padding: 40px;
  font-size: 60px;
  font-weight: bolder;
  color: white;
}

header {
  display: flex;
  justify-content: space-between;
  height: 80px;
  background-color: rgba(0, 0, 0, 0.7);
  color: white;
}

a {
  padding: 30px 0;
}

.header_logo {
  padding: 30px 0;
  margin: 0 2em;
  color: white;
  text-align: left;
  font-size: 2em;
  font-weight: 800;
  display: block;
}

header .header_login {
  margin: 0 2em;
  text-align: center;
  color: white;
  font-size: 1em;
  font-weight: 800;
}

header .header_button {
  margin: 0 2em;
  text-align: center;
  color: white;
  font-size: 1em;
  font-weight: 800;
  display: block;
}

.header_cta {
  display: flex;
}
</style>