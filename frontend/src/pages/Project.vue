<template>
    <div>
        <div>
            <ProjectSkeleton v-if="isLoading" />
            <div v-else>
                <h1>{{ project?.title }}</h1>
                <h3>{{ project?.description }}</h3>
                <p>{{ project?.status }}</p>
                <input class="border" v-model="searchText" type="text">
                <div v-if="users">
                    <div v-for="user in users">
                        {{ user.username }}
                    </div>
                </div>
                <p>{{ project?.urgency }}</p>
                <p>{{ project?.notes }}</p>
            </div>
        </div>
        <div>
            <div v-if="isEditPersonalNotes">
                <textarea v-model="newPersonalNotes"></textarea>
                <button @click="createOrEditPersonalNotes">save</button>
                <button @click="toggleIsEditPersonalNotes">cancel</button>
            </div>
            <div v-else>
                <p class="whitespace-pre-line">{{ newPersonalNotes }}</p>
                <button @click="toggleIsEditPersonalNotes">edit</button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from 'vue';
import { projects } from '../../wailsjs/go/models';
import { useRouter, useRoute } from 'vue-router';
import { CreateNotes, EditPersonalNotes, GetNotesByProjectID, GetProjectByID, SearchProjectAssignee } from '../../wailsjs/go/projects/ProjectsHandler';
import ProjectSkeleton from '../components/ProjectSkeleton.vue';

const router = useRouter()
const route = useRoute()

const id = route.params.id.toString()

const project = ref<projects.Project>()
const personalNotes = ref("")
const newPersonalNotes = ref("")
const isEditPersonalNotes = ref(false)
const isLoading = ref(false)
const searchText = ref("")
const users = ref<projects.User[]>()

const getProject = async () => {
    isLoading.value = true
    const res = await GetProjectByID(id)
    project.value = res
    isLoading.value = false
}

onMounted(() => {
    getProject()
})

const toggleIsEditPersonalNotes = () => {
    isEditPersonalNotes.value = !isEditPersonalNotes.value
}

const createPersonalNotes = () => {
    if (project.value) {
        CreateNotes(project.value.id, newPersonalNotes.value)
    }
}

const editPersonalNotes = () => {
    if (project.value) {
        EditPersonalNotes(project.value.id, newPersonalNotes.value)
    }
}

const createOrEditPersonalNotes = () => {
    if (personalNotes.value) {
        editPersonalNotes()
    } else {
        createPersonalNotes()
    }
    toggleIsEditPersonalNotes()
}

watch(project, () => {
    if (project.value) {
        getPersonalNotes()
    }
})

const getPersonalNotes = async () => {
    if (project.value) {
        const res = await GetNotesByProjectID(project.value.id)
        personalNotes.value = res
    }
}

watch(personalNotes, () => {
    if (personalNotes.value) {
        newPersonalNotes.value = personalNotes.value
    }
})

const handleSearch = (value: string) => {
    const timeoutID: number = window.setTimeout(() => {}, 0)

    for (let id:number = timeoutID; id >= 0; id -= 1) {
        window.clearTimeout(id)
    }

    setTimeout(() => {
        searchUser(value)
    }, 300)
}

const searchUser = async (value:string) => {
    const req: projects.SearchReq = {
        text: value
    }
   const res = await SearchProjectAssignee(req)
   users.value = res
}

watch(searchText, () => {
    if(searchText.value){
        handleSearch(searchText.value)
    }
})

</script>

<style scoped></style>