import {UserCertProfileService} from "../generated/user_cert_profile_pb_service";
import {UserCertProfile} from "../generated/user_cert_profile_pb";
import * as common from './common'

export const GetProfile = async (userID: number): Promise<UserCertProfile.GetProfileResponse> => {
  const req = new UserCertProfile.GetProfileRequest()
  req.setId(userID)
  return common.CallAPI(UserCertProfileService.GetProfile, req)
}
