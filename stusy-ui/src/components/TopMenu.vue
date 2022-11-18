<template>
  <ModalWindow v-if="modal">
    <form class="userData" v-on:submit.prevent="putUserData()">
      <label class="input-form">
        <p>Имя</p>
        <input
            v-model="userData.first_name"
            type="text"
            placeholder="Введите имя"
            required
        />
      </label>
      <label class="input-form">
        <p>Фамилия</p>
        <input
            v-model="userData.last_name"
            type="text"
            placeholder="Введите фамилию"
            required
        />
      </label>
      <div class="button-login">
        <input type="submit" value="Сохранить">
      </div>
    </form>
  </ModalWindow>
  <nav>
    <ul id="topMenu">
      <div class="search">
        <input type="search" placeholder="Введите поисковой запрос">
        <a href="#">
          <button type="submit" class="topMenu_icon">
            <img src="@/assets/search.svg" alt="Search">
          </button>
        </a>
      </div>
      <li class="menu-item">
        <a href="#">
          <span class="topMenu_icon">
            <img src="@/assets/bell.svg" alt="Bell">
          </span>
        </a>
      </li>
      <li class="menu-item">
        <a href="#">
              <span>
                <img src="@/assets/avatar.png" alt="Avatar"/>
              </span>
        </a>
      </li>
      <DropdownMenu class="menu-item" :title="``">
        <template v-slot:top-title>
          <a id="profile-title">
                <span>
                  {{
                    userData.first_name && userData.last_name ? `${userData.first_name} ${userData.last_name[0]}.` : "Профиль"
                  }}
                </span>
            <div class="avatar-arrow" v-bind:style="[menu?{'transform':'rotate(90deg)'}:{}]">
              <img src="@/assets/arrow.svg"/>
            </div>
          </a>
        </template>
        <template v-slot:item>
          <li class="sub-menu-item">
            <router-link to="/profile">Профиль</router-link>
          </li>
          <li class="sub-menu-item" v-on:click="logout"><a>Выйти</a></li>
        </template>
      </DropdownMenu>
    </ul>
  </nav>
</template>

<script setup>
import {getCookie, logout, url} from "@/global";
import DropdownMenu from "@/components/DropdownMenu";
import ModalWindow from "@/components/ModalWindow";
import {onMounted, ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter();

let menu = false;
let userData = ref({
  first_name: "",
  last_name: ""
});
let modal = ref(false);

onMounted(() => {
  getUserData();
});

function getUserData() {
  fetch(`${url}/users/${getCookie("ID")}`, {
    headers: {
      "Authorization": `Bearer ${getCookie("TOKEN")}`
    }
  }).then(response => {
    if (response.ok) return response.json();
    if (response.status === 401) {
      logout();
    }
    if (response.status === 404) modal.value = true;
  }).then(data => {
    if (data !== undefined) {
      userData.value.first_name = data.first_name;
      userData.value.last_name = data.last_name;
    }
  }).catch(err => {
    console.error("Cannot fetch", err);
  });
}

function putUserData() {
  fetch(`${url}/users/${getCookie("ID")}`, {
    method: "PUT",
    headers: {
      "Authorization": `Bearer ${getCookie("TOKEN")}`
    },
    body: JSON.stringify({
      "first_name": userData.value.first_name,
      "last_name": userData.value.last_name
    })
  }).then(response => {
    if (response.ok) return response.json();
    switch (response.status) {
      case 400:
        console.log("Неверные данные");
        break;
    }
  }).then(data => {
    router.push("/auth");
    console.log(data);
  }).catch(err => {
    console.error("Cannot fetch" + err);
  });
}

</script>

<style scoped>

#topMenu {
  display: flex;
  width: 100%;
  gap: 16px;
  flex-direction: row;
  justify-content: flex-end;
  list-style: none;
  margin: 0;
  padding: 0;
}

#topMenu li a {
  display: flex;
  color: var(--dark);
  align-items: center;
  justify-content: center;
  text-decoration: none;
  font-size: 18px;
  min-height: 70px;
  height: 100%;
  position: relative;
  white-space: nowrap;
}

#profile-title {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  gap: 5px;
}

.menu-item {
  font-style: normal;
  font-weight: 500;
  font-size: 18px;
  text-align: center;
  float: right;

}

.sub-menu-item {
  background-color: var(--light);
}

.sub-menu-item a {
  padding: 10px 0;
}

.sub-menu-item:last-child {
  border-bottom-right-radius: 10px;
  border-bottom-left-radius: 10px;
}

/*.sub-menu {*/
/*  position: absolute;*/
/*  top: 65px;*/
/*  left: 0;*/
/*  right: 0;*/
/*  z-index: 2;*/
/*  background: var(--light);*/
/*  border-bottom-right-radius: 10px;*/
/*  border-bottom-left-radius: 10px;*/
/*  overflow: hidden;*/
/*  box-shadow: #0000003b 6px 8px 9px 0px;*/
/*}*/

/*.sub-menu li {*/
/*  margin-right: 0 !important;*/
/*  padding: 10px 36px;*/
/*  cursor: pointer;*/
/*}*/

.sub-menu-item:hover {
  background: var(--light-mate);
}

#topMenu > li:last-child {
  background: var(--light);
  border-radius: 10px;
  padding: 0 36px;
  user-select: none;
}

.topMenu_icon {
  background: var(--light);
  border: none;
  border-radius: 15px;
  padding: 20px;
}

.search {
  position: relative;
  padding: 0;
  width: 50%;
}

.search button {
  position: relative;
  display: inline-block;
  cursor: pointer;
  z-index: 2;
  float: right;
}

input[type="search"] {
  position: absolute;
  height: 100%;
  width: 0;
  border: none;
  outline: none;
  top: 0;
  right: 0;
  z-index: 3;
  border-radius: 15px;
  background: none;
  padding: 1% 72px 1% 0;
  transition: width .4s cubic-bezier(0.000, 0.795, 0.000, 1.000);
  cursor: pointer;
  color: var(--dark);
  font-style: normal;
  font-weight: 400;
  font-size: 18px;
  line-height: 21px;
}

input[type="search"]:focus {
  width: 100%;
  z-index: 1;
  background-color: var(--light);
  padding: 1% 72px 1% 20px;
  cursor: text;
}

.userData {
  background: var(--light);
  padding: 12px;
  border-radius: 20px;
}

.userData p {
  color: #000000;
  font-size: 18px;
  font-weight: bold;
}


@media screen and (max-width: 1000px) {
  nav {
    width: 100%;
  }

  #topMenu {
    display: flex;
    flex-direction: row;
    float: right;
    gap: 16px;
    list-style: none;
    margin: 0;
    padding: 0;
    flex-wrap: wrap;
    justify-content: center;
    align-items: center;
  }

  .menu-item {
    text-align: right;
  }

  .search {
    position: relative;
    width: 100%;
  }

  input[type="search"] {
    position: absolute;
    height: 100%;
    width: 0;
    border: none;
    outline: none;
    top: 0;
    right: 0;
    z-index: 3;
    border-radius: 15px;
    background: none;
    padding: 1% 72px 1% 0;
    transition: width .4s cubic-bezier(0.000, 0.795, 0.000, 1.000);
    cursor: pointer;
    color: var(--dark);
    font-style: normal;
    font-weight: 400;
    font-size: 18px;
    line-height: 21px;
  }

  input[type="search"]:focus {
    width: 100%;
    z-index: 1;
    cursor: text;
  }

  .search button {
    position: relative;
    float: right;
    cursor: pointer;
    z-index: 2;
  }
}
</style>