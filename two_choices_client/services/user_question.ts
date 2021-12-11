import {UserQuestionService} from "@/generated/user_question_pb_service";
import {UserQuestion} from "@/generated/user_question_pb";
import * as commonPb from '@/generated/common_pb'
import * as common from './common'

export const Echo = async (msg: string): Promise<UserQuestion.EchoResponse> => {
  const req = new UserQuestion.EchoRequest()
  req.setMsg(msg)
  return common.CallAPI(UserQuestionService.Echo, req)
}

export const GetRandom = async (): Promise<UserQuestion.GetRandomResponse> => {
  const req = new commonPb.Empty()
  return common.CallAPI(UserQuestionService.GetRandom, req)
}

export const Update = async (id: number, isFirstSelected: boolean): Promise<commonPb.Empty> => {
  const req = new UserQuestion.UpdateRequest()
  req.setId(id)
  req.setIsFirstSelected(isFirstSelected)
  return common.CallAPI(UserQuestionService.Update, req)
}
