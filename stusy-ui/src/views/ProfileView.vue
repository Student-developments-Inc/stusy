<template>
  <div class="profile">
    <AsideMenu/>
    <div class="main">
      <header>
        <h1>Ваш профиль</h1>
        <TopMenu/>
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
import AsideMenu from "@/components/AsideMenu";
import TopMenu from "@/components/TopMenu";

export default {
  name: "ProfileView",
  components: {AsideMenu, TopMenu},
};
</script>

<style scoped>
.profile {
  display: flex !important;
  min-height: 100vh;
  background: var(--dark-mate-mate);
}

.main {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  justify-content: flex-start;
  gap: 10px;
  width: 100%;
  padding: 15px 24px 24px;
  height: 100vh;
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
  justify-content: space-evenly;
  gap: 1%;
  flex-wrap: wrap;
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
  border-radius: 15px;
}

.personalData, .contactsData {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: flex-start;
  gap: 45px;
}

/*.personalDataContent {*/
/*  display: flex;*/
/*  flex-direction: row;*/
/*  flex-wrap: wrap;*/
/*  align-items: flex-start;*/
/*  justify-content: space-between;*/
/*  column-gap: 100px;*/
/*}*/

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
  gap: 20px;
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
  max-width: 950px;
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