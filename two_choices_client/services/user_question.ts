import {UserQuestionService} from "@/generated/user_question_pb_service";
import {UserQuestion} from "@/generated/user_question_pb";
import * as commonPb from '@/generated/common_pb'
import * as common from './common'

export const GetRandom = async (): Promise<UserQuestion.GetRandomResponse> => {
  const req = new commonPb.Empty()
  return common.CallAPI(UserQuestionService.GetRandom, req)
}
