import {UserAuthenticationService} from "../generated/user_authentication_pb_service";
import {UserAuthentication} from "../generated/user_authentication_pb";
import * as common from './common'

export const SendVerificationMail = async (name: string, mailAddress: string): Promise<UserAuthentication.SendVerificationMailResponse> => {
  const req = new UserAuthentication.SendVerificationMailRequest()
  req.setName(name)
  req.setEmail(mailAddress)
  return common.CallAPI(UserAuthenticationService.SendVerificationMail, req)
}

export const UserRegistration = async (state: string): Promise<UserAuthentication.UserRegistrationResponse> => {
  const req = new UserAuthentication.UserRegistrationRequest()
  req.setState(state)
  return common.CallAPI(UserAuthenticationService.UserRegistration, req)
}
