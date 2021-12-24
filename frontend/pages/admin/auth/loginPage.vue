<template>
  <div class="flex h-screen justify-center items-center bg-blue-100">
    <div
      class="flex flex-row h-3/4 w-2/4 shadow-2xl bg-slate-50 rounded-2xl overflow-hidden"
    >
      <div class="flex flex-col flex-auto w-3/4 h-full p-14">
        <div class="flex flex-row justify-between items-center mb-11">
          <img class="w-16" src="@/assets/image/csLogo.png" alt="" />
          <p>for Admin</p>
        </div>

        <div class="mb-16">
          <p class="text-3xl mb-2">
            ยินดีต้อนรับเข้าสู่ <span class="text-orange-600">Cs Concept</span>
          </p>
          <p>เว็บบล็อกแนะนำบทเรียนที่สามารถเข้าถึงได้ทุกที่และตลอดเวลา</p>
        </div>

        <div class="mb-9">
          <div class="flex flex-row justify-between mb-2">
            <p>บัญชีผู้ใช้งาน</p>
            <p>username ต้องใส่ @kmitl.ac.th</p>
          </div>
          <input
            v-model="email"
            class="border-solid border rounded-md border-slate-300 h-10 w-full px-4 py-2"
            type="text"
          />
        </div>

        <div class="mb-12">
          <div class="flex flex-row justify-between mb-2">
            <p>รหัสผ่าน</p>
          </div>
          <input
            v-model="password"
            class="border-solid border rounded-md border-slate-300 h-10 w-full px-4 py-2"
            type="password"
          />
        </div>

        <div class="flex flex-row-reverse justify-between">
          <button
            class="w-28 h-10 bg-orange-400 p-2 text-center rounded text-white hover:bg-orange-500"
            @click="userLogin"
          >
            Login
          </button>

          <p class="text-red-600">{{ status }}</p>
        </div>
      </div>

      <div class="shrink h-full">
        <img
          class="object-fill h-full"
          src="@/assets/image/loginPageImage.png"
          alt=""
        />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      email: '',
      password: '',
      status: '',
    }
  },
  methods: {
    async userLogin() {
      try {
        const res = await this.$auth.loginWith('local', {
          data: {
            email: this.email,
            password: this.password,
          },
        })
        this.$auth.$storage.setLocalStorage('email', res.data.email)
        this.$router.push('/admin')
      } catch (err) {
        this.status = 'อีเมลหรือรหัสผ่านไม่ถูกต้อง'
      }
    },
  },
}
</script>
