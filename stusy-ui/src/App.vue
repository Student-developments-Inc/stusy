<template>
  <router-view/>
</template>

<script></script>

<style>
@import url("https://fonts.googleapis.com/css2?family=Ubuntu:ital,wght@0,300;0,400;0,500;0,700;1,300;1,400;1,500;1,700&display=swap");

* {
  margin: 0;
  padding: 0;
}

body {
  display: block;
  font-family: "Ubuntu", sans-serif;
  background: var(--light);

  --dark: #121212;
  --dark-mate: #292a2d;
  --dark-mate-mate: #dbdee8;

  --light: #fff;
  --light-mate: #dcfffd;

  --blue: #2255f4;
  --blue-mate: #1e42bb;

  overflow: hidden;
}

.input-form {
  margin: 10px 0;
  text-align: left;
}

.input-form input {
  margin: 10px 0;
  padding: 25px 25px;
  width: 100%;
  background: rgba(69, 70, 79, 0.1);
  font-size: 18px;
  border: none;
  border-radius: 15px;
  box-sizing: border-box;
  outline: none;
}

.button-login input[type='submit'] {
  color: #ffffff;
  background: var(--blue);
  border: none;
  border-radius: 15px;
  font-size: 18px;
  box-sizing: border-box;
  width: 100%;
  padding: 25px;
  transition: 0.3s;
  cursor: pointer;
}

.button-login input[type='submit']:hover {
  background: var(--blue-mate);
}
</style>
