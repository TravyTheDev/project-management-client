<template>
    <div class="divide-y flex flex-col gap-1">
        <div>
            <AnimatedPlaceHolder class="mb-2" height="28px" width="100%" />
            <AnimatedPlaceHolder height="24px" width="40%" />
        </div>
        <div class="second">
            <AnimatedPlaceHolder class="mb-2" height="28px" width="100%" />
            <AnimatedPlaceHolder height="24px" width="40%" />
        </div>
        <div class="third">
            <AnimatedPlaceHolder class="mb-2" height="28px" width="100%" />
            <AnimatedPlaceHolder height="24px" width="40%" />
        </div>
        <div class="forth">
            <AnimatedPlaceHolder class="mb-2" height="28px" width="100%" />
            <AnimatedPlaceHolder height="24px" width="40%" />
        </div>
        <div class="fifth">
            <AnimatedPlaceHolder class="mb-2" height="28px" width="100%" />
            <AnimatedPlaceHolder height="24px" width="40%" />
        </div>
    </div>
</template>

<script setup lang="ts">
import AnimatedPlaceHolder from './AnimatedPlaceHolder.vue';


</script>

<style scoped>

.second {
    opacity: 0;
    animation: fadeIn 1s;
    animation-delay: 0.25s;
    animation-fill-mode: forwards;
}
.third {
    opacity: 0;
    animation: fadeIn 1s;
    animation-delay: 0.5s;
    animation-fill-mode: forwards;
}
.forth {
    opacity: 0;
    animation: fadeIn 1s;
    animation-delay: 0.75s;
    animation-fill-mode: forwards;
}
.fifth {
    opacity: 0;
    animation: fadeIn 1s;
    animation-delay: 1s;
    animation-fill-mode: forwards;
}

@keyframes fadeIn {
    from { opacity: 0; }
    to { opacity: 1; }
}

</style>