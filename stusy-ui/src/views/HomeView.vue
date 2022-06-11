<template>
  <ModalWindow v-if="modal">
    <form v-on:submit="putUserData()">
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
  <div class="home">
    <aside id="menuVertical">
      <h1>StudentSystem</h1>
      <ol>
        <li id="li_1"><a href="#">
          <div class="LeftMenu_icon">
            <img src="@/assets/home.svg">
          </div>
          <span style="color: var(--light)">Домашняя страница</span></a></li>
        <li><a href="#">
          <div class="LeftMenu_icon">
            <img src="@/assets/book.svg">
          </div>
          <span>Мои курсы</span></a></li>
        <li><a href="#">
          <div class="LeftMenu_icon">
            <img src="@/assets/calendar.svg">
          </div>
          <span>Расписание</span></a></li>
      </ol>
    </aside>
    <div class="content">
      <header>
        <h1>Доброе утро{{ userData.first_name && userData.last_name ? `, ${userData.first_name}` : "" }}</h1>
        <nav>
          <ul id="topMenu">
            <li>
              <a href="#">
              <span class="topMenu_icon">
                <img src="@/assets/search.svg" alt="Search">
              </span>
              </a>
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
            <li>
              <a v-on:click="menuAction">
                <span>{{
                    userData.first_name && userData.last_name ? `${userData.first_name} ${userData.last_name[0]}.` : "Меню"
                  }}</span>
                <div class="avatar-arrow" v-bind:style="[menu?{'transform':'rotate(90deg)'}:{}]">
                  <img src="@/assets/arrow.svg"/>
                </div>
              </a>
              <transition name="slide-fade">
                <ul class="sub-menu" v-if="menu">
                  <li><a href="#">Профиль</a></li>
                  <li v-on:click="logout()"><a href="#">Выйти</a></li>
                </ul>
              </transition>
            </li>
          </ul>
        </nav>
      </header>

      <div class="mainContent">
        <WeatherWidget/>
        <div class="block" id="schedule">
          <h1>Понедельник, 16.05</h1>
          <div class="block-schedule">
            <div>
              <ul class="block-schedule_topLeft-content">
                <li>8:40–10:15</li>
                <li>лекц.</li>
                <li>3-612</li>
              </ul>
              <ul class="block-schedule_topRight-content">
                <li class="right-post">доцент</li>
                <li>Смирин О.А.</li>
              </ul>
            </div>
            <p>Теория вероятностей. Математическая статистика.</p>
          </div>
          <div class="block-schedule">
            <div>
              <ul class="block-schedule_topLeft-content">
                <li>8:40–10:15</li>
                <li>лекц.</li>
                <li>3-612</li>
              </ul>
              <ul class="block-schedule_topRight-content">
                <li class="right-post">доцент</li>
                <li>Смирин О.А.</li>
              </ul>
            </div>
            <p>Теория вероятностей. Математическая статистика.</p>
          </div>
          <div class="block-schedule">
            <div>
              <ul class="block-schedule_topLeft-content">
                <li>8:40–10:15</li>
                <li>лекц.</li>
                <li>3-612</li>
              </ul>
              <ul class="block-schedule_topRight-content">
                <li class="right-post">доцент</li>
                <li>Смирин О.А.</li>
              </ul>
            </div>
            <p>Теория вероятностей. Математическая статистика.</p>
          </div>
          <div class="block-schedule">
            <div>
              <ul class="block-schedule_topLeft-content">
                <li>8:40–10:15</li>
                <li>лекц.</li>
                <li>3-612</li>
              </ul>
              <ul class="block-schedule_topRight-content">
                <li class="right-post">доцент</li>
                <li>Смирин О.А.</li>
              </ul>
            </div>
            <p>Теория вероятностей. Математическая статистика.</p>
          </div>
        </div>
        <div class="block" id="courses">
          <div>
            <h1>Ваши курсы</h1>
            <a href="#">Смотреть все</a>
          </div>
          <div class="block-courses">
            <div>
              <img src="@/assets/data_archiving_and_compression.svg">
              <p>Архивация и сжатие данных</p>
            </div>
            <div>
              <img src="@/assets/botany_and_plant_physiology.svg">
              <p>Ботаника и физиология растений</p>
            </div>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>
<script>
import {url, getCookie} from "@/global";
import WeatherWidget from "@/components/WeatherWidget";
import ModalWindow from "@/components/ModalWindow";

