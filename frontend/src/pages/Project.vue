<template>
    <div class="overflow-hidden grid grid-cols-4 whitespace-pre-line border-l-2 pl-1">
        <div class="col-span-3">
            <ProjectSkeleton v-if="isLoading" />
            <div class="flex flex-col gap-1 items-start divide-y-2 overflow-y-scroll h-[87vh] pb-2" v-else-if="project">
                <div class="flex flex-col gap-1">
                    <div class="relative">
                        <span @click="editProjectUser" class="whitespace-pre">Assignee: </span>
                        <input ref="projectUserInputRef" @blur="endEditProjectUser" v-if="isEditProjectUser"
                            v-model="searchText" type="text">
                        <div v-if="users && isEditProjectUser && !isSetAssignee">
                            <div class="absolute ml-[70.84px] bg-white z-10 mt-1">
                                <div @mousedown="setAssignee(user)"
                                    class="w-48 px-2 truncate border border-black divide-y hover:cursor-pointer hover:bg-slate-300"
                                    v-for="user in users">
                                    {{ user.username }}
                                </div>
                            </div>
                        </div>
                        <span @click="editProjectUser" v-else-if="!isEditProjectUser">{{ projectUser?.username }}</span>
                    </div>
                    <div class="flex gap-1 items-center">
                        <p @click="!project.startDate ? handleIsEdit('start_date') : null">Start date:</p>
                        <input ref="startDateRef" @blur="handleIsEdit('start_date')" v-if="isEditStartDate"
                            v-model="project.startDate" type="date">
                        <p @click="handleIsEdit('start_date')" v-else>{{ project.startDate }}</p>
                        <p @click="!project.endDate ? handleIsEdit('end_date') : null">End date:</p>
                        <input ref="endDateRef" @blur="handleIsEdit('end_date')" v-if="isEditEndDate"
                            v-model="project.endDate" type="date">
                        <p @click="handleIsEdit('end_date')" v-else>{{ project.endDate }}</p>
                    </div>

                    <div class="flex gap-1 items-center">
                        <p>Status:</p>
                        <select @change="saveEdit" class="text-black" v-model="project.status">
                            <option v-for="todo in todoStatus" :value="todo.value">{{ todo.label }}</option>
                        </select>
                        <p>Urgency:</p>
                        <select @change="saveEdit" class="text-black" v-model="project.urgency" :disabled="!isEditUrgency">
                            <option v-for="status in urgencyStatus" :value="status.value">{{ status.label }}</option>
                        </select>
                    </div>
                </div>

                <div class="w-full pt-2 py-1">
                    <button @click="toggeNewProjectModal" class="border px-2 py-1">Add subtask</button>
                </div>
                <div class="w-full">
                    <textarea @blur="handleIsEdit('title')" @input="resize(ref(titleTextArea)); handleShareNotes(project.title, 'title')" ref="titleTextArea"
                        class="text-xl font-semibold w-full" rows="1" v-if="isEditTitle"
                        v-model="project.title"></textarea>
                    <h1 @click="handleIsEdit('title')" class="text-xl font-semibold" v-else>{{ project.title }}</h1>
                </div>

                <div class="w-full">
                    <p @click="!project.description ? handleIsEdit('description') : null" class="font-semibold">
                        Description:</p>
                    <textarea @blur="handleIsEdit('description')" @input="resize(ref(descriptionTextArea)); handleShareNotes(project.description, 'description')"
                        class="w-full" ref="descriptionTextArea" rows="1" v-if="isEditDescription"
                        v-model="project.description"></textarea>
                    <h3 @click="handleIsEdit('description')" v-else>{{ project.description }}</h3>
                </div>

                <div class="w-full">
                    <p @click="!project.notes ? handleIsEdit('notes') : null" class="font-semibold">Notes:</p>
                    <textarea @blur="handleIsEdit('notes')" class="w-full" rows="1"
                        @input="resize(ref(notesTextArea)); handleShareNotes(project.notes, 'notes')" ref="notesTextArea"
                        v-if="isEditNotes" v-model="project.notes"></textarea>
                    <p @click="handleIsEdit('notes')" v-else>{{ project.notes }}</p>
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

