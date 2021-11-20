<template>
  <div :class="'main-field bg ' + addClassVisible" :style="style">
    <!-- {{ question.getTitle() }}<br>
    {{ question.getFirstAnswer() }}<br>
    {{ question.getSecondAnswer() }}<br> -->
    <span class="content">
      <span>
        <span class="question">あなたはどっち派？</span>
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
    </span>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'nuxt-property-decorator'
import * as userQuestionPb from '@/generated/user_question_pb'
import * as userQuestionService from '@/services/user_question'

const CLASSNAME = '-visible'
const AFTER_PERCENT_COLOR = 'rgba(255,255,255,100)'

@Component
export default class Home extends Vue {
  private question: userQuestionPb.UserQuestion.GetRandomResponse | null = null
  private addClassVisible = ''
  private selected = 0
  private percent = 50
  private afterPercent = 50
  private timer: NodeJS.Timeout | null = null
  private percentColor = 'rgba(255,255,255,0)'

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
      '--percent-color': this.percentColor
    }
  }

  @Watch('afterPercent')
  private watchPercent () {
    this.timer = setInterval(this.updatePercent, 10);
  }

  private updatePercent () {
    if (this.percent === this.afterPercent) {
      clearInterval(this.timer as NodeJS.Timeout)
      this.percentColor = AFTER_PERCENT_COLOR
    } else if (this.percent > this.afterPercent) {
      this.percent = this.percent - 0.5
    } else if (this.percent < this.afterPercent) {
      this.percent = this.percent + 0.5
    }
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
