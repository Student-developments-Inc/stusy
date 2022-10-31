<template>
  <aside id="menuVertical">
    <h1>
      <router-link to="/home" v-on:click="[status='home']">StudentSystem</router-link>
    </h1>
    <ol>
      <li>
        <router-link to="/home" v-bind:id="[status==='home'?'active':'']" v-on:click="[status='home']">
          <div class="LeftMenu_icon" id="home-ico">
          </div>
          <span>Домашняя страница</span>
        </router-link>
      </li>
      <li>
        <router-link to="/mycourses" v-bind:id="[status==='mycourses'?'active':'']" v-on:click="[status='mycourses']">
          <div class="LeftMenu_icon" id="book-ico">
          </div>
          <span>Мои курсы</span>
        </router-link>
      </li>
      <li>
        <router-link to="#" v-bind:id="[status==='timetable'?'active':'']" v-on:click="[status='timetable']">
          <div class="LeftMenu_icon" id="calendar-ico">
          </div>
          <span>Меню для тестов</span>
        </router-link>
      </li>
      <li>
        <router-link to="/courses" v-bind:id="[status==='timetable'?'active':'']" v-on:click="[status='timetable']">
          <div class="LeftMenu_icon" id="calendar-ico">
          </div>
          <span>Курсы</span>
        </router-link>
      </li>
      <li>
        <router-link to="/creatingCourses" v-bind:id="[status==='timetable'?'active':'']"
                     v-on:click="[status='timetable']">
          <div class="LeftMenu_icon" id="calendar-ico">
          </div>
          <span>Создание курса</span>
        </router-link>
      </li>
      <li>
        <router-link to="/creatingTest" v-bind:id="[status==='timetable'?'active':'']"
                     v-on:click="[status='timetable']">
          <div class="LeftMenu_icon" id="calendar-ico">
          </div>
          <span>Создание теста</span>
        </router-link>
      </li>
      <li>
        <router-link to="/asdasd" v-bind:id="[status==='timetable'?'active':'']" v-on:click="[status='timetable']">
          <div class="LeftMenu_icon" id="calendar-ico">
          </div>
          <span>Не найдено</span>
        </router-link>
      </li>
    </ol>
  </aside>
</template>

<script>
export default {
  name: "AsideMenu",
  data() {
    return {
      status: 'home'
    };
  }
}
</script>

<style scoped>
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


#menuVertical > h1 {
  margin-bottom: 30px;
}

#menuVertical > h1 > a {
  font-weight: 500;
  font-size: 28px;
  color: var(--blue);
  margin-top: 0;
  text-decoration: none;
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
  width: 252px;
  height: 54px;
}

#active {
  background: var(--blue);
}

#active img {
  filter: invert();
}

#active > span {
  color: var(--light)
}


#menuVertical ol li a {
  display: flex;
  color: var(--dark);
  text-decoration: none;
  font-size: 18px;
  padding: 16px 21px;
  border-radius: 15px;
  margin-bottom: 15px;
}

.LeftMenu_icon {
  margin: 0 13px 0 0;
}

#home-ico {
  background: url("@/assets/home.svg");
  width: 21px;
}
#book-ico {
  background: url("@/assets/book.svg");
  width: 19px;
}
#calendar-ico{
  background: url("@/assets/calendar.svg");
  width: 19px;
}
</style>