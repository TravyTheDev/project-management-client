<template>
    <div>
        <div class="flex py-4 mx-2 justify-between items-center relative">
            <span v-if="isShowNotifications" @click="closeNotifications" class="absolute top-0 right-0 left-0 h-[100vh] w-[100vw]"></span>
            <div class="flex gap-2">
                <div @click="goBack" class="flex flex-col border hover:cursor-pointer w-20 items-center">
                    <span>Back</span>
                    <span><<<</span>
                </div>
                <div @click="goForward" class="flex flex-col border hover:cursor-pointer w-20 items-center">
                    <span>Forward</span>
                    <span>>>></span>
                </div>
            </div>
            <div class="flex flex-row justify-around gap-1">
                <div v-if="messageData.length" class="relative">
                    <span @click="showNotifications" class="hover:cursor-pointer bg-red-500 rounded-full text-white px-1 text-sm">{{ messageData.length }}</span>
                    <div @click.stop class="absolute w-44 z-10 bg-white max-h-40 overflow-y-auto translate-x-4" v-if="isShowNotifications">
                        <div class="border" v-for="message in messageData">
                            <Notifications :message="message" />
                        </div>
                    </div>
                </div>
                <div>
                    <span>Search:</span>
                    <input class="border" type="text" name="" id="">
                </div>
                <div>
                    <label v-if="colorTheme == 'light' || !colorTheme" for="dark_mode">{{ t('navbar.dark_mode') }}</label>
                    <input class="hidden" value="dark" v-model="colorTheme" name="theme_choice" id="dark_mode" type="radio">
                    <label v-if="colorTheme == 'dark'" for="light_mode">{{t('navbar.light_mode')}}</label>
                    <input class="hidden" value="light" v-model="colorTheme" name="theme_choice" id="light_mode"
                        type="radio">
                </div>
                <select class="text-black" v-model="selectedLanguage">
                    <option v-for="lang in languages" :value="lang.value">{{ lang.label }}</option>
                </select>
                <div>
                    <button @click="toggleCreateProjectModal">New Project</button>
                </div>
                <div>
                    <button @click="logout">Logout</button>
                </div>
            </div>
        </div>
        <NewProjectModal :parent="undefined" :parent-i-d="undefined" v-if="isCreateNewProject" @close-modal="toggleCreateProjectModal" />
    </div>
</template>

<script setup lang="ts">
import { nextTick, onMounted, onUnmounted, ref, watch } from 'vue';
import { languages } from '../consts';
import { useI18n } from 'vue-i18n';
import { EventsOff, EventsOn } from '../../wailsjs/runtime/runtime';
import { useRouter } from 'vue-router';
import NewProjectModal from './NewProjectModal.vue';
import { Logout } from '../../wailsjs/go/main/App';
import Notifications from './Notifications.vue';
import { main } from '../../wailsjs/go/models';


const { t } = useI18n()
const { locale } = useI18n()
const router = useRouter()

const selectedLanguage = ref("en")
const messageData = ref<main.Notification[]>([])
const colorTheme = ref("light")
//TODO SETTINGS DB TABLE
//TODO SEARCH PROJECTS
//TODO NOTIFICATIONS
document.body.classList.add(colorTheme.value)
const isCreateNewProject = ref(false)
const isShowNotifications = ref(false)

const toggleCreateProjectModal = () => {
    isCreateNewProject.value = !isCreateNewProject.value
}

const goForward = () => {
    router.go(1)
}

const goBack = () => {
    router.go(-1)
}


onMounted(() => {
  EventsOn("notification", (msg: main.Notification) => {
    messageData.value.push(msg) 
    
  })
})

onUnmounted(() => {
    EventsOff("notification")
})

watch(selectedLanguage, () => {
    if(selectedLanguage.value){
        locale.value = selectedLanguage.value
    }
})

watch(colorTheme, (newVal, oldVal) => {
    if (colorTheme.value) {
        document.body.classList.add(newVal)
        if (document.body.classList.contains(oldVal)) {
            document.body.classList.remove(oldVal)
        }
    }
})

const logout = async () => {
    await Logout()
    nextTick(() => {
        router.push('/login')
    })
}

const showNotifications = () => {
    isShowNotifications.value = true
}

const closeNotifications = () => {
    isShowNotifications.value = false
    messageData.value.length = 0
}
</script>

<style scoped></style>