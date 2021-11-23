<template>
  <div :class="'main-field bg ' + addClassVisible" :style="style">
    <span class="content">
      <p class="next-count">{{ nextCount }}</p>
      <p class="question">{{ question.getTitle() }}</p>
      <v-row>
        <v-col
          class="percent"
          cols="6"
        >
          {{ afterFirstPercent }}%({{ firstCount }}票)
        </v-col>
        <v-col
          class="percent"
          cols="6"
        >
          {{ afterSecondPercent }}%({{ secondCount }}票)
        </v-col>
      </v-row>
      <v-row>
        <v-col
          class="answer"
          cols="6"
        >
          {{ question.getFirstAnswer() }}
        </v-col>
        <v-col
          class="answer"
          cols="6"
        >
          {{ question.getSecondAnswer() }}
        </v-col>
      </v-row>
      <v-row>
        <v-col
          class="answer"
          cols="6"
        >
          <img 
            :src="question.getFirstImgUrl()"
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
            :src="question.getSecondImgUrl()"
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
const OPAQUE = '100'
const TRANSPARENT = '0'

@Component
export default class Home extends Vue {
  private question: userQuestionPb.UserQuestion.GetRandomResponse | null = null
  private selected = 0
  private firstCount = 0
  private secondCount = 0
  private firstPercent = 50
  private afterFirstPercent = 50
  private afterSecondPercent = 50

  // タイマ関係
  private nextCount = 5
  private resultTimer: NodeJS.Timeout | null = null
  private nextTimer: NodeJS.Timeout | null = null

  // スタイル操作
  private addClassVisible = ''
  private transparency = TRANSPARENT

  private async asyncData () {
    try {
      const res = await userQuestionService.GetRandom()
      return {
        question: res
      }
    } catch (err: any) {
    }
  }

  private mounted () {
    this.setClass()
  }

  private get style (): any {
    return {
      '--percent': String(this.firstPercent) + '%',
      '--transparency': this.transparency,
    }
  }

  private selectFirst () {
    this.selected = 1
    this.updateResult()
  }

  private selectSecond () {
    this.selected = 2
    this.updateResult()
  }

  private updateResult () {
    // 質問データの更新

    this.calculationResult()
  }

  // 表示結果の計算
  private calculationResult () {
    if (this.selected === 1) {
      this.firstCount = this.question!.getFirstCount() + 1
      this.secondCount = this.question!.getSecondCount()
    } else if (this.selected === 2) {
      this.firstCount = this.question!.getFirstCount()
      this.secondCount = this.question!.getSecondCount() + 1
    }
    const totalCount = this.firstCount + this.secondCount
    this.afterFirstPercent = Math.round(this.firstCount / totalCount * 100)
    this.afterSecondPercent = 100 - this.afterFirstPercent
    this.resultTimer = setInterval(this.dispResult, 10);
  }

  // 結果表示
  private dispResult () {
    if (this.firstPercent === this.afterFirstPercent) {
      clearInterval(this.resultTimer as NodeJS.Timeout)
      this.transparency = OPAQUE
      this.nextTimer = setInterval(this.nextQuestion, 1000);
    } else if (this.firstPercent > this.afterFirstPercent) {
      this.firstPercent -= 0.5
    } else if (this.firstPercent < this.afterFirstPercent) {
      this.firstPercent += 0.5
    }
  }

  // 5秒後に次の質問へ行く為のタイマ
  private async nextQuestion () {
    if (this.nextCount < 1) {
      clearInterval(this.nextTimer as NodeJS.Timeout)
      this.refresh()
      return
    }
    this.nextCount -= 1
  }

  private async refresh () {
    this.addClassVisible = ''
    this.question = await userQuestionService.GetRandom()
    this.firstPercent = 50
    this.afterFirstPercent = 50
    this.transparency = TRANSPARENT
    this.nextCount = 5
    setTimeout(this.setClass, 700)
  }

  private setClass () {
    this.addClassVisible = CLASSNAME
  }
}
</script>
