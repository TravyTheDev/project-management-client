<template>
    <div class="h-screen flex flex-col text-fontColor">
        <div class="mx-auto mt-16">
            <div>
                <p>{{t('auth.username')}}: </p>
                <input class="border text-black" v-model="registerUser.username" type="text">
            </div>
            <div>
                <p>{{t('auth.email')}}: </p>
                <input class="border text-black" v-model="registerUser.email" type="text">
            </div>
            <div>
                <p>{{t('auth.password')}}: </p>
                <input class="border text-black" v-model="registerUser.password" type="password">
            </div>
            <div>
                <p>{{t('auth.confirm_password')}}: </p>
                <input class="border text-black" v-model="registerUser.passwordConfirm" type="password">
            </div>
            <div class="flex items-center gap-2">
                <p>{{ t('auth.is_admin') }}:</p>
                <input v-model="registerUser.isAdmin" type="checkbox">
            </div>
            <button class="bg-myMessage px-2 py-1 rounded-lg hover:border border-slate-400 mt-2" @click="register">
                {{t('auth.register')}}
            </button>
            <p class="text-sm text-red-500 whitespace-pre-line" v-if="isShowError">{{ errorValue }}</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import { useI18n } from 'vue-i18n';
import { Register } from '../../wailsjs/go/main/App';
import { types } from '../../wailsjs/go/models';

const {t} = useI18n()

const defaults: types.RegisterReq = {
    email: '',
    username: '',
    password: '',
    passwordConfirm: '',
    isAdmin: false,
}

const registerUser = <types.RegisterReq>reactive({
    email: '',
    username: '',
    password: '',
    passwordConfirm: '',
    isAdmin: false,
})

const isShowError = ref(false)
const errorValue = ref('')

const register = async () => {
    try {
        await Register(registerUser)
        Object.assign(registerUser, defaults)
    } catch (error: any) {
        isShowError.value = true
        errorValue.value = error
    }
}
</script>

<style scoped>

</style>