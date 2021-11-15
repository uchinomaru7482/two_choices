<template>
  <div class="container">
    <h1>ユーザ登録完了</h1>
    <p>ユーザ登録が完了しました</p>
    <p>お楽しみ下さい</p>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import * as firebaseAuth from '@/plugins/firebase_auth'
import * as service from '@/services/user_authentication'

const ALREADY_EXISTS = '既に登録済です'

@Component
export default class RegistUserComplete extends Vue {
  private errorMsg = ''

  private async created () {
    // 登録済確認
    const isAuthed = await firebaseAuth.isAuthed()
    if (isAuthed) {
      this.errorMsg = ALREADY_EXISTS
      return
    }

    // サーバー側に登録する
    try {
      const state = this.$route.query.state as string
      const res = service.UserRegistration(state)
    } catch (err: any) {
      console.log(err)
    }
  }
}
</script>