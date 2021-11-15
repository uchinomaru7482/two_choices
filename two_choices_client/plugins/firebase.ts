//import firebase from 'firebase/app'
//import 'firebase/auth'
import firebase from 'firebase/compat/app'
import 'firebase/compat/auth'
import 'firebase/compat/firestore'

// Your web app's Firebase configuration
const firebaseConfig = {
  apiKey: "AIzaSyBLQ8ahKfEfJ1b3lDoVPuOF5uYfV_NwtSA",
  authDomain: "two-choices.firebaseapp.com",
  projectId: "two-choices",
  storageBucket: "two-choices.appspot.com",
  messagingSenderId: "1044056550936",
  appId: "1:1044056550936:web:3f94e21575e66d7d115d6d"
}

// Appの初期化
if (!firebase.apps.length) {
  firebase.initializeApp(firebaseConfig)
}

export default firebase
