<template>
    <div>
        <div class="flex py-4 mx-2 justify-between items-center">
            <div class="flex gap-2">
                <div @click="goBack" class="flex flex-col border hover:cursor-pointer">
                    <span>Back</span>
                    <span><-</span>
                </div>
                <div @click="goForward" class="flex flex-col border hover:cursor-pointer">
                    <span>Forward</span>
                    <span>-></span>
                </div>
            </div>
            <div class="flex flex-row justify-around gap-1">
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
            </div>
        </div>
        <NewProjectModal v-if="isCreateNewProject" @close-modal="toggleCreateProjectModal" />
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from 'vue';
import { languages } from '../consts';
import { useI18n } from 'vue-i18n';
import { EventsOn } from '../../wailsjs/runtime/runtime';
import { useRouter } from 'vue-router';
import NewProjectModal from './NewProjectModal.vue';
const { t } = useI18n()
const { locale } = useI18n()
const router = useRouter()

type Message = {
  message: string;
}

const selectedLanguage = ref("en")
const messageData = ref("")
const colorTheme = ref("light")
//TODO SETTINGS DB TABLE
document.body.classList.add(colorTheme.value)
const isCreateNewProject = ref(false)

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
  EventsOn("notification", (msg: Message) => {
    messageData.value = msg.message
  })
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
</script>

<style scoped></style>