<template>
  <v-app>
    <v-btn
      color="primary"
      elevation="2"
      outlined
      large
    >
      <nuxt-link to="/auth/registration/">ユーザ登録画面</nuxt-link>
    </v-btn>
  </v-app>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

@Component
export default class Home extends Vue {
}
</script>
