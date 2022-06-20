<template>
  <div class="profile">
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
    <div class="main">
      <header>
        <h1>Ваш профиль</h1>
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
                <span>Элеонора А.</span>
                <div class="avatar-arrow" v-bind:style="[menu?{'transform':'rotate(90deg)'}:{}]">
                  <img src="@/assets/arrow.svg"/>
                </div>
              </a>
              <transition name="slide-fade">
                <ul class="sub-menu" v-if="menu">
                  <li><a href="#">Профиль</a></li>
                  <li><a href="#">Выйти</a></li>
                </ul>
              </transition>
            </li>
          </ul>
        </nav>
      </header>
      <div class="mainContent">
        <div class="AvatarAndPersonalDataAndContacts">
          <div class="AvatarProfile">
            <div id="AvatarProfileMedia">АВАТАРКИ НЕТУ</div>
            <div class="button-login">
              <input type="submit" value="Загрузить новое фото"/>
            </div>
          </div>
          <div class="personalData">
            <h1>Личные данные</h1>
            <div class="personalDataContent">
              <form class="personalDataContentForm">
                <label class="input-form">
                  <h2>Имя</h2>
                  <input
                      type="text"
                      placeholder="Введите имя"
                      required
                  />
                </label>
                <label class="input-form">
                  <h2>Фамилия</h2>
                  <input
                      type="text"
                      placeholder="Введите фамилию"
                      required
                  />
                </label>
                <label class="input-form">
                  <h2>Отчество</h2>
                  <input
                      type="text"
                      placeholder="Введите отчество"
                  />
                </label>
              </form>
              <form class="personalDataContentForm">
                <label class="input-form">
                  <h2>Дата рождения</h2>
                  <input
                      type="text"
                      placeholder="хх.хх.хххх"
                      required
                  />
                </label>
                <label class="input-form">
                  <h2>Пол</h2>
                  <label class="radio-button">
                    <input
                        name="gender"
                        type="radio"
                        value="male"
                    />
                    <span>Мужской</span>
                  </label>
                  <label class="radio-button">
                    <input
                        name="gender"
                        type="radio"
                        value="female"
                    />
                    <span>Женский</span>
                  </label>
                </label>
              </form>
            </div>
          </div>
          <div class="contactsData">
            <h1>Контакты</h1>
            <form class="contactsDataForm">
              <label class="input-form">
                <h2>Номер телефона</h2>
                <input
                    type="text"
                    placeholder="Введите номер телефона"
                />
              </label>
              <label class="input-form">
                <h2>Адрес электронной почты</h2>
                <input
                    type="text"
                    placeholder="name@gmail.com"
                    v-model="email"
                    required
                />
              </label>
            </form>
          </div>
        </div>
        <div class="aboutMe">
          <h1>Обо мне</h1>
          <textarea class="aboutMeInput" placeholder="Расскажите о себе вкратце"></textarea>
        </div>
        <div class="linksToSocialNetworks">
          <h1>Ссылки на соц. сети</h1>
          <form class="linksToSocialNetworksForm">
            <label class="input-form">
              <h2>Vk.com</h2>
              <input
                  type="text"
                  placeholder="vk.com"
              />
            </label>
            <label class="input-form">
              <h2>Facebook</h2>
              <input
                  type="text"
                  placeholder="facebook.com"
              />
            </label>
            <label class="input-form">
              <h2>Instagram</h2>
              <input
                  type="text"
                  placeholder="instagram.com"
              />
            </label>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ProfileView"
};
</script>

<style scoped>
body {
  overflow: auto !important;
}

.profile {
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

.LeftMenu_icon {
  padding: 0 13px 0 0;
}

.main {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  justify-content: flex-start;
  gap: 10px;
  width: 100%;
  padding: 15px 24px 24px;
  overflow: auto;

}

.main header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.main header h1 {
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

.mainContent {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  gap: 30px;
  padding: 57px 29px;
}

.AvatarAndPersonalDataAndContacts {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: stretch;
  gap: 1%;
  padding: 0 200px 0 0;
}

.AvatarProfile {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 10px;
}

#AvatarProfileMedia {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 349px;
  height: 350px;
  background: #2255f4;
}

.personalData, .contactsData {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: flex-start;
  gap: 45px;
}

.personalDataContent {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  align-items: flex-start;
  justify-content: space-between;
  column-gap: 100px;
}

.personalDataContentForm, .contactsDataForm {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.input-form input {
  background: var(--light);
  width: auto;
}

.radio-button {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: flex-start;
  gap: 20px ;
}

.input-form label span {
  font-size: 18px;
}

h2 {
  font-weight: 400;
  font-size: 20px;
}

h1 {
  font-weight: 500;
  font-size: 28px;
}

.aboutMe, .linksToSocialNetworks {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  gap: 30px;
}

.aboutMeInput {
  resize: none;
  height: 10rem;
  overflow: auto;
  width: 50%;
  padding: 20px;
  font-weight: 300;
  font-size: 24px;
  border: 1px solid var(--blue);
  border-radius: 10px;
}

.aboutMeInput:focus {
  outline: none !important;
  border: 1px solid var(--blue-mate);
}

.linksToSocialNetworksForm {
  display: flex;
  justify-content: space-around;
  align-items: flex-start;
  flex-direction: row;
  flex-wrap: wrap;
  gap: 10px;
}
</style>