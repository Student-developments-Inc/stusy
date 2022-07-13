<template>
  <aside id="menuVertical">
    <h1>
      <router-link to="/home">StudentSystem</router-link>
    </h1>
    <ol>
      <li>
        <router-link to="/home" v-bind:id="[status==='home'?'active':'']" v-on:click="[status='home']">
          <div class="LeftMenu_icon">
            <img src="@/assets/home.svg">
          </div>
          <span>Домашняя страница</span>
        </router-link>
      </li>
      <li>
        <router-link to="/mycourses" v-bind:id="[status==='mycourses'?'active':'']" v-on:click="[status='mycourses']">
          <div class="LeftMenu_icon">
            <img src="@/assets/book.svg">
          </div>
          <span>Мои курсы</span>
        </router-link>
      </li>
      <li>
        <router-link to="#" v-bind:id="[status==='timetable'?'active':'']" v-on:click="[status='timetable']">
          <div class="LeftMenu_icon">
            <img src="@/assets/calendar.svg">
          </div>
          <span>Расписание</span>
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
  padding: 0 13px 0 0;
}
</style>