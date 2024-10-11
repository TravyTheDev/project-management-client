<template>
    <div class="overflow-hidden grid grid-cols-4 whitespace-pre-line">
        <div class="col-span-3">
            <ProjectSkeleton v-if="isLoading" />
            <div class="flex flex-col gap-1 items-start overflow-y-scroll h-[87vh] pb-2" v-else-if="project">
                <div class="flex flex-col gap-1">
                    <div class="relative">
                        <span class="whitespace-pre">Assignee: </span>
                        <input v-if="isEdit" v-model="searchText" type="text">
                        <div v-if="users && isEdit && !isSetAssignee">
                            <div class="absolute ml-[70.84px] bg-white z-10 mt-1">
                                <div @click="setAssignee(user)"
                                    class="w-48 px-2 truncate border border-black divide-y hover:cursor-pointer hover:bg-slate-300"
                                    v-for="user in users">
                                    {{ user.username }}
                                </div>
                            </div>
                        </div>
                        <span v-else-if="!isEdit">{{ projectUser?.username }}</span>
                    </div>
                    <div class="flex gap-1 items-center">
                        <p>Start date:</p>
                        <input v-if="isEdit" v-model="project.startDate" type="date">
                        <p v-else>{{ project.startDate }}</p>
                        <p>End date:</p>
                        <input v-if="isEdit" v-model="project.endDate" type="date">
                        <p v-else>{{ project.endDate }}</p>
                    </div>

                    <div class="flex gap-1 items-center">
                        <p>Status:</p>
                        <select class="text-black" v-model="project.status" :disabled="!isEdit">
                            <option v-for="todo in todoStatus" :value="todo.value">{{ todo.label }}</option>
                        </select>
                        <p>Urgency:</p>
                        <select class="text-black" v-model="project.urgency" :disabled="!isEdit">
                            <option v-for="status in urgencyStatus" :value="status.value">{{ status.label }}</option>
                        </select>
                        <div>
                            <button class="border px-2 py-1" v-if="!isEdit" @click="toggleIsEdit">EDIT</button>
                            <div v-else>
                                <button class="border px-2 py-1" @click="saveEdit">SAVE</button>
                                <button class="border px-2 py-1 ml-2" @click="toggleIsEdit">CANCEL</button>
                            </div>
                        </div>
                    </div>
                </div>

                <div>
                    <button @click="toggeNewProjectModal" class="border px-2 py-1">Add subtask</button>
                </div>
                <textarea @input="resize(ref(titleTextArea))" ref="titleTextArea" class="text-xl font-semibold w-full"
                    rows="1" v-if="isEdit" v-model="project.title"></textarea>
                <h1 class="text-xl font-semibold" v-else>{{ project.title }}</h1>

                <p>Description:</p>
                <textarea @input="resize(ref(descriptionTextArea))" class="w-full" ref="descriptionTextArea" rows="1"
                    v-if="isEdit" v-model="project.description"></textarea>
                <h3 v-else>{{ project.description }}</h3>

                <p>Notes:</p>
                <textarea class="w-full" rows="1" @input="resize(ref(notesTextArea))" ref="notesTextArea" v-if="isEdit"
                    v-model="project.notes"></textarea>
                <p v-else>{{ project.notes }}</p>

                <div class="w-full border-t-2 mt-2">
                    <p>Personal notes:</p>
                    <div v-if="isEditPersonalNotes">
                        <textarea class="w-full" rows="1" @input="resize(ref(personalNotesRef))" ref="personalNotesRef"
                            v-model="newPersonalNotes"></textarea>
                        <div>
                            <button class="border px-2 py-1" @click="createOrEditPersonalNotes">SAVE</button>
                            <button class="border px-2 py-1 ml-2" @click="toggleIsEditPersonalNotes">CANCEL</button>
                        </div>
                    </div>
                    <div v-else>
                        <p class="whitespace-pre-line">{{ newPersonalNotes }}</p>
                        <button class="border px-2 py-1" @click="toggleIsEditPersonalNotes">EDIT</button>
                    </div>
                </div>
            </div>
        </div>
        <div v-if="childProjects?.length" class="border-l-2 px-1 h-[87vh] pb-2 overflow-y-auto divide-y">
            <div v-for="child in childProjects">
                <div @click="goToChild(child.id)" class="hover:cursor-pointer">
                    <p class="font-semibold">{{ child.title }}</p>
                    <p>Status:</p>
                    <p>{{ todoStatus[child.status].label }}</p>
                    <p>Urgency:</p>
                    <p>{{ urgencyStatus[child.urgency].label }}</p>
                </div>
            </div>
        </div>
        <NewProjectModal v-if="isShowNewProjectModal" :parent="project" :parent-i-d="id"
            @close-modal="toggeNewProjectModal" @load-children="getChildProjects" />
    </div>
