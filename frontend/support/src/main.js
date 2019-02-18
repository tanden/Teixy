import Vue from 'vue';
import App from './App.vue';
import store from './store'
import router from './router'
import firebase from 'firebase/app'

Vue.config.productionTip = false;

// Initialize Firebase
const config = {
  apiKey: "AIzaSyCz5_h_UjxU5s-pmEuTHOQ5HkaogTexOLQ",
  authDomain: "teixy-support.firebaseapp.com",
  databaseURL: "https://teixy-support.firebaseio.com",
  projectId: "teixy-support",
  storageBucket: "teixy-support.appspot.com",
  messagingSenderId: "891516084242"
};
firebase.initializeApp(config);

new Vue({
  store,
  router,
  render: h => h(App)
}).$mount('#app');
