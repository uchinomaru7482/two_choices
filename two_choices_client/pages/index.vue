<template>
  <div :class="'main-field bg ' + addClassVisible" :style="style">
    <!-- {{ question.getTitle() }}<br>
    {{ question.getFirstAnswer() }}<br>
    {{ question.getSecondAnswer() }}<br> -->
    <span class="content">
      <p class="next-count">{{ nextCount }}</p>
      <p class="question">あなたはどっち派？</p>
      <v-row>
        <v-col
          class="percent"
          cols="6"
        >
          20%(20票)
        </v-col>
        <v-col
          class="percent"
          cols="6"
        >
          80%(80票)
        </v-col>
      </v-row>
      <v-row>
        <v-col
          class="answer"
          cols="6"
        >
          いぬ
        </v-col>
        <v-col
          class="answer"
          cols="6"
        >
          ねこ
        </v-col>
      </v-row>
      <v-row>
        <v-col
          class="answer"
          cols="6"
        >
          <img 
            src="@/static/images/answer_image/dog.png"
            class="ans-img"
            alt="回答1"
            @click="selectFirst()"
          >
        </v-col>
        <v-col
          class="answer"
          cols="6"
        >
          <img 
            src="@/static/images/answer_image/cat.png"
            class="ans-img"
            alt="回答2"
            @click="selectSecond()"
          >
        </v-col>
      </v-row>
    </span>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'nuxt-property-decorator'
import * as userQuestionPb from '@/generated/user_question_pb'
import * as userQuestionService from '@/services/user_question'

const CLASSNAME = '-visible'
const OPAQUE_WHITE_COLOR = 'rgba(255,255,255,100)'
const OPAQUE_GRAY_COLOR = 'rgba(200,200,200,100)'

@Component
export default class Home extends Vue {
  private question: userQuestionPb.UserQuestion.GetRandomResponse | null = null
  private addClassVisible = ''
  private selected = 0
  private percent = 50
  private afterPercent = 50
  private resultTimer: NodeJS.Timeout | null = null
  private whiteColor = 'rgba(255,255,255,0)'
  private grayColor = 'rgba(200,200,200,0)'
  private nextCount = 3
  private nextTimer: NodeJS.Timeout | null = null

  private async created () {
    try {
      this.question = await userQuestionService.GetRandom()
    } catch (err: any) {
    }
  }

  private mounted () {
    this.addClassVisible = CLASSNAME
  }

  private get style (): any {
    return {
      '--percent': String(this.percent) + '%',
      '--white': this.whiteColor,
      '--gray': this.grayColor
    }
  }

  @Watch('afterPercent')
  private watchPercent () {
    this.resultTimer = setInterval(this.dispResult, 10);
  }

  // 結果表示
  private dispResult () {
    if (this.percent === this.afterPercent) {
      clearInterval(this.resultTimer as NodeJS.Timeout)
      this.whiteColor = OPAQUE_WHITE_COLOR
      this.grayColor = OPAQUE_GRAY_COLOR
      this.nextTimer = setInterval(this.nextQuestion, 1000);
    } else if (this.percent > this.afterPercent) {
      this.percent -= 0.5
    } else if (this.percent < this.afterPercent) {
      this.percent += 0.5
    }
  }

  // 3秒後に次の質問へ行く為のタイマ
  private nextQuestion () {
    if (this.nextCount <= 0) {
      clearInterval(this.nextTimer as NodeJS.Timeout)
      location.reload()
    }
    this.nextCount -= 1
  }

  private selectFirst () {
    this.selected = 1
    this.afterPercent = 20
    this.updateResult()
  }

  private selectSecond () {
    this.selected = 2
    this.updateResult()
  }

  private updateResult () {

  }
}
</script>
