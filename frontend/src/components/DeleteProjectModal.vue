<template>
    <div @click="emit('close-modal')"
        class="absolute w-full bg-black bg-opacity-30 h-screen top-0 left-0 flex flex-col px-8">
        <div @click.stop class="relative p-4 bg-white text-black self-start m-auto top-0 bottom-0 w-[60vw]">
            <div>
                <p>Deleting</p>
                <h1 class="font-semibold text-2xl">{{ title }}</h1>
            </div>
            <div class="flex justify-around mt-8">
                <button @click="deleteProject" class="text-red-500 border border-red-500">DELETE PROJECT</button>
                <button @click="emit('close-modal')" class="text-sky-300 border border-sky-300">CANCEL</button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { DeleteProject } from '../../wailsjs/go/projects/ProjectsHandler';
import { useRouter } from 'vue-router';
const router = useRouter()

interface props {
    title: string | undefined;
    parentID: string | undefined;
}

const props = defineProps<props>()
const emit = defineEmits(['close-modal'])

const deleteProject = async () => {
   await DeleteProject(Number(props.parentID))
   router.push('/main')
}
</script>

<style scoped>

</style>