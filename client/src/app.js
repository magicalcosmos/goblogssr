import Vue from 'vue';
import './theme/index.scss'; // global css
import App from './App.vue';
import { createRouter } from './router'
import { createStore } from './store'
import axios from 'axios';
import i18n from '@/i18n';


export function createApp (context) {
  const store = createStore(context);
  const router = createRouter();

  if (typeof window == 'undefined') {
    if (context && context.ssrCtx) {
      axios.defaults.headers.common['SSR-Ctx'] = JSON.stringify(context.ssrCtx);
    }
  }
  axios.defaults.baseURL = API_BASE_URL;

  const app = new Vue({
      i18n,
      router,
      store,
      render: h => h(App)
    });

  return { app, router, store }
}

export function callComponentsHookWith(compoList, hookName, context) {
  try {
    return compoList.map((component) => {
      const hook = component[hookName];
      if (hook) {
          return hook(context);
      }
    }).filter(_ => _);
  } catch (e) {
    console.error(e);
    return undefined
  }
}
