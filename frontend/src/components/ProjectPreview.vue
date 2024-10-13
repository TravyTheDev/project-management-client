<template>
    <div class="hover:cursor-pointer hover:bg-slate-300" @click="viewProject(project?.id)">
        <p class="font-semibold text-xl whitespace-pre-line">{{ project?.title }}</p>
        <p>Assignee: {{ user?.username }}</p>
    </div>
</template>

<script setup lang="ts">
import { projects } from '../../wailsjs/go/models';

const emit = defineEmits(['view-project'])

interface props {
    project: projects.Project | undefined;
    user: projects.User | undefined;
}

const viewProject = (id: number | undefined) => {
    emit('view-project', id)
}

const props = defineProps<props>()

</script>

<style scoped></style>