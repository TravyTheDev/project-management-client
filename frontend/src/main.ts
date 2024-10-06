import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import './style.css';
import jp from "./locales/jp.json" 
import en from "./locales/en.json" 
import { createI18n } from 'vue-i18n'

const i18n = createI18n({ 
    locale: navigator.language, 
    fallbackLocale: "en", 
    messages: { jp, en }, 
    missingWarn: false,
    fallbackWarn: false,
    legacy: false 
  })

createApp(App).use(router).use(i18n).mount('#app')