</template>

<script setup lang="ts">
import { nextTick, onMounted, Ref, ref, watch, inject } from 'vue';
import { main, projects } from '../../wailsjs/go/models';
import { useRouter, useRoute } from 'vue-router';
import { CreateNotes, EditPersonalNotes, EditProject, GetChildProjectsByParentID, GetNotesByProjectID, GetProjectByID, SearchProjectAssignee } from '../../wailsjs/go/projects/ProjectsHandler';
import ProjectSkeleton from '../components/ProjectSkeleton.vue';
import { todoStatus, urgencyStatus } from '../consts';
import NewProjectModal from '../components/NewProjectModal.vue';
import { SendNotification } from '../../wailsjs/go/main/App';

const loginUser = inject<Ref<projects.User | undefined>>("loginUser")
const router = useRouter()
const route = useRoute()

const id = route.params.id.toString()

const projectRes = ref<projects.ProjectRes>()
const project = ref<projects.Project>()
const projectUser = ref<projects.User>()
const prevAssignee = ref(projectUser.value)
const users = ref<projects.User[]>()
const personalNotes = ref("")
const newPersonalNotes = ref("")
const isEditPersonalNotes = ref(false)
const isLoading = ref(false)
const searchText = ref("")
const isEdit = ref(false)
const isSetAssignee = ref(false)
const childProjects = ref<projects.Project[]>()
const titleTextArea = ref()
const descriptionTextArea = ref()
const notesTextArea = ref()
const personalNotesRef = ref()
const isShowNewProjectModal = ref(false)
const notificationMessage = ref<main.Notification>({
    id: Number(id),
    message: ''
})

const getProject = async () => {
    isLoading.value = true
    const res = await GetProjectByID(id)
    projectRes.value = res
    project.value = res.project
    projectUser.value = res.user
    searchText.value = projectUser.value?.username ?? ''
    isLoading.value = false
}

const toggeNewProjectModal = () => {
    isShowNewProjectModal.value = !isShowNewProjectModal.value
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
    nextTick(() => {
        resize(personalNotesRef)
    })
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
    nextTick(() => {
        resize(titleTextArea)
        resize(descriptionTextArea)
        resize(notesTextArea)
    })
}

const saveEdit = async () => {
    if (project.value) {
        project.value.id = Number(id)
        await EditProject(project.value)
        isEdit.value = false
        if(prevAssignee.value?.id !== project.value.assigneeID && project.value.assigneeID !== loginUser?.value?.id){
            sendNotification()
        }
    }
}

const sendNotification = () => {
    notificationMessage.value = {
        id: Number(id),
        message: `You've been assigned to ${project.value?.title}`
    }
    SendNotification(Number(project.value?.assigneeID), notificationMessage.value)
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

const resize = (refName: Ref) => {
    refName.value.style.height = 'auto'
    if (refName.value.scrollHeight < 104) {
        refName.value.style.height = refName.value.scrollHeight + 'px'
    } else {
        refName.value.style.height = '6.5rem'
    }
    if (refName.value.scrollHeight > 40) {
        refName.value.classList.add('expanded-text')
    } else {
        refName.value.classList.remove('expanded-text')
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

.expanded-text {
    overflow-y: auto;
}
</style>