<template>
  <div class="homepage">
    <AsideMenu/>
    <div class="home">
      <div class="content">
        <header>
          <h1>{{ localeHours }}{{ name }}</h1>
          <TopMenu/>
        </header>
        <router-view/>
      </div>
    </div>
  </div>
</template>

<script setup>
import AsideMenu from "@/components/AsideMenu";
import TopMenu from "@/components/TopMenu";
import {getCookie, logout, url} from "@/global";
import {computed, onMounted, ref} from "vue";

onMounted(() => {
  getUserData();
})

const userData = ref(null)

const localeHours = computed(() => {
  let localeHours = new Date().getHours();
  if (localeHours > 3 && localeHours < 12) return "Доброе утро";
  else if (localeHours > 11 && localeHours < 19) return "Добрый день";
  else if (localeHours > 18 && localeHours < 24) return "Добрый вечер";
  else if (localeHours > 23 || localeHours < 4) return "Привет";
  throw Error('localHours error')
})

const name = computed(() => {
  if (userData.value === null) {
    return '';
  }
  return `, ${userData.value.first_name}`;
})

function getUserData() {
  if (getCookie("ID") === undefined) logout();
  fetch(`${url}/users/${getCookie("ID")}`, {
    headers: {
      "Authorization": `Bearer ${getCookie("TOKEN")}`
    }
  }).then(response => {
    if (response.ok) return response.json();
    if (response.status === 401) {
      logout();
    }
  }).then(data => {
    if (data !== undefined) {
      userData.value = data;
    }
  }).catch(err => {
    console.error("Cannot fetch", err);
  });
}

</script>

<style scoped>
.homepage {
  display: flex !important;
}

.home {
  flex-grow: 1;
  background: var(--dark-mate-mate);
}

.content header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.content header h1 {
  font-weight: 500;
  font-size: 36px;
  margin: 0;
}

.content {
  padding: 15px 24px 24px;
  height: 100vh;
  overflow: auto;
}

header {
  margin: 12px 0;
}
</style>