import firebase from '@/plugins/firebase'

// FirebaseUserを取得(サービス用カスタムクレーム判定を行わない)
export const getFirebaseUser = (): Promise<firebase.User | null> => {
  return new Promise((resolve, reject) => {
    const unsubscribe = firebase.auth().onAuthStateChanged((user) => {
      unsubscribe()
      return resolve(user)
    }, (err: firebase.auth.Error) => {
      unsubscribe()
      return reject(err)
    })
  })
}

// JWT+カスタムクレーム取得(サービス用カスタムクレーム判定を行わない)
export const getIDTokenResult = async (forceRefresh?: boolean | undefined): Promise<firebase.auth.IdTokenResult | null> => {
  try {
    const user = await getFirebaseUser()
    if (user !== null) {
      const tokenResult = await user.getIdTokenResult(forceRefresh)
      return tokenResult
    }
  } catch (error: any) {}
  return null
}

export const getIDToken = async (): Promise<string> => {
  try {
    const tokenResult = await getIDTokenResult()
    if (tokenResult !== null) {
      return tokenResult.token
    }
  } catch (error: any) {}
  return ''
}

// カスタムクレーム込で認証判定
export const isAuthed = async (): Promise<boolean> => {
  try {
    const tokenResult = await getIDTokenResult()
    if (tokenResult !== null && tokenResult.claims.RegistComplete) {
      return true
    }
  } catch (error: any) {}
return false
}
