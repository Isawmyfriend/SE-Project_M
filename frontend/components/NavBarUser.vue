<template>
  <div class="flex bg-white w-auto shadow-lg h-20">
    <button
      class="flex flex-row w-[22.1rem] pt-8 pb-8 pl-8 justify-center items-center"
      @click="onClick('/')"
    >
      <img class="w-14" src="@/assets/image/csLogo.png" alt="" /><span
        class="text-2xl text-gray-600 font-bold pl-2 w-80 flex"
      >
        CS-Concept
      </span>
      <img src="@/assets/image/NavLine.svg" alt="" class="pr-8 h-12" />
    </button>
  </div>
</template>

<script>
export default {
  method: {
    onClick(path) {
      this.$router.push(path)
    },
  },
}
</script>
