<template>
  <div
    class="fixed inset-0 bg-slate-300 bg-opacity-70 flex items-center justify-center"
  >
    <div class="w-[40rem] h-[26rem] bg-slate-100 p-14 rounded-lg shadow-xl">
      <p class="text-lg font-semibold mb-8">เพิ่มบทเรียน</p>
      <div class="mb-6 w-full">
        <div class="flex flex-row justify-between mb-2">
          <p>บทที่</p>
        </div>
        <input
          v-model="lesson_no"
          class="border-solid border rounded-md border-slate-300 h-10 w-full px-4 py-2"
          type="number"
          min="1"
        />
      </div>
      <div class="mb-6 w-full">
        <div class="flex flex-row justify-between mb-2">
          <p>ชื่อบท</p>
        </div>
        <input
          v-model="lesson_name"
          class="border-solid border rounded-md border-slate-300 h-10 w-full px-4 py-2"
          type="text"
        />
      </div>

      <div class="flex flex-row-reverse gap-5">
        <button class="w-20 h-10 text-black" @click="onPopup">
          <span>ยกเลิก</span>
        </button>
        <button
          class="w-20 h-10 bg-orange-500 rounded-lg text-white"
          @click="createClick"
        >
          <span>เพิ่ม</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      lesson_no: '',
      lesson_name: '',
    }
  },

  methods: {
    onPopup() {
      this.$emit('notify')
    },

    async createClick() {
      try {
        await this.$axios.post(
          '/lesson/sign-up',
          {
            lesson_no: parseInt(this.lesson_no),
            lesson_name: this.lesson_name,
          },
          {
            params: {
              classroom: this.$route.query.classroom,
            },
          }
        )
        alert('เพิ่มบทเรียนเรียบร้อยแล้ว')
        location.reload()
      } catch (error) {
        alert('เกิดข้อผิดพลาด')
      }
    },
  },
}
</script>
