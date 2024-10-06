<template>
    <div @click="emit('close-modal')" class="absolute w-full bg-black bg-opacity-30 h-screen top-0 left-0 flex flex-col px-8">
        <div @click.stop class="relative p-4 bg-white text-black self-start m-auto top-0 bottom-0 w-auto">
            <div>
                <p>Parent</p>
                <input v-model="project.parentID" type="number">
                <p>Title</p>
                <input v-model="project.title" type="text">
                <p>Description</p>
                <textarea v-model="project.description"></textarea>
                <p>Status</p>
                <input v-model="project.status" type="number">
                <p>Assignee</p>
                <input v-model="project.assigneeID" type="number">
                <p>Urgency</p>
                <input v-model="project.urgency" type="number">
                <p>Start</p>
                <input v-model="project.startDate" type="datetime">
                <p>End</p>
                <input v-model="project.endDate" type="datetime">
                <button @click="createProject">Create project</button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import { CreateProject } from '../../wailsjs/go/main/ProjectsHandler';
import { main } from '../../wailsjs/go/models';

const emit = defineEmits(['close-modal'])

const project: main.ProjectReq = reactive({
    parentID: 0,
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
   await CreateProject(project)
   window.location.href = "/"
}
</script>

<style scoped>

</style>