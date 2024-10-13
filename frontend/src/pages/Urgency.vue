<template>
    <div class="grid grid-cols-3 gap-4">
        <div>
            <h1 class="rounded-tr-lg rounded-tl-lg mx-auto text-center bg-green-600 text-white w-2/3">LOW</h1>
            <div class="h-[87vh] pb-2 overflow-y-auto divide-y" v-if="!isLoadingLowUrgencyProjects">
                <div v-for="project in lowUrgencyProjects">
                    <ProjectPreview :project="project.project" :user="project.user" @view-project="viewProject" />
                </div>
            </div>
            <div v-else>
                <ProjectPreviewSkeleton />
            </div>
        </div>
        <div>
            <h1 class="rounded-tr-lg rounded-tl-lg mx-auto text-center bg-cyan-600 text-white w-2/3">NORMAL</h1>
            <div class="h-[87vh] pb-2 overflow-y-auto divide-y" v-if="!isLoadingNormalUrgencyProjects">
                <div v-for="project in normalUrgencyProjects">
                    <ProjectPreview :project="project.project" :user="project.user" @view-project="viewProject" />
                </div>
            </div>
            <div v-else>
                <ProjectPreviewSkeleton />
            </div>
        </div>
        <div>
            <h1 class="rounded-tr-lg rounded-tl-lg mx-auto text-center bg-red-500 text-white w-2/3">HIGH</h1>
            <div class="h-[87vh] pb-2 overflow-y-auto divide-y" v-if="!isLoadingHighUrgencyProjects">
                <div v-for="project in highUrgencyProjects">
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
import { GetProjectsByUrgency } from '../../wailsjs/go/projects/ProjectsHandler';
import { types } from '../../wailsjs/go/models';
import router from '../router';
import { urgencyMap } from '../consts';
import ProjectPreview from '../components/ProjectPreview.vue';
import ProjectPreviewSkeleton from '../components/ProjectPreviewSkeleton.vue';

const lowUrgencyProjects = ref<types.ProjectRes[]>([])
const normalUrgencyProjects = ref<types.ProjectRes[]>([])
const highUrgencyProjects = ref<types.ProjectRes[]>([])

const isLoadingLowUrgencyProjects = ref(false)
const isLoadingNormalUrgencyProjects = ref(false)
const isLoadingHighUrgencyProjects = ref(false)

const getTodoProjects = async (status: number) => {
    isLoadingLowUrgencyProjects.value = true
    const res = await GetProjectsByUrgency(status)
    lowUrgencyProjects.value = res
    isLoadingLowUrgencyProjects.value = false
}
const getInProgressProjects = async (status: number) => {
    isLoadingNormalUrgencyProjects.value = true
    const res = await GetProjectsByUrgency(status)
    normalUrgencyProjects.value = res
    isLoadingNormalUrgencyProjects.value = false
}
const getDoneProjects = async (status: number) => {
    isLoadingHighUrgencyProjects.value = true
    const res = await GetProjectsByUrgency(status)
    highUrgencyProjects.value = res
    isLoadingHighUrgencyProjects.value = false
}

onMounted(() => {
    getTodoProjects(urgencyMap["low"])
    getInProgressProjects(urgencyMap["normal"])
    getDoneProjects(urgencyMap["high"])
})

const viewProject = (id: number | undefined) => {
    router.push(`/project/${id}`)
}

</script>

<style scoped>

</style>