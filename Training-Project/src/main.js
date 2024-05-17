import { createApp } from 'vue'
// import './style.css'
import App from './App.vue'
import Antd from 'ant-design-vue'
import 'ant-design-vue/dist/reset.css'
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  { path: '/Login', component: () => import('./components/Login.vue') },
  { path: '/Hello', component: () => import('./components/HelloWorld.vue') },
  { path: '/', component: () => import('./components/main/src/App.vue') }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token');
  let verified = false;
  if (token === 'admin_logged_in' || token === 'user_logged_in') {
    verified = true;
  }
  if (verified) {
    if (to.path === '/Login' || to.path === '/login') {
      next('/');
    } else {
      next();
    }
  } else {
    if (to.path === '/Login' || to.path === 'login') {
      next();
    } else {
      next('/Login');
    }
  }
});


createApp(App).use(Antd).use(router).mount('#app')
