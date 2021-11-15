import {grpc} from "@improbable-eng/grpc-web"
import * as firebaseAuth from '@/plugins/firebase_auth'

// Import code-generated data structures.
import { ProtobufMessage } from "@improbable-eng/grpc-web/dist/typings/message";

const timeoutInterval: number = 3000

const getMetadata = async (): Promise<grpc.Metadata> => {
  const idToken = await firebaseAuth.getIDToken()
  const metadata = new grpc.Metadata()
  metadata.set('authorization', `bearer ${idToken}`)
  return metadata
}

export const CallAPI = async <TReq extends ProtobufMessage, TRes extends ProtobufMessage, Method extends grpc.MethodDefinition<TReq, TRes>>(method: Method, req: TReq): Promise<TRes> => {
  const metadata = await getMetadata()
  const fun = new Promise((resolve, reject) => {
    grpc.invoke(method, {
      request: req,
      host: "http://localhost:8080",
      metadata: metadata,
      onMessage: (res: TRes) => {
        resolve(res)
      },
      onEnd: (code: grpc.Code, msg: string | undefined) => {
        if (code !== grpc.Code.OK) {
          reject(code)
          console.log(msg)
        }
      }
    })
  })
  return Promise.race([
    fun,
    grpcTimeout().then(() => { throw new Error('timeout error') })
  ]) as Promise<TRes>
}

export const grpcTimeout = (ms?: number) => {
  return new Promise(resolve => setTimeout(resolve, (ms) || timeoutInterval))
}
