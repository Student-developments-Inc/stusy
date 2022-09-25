<template>
  <div class="homepage">
    <AsideMenu/>
    <div class="home">
      <div class="content">
        <header>
          <h1>Доброе утро{{ userData.first_name && userData.last_name ? `, ${userData.first_name}` : "" }}</h1>
          <TopMenu/>
        </header>
        <router-view/>
      </div>
    </div>
  </div>
</template>

<script>
import AsideMenu from "@/components/AsideMenu";
import TopMenu from "@/components/TopMenu";
import {getCookie, logout, url} from "@/global";

export default {
  name: "HomeLayout",
  components: {
    AsideMenu, TopMenu
  },
  data() {
    return {
      userData: {
        first_name: "",
        last_name: ""
      }
    };
  },
  methods: {
    getUserData() {
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
    if (getCookie("ID") !== undefined && this.userData.first_name === "") this.getUserData();
  }
}
</script>

<style scoped>
.homepage {
  display: flex !important;
}

.home {
  display: flex !important;
  min-height: 100vh;
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