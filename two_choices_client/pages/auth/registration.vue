<template>
  <div class="main-field">
    <h1 class="page-title">ユーザ新規登録</h1>
    <p v-if="errMsg">{{ errMsg }}</p>
    <v-form>
      <v-text-field
        v-model="name"
        counter="20"
        label="名前"
      />
      <v-text-field
        v-model="mailAddress"
        counter="50"
        label="メールアドレス"
      />
      <v-text-field
        v-model="password"
        counter="20"
        label="パスワード"
      />
      <v-btn
        color="primary"
        elevation="2"
        outlined
        @click="signUp"
      >
        新規登録
      </v-btn>
    </v-form>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import firebase from '@/plugins/firebase'
import * as userAuthenticationService from '@/services/user_authentication'
import * as userAuthenticationPb from '@/generated/user_authentication_pb'

const ALREADY_EXISTS_ADDRESS_ERROR_MESSAGE = 'このメールアドレスは既に登録済です'

@Component
export default class UserRegistration extends Vue {
  private name = ''
  private mailAddress = ''
  private password = ''
  private errMsg = ''

  private async signUp () {
    // firebaseユーザの作成
    let fbCredential: firebase.auth.UserCredential | null = null
    try {
      fbCredential = await firebase.auth().createUserWithEmailAndPassword(this.mailAddress, this.password)
    } catch (err: any) {
      console.log(err)
    }

    // 確認メール送信
    if (fbCredential !== null) {
      try {
        const result = await this.sendVerifyMail(fbCredential)
        if (result === userAuthenticationPb.UserAuthentication.SendVerificationMailResult.ALREADY_EXISTS_ADDRESS) {
          this.errMsg = ALREADY_EXISTS_ADDRESS_ERROR_MESSAGE
          return
        }
        // メール送信ページへリダイレクト
        this.$router.push('/auth/send_mail/')
        return
      } catch (err: any) {
        console.log(err)
      }
    }
  }

  private async sendVerifyMail (credential: firebase.auth.UserCredential):Promise<number> {
    try {
      const res = await userAuthenticationService.SendVerificationMail(this.name, this.mailAddress)
      if (
        res.getResult() === userAuthenticationPb.UserAuthentication.SendVerificationMailResult.SEND_MAIL_SUCCESS ||
        res.getResult() === userAuthenticationPb.UserAuthentication.SendVerificationMailResult.ALREADY_EXISTS_ADDRESS
      ) {
        return res.getResult()
      }
    } catch (err: any) {
      console.log(err)
    }
    return userAuthenticationPb.UserAuthentication.SendVerificationMailResult.SEND_MAIL_SUCCESS
  }
}
</script>
