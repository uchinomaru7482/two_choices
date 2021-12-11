<template>
  <v-container>
    <v-sheet
      rounded="lg"
      min-height="200"
      max-width="700"
      outlined
      class="p-5 my-3 mx-auto"
    >
      <div class="main-field">
        <v-form>
          <v-text-field
            class="mt-5"
            v-model="reqMsg"
            counter="100"
            label="msg"
          />
          <v-btn
            class="mt-5"
            color="primary"
            elevation="2"
            outlined
            @click="echo"
          >
            Echo
          </v-btn>
        </v-form>
        <div class="py-5">
          res: {{ resMsg }}
        </div>
      </div>
    </v-sheet>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import * as userQuestionService from '@/services/user_question'

@Component
export default class Echo extends Vue {
  private reqMsg = ''
  private resMsg = ''

  private async echo () {
    const res = await userQuestionService.Echo(this.reqMsg)
    this.resMsg = res.getMsg()
  }
}

</script>