const isEditTitle = ref(false)
const isEditDescription = ref(false)
const isEditNotes = ref(false)
const isEditProjectUser = ref(false)
const isEditStartDate = ref(false)
const isEditEndDate = ref(false)
const isEditStatus = ref(false)
const isEditUrgency = ref(false)

const isSetAssignee = ref(false)
const childProjects = ref<types.Project[]>()
const titleTextArea = ref()
const descriptionTextArea = ref()
const notesTextArea = ref()
const personalNotesRef = ref()
const projectUserInputRef = ref()
const startDateRef = ref()
const endDateRef = ref()

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

const handleIsEdit = (section: string) => {
    switch (section) {
        case "title":
            isEditTitle.value = !isEditTitle.value
            nextTick(() => {
                handleResizeAndFocus(isEditTitle, titleTextArea)
            })
            return
        case "description":
            isEditDescription.value = !isEditDescription.value
            nextTick(() => {
                handleResizeAndFocus(isEditDescription, descriptionTextArea)
            })
            return
        case "notes":
            isEditNotes.value = !isEditNotes.value
            nextTick(() => {
                handleResizeAndFocus(isEditNotes, notesTextArea)
            })
            return
        case "start_date":
            isEditStartDate.value = !isEditStartDate.value
            nextTick(() => {
                handleEditNonText(isEditStartDate, startDateRef)
            })
            return
        case "end_date":
            isEditEndDate.value = !isEditEndDate.value
            nextTick(() => {
                handleEditNonText(isEditEndDate, endDateRef)
            })
            return
        case "end_date":
            isEditEndDate.value = !isEditEndDate.value
            nextTick(() => {
                handleEditNonText(isEditEndDate, endDateRef)
            })
            return
    }
}

const handleResizeAndFocus = async (isEditRefVal: Ref, textAreaRefVal: Ref) => {
    if (isEditRefVal.value) {
        resize(textAreaRefVal)
        textAreaRefVal.value.focus()
    } else {
        await saveEdit()
    }
}

const editProjectUser = () => {
    isEditProjectUser.value = true
    nextTick(() => {
        projectUserInputRef.value.focus()
    })
}

const endEditProjectUser = () => {
    if (!isSetAssignee.value) {
        searchText.value = projectRes.value?.user?.username ?? ''
    }
    isEditProjectUser.value = false
}

const handleEditNonText = async (isEditRefVal: Ref, refVal: Ref) => {
    if (isEditRefVal.value) {
        refVal.value.focus()
    } else {
        await saveEdit()
    }
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
            //This is disgusting
            const message = JSON.parse(JSON.parse(msg.body))
            if(message.area === "title"){
                project.value.title = message.body
            }
            if(message.area === "description"){
                project.value.description = message.body
            }
            if(message.area === "notes"){
                project.value.notes = message.body
            }
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

const saveEdit = async () => {
    if (project.value) {
        project.value.id = Number(id)
        await EditProject(project.value)
    }
}

const sendNotification = () => {
    notificationMessage.value = {
        id: Number(id),
        message: `You've been assigned to ${project.value?.title}`
    }
    SendNotification(Number(project.value?.assigneeID), notificationMessage.value)
}

const setAssignee = async (user: types.User) => {
    if (project.value) {
        project.value.assigneeID = user.id
        searchText.value = user.username
        projectUser.value = user
        isSetAssignee.value = true
        if (
        project.value.assigneeID !== 0 &&
        prevAssignee.value?.id !== 
        project.value.assigneeID && 
        project.value.assigneeID !== 
        loginUser?.value?.id
        ) {
            sendNotification()
        }
        await nextTick(() => {
            saveEdit()
        })
    }
    if (users.value) {
        users.value.length = 0
    }
}

const resize = (refName: Ref) => {
    if (!refName.value) {
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

const handleShareNotes = (value: string, area: string) => {
    const messageBody = {
            area: area,
            body: value,
        }
    const json = JSON.stringify(messageBody)
    debounceFunc(json, sendWebSocketMessage)
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