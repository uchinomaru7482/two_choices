<template>
  <v-app id="inspire">
    <v-navigation-drawer
      v-model="drawer"
      app
    >
      <nuxt-link to="/">ホーム</nuxt-link>
      <nuxt-link to="/auth/registration/">ユーザ登録画面</nuxt-link>
    </v-navigation-drawer>

    <v-app-bar app>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>Two Choices</v-toolbar-title>
    </v-app-bar>

    <v-main class="grey lighten-3">
      <!-- <v-container> -->
        <v-sheet
          min-height="87vh"
          rounded="lg"
        >
          <Nuxt />
        </v-sheet>    
      <!-- </v-container> -->
    </v-main>
  </v-app>
</template>

<script>
  export default {
    data: () => ({ drawer: null }),
  }
</script>
