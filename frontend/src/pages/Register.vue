<template>
    <div class="h-screen flex flex-col text-fontColor">
        <div class="m-auto">
            <div>
                <p>{{t('auth.username')}}: </p>
                <input class="border text-black" v-model="username" type="text">
            </div>
            <div>
                <p>{{t('auth.email')}}: </p>
                <input class="border text-black" v-model="email" type="text">
            </div>
            <div class="last-name">
                <p>last name: </p>
                <input class="border text-black" v-model="lastName" type="text">
            </div>
            <div>
                <p>{{t('auth.password')}}: </p>
                <input class="border text-black" v-model="password" type="password">
            </div>
            <div>
                <p>{{t('auth.confirm_password')}}: </p>
                <input class="border text-black" v-model="passwordConfirm" type="password">
            </div>
            <button class="bg-myMessage px-2 py-1 rounded-lg hover:border border-slate-400 mt-2" @click="register">
                {{t('auth.register')}}
            </button>
            <p class="text-sm text-red-500 whitespace-pre-line" v-if="isShowError">{{ errorValue }}</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useI18n } from 'vue-i18n';

const {t} = useI18n()

const email = ref('')
const username = ref('')
const password = ref('')
const passwordConfirm = ref('')
const lastName = ref('')
const isShowError = ref(false)
const errorValue = ref('')

const register = () => {
    
}
</script>

<style scoped>

</style>