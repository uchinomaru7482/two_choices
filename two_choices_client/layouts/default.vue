<template>
  <v-app id="inspire">
    <v-navigation-drawer
      v-model="drawer"
      app
    >
      <v-list>
        <v-list-item
          v-for="[icon, text, link] in links"
          :key="icon"
          link
        >
          <nuxt-link :to="link" class="link">
            <v-list-item-icon>
              <v-icon>{{ icon }}</v-icon>
            </v-list-item-icon>

            <v-list-item-content>
              <v-list-item-title>{{ text }}</v-list-item-title>
            </v-list-item-content>
          </nuxt-link>
        </v-list-item>
      </v-list>
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
    data: () => ({ 
      drawer: null,
      links: [
        ['mdi-home', 'ホーム', '/'],
        ['mdi-account', '会員登録', '/auth/registration/'],
        ['mdi-bell', 'お知らせ', '/notification/list'],
      ],
    }),
  }
</script>
