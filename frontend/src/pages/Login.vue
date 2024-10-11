<template>
    <div class="h-screen flex flex-col items-center text-fontColor">
        <div class="m-auto">
            <div>
                <p>{{t('auth.email')}}: </p>
                <input class="border text-black" v-model="email" type="text">
            </div>
            <div>
                <p>{{t('auth.password')}}: </p>
                <input class="border text-black" v-model="password" type="password">
            </div>
            <button class="bg-myMessage px-2 py-1 rounded-lg hover:border border-slate-400 mt-2" @click="login">
                {{t('auth.login')}}
            </button>
            <br />
            <p class="text-sm text-red-500 whitespace-pre-line" v-if="isShowError">{{ errorValue }}</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { nextTick, ref } from 'vue';
import { Login } from '../../wailsjs/go/main/App';
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router';
import { getUser } from '../functions';
import { main } from '../../wailsjs/go/models';

const router = useRouter()
const {t} = useI18n()

const email = ref('')
const password = ref('')
const errorValue = ref('')
const isShowError = ref(false)

const login = async () => {
    try {
       const res = await Login(email.value, password.value)
       if (res && res!== 0){
          let user:main.User | undefined = await getUser()
          if (user?.id){
            window.location.href = '/#/main'
          }
       }else {
        isShowError.value = true
        errorValue.value = "invalid email or password"
       }
    } catch (error:any) {
        isShowError.value = true
        errorValue.value = error        
    }
  
}
</script>

<style scoped></style>