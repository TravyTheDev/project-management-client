<template>
    <div class="overflow-hidden grid grid-cols-4 whitespace-pre-line border-l-2 pl-1">
        <div class="col-span-3">
            <ProjectSkeleton v-if="isLoading" />
            <div class="flex flex-col gap-1 items-start divide-y-2 overflow-y-scroll h-[87vh] pb-2" v-else-if="project">
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
                                <button class="border px-2 py-1 ml-2" @click="cancelEdit">CANCEL</button>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="w-full pt-2 py-1">
                    <button @click="toggeNewProjectModal" class="border px-2 py-1">Add subtask</button>
                </div>
                <div class="w-full">
                    <textarea @input="resize(ref(titleTextArea))" ref="titleTextArea" class="text-xl font-semibold w-full"
                        rows="1" v-if="isEdit" v-model="project.title"></textarea>
                    <h1 class="text-xl font-semibold" v-else>{{ project.title }}</h1>
                </div>

                <div class="w-full">
                    <p class="font-semibold">Description:</p>
                    <textarea @input="resize(ref(descriptionTextArea))" class="w-full" ref="descriptionTextArea" rows="1"
                        v-if="isEdit" v-model="project.description"></textarea>
                    <h3 v-else>{{ project.description }}</h3>
                </div>

                <div class="w-full">
                    <p class="font-semibold">Notes:</p>
                    <textarea class="w-full" rows="1" @input="resize(ref(notesTextArea)); handleShareNotes(project.notes)" ref="notesTextArea" v-if="isEdit"
                        v-model="project.notes"></textarea>
                    <p v-else>{{ project.notes }}</p>
                </div>

                <div class="w-full border-t-2 mt-2">
                    <p class="font-semibold">Personal notes:</p>
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
                <div class="w-full">
                    <h1 class="text-red-500 border-t-2 border-red-500 w-full">DANGER ZONE</h1>
                    <button @click="toggleDeleteModal" class="text-red-500 border border-red-500 mt-4">DELETE
                        PROJECT</button>
                </div>
            </div>
        </div>
        <div v-if="childProjects?.length" class="border-l-2 px-1 h-[87vh] pb-2 overflow-y-auto divide-y">
            <div v-for="child in childProjects">
                <div @click="goToChild(child.id)" class="hover:cursor-pointer">
                    <p class="font-semibold">{{ child.title }}</p>
                    <div>
                        <span>Status: </span>
                        <span>{{ todoStatus[child.status].label }}</span>
                    </div>
                    <div>
                        <span>Urgency: </span>
                        <span>{{ urgencyStatus[child.urgency].label }}</span>
                    </div>
                </div>
            </div>
        </div>
        <NewProjectModal v-if="isShowNewProjectModal" :parent="project" :parent-i-d="id"
            @close-modal="toggeNewProjectModal" @load-children="getChildProjects" />
        <DeleteProjectModal @close-modal="toggleDeleteModal" v-if="isShowDeleteModal" :title="project?.title"
            :parent-i-d="id" />
    </div>
</template>

<script setup lang="ts">
import { nextTick, onMounted, Ref, ref, watch, inject, onUnmounted } from 'vue';
import { types } from '../../wailsjs/go/models';
import { useRouter, useRoute } from 'vue-router';
import { CreateNotes, EditPersonalNotes, EditProject, GetChildProjectsByParentID, GetNotesByProjectID, GetProjectByID, SearchProjectAssignee } from '../../wailsjs/go/projects/ProjectsHandler';
import ProjectSkeleton from '../components/ProjectSkeleton.vue';
import { todoStatus, urgencyStatus } from '../consts';
import NewProjectModal from '../components/NewProjectModal.vue';
import { DisconnectWebSocketRoom, JoinWebSocketRoom, SendNotification, WriteSocketMessage } from '../../wailsjs/go/main/App';
import DeleteProjectModal from '../components/DeleteProjectModal.vue';
import { EventsOn } from '../../wailsjs/runtime/runtime';
import { debounceFunc } from '../functions';

const loginUser = inject<Ref<types.User | undefined>>("loginUser")
const router = useRouter()
const route = useRoute()

const id = route.params.id.toString()

const projectRes = ref<types.ProjectRes>()
const project = ref<types.Project>()
const projectUser = ref<types.User>()
const prevAssignee = ref(projectUser.value)
const users = ref<types.User[]>()
const personalNotes = ref("")
const newPersonalNotes = ref("")
const isEditPersonalNotes = ref(false)
const isLoading = ref(false)
const searchText = ref("")
const isEdit = ref(false)
const isSetAssignee = ref(false)
const childProjects = ref<types.Project[]>()
const titleTextArea = ref()
const descriptionTextArea = ref()
const notesTextArea = ref()
const personalNotesRef = ref()
const isShowNewProjectModal = ref(false)
const notificationMessage = ref<types.Notification>({
    id: Number(id),
    message: ''
})
const isShowDeleteModal = ref(false)

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
    if (loginUser?.value) {
        joinWebSocketRoom(loginUser.value)
    }
    EventsOn("websocket", (msg: types.SocketMessage) => {
        if (project.value) {
            project.value.notes = JSON.parse(msg.body)
        }
    })
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
    debounceFunc(value, searchUser)
}

const searchUser = async (value: string) => {
    const req: types.SearchReq = {
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
        if (prevAssignee.value?.id !== project.value.assigneeID && project.value.assigneeID !== loginUser?.value?.id) {
            sendNotification()
        }
    }
}

const cancelEdit = () => {
    //todo reset values
    getProject()
    isEdit.value = false
}

const sendNotification = () => {
    notificationMessage.value = {
        id: Number(id),
        message: `You've been assigned to ${project.value?.title}`
    }
    SendNotification(Number(project.value?.assigneeID), notificationMessage.value)
}

const setAssignee = (user: types.User) => {
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
    if (!refName.value){
        return
    }
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

const toggleDeleteModal = () => {
    isShowDeleteModal.value = !isShowDeleteModal.value
}

const joinWebSocketRoom = (user: types.User) => {
    JoinWebSocketRoom(Number(id), user.id, user.username)
}

//TODO
const sendWebSocketMessage = (value: string) => {
    if (!loginUser || loginUser.value === undefined) {
        return
    }
    if (project.value) {
        const message: types.SocketMessage = {
            body: value,
            roomID: id,
            userID: String(loginUser?.value.id),
            username: loginUser.value.username,
        }
        WriteSocketMessage(message)
    }
}

const handleShareNotes = (value: string) => {
    debounceFunc(value, sendWebSocketMessage)
}

const disconnectWebSocket = async () => {
    await DisconnectWebSocketRoom()
}

onUnmounted(() => {
    disconnectWebSocket()
})

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