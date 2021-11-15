<template>
  <div class="container">
    <h1>ユーザ新規登録</h1>
    <p>確認メールを送信しました</p>
    <p>URLをクリックして本人確認を行って下さい</p>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

@Component
export default class SendVerificationMail extends Vue {
}
</script>
