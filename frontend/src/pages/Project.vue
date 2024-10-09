<template>
    <div class="overflow-hidden grid grid-cols-4 whitespace-pre-line">
        <div class="col-span-3">
            <ProjectSkeleton v-if="isLoading" />
            <div class="flex flex-col gap-1 items-start overflow-y-scroll h-[87vh] pb-2" v-else-if="project">
                <p>Title:</p>
                <textarea class="text-xl font-semibold w-full" rows="3" v-if="isEdit"
                    v-model="project.title"></textarea>
                <h1 class="text-xl font-semibold mb-14" v-else>{{ project.title }}</h1>

                <p>Description:</p>
                <textarea class="w-full" rows="3" v-if="isEdit" v-model="project.description"></textarea>
                <h3 class="mb-12" v-else>{{ project.description }}</h3>

                <div class="flex gap-1 mt-4">
                    <p>Status:</p>
                    <select class="text-black" v-model="project.status" :disabled="!isEdit">
                        <option v-for="todo in todoStatus" :value="todo.value">{{ todo.label }}</option>
                    </select>
                    <p>Urgency:</p>
                    <select class="text-black" v-model="project.urgency" :disabled="!isEdit">
                        <option v-for="status in urgencyStatus" :value="status.value">{{ status.label }}</option>
                    </select>
                </div>

                <div class="relative">
                    <p>Assignee:</p>
                    <input v-if="isEdit" v-model="searchText" type="text">
                    <div v-if="users && isEdit && !isSetAssignee">
                        <div class="absolute bg-white z-10 mt-1">
                            <div @click="setAssignee(user)"
                                class="w-52 truncate border border-black divide-y hover:cursor-pointer hover:bg-slate-300"
                                v-for="user in users">
                                {{ user.username }}
                            </div>
                        </div>
                    </div>
                    <p v-else-if="!isEdit">{{ projectUser?.username }}</p>
                </div>

                <div class="flex gap-1 mt-4">
                    <p>Start date:</p>
                    <input v-if="isEdit" v-model="project.startDate" type="date">
                    <p v-else>{{ project.startDate }}</p>
                    <p>End date:</p>
                    <input v-if="isEdit" v-model="project.endDate" type="date">
                    <p v-else>{{ project.endDate }}</p>
                </div>

                <p>Notes:</p>
                <textarea v-if="isEdit" v-model="project.notes"></textarea>
                <p v-else>{{ project.notes }}</p>
                <div>
                    <button v-if="!isEdit" @click="toggleIsEdit">EDIT</button>
                    <button v-else @click="saveEdit">SAVE</button>
                </div>
                <div>
                    <p>Personal notes:</p>
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
        </div>
        <div v-if="childProjects?.length" class="border-l-2 px-1 h-[87vh] pb-2 overflow-y-auto">
            <div v-for="child in childProjects">
                <div @click="goToChild(child.id)" class="hover:cursor-pointer">
                    <p>Title:</p>
                    <p class="font-semibold">{{ child.title }}</p>
                    <p>Status:</p>
                    <p>{{ todoStatus[child.status].label }}</p>
                    <p>Urgency:</p>
                    <p>{{ urgencyStatus[child.urgency].label }}</p>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from 'vue';
import { projects } from '../../wailsjs/go/models';
import { useRouter, useRoute } from 'vue-router';
import { CreateNotes, EditPersonalNotes, EditProject, GetChildProjectsByParentID, GetNotesByProjectID, GetProjectByID, SearchProjectAssignee } from '../../wailsjs/go/projects/ProjectsHandler';
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
const isSetAssignee = ref(false)
const childProjects = ref<projects.Project[]>()

const getProject = async () => {
    isLoading.value = true
    const res = await GetProjectByID(id)
    projectRes.value = res
    project.value = res.project
    projectUser.value = res.user
    searchText.value = projectUser.value?.username ?? ''
    isLoading.value = false
}

const getChildProjects = async () => {
    const res = await GetChildProjectsByParentID(Number(id))
    childProjects.value = res
}

const goToChild = (id: number) => {
    router.push(`/project/${id}`)
}

onMounted(() => {
    getProject()
    getChildProjects()
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
    if (value === projectUser.value?.username) {
        return
    }
    isSetAssignee.value = false
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

const setAssignee = (user: projects.User) => {
    if (project.value) {
        project.value.assigneeID = user.id
        searchText.value = user.username
        projectUser.value = user
        isSetAssignee.value = true
    }
    if (users.value) {
        users.value.length = 0
    }
}

</script>

<style scoped>
input {
    border: 1px solid black;
}

textarea,
input {
    color: black;
}
</style>