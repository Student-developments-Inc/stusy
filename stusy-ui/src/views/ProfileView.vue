<template>
  <div class="mainContent">
    <div class="personalData">
      <div class="photoName">
        <div class="photo">
          <img src="@/assets/avatar.jpg" alt="Фото профиля" draggable="false">
        </div>
        <div class="nameNum">
          <p class="name">{{ userData.first_name }} {{ userData.last_name }}</p>
          <p>Номер студенческого: {{ userData.id }}</p>
        </div>
      </div>
      <form v-on:submit.prevent="submit(userData)">
        <div class="personalDataForm">
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
          <label class="input-form">
            <p>Отчество</p>
            <input
                type="text"
                placeholder="Введите отчество"
            />
          </label>
        </div>
        <div class="button-login">
          <input type="submit" value="Сохранить"/>
        </div>
      </form>

    </div>
    <!--        <div class="AvatarAndPersonalDataAndContacts">-->
    <!--          <div class="AvatarProfile">-->
    <!--            <div id="AvatarProfileMedia">АВАТАРКИ НЕТУ</div>-->
    <!--            <div class="button-login">-->
    <!--              <input type="submit" value="Загрузить новое фото"/>-->
    <!--            </div>-->
    <!--          </div>-->
    <!--          <div class="personalData">-->
    <!--            <h1>Личные данные</h1>-->
    <!--            <div class="personalDataContent">-->
    <!--              <form class="personalDataContentForm">-->
    <!--                <label class="input-form">-->
    <!--                  <h2>Имя</h2>-->
    <!--                  <input-->
    <!--                      type="text"-->
    <!--                      placeholder="Введите имя"-->
    <!--                      required-->
    <!--                  />-->
    <!--                </label>-->
    <!--                <label class="input-form">-->
    <!--                  <h2>Фамилия</h2>-->
    <!--                  <input-->
    <!--                      type="text"-->
    <!--                      placeholder="Введите фамилию"-->
    <!--                      required-->
    <!--                  />-->
    <!--                </label>-->
    <!--                <label class="input-form">-->
    <!--                  <h2>Отчество</h2>-->
    <!--                  <input-->
    <!--                      type="text"-->
    <!--                      placeholder="Введите отчество"-->
    <!--                  />-->
    <!--                </label>-->
    <!--                <div class="button-login">-->
    <!--                  <input type="submit" value="Сохранить"/>-->
    <!--                </div>-->
    <!--              </form>-->
    <!--              <form class="personalDataContentForm">-->
    <!--                <label class="input-form">-->
    <!--                  <h2>Дата рождения</h2>-->
    <!--                  <input-->
    <!--                      type="text"-->
    <!--                      placeholder="хх.хх.хххх"-->
    <!--                      required-->
    <!--                  />-->
    <!--                </label>-->
    <!--                <label class="input-form">-->
    <!--                  <h2>Пол</h2>-->
    <!--                  <label class="radio-button">-->
    <!--                    <input-->
    <!--                        name="gender"-->
    <!--                        type="radio"-->
    <!--                        value="male"-->
    <!--                    />-->
    <!--                    <span>Мужской</span>-->
    <!--                  </label>-->
    <!--                  <label class="radio-button">-->
    <!--                    <input-->
    <!--                        name="gender"-->
    <!--                        type="radio"-->
    <!--                        value="female"-->
    <!--                    />-->
    <!--                    <span>Женский</span>-->
    <!--                  </label>-->
    <!--                </label>-->
    <!--              </form>-->
    <!--            </div>-->
    <!--          </div>-->
    <!--          <div class="contactsData">-->
    <!--            <h1>Контакты</h1>-->
    <!--            <form class="contactsDataForm">-->
    <!--              <label class="input-form">-->
    <!--                <h2>Номер телефона</h2>-->
    <!--                <input-->
    <!--                    type="text"-->
    <!--                    placeholder="Введите номер телефона"-->
    <!--                />-->
    <!--              </label>-->
    <!--              <label class="input-form">-->
    <!--                <h2>Адрес электронной почты</h2>-->
    <!--                <input-->
    <!--                    type="text"-->
    <!--                    placeholder="name@gmail.com"-->
    <!--                    required-->
    <!--                />-->
    <!--              </label>-->
    <!--            </form>-->
    <!--          </div>-->
    <!--        </div>-->
    <!--        <div class="aboutMe">-->
    <!--          <h1>Обо мне</h1>-->
    <!--          <textarea class="aboutMeInput" placeholder="Расскажите о себе вкратце"></textarea>-->
    <!--        </div>-->
    <!--        <div class="linksToSocialNetworks">-->
    <!--          <h1>Ссылки на соц. сети</h1>-->
    <!--          <form class="linksToSocialNetworksForm">-->
    <!--            <label class="input-form">-->
    <!--              <h2>Vk.com</h2>-->
    <!--              <input-->
    <!--                  type="text"-->
    <!--                  placeholder="vk.com"-->
    <!--              />-->
    <!--            </label>-->
    <!--            <label class="input-form">-->
    <!--              <h2>Facebook</h2>-->
    <!--              <input-->
    <!--                  type="text"-->
    <!--                  placeholder="facebook.com"-->
    <!--              />-->
    <!--            </label>-->
    <!--            <label class="input-form">-->
    <!--              <h2>Instagram</h2>-->
    <!--              <input-->
    <!--                  type="text"-->
    <!--                  placeholder="instagram.com"-->
    <!--              />-->
    <!--            </label>-->
    <!--          </form>-->
    <!--        </div>-->
  </div>
