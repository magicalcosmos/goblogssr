import Vue from 'vue';
import VueI18n from 'vue-i18n';
import zhLocal from './zh-CN.js';
import enLocal from './en-US.js';
import jaLocal from './ja-JP.js';
import { getLang } from '../utils/auth';
import { getTimezoneName } from '../utils/time';

const messages = {
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
Vue.use(VueI18n);

const lang = getLang() ? getLang() : getTimezoneName();
const i18n = new VueI18n({
  locale: getLang() || 'zh-CN',
  fallbackLocale: 'zh-CN',
  messages
});
export default i18n;

