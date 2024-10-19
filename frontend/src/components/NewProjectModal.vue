<template>
    <div @click="emit('close-modal')"
        class="absolute w-full bg-black bg-opacity-30 h-screen top-0 left-0 flex flex-col px-8">
        <div @click.stop class="z-20 relative p-4 bg-white text-black self-start m-auto top-0 bottom-0 w-[85vw]">
            <div>
                <div class="flex flex-col gap-1">
                    <div v-if="parent && parentID">
                        <p>Parent:</p>
                        <h2 class="whitespace-pre-line font-semibold">{{ parent.title }}</h2>
                    </div>
                    <div class="relative">
                        <span class="whitespace-pre">Assignee: </span>
                        <input v-model="searchText" type="text">
                        <div v-if="users && !isSetAssignee">
                            <div class="absolute ml-[70.84px] bg-white z-10 mt-1">
                                <div @click="setAssignee(user)"
                                    class="w-48 px-2 truncate border border-black divide-y hover:cursor-pointer hover:bg-slate-300"
                                    v-for="user in users">
                                    {{ user.username }}
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="flex gap-1 items-center">
                        <p>Start date:</p>
                        <input v-model="project.startDate" type="date">
                        <p>End date:</p>
                        <input v-model="project.endDate" type="date">
                    </div>

                    <div class="flex gap-1 items-center">
                        <p>Status:</p>
                        <select class="text-black" v-model="project.status">
                            <option v-for="todo in todoStatus" :value="todo.value">{{ todo.label }}</option>
                        </select>
                        <p>Urgency:</p>
                        <select class="text-black" v-model="project.urgency">
                            <option v-for="status in urgencyStatus" :value="status.value">{{ status.label }}</option>
                        </select>
                    </div>
                </div>

                <p>Title:</p>
                <textarea @input="resize(ref(titleTextArea))" ref="titleTextArea" class="text-xl font-semibold w-full"
                    rows="1" v-model="project.title"></textarea>
                <p v-if="noTitle" class="text-red-500 text-sm">title required</p>
                <p>Description:</p>
                <textarea @input="resize(ref(descriptionTextArea))" class="w-full" ref="descriptionTextArea" rows="1"
                    v-model="project.description"></textarea>

                <p>Notes:</p>
                <textarea class="w-full" rows="1" @input="resize(ref(notesTextArea))" ref="notesTextArea"
                    v-model="project.notes"></textarea>

                <div>
                    <div>
                        <button :disabled="project.title === ''" class="border px-2 py-1" @click="createProject">SAVE</button>
                        <button class="border px-2 py-1 ml-2" @click="emit('close-modal')">CANCEL</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { inject, reactive, ref, Ref, watch } from 'vue';
import { CreateProject, SearchProjectAssignee } from '../../wailsjs/go/projects/ProjectsHandler';
import { types } from '../../wailsjs/go/models';
import { useRouter, useRoute } from 'vue-router';
import { todoStatus, urgencyStatus } from '../consts';
import { debounceFunc } from '../functions';

interface props {
    parent: types.Project | undefined;
    parentID: string | undefined;
}

const props = defineProps<props>()

const route = useRoute()
const router = useRouter()

//TODO ASSIGN SELF
const loginUser = inject<Ref<types.User | undefined>>("loginUser")

const emit = defineEmits(['close-modal', 'load-children'])
const titleTextArea = ref()
const descriptionTextArea = ref()
const notesTextArea = ref()
const searchText = ref("")
const users = ref<types.User[]>()
const isSetAssignee = ref(false)
const chosenUser = ref<types.User>()
const noTitle = ref(false)

const project: types.ProjectReq = reactive({
    parentID: Number(props.parentID) ?? 0,
    title: "",
    description: "",
    status: 0,
    assigneeID: 0,
    urgency: 0,
    notes: "",
    startDate: "",
    endDate: "",
})

const createProject = async () => {
    if(project.title === ''){
        noTitle.value = true
        return
    }
    await CreateProject(project)
    emit('close-modal')
    if (props.parentID){
        emit('load-children')
    }else{
        if (route.path === '/main'){
            router.push('/main/reload')
        }else {
            router.push('/main')
        }
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

const handleSearch = (value: string) => {
    if (value === chosenUser.value?.username) {
        return
    }
    isSetAssignee.value = false
    debounceFunc(value, searchUser)
}

const searchUser = async (value: string) => {
    if (isSetAssignee.value) {
        return
    }
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

const setAssignee = (user: types.User) => {
    isSetAssignee.value = true
    project.assigneeID = user.id
    searchText.value = user.username
    chosenUser.value = user
    if (users.value) {
        users.value.length = 0
    }
}
</script>

<style scoped>
input,
textarea {
    border: 1px solid black;
}
</style>