export default {
  name: "HomeView",
  components: {WeatherWidget, ModalWindow},
  data() {
    return {
      menu: false,
      userData: {
        first_name: "",
        last_name: ""
      },
      modal: false
    };
  },
  methods: {
    menuAction() {
      this.menu = !this.menu;
    },
    getUserData() {
      fetch(`${url}/users/${getCookie("ID")}`, {
        headers: {
          "Authorization": `Bearer ${getCookie("TOKEN")}`
        }
      }).then(response => {
        if (response.ok) return response.json();
        if (response.status === 401) {
          this.logout();
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
    },
    putUserData() {
      fetch(`${url}/users/${getCookie("ID")}`, {
        method: "PUT",
        headers: {
          "Authorization": `Bearer ${getCookie("TOKEN")}`
        },
        body: JSON.stringify({
          "first_name": this.userData.first_name,
          "last_name": this.userData.last_name
        })
      }).then(response => {
        if (response.ok) return response.json();
      }).then(data => {
        console.log(data);
      }).catch(err => {
        console.error("Cannot fetch", err);
      });
    },
    logout() {
      document.cookie = "TOKEN=null;max-age=0";
      document.cookie = "ID=null;max-age=0";
      this.$router.push("/auth");
    }
  },
  mounted() {
    if (getCookie("ID") !== undefined && this.userData.first_name === "") this.getUserData();

    if (!getCookie("TOKEN")) {
      this.$router.push("/auth");
    }
  }
};
</script>

<style scoped>

.slide-fade-leave-active {
  transition: all .3s ease;
}

.slide-fade-leave-to {
  transform: translateY(-10px);
  opacity: 0;
}

.home {
  display: flex !important;
  min-height: 100vh;
  background: var(--dark-mate-mate);
}

#menuVertical {
  display: block;
  text-align: center;
  padding: 39px 24px;
  width: 300px;
  min-height: 100%;
  box-sizing: border-box;
  border-radius: 0 25px 0 0;
  background: var(--light);
}

#menuVertical h1 {
  font-style: normal;
  font-weight: 500;
  font-size: 28px;
  color: var(--blue);
  margin-top: 0;
  margin-bottom: 30px;
}

#menuVertical ol {
  display: block;
  margin: 0;
  padding: 0;
  list-style: none;
  position: relative;
}

#menuVertical ol li {
  display: list-item;
  box-sizing: border-box;
  padding: 16px 21px;
  width: 252px;
  height: 54px;
  border-radius: 15px;
}

#li_1 {
  background: var(--blue);
}

#menuVertical ol li a {
  display: flex;
  color: var(--dark);
  text-decoration: none;
  font-size: 18px;
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
  border-radius: 15px;
  padding: 20px;
}

.LeftMenu_icon {
  padding: 0 13px 0 0;
}

.content {
  padding: 15px 24px 24px;
  height: 100vh;
  overflow: auto;
}

.mainContent {
  display: flex;
  width: 100%;
  margin-top: 12px;
  flex-wrap: wrap;
  align-items: flex-start;
  justify-content: space-evenly;
  margin-bottom: 30px;
}

.block {
  position: relative;
  overflow: hidden;
  padding: 38px 39px;
  background: var(--light);
  border-radius: 25px;
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;
  margin: 0 10px 10px;
}

.block-courses {
  display: flex;
  justify-content: space-between;
  margin: 14px -5px 0 0;
}

.block-courses div {
  flex: 1 1;
  margin: 0 5px;
  padding: 21px 10px;
  border: 1px solid var(--blue);
  border-radius: 25px;
}

.block-courses div img {
  display: block;
  margin: auto;
}

.block-courses div p {
  margin: auto;
  max-width: 258px;
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  text-align: center;
}

#courses {
  height: 251px;
  padding: 18px 41px 23px;
  display: flex;
  flex-direction: column;
}

#courses div h1 {
  margin: 0;
  font-style: normal;
  font-weight: 500;
  font-size: 24px;
  float: left;
}

#courses div a {
  float: right;
  font-style: normal;
  font-weight: 400;
  font-size: 18px;
  color: var(--blue);
  text-decoration: none;
  text-align: right;
}

#schedule {
  display: flex;

  flex-direction: column;
  padding: 18px 23px;
}

#schedule h1 {
  margin: 0;
  font-style: normal;
  font-weight: 500;
  font-size: 24px;
  color: var(--dark);
}

.block-schedule {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 7px 10px;
  margin-top: 15px;
  margin-bottom: -5px;
  box-shadow: 0 0 4px rgba(0, 0, 0, 0.25);
  border-radius: 15px;
}

.block-schedule div {
  display: inline-block;
  width: 100%;
  flex: 1 1 auto;
}

.block-schedule_topLeft-content, .block-schedule_topRight-content {
  display: flex;
  align-items: center;
  margin: 0;
  float: left;
  padding: 0;
  list-style: none;
}

.block-schedule_topLeft-content {
  float: left;
}

.block-schedule_topLeft-content li {
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  padding: 4px 20px;
  font-style: normal;
  font-weight: 400;
  font-size: 14px;
  line-height: 16px;
  margin: 0 8px 0 0;
  border: 1px solid var(--dark);
  border-radius: 25px;
}

.block-schedule_topRight-content {
  float: right;
}

.right-post {
  text-align: left;
  margin: 0 -20px 0 0;
  padding: 4px 35px 4px 20px;
}

.block-schedule_topRight-content li {
  flex: 1 1 auto;
  margin: 0 0 0 4px;
  text-align: center;
  padding: 4px 20px;
  font-style: normal;
  font-weight: 400;
  font-size: 14px;
  line-height: 16px;
  box-shadow: 0 0 4px rgba(0, 0, 0, 0.25);
  border-radius: 25px;
  background: var(--light);
}

.block-schedule p {
  margin: 15px 0;
}

form {
  background: var(--light);
  padding: 12px;
  border-radius: 20px;
}

form p {
  color: #000000;
  font-size: 18px;
  font-weight: bold;
}

@media (max-width: 1600px) {
  .block {
    width: 100%;
  }

  .block-schedule div {
    width: auto;
  }
}
</style>
