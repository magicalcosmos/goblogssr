import Vue from 'vue';
import * as VueI18n from 'vue-i18n';
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

const lang = getLang() ? getLang() : getTimezoneName();
const i18n = VueI18n.createI18n({
  locale: lang || 'zh-CN',
  fallbackLocale: 'zh-CN',
  messages
});

Vue.use(i18n);

export default i18n;

