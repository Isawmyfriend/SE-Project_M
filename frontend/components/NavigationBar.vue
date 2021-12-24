<template>
  <div class="flex bg-white w-auto shadow-lg h-20">
    <button
      class="flex flex-row w-[22.1rem] pt-8 pb-8 pl-8 justify-center items-center"
      @click="onClick('/admin')"
    >
      <img class="w-14" src="@/assets/image/csLogo.png" alt="" /><span
        class="text-2xl text-gray-600 font-bold pl-2 w-80 flex"
      >
        CS-Concept
      </span>
      <img src="@/assets/image/NavLine.svg" alt="" class="pr-8 h-12" />
    </button>
    <div class="flex w-1/3 items-center"></div>
    <div class="flex w-2/4 flex-row-reverse items-center pr-10">
      <button
        class="flex items-center text-2xl bg-gray-200 rounded-full w-auto h-10 text-gray-600 justify-center p-5 hover:bg-gray-300"
        @click="logout"
      >
        <span class="text-lg">ออกจากระบบ</span>
      </button>
      <div class="flex items-center justify-center pr-10">
        <button
          class="flex flex-row justify-center items-center bg-gray-200 rounded-full w-auto h-10 text-2xl text-gray-600 p-5 hover:bg-gray-300"
          @click="onClick('/admin/classroom')"
        >
          <span class="text-lg"> {{ email }} </span>
          <span>
            <img src="@/assets/image/email_icon.svg" class="pl-3 w-8" alt="" />
          </span>
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      email: '',
    }
  },

  async created() {
    const email = await this.$auth.$storage.getLocalStorage('email')
    this.email = email
  },

  methods: {
    async logout() {
      await this.$auth.logout()
      this.$router.push('/admin/auth/loginPage')
      this.$auth.$storage.removeLocalStorage('email')
    },

    onClick(path) {
      this.$router.push(path)
    },
  },
}
</script>
