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
}
</style>