</template>

<script setup>
import {getUserData} from "@/composables/getUserData";
import {getCookie, url} from "@/global";

const {userData} = getUserData();

function submit(userData) {
  fetch(`${url}/users/${getCookie("ID")}`, {
    method: "PUT",
    headers: {
      "Authorization": `Bearer ${getCookie("TOKEN")}`
    },
    body: JSON.stringify({
      "first_name": userData.first_name,
      "last_name": userData.last_name
    })
  }).then(response => {
    if (response.ok) return response.json();
    switch (response.status) {
      case 400:
        console.log("Неверные данные");
        break;
      case 403:
        console.log("Нет доступа")
        break;
    }
  })
      // .then(data => {
    // console.log(data);
  // }).catch(err => {
  //   console.error("Cannot fetch" + err);
  // });
}
</script>

<style scoped>

.mainContent {
  display: flex;
  justify-content: space-evenly;
  align-items: stretch;
  gap: 30px;
  padding: 57px 29px;
  flex-wrap: wrap;
}

.personalData {
  background: var(--light);
  padding: 27px 32px;
  border-radius: 25px;
  width: 553px;
}

.photoName {
  display: flex;
  gap: 25px;
}

.photo img {
  width: 72px;
  border-radius: 15px;
  user-select: none;
}

.personalDataForm {
  display: flex;
  gap: 15px;
  flex-wrap: wrap;
}

input:disabled,
input[disabled] {
  background: #000 !important;
}

.name {
  margin-bottom: 10px;
  font-size: 17px;
  font-weight: bold;
}

/*.AvatarAndPersonalDataAndContacts {*/
/*  display: flex;*/
/*  flex-direction: row;*/
/*  gap: 1%;*/
/*  width: 100%;*/
/*  justify-content: space-evenly;*/
/*  flex-wrap: wrap;*/
/*}*/

/*.AvatarProfile {*/
/*  display: flex;*/
/*  flex-direction: column;*/
/*  justify-content: center;*/
/*  align-items: center;*/
/*  gap: 10px;*/
/*}*/

/*#AvatarProfileMedia {*/
/*  display: flex;*/
/*  align-items: center;*/
/*  justify-content: center;*/
/*  width: 349px;*/
/*  height: 350px;*/
/*  background: #2255f4;*/
/*  border-radius: 15px;*/
/*}*/

/*.personalData, .contactsData {*/
/*  display: flex;*/
/*  flex-direction: column;*/
/*  justify-content: flex-start;*/
/*  align-items: flex-start;*/
/*  gap: 45px;*/
/*}*/

/*.personalDataContent {*/
/*  display: flex;*/
/*  flex-direction: row;*/
/*  flex-wrap: wrap;*/
/*  align-items: flex-start;*/
/*  justify-content: space-between;*/
/*  column-gap: 100px;*/
/*}*/

/*.personalDataContentForm, .contactsDataForm {*/
/*  display: flex;*/
/*  flex-direction: column;*/
/*  gap: 10px;*/
/*}*/

.input-form input {
  background: var(--dark-mate-mate);
  width: auto;
}

/*.radio-button {*/
/*  display: flex;*/
/*  flex-direction: row;*/
/*  align-items: center;*/
/*  justify-content: flex-start;*/
/*  gap: 20px;*/
/*}*/

.input-form label span {
  font-size: 18px;
}

/*h2 {*/
/*  font-weight: 400;*/
/*  font-size: 20px;*/
/*}*/

/*h1 {*/
/*  font-weight: 500;*/
/*  font-size: 28px;*/
/*}*/

/*.aboutMe, .linksToSocialNetworks {*/
/*  display: flex;*/
/*  flex-direction: column;*/
/*  justify-content: flex-start;*/
/*  gap: 30px;*/
/*}*/

/*.aboutMeInput {*/
/*  resize: none;*/
/*  height: 10rem;*/
/*  overflow: auto;*/
/*  max-width: 950px;*/
/*  padding: 20px;*/
/*  font-weight: 300;*/
/*  font-size: 24px;*/
/*  border: 1px solid var(--blue);*/
/*  border-radius: 10px;*/
/*}*/

/*.linksToSocialNetworksForm {*/
/*  display: flex;*/
/*  justify-content: space-around;*/
/*  align-items: flex-start;*/
/*  flex-direction: column;*/
/*  flex-wrap: wrap;*/
/*  gap: 10px;*/
/*}*/
</style>