import Vue from 'vue';
import VueI18n from 'vue-i18n';
import zhLocal from './zh-CN.js';
import enLocal from './en-US.js';
import jaLocal from './ja-JP.js';
import ToastService from 'primevue/toastservice';
import PrimeVue from 'primevue/config';
// import { getLang } from '../utils/auth';
// import { getTimezoneName } from '../utils/time';

Vue.use(VueI18n);
Vue.use(PrimeVue);
Vue.use(ToastService);

export const messages = {
  'zh-CN': {
    ...zhLocal
  },
  'en-US': {
    ...enLocal
  },
  'ja-JP': {
    ...jaLocal
  }
};


// export const lang = getLang() ? getLang() : getTimezoneName();
export const i18n = new VueI18n({
  locale: 'zh-CN',
  fallbackLocale: 'zh-CN',
  messages
});


export default {};
