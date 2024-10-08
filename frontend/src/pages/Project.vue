<template>
    <div>
        <div>
            <ProjectSkeleton v-if="isLoading" />
            <div class="flex flex-col gap-1 items-start" v-else-if="project">
                <input v-if="isEdit" v-model="project.title" type="text">
                <h1 v-else>{{ project.title }}</h1>
                <input v-if="isEdit" v-model="project.description" type="text">
                <h3 v-else>{{ project.description }}</h3>
                <select v-model="project.status" :disabled="!isEdit">
                    <option v-for="todo in todoStatus" :value="todo.value">{{ todo.label }}</option>
                </select>
                <input v-if="isEdit" v-model="searchText" type="text">
                <div v-if="users && isEdit">
                    <div v-for="user in users">
                        {{ user.username }}
                    </div>
                </div>
                <p v-else-if="!isEdit">{{ projectUser?.username }}</p>
                <select v-model="project.urgency" :disabled="!isEdit">
                    <option v-for="status in urgencyStatus" :value="status.value">{{ status.label }}</option>
                </select>
                <textarea v-if="isEdit" v-model="project.notes"></textarea>
                <p v-else>{{ project.notes }}</p>
            </div>
            <div>
                <button v-if="!isEdit" @click="toggleIsEdit">EDIT</button>
                <button v-else @click="saveEdit">SAVE</button>
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
import { CreateNotes, EditPersonalNotes, EditProject, GetNotesByProjectID, GetProjectByID, SearchProjectAssignee } from '../../wailsjs/go/projects/ProjectsHandler';
import ProjectSkeleton from '../components/ProjectSkeleton.vue';
import { todoStatus, urgencyStatus } from '../consts';

const router = useRouter()
const route = useRoute()

const id = route.params.id.toString()

const projectRes = ref<projects.ProjectRes>()
const project = ref<projects.Project>()
const projectUser = ref<projects.User>()
const users = ref<projects.User[]>()
const personalNotes = ref("")
const newPersonalNotes = ref("")
const isEditPersonalNotes = ref(false)
const isLoading = ref(false)
const searchText = ref("")
const isEdit = ref(false)

const getProject = async () => {
    isLoading.value = true
    const res = await GetProjectByID(id)
    projectRes.value = res
    project.value = res.project
    projectUser.value = res.user
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
        CreateNotes(Number(id), newPersonalNotes.value)
    }
}

const editPersonalNotes = () => {
    if (project.value) {
        EditPersonalNotes(Number(id), newPersonalNotes.value)
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
        const res = await GetNotesByProjectID(Number(id))
        personalNotes.value = res
    }
}

watch(personalNotes, () => {
    if (personalNotes.value) {
        newPersonalNotes.value = personalNotes.value
    }
})

const handleSearch = (value: string) => {
    const timeoutID: number = window.setTimeout(() => { }, 0)

    for (let id: number = timeoutID; id >= 0; id -= 1) {
        window.clearTimeout(id)
    }

    setTimeout(() => {
        if (value.length < 2) {
            value = ''
        }
        searchUser(value)
    }, 300)
}

const searchUser = async (value: string) => {
    const req: projects.SearchReq = {
        text: value
    }
    const res = await SearchProjectAssignee(req)
    users.value = res
}

watch(searchText, () => {
    if (searchText.value) {
        handleSearch(searchText.value)
    }
})

const toggleIsEdit = () => {
    isEdit.value = !isEdit.value
}

const saveEdit = async () => {
    if (project.value) {
        project.value.id = Number(id)
        await EditProject(project.value)
        isEdit.value = false
    }
}

</script>

<style scoped>
input {
    border: 1px solid black;
}
</style>