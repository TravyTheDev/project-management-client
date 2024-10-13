<template>
    <div class="grid grid-cols-3 gap-4">
        <div>
            <h1 class="rounded-tr-lg rounded-tl-lg mx-auto text-center bg-orange-300 text-white w-2/3">TODO</h1>
            <div class="h-[87vh] pb-2 overflow-y-auto divide-y" v-if="!isLoadingTodoProjects">
                <div v-for="project in todoProjectsData">
                    <ProjectPreview :project="project.project" :user="project.user" @view-project="viewProject" />
                </div>
            </div>
            <div v-else>
                <ProjectPreviewSkeleton />
            </div>
        </div>
        <div>
            <h1 class="rounded-tr-lg rounded-tl-lg mx-auto text-center bg-cyan-600 text-white w-2/3">IN PROGRESS</h1>
            <div class="h-[87vh] pb-2 overflow-y-auto divide-y" v-if="!isLoadingInProgressProjects">
                <div v-for="project in inProgressProjectsData">
                    <ProjectPreview :project="project.project" :user="project.user" @view-project="viewProject" />
                </div>
            </div>
            <div v-else>
                <ProjectPreviewSkeleton />
            </div>
        </div>
        <div>
            <h1 class="rounded-tr-lg rounded-tl-lg mx-auto text-center bg-green-600 text-white w-2/3">DONE</h1>
            <div class="h-[87vh] pb-2 overflow-y-auto divide-y" v-if="!isLoadingDoneProjects">
                <div v-for="project in doneProjectsData">
                    <ProjectPreview :project="project.project" :user="project.user" @view-project="viewProject" />
                </div>
            </div>
            <div v-else>
                <ProjectPreviewSkeleton />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { GetProjectsByStatus } from '../../wailsjs/go/projects/ProjectsHandler';
import { projects } from '../../wailsjs/go/models';
import router from '../router';
import { todoMap } from '../consts';
import ProjectPreview from '../components/ProjectPreview.vue';
import ProjectPreviewSkeleton from '../components/ProjectPreviewSkeleton.vue';

const todoProjectsData = ref<projects.ProjectRes[]>([])
const inProgressProjectsData = ref<projects.ProjectRes[]>([])
const doneProjectsData = ref<projects.ProjectRes[]>([])

const isLoadingTodoProjects = ref(false)
const isLoadingInProgressProjects = ref(false)
const isLoadingDoneProjects = ref(false)

const getTodoProjects = async (status: number) => {
    isLoadingTodoProjects.value = true
    const res = await GetProjectsByStatus(status)
    todoProjectsData.value = res
    isLoadingTodoProjects.value = false
}
const getInProgressProjects = async (status: number) => {
    isLoadingInProgressProjects.value = true
    const res = await GetProjectsByStatus(status)
    inProgressProjectsData.value = res
    isLoadingInProgressProjects.value = false
}
const getDoneProjects = async (status: number) => {
    isLoadingDoneProjects.value = true
    const res = await GetProjectsByStatus(status)
    doneProjectsData.value = res
    isLoadingDoneProjects.value = false
}

onMounted(() => {
    getTodoProjects(todoMap["todo"])
    getInProgressProjects(todoMap["in_progress"])
    getDoneProjects(todoMap["done"])
})

const viewProject = (id: number | undefined) => {
    router.push(`/project/${id}`)
}

</script>

<style scoped>

</style>