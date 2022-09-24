<template>
  <nav>
    <ul id="topMenu">
      <li>
        <div class="search">
          <input type="search" placeholder="Введите поисковой запрос">
          <a href="#">
            <button type="submit" class="topMenu_icon">
              <img src="@/assets/search.svg" alt="Search">
            </button>
          </a>
        </div>
      </li>
      <li>
        <a href="#">
          <span class="topMenu_icon">
            <img src="@/assets/bell.svg" alt="Bell">
          </span>
        </a>
      </li>
      <li>
        <a href="#">
              <span>
                <img src="@/assets/avatar.png" alt="Avatar"/>
              </span>
        </a>
      </li>
      <li v-on:click="menuAction">
        <a>
          <span>{{ userData.first_name && userData.last_name ? `${userData.first_name} ${userData.last_name[0]}.` : "Профиль" }}</span>
          <div class="avatar-arrow" v-bind:style="[menu?{'transform':'rotate(90deg)'}:{}]">
            <img src="@/assets/arrow.svg"/>
          </div>
        </a>
        <transition name="slide-fade">
          <ul class="sub-menu" v-if="menu">
            <li><a href="/profile">Профиль</a></li>
            <li v-on:click="logout"><a>Выйти</a></li>
          </ul>
        </transition>
      </li>
    </ul>
  </nav>
</template>

<script>
import {getCookie, logout, url} from "@/global";

export default {
  name: "TopMenu",
  data() {
    return {
      menu: false,
      userData: {
        first_name: "",
        last_name: ""
      }
    };
  },
  methods: {
    menuAction() {
      this.menu = !this.menu;
    },
    logout() {
      logout();
    },
    getUserData() {
      fetch(`${url}/users/${getCookie("ID")}`, {
        headers: {
          "Authorization": `Bearer ${getCookie("TOKEN")}`
        }
      }).then(response => {
        if (response.ok) return response.json();
        if (response.status === 401) {
          logout();
        }
        if (response.status === 404) this.modal = true;
      }).then(data => {
        if (data !== undefined) {
          this.userData.first_name = data.first_name;
          this.userData.last_name = data.last_name;
        }
      }).catch(err => {
        console.error("Cannot fetch", err);
      });
    }
  },
  mounted() {
    this.getUserData();
  }
};
</script>

<style scoped>
.search button {
  height: 100%;
  border: none;
  top: 0;
  right: 0;
  z-index: 2;
  cursor: pointer;
}

#topMenu {
  display: flex;
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
  height: 100%;
  position: relative;
  white-space: nowrap;
}

#topMenu li {
  font-style: normal;
  font-weight: 500;
  font-size: 18px;
  display: flex;
  margin-right: 16px;
  text-align: center;
}

.sub-menu {
  position: absolute;
  top: 65px;
  left: 0;
  right: 0;
  z-index: 2;
  background: var(--light);
  border-bottom-right-radius: 10px;
  border-bottom-left-radius: 10px;
  overflow: hidden;
  box-shadow: #0000003b 6px 8px 9px 0px;
}

.sub-menu li {
  margin-right: 0 !important;
  padding: 10px 36px;
  cursor: pointer;
}

.sub-menu li:hover {
  background: var(--light-mate);
}

#topMenu > li:last-child {
  background: var(--light);
  border-radius: 10px;
  padding: 0 36px;
  position: relative;
  user-select: none;
}

.avatar-arrow {
  margin-left: 7px;
  transition: .3s;
}

.topMenu_icon {
  background: var(--light);
  border: none;
  border-radius: 15px;
  padding: 20px;
}

.search {
  padding: 0;
  position: relative;
  display: inline-block;
}

input[type="search"] {
  position: absolute;
  display: inline-block;
  border-radius: 15px;
  border: none;
  outline: none;
  background: none;
  padding: 24px 100% 24px 0;
  width: 0;
  height: 100%;
  top: 0;
  right: 0;
  z-index: 3;
  transition: width .4s cubic-bezier(0.000, 0.795, 0.000, 1.000);
  cursor: pointer;
  color: var(--dark);
  font-style: normal;
  font-weight: 400;
  font-size: 18px;
  line-height: 21px;
}

input[type="search"]:focus {
  padding-left: 25px;
  width: 458px;
  z-index: 1;
  background-color: var(--light);
  cursor: text;
}

</style>